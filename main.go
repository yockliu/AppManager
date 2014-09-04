package main

import (
	"./appmanager"
	"./mongodb"
	"./route"
	//	"encoding/base64"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	"github.com/martini-contrib/render"
	"time"
)

func main() {
	fmt.Println("Hello, let's go!")

	//	aaa := base64.StdEncoding.EncodeToString([]byte("admin:guessme"))
	//	fmt.Println(aaa)

	mongodb.Init()

	m := martini.Classic()

	m.Use(auth.BasicFunc(func(username, password string) bool {
		fmt.Println("username = " + username)
		fmt.Println("password = " + password)
		return auth.SecureCompare(username, "admin") && auth.SecureCompare(password, "guessme")
	}))

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
	}))

	route.Route(m)

	appmanager.Init()
	appmanager.RouteApi(m)

	m.Run()

	var ab = appmanager.NewAppBuilder()
	go ab.Run2()
	for {
		time.Sleep(time.Second)
	}
}
