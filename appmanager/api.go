package appbuild

import (
	"fmt"
	"github.com/go-martini/martini"
	//	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func RouteApi(m *martini.ClassicMartini) {
	m.Group("/api/appbuild", func(r martini.Router) {
		r.Get("/app/list", api_app_list)
		r.Get("/app/:id", api_app_get)
		r.Post("/app/add", api_app_post)
		r.Put("/app/:id", api_app_put)
		r.Delete("/app/:id", api_app_delete)

		r.Get("/app/:appid/:platform/channel/list", api_channel_list)
		r.Get("/app/:appid/:platform/channel/:id", api_channel_get)
		r.Post("/app/:appid/:platform/channel/add", api_channel_post)
		r.Put("/app/:appid/:platform/channel/:id", api_channel_put)
		r.Delete("/app/:appid/:platform/channel/:id", api_channel_delete)

		r.Get("/app/:appid/:platform/version/list", api_version_list)
		r.Get("/app/:appid/:platform/version/:id", api_version_get)
		r.Post("/app/:appid/:platform/version/add", api_version_post)
		r.Put("/app/:appid/:platform/version/:id", api_version_put)
		r.Delete("/app/:appid/:platform/version/:id", api_version_delete)
	})
}
