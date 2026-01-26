package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

type MdContent struct {
	Tags    []string
	Date    string
	Lastmod string
	Content template.HTML
}

func ResolveMdFile(fName string) (string, error) {
	var file string
	files, err := os.ReadDir("content")
	if err != nil {
		return "", err
	}

	for _, v := range files {
		if v.Name() == fName {
			file = v.Name()
		}
	}

	return file, nil
}

func MarkdownParsed(fName string) (MdContent, error) {
	var data MdContent
	var buf bytes.Buffer

	md := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
		),
	)
	file, err := os.ReadFile(fName)

	if err != nil {
		return MdContent{}, err
	}

	context := parser.NewContext()
	if err := md.Convert(file, &buf, parser.WithContext(context)); err != nil {
		return MdContent{}, err
	}
	metaData := meta.Get(context)

	if v, ok := metaData["tags"].([]interface{}); ok {
		for _, tag := range v {
			if t, ok := tag.(string); ok {
				data.Tags = append(data.Tags, t)
			}
		}
	}

	if v, ok := metaData["created_at"].(string); ok {
		data.Date = v
	}
	if v, ok := metaData["updated_at"].(string); ok {
		data.Lastmod = v
	}

	data.Content = template.HTML(buf.String())

	return data, nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./view/assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fileServer))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	r.Get("/content/", func(w http.ResponseWriter, r *http.Request) {
		files, _ := os.ReadDir("content")

		for i, v := range files {
			w.Write([]byte(strconv.Itoa(i) + "	" + v.Name() + "\n"))
		}
	})

	r.Get("/content/{file_name}", func(w http.ResponseWriter, r *http.Request) {
		fileNameParam := chi.URLParam(r, "file_name")
		fileName, err := ResolveMdFile(fileNameParam)

		if err != nil {
			w.Write([]byte(err.Error()))
		}

		content, err := MarkdownParsed("content/" + fileName)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		tmpl := template.Must(template.ParseFiles("view/pages/index.html"))
		tmpl.Execute(w, content)
	})

	log.Default().Println("Server running: http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
