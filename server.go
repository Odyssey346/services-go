package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levigross/grequests"
	"github.com/patrickmn/go-cache"

	"time"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	cash := cache.New(5*time.Minute, 10*time.Minute)

	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		ms, err := grequests.Get("https://ms.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("ms", ms.StatusCode, cache.DefaultExpiration)
		inv, err := grequests.Get("https://inv.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("inv", inv.StatusCode, cache.DefaultExpiration)
		libreddit, err := grequests.Get("https://lr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("libreddit", libreddit.StatusCode, cache.DefaultExpiration)
		quetre, err := grequests.Get("https://qtr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("quetre", quetre.StatusCode, cache.DefaultExpiration)
		breezewiki, err := grequests.Get("https://bw.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("breezewiki", breezewiki.StatusCode, cache.DefaultExpiration)
		thiswebsite, err := grequests.Get("https://services.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("thiswebsite", thiswebsite.StatusCode, cache.DefaultExpiration)
		rimgo, err := grequests.Get("https://rim.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("rimgo", rimgo.StatusCode, cache.DefaultExpiration)
		proxitok, err := grequests.Get("https://proxitok.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("proxitok", proxitok.StatusCode, cache.DefaultExpiration)
		nitter, err := grequests.Get("https://ntr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		cash.Set("nitter", nitter.StatusCode, cache.DefaultExpiration)
		dt := time.Now()
		// cache stuff
		cash.Set("dt", dt.Format("2006-01-02 15:04:05"), cache.DefaultExpiration)
		msCache, found := cash.Get("ms")
		if found {
			fmt.Println("Found ms in cache", msCache.(int))
		}
		invCache, found := cash.Get("inv")
		if found {
			fmt.Println("Found inv in cache", invCache.(int))
		}
		libredditCache, found := cash.Get("libreddit")
		if found {
			fmt.Println("Found libreddit in cache", libredditCache.(int))
		}
		quetreCache, found := cash.Get("quetre")
		if found {
			fmt.Println("Found quetre in cache", quetreCache.(int))
		}
		breezewikiCache, found := cash.Get("breezewiki")
		if found {
			fmt.Println("Found breezewiki in cache", breezewikiCache.(int))
		}
		thiswebsiteCache, found := cash.Get("thiswebsite")
		if found {
			fmt.Println("Found thiswebsite in cache", thiswebsiteCache.(int))
		}
		rimgoCache, found := cash.Get("rimgo")
		if found {
			fmt.Println("Found rimgo in cache", rimgoCache.(int))
		}
		proxitokCache, found := cash.Get("proxitok")
		if found {
			fmt.Println("Found proxitok in cache", proxitokCache.(int))
		}
		nitterCache, found := cash.Get("nitter")
		if found {
			fmt.Println("Found nitter in cache", nitterCache.(int))
		}
		timeCache, found := cash.Get("dt")
		if found {
			fmt.Println("Found dt in cache", timeCache.(string))
		}
		return c.Render(http.StatusOK, "root.html", map[string]interface{}{
			"memestream":  msCache.(int),
			"invidious":   invCache.(int),
			"thiswebsite": thiswebsiteCache.(int),
			"libreddit":   libredditCache.(int),
			"quetre":      quetreCache.(int),
			"breezewiki":  breezewikiCache.(int),
			"rimgo":       rimgoCache.(int),
			"proxitok":    proxitokCache.(int),
			"nitter":      nitterCache.(int),
			"time":        timeCache.(string),
		})
	})

	e.GET("/502", func(c echo.Context) error {
		return c.Render(http.StatusOK, "502.html", nil)
	})

	e.File("/style.css", "templates/style.css")

	e.Logger.Fatal(e.Start(":8000"))
}
