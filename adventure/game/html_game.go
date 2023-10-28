package game

import (
	"html/template"
	"log"
	"net/http"
)

type storyHandler struct {
	adventure *Adventure
}

func (h storyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	p := req.URL.Path
	if p == "" || p == "/" {
		p = "/intro"
	}
	p = p[1:]
	entry, ok := h.adventure.EntryMap[p]
	if !ok {
		http.Error(res, "Uh Oh", http.StatusInternalServerError)
		return
	}
	entry.ServeHTTP(res, req)
}

func (a *Adventure) NewHandler() http.Handler {
	return storyHandler{a}
}

func (e Entry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var tmplFile = "adventure_template.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, e)
	if err != nil {
		panic(err)
	}
}

func (a *Adventure) PlayHtmlAdventure() {
	handler := a.NewHandler()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
