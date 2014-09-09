package appmanager

import (
	"encoding/json"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"net/http"
)

func RouteApi(m *martini.ClassicMartini) {
	m.Group("/api", func(r martini.Router) {
		r.Get("/app", api_app_list)
		r.Post("/app", binding.Json(App{}), api_app_post)
		r.Get("/app/:id", api_app_get)
		r.Put("/app/:id", api_app_put)
		r.Delete("/app/:id", api_app_delete)

		r.Get("/app/:appid/version", api_version_list)
		r.Post("/app/:appid/version", binding.Json(Version{}), api_version_post)
		r.Get("/app/:appid/version/:id", api_version_get)
		r.Put("/app/:appid/version/:id", api_version_put)
		r.Delete("/app/:appid/version/:id", api_version_delete)

		r.Get("/app/:appid/channel", api_channel_list)
		r.Post("/app/:appid/channel", binding.Json(Channel{}), api_channel_post)
		r.Get("/app/:appid/channel/:id", api_channel_get)
		r.Put("/app/:appid/channel/:id", api_channel_put)
		r.Delete("/app/:appid/channel/:id", api_channel_delete)
	})
}

func api_app_list(r render.Render) {
	fmt.Println("api_app_list")
	var apps []App
	var err error
	apps, err = ListApp()
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		if apps == nil || len(apps) == 0 {
			r.JSON(200, "[]")
		} else {
			jsonbyte, err := json.Marshal(apps)
			if err != nil {
				r.JSON(500, err.Error())
			} else {
				r.JSON(200, string(jsonbyte))
			}
		}
	}
}

func api_app_get(params martini.Params, r render.Render) {
	id := params["id"]
	app, err := ReadApp(bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	}
	jsonbyte, err := json.Marshal(app)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, string(jsonbyte))
	}
}

func api_app_post(app App, r render.Render) {
	fmt.Println(app)
	err := CreateApp(&app)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, "{}")
	}
}

func api_app_put(params martini.Params, req *http.Request, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Println(m)
	id := params["id"]
	putedApp, err := UpdateApp(bson.ObjectIdHex(id), m)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, putedApp)
	}
}

func api_app_delete(params martini.Params, r render.Render) {
	id := params["id"]
	err := DeleteApp(bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(204, "{}")
	}
}

func api_version_list(params martini.Params, r render.Render) {
	appid := params["appid"]
	versions, err := ListVersion(appid, "")
	if err != nil {
		r.JSON(500, err)
	} else {
		jsonbyte, err := json.Marshal(versions)
		if err != nil {
			r.JSON(500, err)
		} else if jsonbyte == nil {
			r.JSON(500, "[]")
		} else {
			r.JSON(200, string(jsonbyte))
		}
	}
}

func api_version_post(version Version, r render.Render) {

}
func api_version_get(params martini.Params, r render.Render) {
}
func api_version_put(params martini.Params, req *http.Request, r render.Render) {
}
func api_version_delete(params martini.Params, r render.Render) {
}

func api_channel_list(params martini.Params, r render.Render)                   {}
func api_channel_post(version Version, r render.Render)                         {}
func api_channel_get(params martini.Params, r render.Render)                    {}
func api_channel_put(params martini.Params, req *http.Request, r render.Render) {}
func api_channel_delete(params martini.Params, r render.Render)                 {}
