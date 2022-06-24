package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/blaqollar/blog-post/app/data"
	"github.com/gorilla/mux"
)

// Articles is a http.Handler
type Articles struct {
	l *log.Logger
}

// NewArticles creates a method for the Articles handler with the given logger
func NewArticles(l *log.Logger) *Articles {
	return &Articles{l}
}

// handle the request for a list of Articles
// fetch the Articles from the datastore with lp
// getArticles returns the Articles from the data store
// serialize the list to JSON This marshalls the JSON into your Article object
//You use json encoder instead of marshal is to avoid allocatimg memory to the slice of data
//You parse writer to your encoder to write direct to it

func (a *Articles) GetArticles(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle Post Article")

	la := data.GetArticles()
	err := la.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to Marshal JSON", http.StatusInternalServerError)
	}
}

//Create a post method function logic for addArticle
//We need to take the data from the post and convert it to the Article object
func (a *Articles) AddArticle(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Handle Post Article")

	//create new Article object then we need to pass it that reader from the response body in the httprequest
	content := &data.Article{}

	err := content.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to Marshall Json", http.StatusBadRequest)
	}

	data.AddArticle(content)

}

func (a *Articles) UpdateArticles(rw http.ResponseWriter, r *http.Request) {
	//mux.vars picks out the variable and defines the placeholder in the mux URI

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	a.l.Println("Handle Post Article", id)

	//create new Article object then we need to pass it that reader from the response body in the httprequest
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
