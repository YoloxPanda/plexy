package plexy

import (
	"fmt"
	"log"
	"net/http"
)

type Params struct {
	params map[string]string
}

func newParams() *Params {
	return &Params{
		params: map[string]string{},
	}
}

func (p *Params) Get(key string) string {
	return p.params[key]
}

type PlexyHandler interface {
	Handle(w http.ResponseWriter, r *http.Request, params *Params)
}

type Plexy interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Handle(path string, handler PlexyHandler)
	HandleFunc(path string, handler func(w http.ResponseWriter, r *http.Request, params *Params))
}

type funcPlexyHandler struct {
	handler func(w http.ResponseWriter, r *http.Request, params *Params)
}

func (f funcPlexyHandler) Handle(w http.ResponseWriter, r *http.Request, params *Params) {
	f.handler(w, r, params)
}

type defaultPlexy struct {
	ph *pathHandler
}

func (d *defaultPlexy) Handle(path string, handler PlexyHandler) {
	n := constructNode(path, handler)
	d.ph.addPath(n)
}

func (d *defaultPlexy) HandleFunc(path string, handler func(w http.ResponseWriter, r *http.Request, params *Params)) {
	n := constructNode(path, funcPlexyHandler{handler})
	d.ph.addPath(n)
}

type loggingResponseWriter struct {
	http.ResponseWriter
	status int
}

func (l *loggingResponseWriter) WriteHeader(status int) {
	l.status = status
	l.ResponseWriter.WriteHeader(status)
}

func (p *defaultPlexy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// we want to log the request code so we will use our own ResponseWriter to capture the status
	lrw := &loggingResponseWriter{w, http.StatusOK}

	n, params := p.ph.matchPath(r.URL.Path)

	if n == nil || n.handler == nil {
		http.NotFound(lrw, r)
		return
	}

	n.handler.Handle(lrw, r, params)

	log.Printf("%s - %d - %s\n", http.StatusText(lrw.status), lrw.status, r.URL.Path)
}

func NewPlexy() Plexy {
	return &defaultPlexy{
		ph: &pathHandler{},
	}
}

func main() {
	plexy := NewPlexy()

	plexy.HandleFunc("/auth", auth)
	plexy.HandleFunc("/user/:username/settings", func(w http.ResponseWriter, r *http.Request, params *Params) {
		username := params.Get("username")
		fmt.Println("Username:", username)
	})
	plexy.Handle("/test", authHandler{})

	http.ListenAndServe(":8080", plexy)
}

func auth(w http.ResponseWriter, r *http.Request, params *Params) {

}

type authHandler struct {
}

func (authHandler) Handle(w http.ResponseWriter, r *http.Request, params *Params) {

}
