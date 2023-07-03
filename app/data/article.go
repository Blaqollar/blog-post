package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Article defines the structure for an API Article
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a collection of Article
type Articles []*Article

func (a *Articles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

// This encodes the JSON starightout to the responsewriter
func (a *Article) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

// This adds a method to your data object (Article)
func GetArticles() Articles {
	return articleList
}

// This save the data to our datastore
// And also pass in your Article
func AddArticle(a *Article) {
	a.ID = getNextID()
	articleList = append(articleList, a)
}


// This gets the last Article on the list then return by adding the Article to the list
func getNextID() int {
	la := articleList[len(articleList)-1]
	return la.ID + 1
}

//This function updates the list of articles
func UpdateArticles(id int, a *Article) error {
	_, pos, err := findArticle(id)
	if err != nil {
		return err
	}
	a.ID = id
	articleList[pos] = a
	return nil
}

var ErrArticleNotFound = fmt.Errorf("Article not found")

// This function finds the article with the given ID and returns it.
func findArticle(id int) (*Article, int, error) {
	for i, a := range articleList {
		if a.ID == id {
			return a, i, nil
		}
	}
	return nil, -1, ErrArticleNotFound
}

// ArticleList is a hard coded list of Articles for this
// example data source
var articleList = []*Article{
	&Article{
		ID:      1,
		Title:   "Hello",
		Desc:    "Article Description",
		Content: "Article Content",
	},
	&Article{
		ID:      2,
		Title:   "Hello 2",
		Desc:    "Article Description",
		Content: "Article Content",
	},
}
