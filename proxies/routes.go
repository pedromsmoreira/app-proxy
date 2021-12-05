package proxies

import "github.com/gofiber/fiber/v2"

// AddRoutesTo method configures mentions routes for the api
func AddRoutesTo(app fiber.Router, proxies *Proxies) {
	app.Get("/:endpoint", get(proxies))
	// app.Post("/:endpoint", post(service))
	// app.Put("/:endpoint", put(service))
	// app.Delete("/:endpoint", delete(service))
}

func get(proxies *Proxies) fiber.Handler {
	return func(c *fiber.Ctx) error {
		/*
			map
				key: GET /endpoint -> content
		*/

		return c.JSON(&fiber.Map{
			"success": true,
			"error":   nil,
		})
	}
}
