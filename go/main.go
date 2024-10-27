package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
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

func (u User) Log() string {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo %s", u.Name))
	out, _ := cmd.CombinedOutput()
	return string(out)
}

type session struct {
	user     User
	template string
}

type server struct {
	template string
	sessions map[string]session
	mutex    sync.Mutex
}

func createKey() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func createCookie(k string) *http.Cookie {
	return &http.Cookie{
		Name:     "session",
		Value:    k,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   60 * 60,
	}
}

func (s *server) createSession() (session, string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	k := createKey()
	ss := session{template: "CHANGE ME"}
	s.sessions[k] = ss
	return ss, k
}

func (s *server) getSession(k string) (session, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	ss, ok := s.sessions[k]
	return ss, ok
}

func (s *server) editSession(k string, ss session) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, ok := s.sessions[k]; !ok {
		return false
	}
	s.sessions[k] = ss
	return true
}

func (s *server) handler(w http.ResponseWriter, r *http.Request) {
	var ss session
	var ok bool
	c, err := r.Cookie("session")
	if err != nil {
		ok = false
	} else {
		ss, ok = s.getSession(c.Value)
	}
	if !ok {
		var k string
		ss, k = s.createSession()
		c = createCookie(k)
		http.SetCookie(w, c)
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		s.editSession(
			c.Value,
			session{User{r.FormValue("name")}, r.FormValue("template")},
		)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t, err := template.New("index.html").Parse(
		strings.Replace(s.template, "CHANGE ME", ss.template, 1),
	)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	if err = t.Execute(w, ss.user); err != nil {
		panic(err)
	}
}

func main() {
	t, err := os.ReadFile("templates/index.html")
	if err != nil {
		panic(err)
	}
	s := &server{template: string(t), sessions: make(map[string]session)}
	http.ListenAndServe(":4444", http.HandlerFunc(s.handler))
}
