package proxies

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

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

		for _, p := range proxies.Proxies {
			if p.Endpoint != fmt.Sprintf("/%v", c.Params("endpoint")) {
				continue
			}
			if !contains(p.Methods, "GET") {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   fmt.Sprintf("%v is not configured", c.Method()),
				})
			}

			return c.Status(p.Http_result).JSON(p.Body)

		}
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"error":   "route does not exist",
		})
	}
}

func contains(methods []string, method string) bool {
	for _, m := range methods {
		if m == method {
			return true
		}
	}

	return false
}
