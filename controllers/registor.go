package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"encoding/json"
)

var dynamincs = map[string]string{}

type RegistorController struct {
	beego.Controller
}

type Body struct {
	Path string
	file string
}

func (c *RegistorController) Post() {
	fmt.Println("in post")
	body := Body{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	if err != nil {
		c.Data["json"] = "Error:" + err.Error()
		c.ServeJSON()
		return
	}

	if body.file == "" || body.Path == "" || body.Path[0] != '/' {
		c.Data["json"] = "File or path can not be null.  And path should be start with '/'."
		c.ServeJSON()
		return
	}

	dynamincs[body.Path] = body.file
}