package appmanager

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"os"
	"os/exec"
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
	r.HTML(200, "appbuild/index", "AppBuild")
}

func page_channel_list(r render.Render) {
	//	channels, err := ListChannels()
	//	if err != nil {
	//	} else {
	//		fmt.Println(channels)
	//		r.HTML(200, "appbuild/channel_list", channels)
	//	}
}

func page_version_add(r render.Render) {
	r.HTML(200, "appbuild/version_add", "")
}

func version_add(r render.Render) {

}

func page_channel_add(r render.Render) {
	r.HTML(200, "appbuild/channel_add", "")
}

func channel_add(caForm ChannelAddForm, r render.Render) {
	fmt.Println("channel add")
	//ch := Channel{caForm.Code, caForm.Name}
	//fmt.Println(ch)
	//SaveChannel(&ch)
	r.Redirect("/appbuild/channel/list")
	return
}

type ChannelAddForm struct {
	Name string `form:"channel_name" binding:"required"`
	Code string `form:"channel_code" binding:"required"`
}

type AppBuilder struct {
	cmd *exec.Cmd
}

func (ab *AppBuilder) Run1() {
	stdout, err := ab.cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	if err := ab.cmd.Start(); err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	d, _ := ioutil.ReadAll(stdout)
	if err := ab.cmd.Wait(); err != nil {
		fmt.Println("Error: %s\n", err)
		return
	}
	fmt.Println(string(d))
}

func (ab *AppBuilder) Run2() {
	ab.cmd.Stdout = os.Stdout
	ab.cmd.Stderr = os.Stderr
	ab.cmd.Run()
}

func NewAppBuilder() AppBuilder {
	fmt.Println("RunBuild")
	cmd := exec.Command("/bin/sh", "-c", "cd /Users/yinxiaoliu/android/zhoumo_android/\n./gradlew tasks")
	ab := AppBuilder{cmd}
	return ab
}
