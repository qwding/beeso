package controllers

import (
	"beeso/models"
	"fmt"
	"github.com/astaxie/beego"
	"plugin"
	"strings"
)

type DynamicController struct {
	beego.Controller
}

//var uris = map[string]beego.ControllerInterface{}
func (c *DynamicController) Router() {
	uri := c.Ctx.Input.URI()
	uri = strings.TrimPrefix(uri, "/dynamic")

	file, ok := dynamincs[uri]
	if !ok {
		c.Abort("404")
	}

	p, err := plugin.Open(file)
	if err != nil {
		fmt.Println("error open:",err)
		c.Abort("404")
	}
	data, err := p.Lookup("Data")
	if err != nil {
		fmt.Println("error lookup:",err)
		panic(err)
	}

	*data.(*models.Data) = models.Data{
		Ctx: *c.Ctx,
	}

	run, err := p.Lookup(c.Ctx.Input.Method())
	if err != nil {
		fmt.Println("error lookup:",err)
		c.Abort("404")
	}

	resp, err := run.(func() (string, error))()
	if err != nil {
		c.Data["json"] = "Error:" + err.Error()
		c.ServeJSON()
		return
	}

	c.Data["json"] = resp
	c.ServeJSON()
	return
}
