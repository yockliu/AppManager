package appmanager

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"strconv"
)

func RouteApi(m *martini.ClassicMartini) {
	m.Group("/api/appmanager", func(r martini.Router) {
		r.Get("/app/list", api_app_list)
		r.Get("/app/:id", api_app_get)
		r.Post("/app/add", binding.Json(App{}), api_app_post)
		//r.Post("/app/add", api_app_post)
		r.Put("/app/:id", api_app_put)
		r.Delete("/app/:id", api_app_delete)

		//		r.Get("/app/:appid/:platform/channel/list", api_channel_list)
		//		r.Get("/app/:appid/:platform/channel/:id", api_channel_get)
		//		r.Post("/app/:appid/:platform/channel/add", api_channel_post)
		//		r.Put("/app/:appid/:platform/channel/:id", api_channel_put)
		//		r.Delete("/app/:appid/:platform/channel/:id", api_channel_delete)
		//
		//		r.Get("/app/:appid/:platform/version/list", api_version_list)
		//		r.Get("/app/:appid/:platform/version/:id", api_version_get)
		//		r.Post("/app/:appid/:platform/version/add", api_version_post)
		//		r.Put("/app/:appid/:platform/version/:id", api_version_put)
		//		r.Delete("/app/:appid/:platform/version/:id", api_version_delete)
	})
}

func api_app_list(r render.Render) {
	fmt.Println("api_app_list")
	var apps []App
	var err error
	apps, err = ListApp()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	jsonbyte, err := json.Marshal(apps)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("----")
	r.JSON(200, string(jsonbyte))
}

func api_app_get(params martini.Params, r render.Render) {
	idstr := params["id"]
	id, err := strconv.Atoi(idstr)
	if err != nil {
		panic(err)
	}
	app, err := ReadApp(id)
	if err != nil {
		panic(err)
	}
	jsonbyte, err := json.Marshal(app)
	if err != nil {
		panic(err)
	}
	r.JSON(200, string(jsonbyte))
}

func api_app_post(app App, r render.Render) {
	fmt.Println(app)
	err := CreateApp(&app)
	if err != nil {
		fmt.Println(err)
	}
}

func api_app_put(params martini.Params, r render.Render) {

}

func api_app_delete(params martini.Params, r render.Render) {

}
