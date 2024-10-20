package main

import (
	"crypto/rand"
	"encoding/hex"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type User struct {
	Name string
}

func (u User) Exec(name string, arg ...string) string {
	out, _ := exec.Command(name, arg...).CombinedOutput()
	return string(out)
}

type Session struct {
	User     User
	Template string
}

type Server struct {
	template string
	sessions map[string]Session
	mutex    sync.Mutex
}

func (s *Server) CreateSession() (Session, string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	k := hex.EncodeToString(b)
	sess := Session{}
	s.sessions[k] = sess
	return sess, k
}

func (s *Server) GetSession(key string) (Session, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	sess, ok := s.sessions[key]
	return sess, ok
}

func (s *Server) EditSession(key string, sess Session) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.sessions[key]; !ok {
		return false
	}
	s.sessions[key] = sess
	return true
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	var sess Session
	c, err := r.Cookie("session")
	if err != nil {
		var key string
		sess, key = s.CreateSession()
		c = &http.Cookie{
			Name:     "session",
			Value:    key,
			Path:     "/",
			HttpOnly: true,
			MaxAge:   60 * 60,
		}
		http.SetCookie(w, c)
	} else {
		var ok bool
		sess, ok = s.GetSession(c.Value)
		if !ok {
			var key string
			sess, key = s.CreateSession()
			c = &http.Cookie{
				Name:     "session",
				Value:    key,
				Path:     "/",
				HttpOnly: true,
				MaxAge:   60 * 60,
			}
			http.SetCookie(w, c)
		}
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.EditSession(
			c.Value,
			Session{User{r.FormValue("name")}, r.FormValue("template")},
		)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t, err := template.New("index.html").Parse(strings.Replace(s.template, "CHANGE ME", sess.Template, 1))
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	if err = t.Execute(w, sess.User); err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}
	s := &Server{template: string(f), sessions: make(map[string]Session)}
	http.ListenAndServe(":4444", http.HandlerFunc(s.Handler))
}
