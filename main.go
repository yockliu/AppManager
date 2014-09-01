package main

import (
	"./appbuild"
	"./mongodb"
	"./route"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"time"
)

func main() {
	fmt.Println("Hello, let's go!")

	mongodb.Init()

	m := martini.Classic()

	m.Use(auth.BasicFunc(func(username, password string) bool {
		return auth.SecureCompare(username, "admin") && auth.SecureCompare(password, "guessme")
	}))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
	}))

	route.Route(m)
	appbuild.Route(m)

	m.Run()

	var ab = appbuild.NewAppBuilder()
	go ab.Run2()
	for {
		time.Sleep(time.Second)
	}
}
