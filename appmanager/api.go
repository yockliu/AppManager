package appmanager

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
)

func RouteApi(m *martini.ClassicMartini) {
	m.Group("/api", func(r martini.Router) {
		r.Get("/app", api_app_list)
		r.Post("/app", binding.Json(App{}), api_app_post)
		r.Get("/app/:id", api_app_get)
		r.Put("/app/:id", binding.Json(App{}), api_app_put)
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
	if apps == nil || len(apps) == 0 {
		r.JSON(200, "[]")
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
	id := params["id"]
	app, err := ReadApp(bson.ObjectIdHex(id))
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
		panic(err)
	}
	r.JSON(201, nil)
}

func api_app_put(app App, params martini.Params, r render.Render) {
	fmt.Println(app)
	id := params["id"]
	err := UpdateApp(bson.ObjectIdHex(id), &app)
	if err != nil {
		panic(err)
	}
	r.JSON(201, nil)
}

func api_app_delete(params martini.Params, r render.Render) {
	id := params["id"]
	err := DeleteApp(bson.ObjectIdHex(id))
	if err != nil {
		panic(err)
	}
	r.JSON(204, nil)
}
