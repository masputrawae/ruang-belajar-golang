package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	_ "github.com/yuin/goldmark"
)

var (
	templates map[string]*template.Template
	contents  map[string]bytes.Buffer
)

type Data struct {
	SiteTitle string
	Content   template.HTML
	Nav       []Navigation
}

type Navigation struct {
	URL  string
	Name string
}

var SiteData = Data{
	Nav: []Navigation{
		{
			URL:  "/",
			Name: "Home",
		},
		{
			URL:  "/about",
			Name: "About",
		},
	},
}

func main() {
	if err := loadTemplate("templates"); err != nil {
		panic(err)
	}
	if err := loadContent("content"); err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		SiteData.SiteTitle = "Home"
		if err := render(w, "home.html", SiteData); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	})

	mux.HandleFunc("GET /{page}", func(w http.ResponseWriter, r *http.Request) {
		page := r.PathValue("page") + ".html"
		SiteData.SiteTitle = "About"
		if err := render(w, page, SiteData); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	})

	mux.HandleFunc("GET /post/{page}", func(w http.ResponseWriter, r *http.Request) {
		file := r.PathValue("page")
		buf, ok := contents[file]
		if !ok {
			http.Error(w, "file not found", http.StatusNotFound)
		}
		SiteData.SiteTitle = file
		SiteData.Content = template.HTML(buf.String())
		if err := render(w, "content.html", SiteData); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
	})

	slog.Info("http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func loadContent(root string) error {
	contents = make(map[string]bytes.Buffer)
	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".md") {
			name := d.Name()
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			var buf bytes.Buffer
			if err := goldmark.Convert(file, &buf); err != nil {
				return err
			}
			contents[name] = buf
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

// ===============
func render(w http.ResponseWriter, page string, data any) error {
	tmpl, ok := templates[page]
	if !ok {
		return errors.New("template not found")
	}
	return tmpl.ExecuteTemplate(w, "index.html", data)
}

func loadTemplate(root string) error {
	templates = make(map[string]*template.Template)
	files := []string{}
	pages := filepath.Join(root, "pages", "*.html")

	if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && !strings.HasPrefix(path, filepath.Join(root, "pages")) {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return err
	}

	pageFiles, err := filepath.Glob(pages)
	if err != nil {
		return err
	}

	for i := range pageFiles {
		name := filepath.Base(pageFiles[i])
		all := append(files, pageFiles[i])
		tmpl, err := template.ParseFiles(all...)
		if err != nil {
			return err
		}
		templates[name] = tmpl
	}

	fmt.Println(files, templates)
	return nil
}
