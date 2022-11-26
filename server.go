package main

import (
	"log"

	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/template/html"

	"os"

	"github.com/containrrr/shoutrrr"
)

// TemplateRenderer is a custom html/template renderer for Echo framework

func main() {
	renderEngine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views:                   renderEngine,
		AppName:                 "services-go",
		EnableTrustedProxyCheck: true,
		TrustedProxies:          []string{"*"},
		ProxyHeader:             fiber.HeaderXForwardedFor,
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
		dt := time.Now()
		// cache stuff
		return c.Render("root", fiber.Map{
			"memestream": Ping("https://ms.odyssey346.dev"),
			"invidious":  Ping("https://inv.odyssey346.dev"),
			"libreddit":  Ping("https://lr.odyssey346.dev"),
			"quetre":     Ping("https://qtr.odyssey346.dev"),
			"breezewiki": Ping("https://bw.odyssey346.dev"),
			"rimgo":      Ping("https://rim.odyssey346.dev"),
			"proxitok":   Ping("https://proxitok.odyssey346.dev"),
			"nitter":     Ping("https://ntr.odyssey346.dev"),
			"time":       dt.Format("2006-01-02 15:04:05"),
		})
	})

	app.Get("/feedback", func(c *fiber.Ctx) error {
		return c.Render("feedback", fiber.Map{
			"title": "Feedback",
		})
	})

	app.Post("/api/feedback", func(c *fiber.Ctx) error {
		service := c.FormValue("service")
		email := c.FormValue("email")
		message := c.FormValue("message")
		if service == "" || email == "" || message == "" {
			return c.Status(400).SendString("Please fill out all fields.")
		}

		url := os.Getenv("SHOUTRRR_URL")
		if url == "" {
			return c.Status(500).SendString("Shoutrrr URL not set. (this is a server-side issue.)")
		}
		err := shoutrrr.Send(url, "Feedback for "+service+" ("+email+", "+c.IP()+"):\n"+message)
		if err != nil {
			return c.Status(500).SendString("Error sending feedback. (this is a server-side issue.)")
		}
		log.Printf("Feedback from %s (%s) for %s: %s", email, c.IP(), service, message)
		return c.SendString("Thanks for your feedback!")
	})

	app.Get("/502", func(c *fiber.Ctx) error {
		return c.Render("502", fiber.Map{
			"Title": "502 Bad Gateway",
		})
	})

	app.Static("/style.css", "templates/style.css")

	log.Fatal(app.Listen(":8000"))
}
