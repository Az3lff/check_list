package http

import (
	"github.com/gofiber/fiber/v2"

	"github.com/Az3lff/check_list_proxy/internal/config"
	check_list2 "github.com/Az3lff/check_list_proxy/internal/delivery/http/check_list"
	"github.com/Az3lff/check_list_proxy/internal/service/check_list"
)

func SetupRoutes(app *fiber.App, cfg config.HTTPServer, svs *check_list.TaskService) {
	h := check_list2.NewTaskHandler(cfg, svs)

	app.Get("api/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("ok")
	})

	app.Post("/api/create_task", h.CreateTask)
	app.Get("/api/get_task", h.GetTask)
	app.Get("/api/list", h.GetList)
	app.Delete("/api/delete_task", h.DeleteTask)
	app.Patch("/api/done_task", h.DoneTask)
}
