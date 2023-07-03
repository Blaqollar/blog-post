package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/blaqollar/blog-post/app/data"
	"github.com/gorilla/mux"
)

type Articles struct {
	l *log.Logger
}

func NewArticles(l *log.Logger) *Articles {
	return &Articles{l}
}


func (a *Articles) GetArticles(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle Post Article")

	la := data.GetArticles()
	err := la.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

func (a *Articles) AddArticle(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle Post Article")

	content := &data.Article{}

	err := content.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Marshall Json", http.StatusBadRequest)
	}

	data.AddArticle(content)

}

func (a *Articles) UpdateArticles(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	a.l.Println("Handle Post Article", id)

	prod := &data.Article{}

	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Marshall Json", http.StatusBadRequest)
	}

	err = data.UpdateArticles(id, prod)
	if err == data.ErrArticleNotFound {
		http.Error(rw, "Article not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Article not found", http.StatusNotFound)
		return
	}

}
