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

		r.Post("/build", api_app_build)
		r.Get("/build/tasks", api_app_build_tasks)
	})
}

func api_app_build(params martini.Params, req *http.Request, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, err.Error())
		return
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Println(m)

	appid := m["appid"]
	platform := m["platform"]
	versionid := m["versionid"]
	channels := m["channels"]

	appidst, ok := appid.(string)
	if !ok {
		r.JSON(500, "appid格式错误")
		return
	}

	platformst, ok := platform.(string)
	if !ok {
		r.JSON(500, "platform格式错误")
		return
	}

	versionidst, ok := versionid.(string)
	if !ok {
		r.JSON(500, "versionid格式错误")
		return
	}

	cia, ok := channels.([]interface{})
	if !ok {
		r.JSON(500, "channels格式错误")
		return
	}
	channelsar := make([]string, len(cia))
	for i, v := range cia {
		channelsar[i], ok = v.(string)
		if !ok {
			r.JSON(500, "channels格式错误")
			return
		}
	}

	appBuilder, err := GetAppBuilder(appidst, platformst)
	if err != nil {
		r.JSON(500, err.Error())
		return
	}

	_, err = appBuilder.AddBuild(versionidst, channelsar)
	if err != nil {
		r.JSON(500, err.Error())
	}

	r.JSON(200, "{}")
}

func api_app_build_tasks(params martini.Params, req *http.Request, r render.Render) {
	query := req.URL.Query()
	appid := query["appid"]
	platform := query["platform"]
	// TODO: check query

	// m := bson.M{"status": bson.M{"$in": []T_ABTask_Status{T_ABTask_ST_RUNNING, T_ABTask_ST_INIT}}}
	if len(appid) > 0 {
		m["appid"] = appid[0]
	}
	if len(platform) > 0 {
		m["platform"] = platform[0]
	}
	tasks, err := ReadAppBuildTaskList(m)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, tasks)
	}
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
			r.JSON(200, apps)
		}
	}
}

func api_app_get(params martini.Params, r render.Render) {
	id := params["id"]
	app, err := ReadApp(bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, app)
	}
}

func api_app_post(app App, r render.Render) {
	fmt.Println(app)
	newApp, err := CreateApp(&app)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, newApp)
	}
}

func api_app_put(params martini.Params, req *http.Request, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, err.Error())
		return
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
		r.JSON(204, "")
	}
}

func api_version_list(params martini.Params, r render.Render) {
	appid := params["appid"]
	platform := params["platform"]
	versions, err := ListVersion(appid, platform)
	if err != nil {
		r.JSON(500, err.Error())
	} else if len(versions) == 0 {
		r.JSON(200, "[]")
	} else {
		r.JSON(200, versions)
	}
}

func api_version_post(version Version, params martini.Params, r render.Render) {
	appid := params["appid"]
	newVersion, err := CreateVersion(appid, version)
	if err != nil {
		r.JSON(500, err.Error())
		return
	}
	r.JSON(201, newVersion)
}

func api_version_get(params martini.Params, r render.Render) {
	appid := params["appid"]
	id := params["id"]
	version, err := ReadVersion(appid, bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
		return
	}
	r.JSON(200, version)
}

func api_version_put(params martini.Params, req *http.Request, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, err.Error())
		return
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Println(m)
	appid := params["appid"]
	id := params["id"]
	newVersion, err := UpdateVersion(appid, bson.ObjectIdHex(id), m)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, newVersion)
	}
}

func api_version_delete(params martini.Params, r render.Render) {
	appid := params["appid"]
	id := params["id"]
	err := DeleteVersion(appid, bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(204, "")
	}
}

func api_channel_list(params martini.Params, r render.Render) {
	appid := params["appid"]
	platform := params["platform"]
	channels, err := ListChannels(appid, platform)
	if err != nil {
		r.JSON(500, err.Error())
	} else if len(channels) == 0 {
		r.JSON(200, "[]")
	} else {
		r.JSON(200, channels)
	}
}

func api_channel_post(channel Channel, params martini.Params, r render.Render) {
	appid := params["appid"]
	newChannel, err := CreateChannel(appid, channel)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, newChannel)
	}
}

func api_channel_get(params martini.Params, r render.Render) {
	appid := params["appid"]
	id := params["id"]
	channel, err := ReadChannel(appid, bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(200, channel)
	}
}

func api_channel_put(params martini.Params, req *http.Request, r render.Render) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, err.Error())
		return
	}
	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Println(m)
	appid := params["appid"]
	id := params["id"]
	newChannel, err := UpdateChannel(appid, bson.ObjectIdHex(id), m)
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(201, newChannel)
	}
}

func api_channel_delete(params martini.Params, r render.Render) {
	appid := params["appid"]
	id := params["id"]
	err := DeleteChannel(appid, bson.ObjectIdHex(id))
	if err != nil {
		r.JSON(500, err.Error())
	} else {
		r.JSON(204, "{}")
	}
}
