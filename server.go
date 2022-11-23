package main

import (
	"log"

	"github.com/levigross/grequests"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"
)

// TemplateRenderer is a custom html/template renderer for Echo framework

func main() {
	renderEngine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views:   renderEngine,
		AppName: "services-go",
	})

	app.Use(cache.New(cache.Config{
		Expiration: 1 * time.Minute,
	}))

	ratelimit := limiter.New(limiter.Config{
		Max:        10,
		Expiration: 5 * time.Minute,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(429).SendString("Please stop refreshing the page so much. Wait 5 minutes and try again.")
		},
	})

	app.Get("/", ratelimit, func(c *fiber.Ctx) error {
		memestream, err := grequests.Get("https://ms.odyssey346.dev", nil)
		log.Println("Memestream status:", memestream.StatusCode)
		if err != nil {
			return c.SendString("Something broke")
		}
		inv, err := grequests.Get("https://inv.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Invidious status:", inv.StatusCode)
		libreddit, err := grequests.Get("https://lr.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Libreddit status:", libreddit.StatusCode)
		quetre, err := grequests.Get("https://qtr.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Quetre status:", quetre.StatusCode)
		breezewiki, err := grequests.Get("https://bw.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Breezewiki status:", breezewiki.StatusCode)
		rimgo, err := grequests.Get("https://rim.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Rimgo status:", rimgo.StatusCode)
		proxitok, err := grequests.Get("https://proxitok.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Proxitok status:", proxitok.StatusCode)
		nitter, err := grequests.Get("https://ntr.odyssey346.dev", nil)
		if err != nil {
			return c.SendString("Something broke")
		}
		log.Println("Nitter status:", nitter.StatusCode)
		dt := time.Now()
		// cache stuff
		return c.Render("root", fiber.Map{
			"memestream": memestream.StatusCode,
			"invidious":  inv.StatusCode,
			"libreddit":  libreddit.StatusCode,
			"quetre":     quetre.StatusCode,
			"breezewiki": breezewiki.StatusCode,
			"rimgo":      rimgo.StatusCode,
			"proxitok":   proxitok.StatusCode,
			"nitter":     nitter.StatusCode,
			"time":       dt.Format("2006-01-02 15:04:05"),
		})
	})

	app.Get("/502", func(c *fiber.Ctx) error {
		return c.Render("502", fiber.Map{
			"Title": "502 Bad Gateway",
		})
	})

	app.Static("/style.css", "templates/style.css")

	log.Fatal(app.Listen(":8000"))
}
