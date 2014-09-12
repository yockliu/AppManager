package appmanager

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func Init() {
	InitApp()
}

func RoutePage(m *martini.ClassicMartini) {
	m.Group("/appmanager", func(r martini.Router) {
		r.Get("", page_index)
		r.Get("/channel/list", page_channel_list)
		r.Get("/version/add", page_version_add)
		r.Get("/channel/add", page_channel_add)
		r.Post("/version/add", version_add)
		r.Post("/channel/add/submit", binding.Bind(ChannelAddForm{}), channel_add)
	})
}

func page_index(r render.Render) {
	r.HTML(200, "appmanager/index", "AppBuild")
}

func page_channel_list(r render.Render) {
	//	channels, err := ListChannels()
	//	if err != nil {
	//	} else {
	//		fmt.Println(channels)
	//		r.HTML(200, "appmanager/channel_list", channels)
	//	}
}

func page_version_add(r render.Render) {
	r.HTML(200, "appmanager/version_add", "")
}

func version_add(r render.Render) {

}

func page_channel_add(r render.Render) {
	r.HTML(200, "appmanager/channel_add", "")
}

func channel_add(caForm ChannelAddForm, r render.Render) {
	fmt.Println("channel add")
	//ch := Channel{caForm.Code, caForm.Name}
	//fmt.Println(ch)
	//SaveChannel(&ch)
	r.Redirect("/appmanager/channel/list")
	return
}

type ChannelAddForm struct {
	Name string `form:"channel_name" binding:"required"`
	Code string `form:"channel_code" binding:"required"`
}
