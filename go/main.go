package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

type Session struct {
	Username string `json:"username"`
	Template string `json:"template"`
}

func (s Session) Log() string {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo %s", s.Username))
	out, _ := cmd.CombinedOutput()
	return string(out)
}

type handler struct {
	template string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var s Session
	c, err := r.Cookie("session")
	if err != nil {
		s = Session{"", "CHANGE ME"}
	} else {
		v, err := url.QueryUnescape(c.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err = json.Unmarshal([]byte(v), &s); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		v, err := json.Marshal(
			Session{r.FormValue("username"), r.FormValue("template")},
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    url.QueryEscape(string(v)),
			Path:     "/",
			SameSite: http.SameSiteStrictMode,
			HttpOnly: true,
			MaxAge:   60 * 60,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	t, err := template.New("index.html").Parse(
		strings.Replace(h.template, "CHANGE ME", s.Template, 1),
	)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	if err = t.Execute(w, s); err != nil {
		panic(err)
	}
}

func main() {
	t, err := os.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	if err = http.ListenAndServe(":4444", &handler{string(t)}); err != nil {
		panic(err)
	}
}
