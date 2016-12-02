package api

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"io/ioutil"

	"encoding/json"

	"strconv"

	"github.com/SlyMarbo/rss"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Feeder struct {
	ID  int    `json:"id"`
	Url string `json:"url"`
}
type Characters []Character
type Feeders []Feeder

func init() {

	m := martini.Classic()
	m.Use(render.Renderer())
	// m.Get("/", func(r render.Render) {
	// 	r.Redirect("static/index.html", 302)
	// })
	m.Use(martini.Static("static/dist"))
	m.Get("/charList", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		// r.JSON(200, map[string]interface{}{"hello": "world"})
		r.JSON(200, getCharList(res, req))
	})

	m.Get("/feed/:id", func(params martini.Params, r render.Render, res http.ResponseWriter, req *http.Request) {
		r.JSON(200, feedParser(params["id"], res, req))
	})

	// for request test
	m.Get("/hoge", func(r render.Render, res http.ResponseWriter, req *http.Request) {
		// c := appengine.NewContext(req)
		// log.Infof(c, "START")
		r.JSON(200, testFeedParser(res, req))
		// r.JSON(200, "hoge")

	})

	http.Handle("/", m)
}

func redirect(w http.ResponseWriter, r *http.Request) {

}

func getCharList(w http.ResponseWriter, r *http.Request) Characters {

	file := loadJSONFile("./json/charlist.json")

	var chars Characters
	jsonErr := json.Unmarshal(file, &chars)
	if jsonErr != nil {
		fmt.Println("Format Error: ", jsonErr)
	}

	// fmt.Fprint(w, chars)
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

func feedParser(paramID string, w http.ResponseWriter, r *http.Request) rss.Feed {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	var feed rss.Feed
	var feeders Feeders
	file := loadJSONFile("./json/feedlist.json")
	jsonErr := json.Unmarshal(file, &feeders)
	if jsonErr != nil {
		fmt.Println("Format Error: ", jsonErr)
	}
	for _, feeder := range feeders {
		id, _ := strconv.Atoi(paramID)
		if feeder.ID == id {
			rowFeed, err := rss.FetchByClient(feeder.Url, client)
			if err != nil {
				fmt.Println("Parse Error: ", err)
			}
			feed = *rowFeed
			break
		}
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	return feed
}

// for feed parsing test
func testFeedParser(w http.ResponseWriter, r *http.Request) []*rss.Item {
	ctx := appengine.NewContext(r)
	client := urlfetch.Client(ctx)

	rowFeed, err := rss.FetchByClient("http://blog.livedoor.jp/ubiquitous777/index.rdf", client)
	if err != nil {
		fmt.Println("Parse Error: ", err)
	}
	// fmt.Fprint(w, rowFeed.Items)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	return rowFeed.Items
}
