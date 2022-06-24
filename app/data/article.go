package data

import (
	"encoding/json"
	"fmt"
	"io"
)


type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []*Article

func (a *Articles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *Article) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func GetArticles() Articles {
	return articleList
}

func AddArticle(a *Article) {
	a.ID = getNextID()
	articleList = append(articleList, a)
}

func getNextID() int {
	la := articleList[len(articleList)-1]
	return la.ID + 1
}

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

func findArticle(id int) (*Article, int, error) {
	for i, a := range articleList {
		if a.ID == id {
			return a, i, nil
		}
	}
	return nil, -1, ErrArticleNotFound
}

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
