package main

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const port = 5555

type Data struct {
	Username string
}

/*
This method is meant to be used later for logging users using our app.
There is nothing wrong with it being here, right?
*/
func (d Data) Log() string {
	cmd := exec.Command("sh", "-c", "echo "+d.Username+" is cool!")
	out, _ := cmd.CombinedOutput()
	return string(out)
}

type handler struct {
	template string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	s := h.template
	if q.Has("template") {
		s = strings.Replace(s, "CHANGE ME", q.Get("template"), 1)
	}
	t, err := template.New("index.html").Parse(s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var b bytes.Buffer
	if err = t.Execute(&b, Data{q.Get("username")}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	io.Copy(w, &b)
}

func main() {
	t, err := os.ReadFile("index.html")
	if err != nil {
		panic(err)
	}
	f := make([]byte, 16)
	if _, err := rand.Read(f); err != nil {
		panic(err)
	}
	if err = os.Setenv("FLAG", fmt.Sprintf("SSTI{%s}\n", hex.EncodeToString(f))); err != nil {
		panic(err)
	}
	fmt.Printf("server is running on port %d\n", port)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), &handler{string(t)}); err != nil {
		panic(err)
	}
}
