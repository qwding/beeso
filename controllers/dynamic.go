package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"plugin"
	"fmt"
	"beeso/models"
)

type DynamicController struct {
	beego.Controller
}



//var uris = map[string]beego.ControllerInterface{}
func (c *DynamicController) Router() {
	uri := c.Ctx.Input.URI()
	uri = strings.TrimPrefix(uri, "/dynamic")

	fmt.Println("method:", c.Ctx.Input.Method())
	file, ok := dynamincs[uri]
	if !ok {
		c.Abort("404")
	}

	p, err := plugin.Open(file)
	if err != nil {
		panic(err)
	}
	data, err := p.Lookup("Data")
	if err != nil {
		panic(err)
	}

	*data.(*models.Data) = models.Data{
		Ctx:*c.Ctx,
	}

	run, err := p.Lookup(c.Ctx.Input.Method())
	if err != nil {
		c.Abort("404")
	}

	run.(func())()
}