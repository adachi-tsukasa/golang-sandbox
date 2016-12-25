package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Feed struct {
	Title       string   `xml:"channel>title"`
	ItemTitle   []string `xml:"channel>item>title"`
	Description []string `xml:"channel>item>description"`
}

type RDFFeed struct {
	Title       string   `xml:"channel>title"`
	ItemTitle   []string `xml:"item>title"`
	Description []string `xml:"item>description"`
}

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type Characters []Character

type FeedList struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}
type FeedLists []FeedList

func init() {

	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(martini.Static("static/dist"))
	m.Get("/charList", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, getCharList(res, req))
	})
	m.Get("/feed/:id", func(params martini.Params, r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, feedParser(params["id"], res, req))
	})

	http.Handle("/", m)
}

func getCharList(w http.ResponseWriter, r *http.Request) Characters {
	file := loadJSONFile("./json/charlist.json")

	var chars Characters
	jsonErr := json.Unmarshal(file, &chars)
	if jsonErr != nil {
		fmt.Println("Format Error: ", jsonErr)
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	return chars
}

func loadJSONFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}
	return file
}

func feedParser(paramID string, w http.ResponseWriter, r *http.Request) []string {
	var retunableStringArrays []string
	var feedlists FeedLists
	file := loadJSONFile("./json/feedlist.json")
	jsonErr := json.Unmarshal(file, &feedlists)
	if jsonErr != nil {
		fmt.Println("Format Error: ", jsonErr)
	}
	for _, feedlist := range feedlists {
		id, _ := strconv.Atoi(paramID)
		if feedlist.ID == id {
			wh, err := getFeed(feedlist.Url, r)
			if err != nil {
				log.Fatalf("Log: %v", err)
				return nil
			}
			// for n, v := range wh.ItemTitle {
			// 	if n > 0 {
			// 		fmt.Printf("%s \n", v)
			// 	}
			// }
			retunableStringArrays = wh.ItemTitle
		}
	}
	// wh, err := getFeed("http://rssblog.ameba.jp/eriko-nakamura-blog/rss20.xml", r)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	return retunableStringArrays
}

func getFeed(feed string, r *http.Request) (p *Feed, err error) {

	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	res, err := client.Get(feed)

	// res, err := http.Get(feed)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	wh := new(Feed)
	err = xml.Unmarshal(b, &wh)

	return wh, err
}
