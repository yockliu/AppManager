package route

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func Route(m *martini.ClassicMartini) {
	m.Get("/hello", func(r render.Render) {
		fmt.Println("main index")
		r.HTML(200, "hello", "Yock")
	})
}
