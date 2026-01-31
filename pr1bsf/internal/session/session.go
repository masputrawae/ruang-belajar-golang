package session

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

var store = map[string]string{}

func generateSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func Create(w http.ResponseWriter, username string) {
	id := generateSessionID()
	store[id] = username

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    id,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}

func Get(r *http.Request) (string, bool) {
	c, err := r.Cookie("session_id")
	if err != nil {
		return "", false
	}
	u, ok := store[c.Value]
	return u, ok
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_id")
	if err == nil {
		delete(store, c.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "session_id",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
}
