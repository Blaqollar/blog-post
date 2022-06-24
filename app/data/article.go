package data

import (
	"encoding/json"
	"fmt"
	"io"
)

// Article defines the structure for an API Article
//You use struct tags to name the json in lowercase
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a collection of Article
// Rather than returning the slice of Article as bytes
//You can define a type which is a slice of Article (You do this so you can the add a method to it in the same way you do it to a struct)
//You add a method called Tojson which contains an io writer which also returns an error
//You then encapsulate all of the logic for encoding the JSON into this method
//By creating a new encoder that you'll parse the writer then encodes the JSON into p
type Articles []*Article

func (a *Articles) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

//In the above we used the encoder to encode the JSON starightout to the responsewriter
//We use the reverse to be able create a Article structure from the JSON we parsed in
//You will add a FromJson func a method to the Article struct were the source would be an io reader
//To do that we use Decoder to decode the body from the request with an IO reader(inserted in the func of Newdecoder) into a struct.
func (a *Article) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

//This adds a method to your data object (Article)
//This sends all of the properties of Article into the reader depending on the http request
//This is a data access module for (databases, textfiles or fixed data)
func GetArticles() Articles {
	return articleList
}

//This save the data to our datastore
//And also pass in your Article
func AddArticle(a *Article) {
	a.ID = getNextID()
	articleList = append(articleList, a)
}

//You have to generate an ID to add Articles to the datastore
//Gets the last Article on the list then return by adding the Article to the list
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
