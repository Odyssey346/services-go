package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/levigross/grequests"

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
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		memestream, err := grequests.Get("https://ms.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		inv, err := grequests.Get("https://inv.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		libreddit, err := grequests.Get("https://lr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		quetre, err := grequests.Get("https://qtr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		breezewiki, err := grequests.Get("https://bw.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		thiswebsite, err := grequests.Get("https://services.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		rimgo, err := grequests.Get("https://rim.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		proxitok, err := grequests.Get("https://proxitok.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		nitter, err := grequests.Get("https://ntr.odyssey346.dev", nil)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		dt := time.Now()
		// cache stuff
		return c.Render(http.StatusOK, "root.html", map[string]interface{}{
			"memestream":  memestream.StatusCode,
			"invidious":   inv.StatusCode,
			"thiswebsite": thiswebsite.StatusCode,
			"libreddit":   libreddit.StatusCode,
			"quetre":      quetre.StatusCode,
			"breezewiki":  breezewiki.StatusCode,
			"rimgo":       rimgo.StatusCode,
			"proxitok":    proxitok.StatusCode,
			"nitter":      nitter.StatusCode,
			"time":        dt.Format("2006-01-02 15:04:05"),
		})
	})

	e.GET("/502", func(c echo.Context) error {
		return c.Render(http.StatusOK, "502.html", nil)
	})

	e.File("/style.css", "templates/style.css")

	e.Logger.Fatal(e.Start(":8000"))
}
