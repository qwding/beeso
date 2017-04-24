package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

var dynamincs = map[string]string{}

type RegistorController struct {
	beego.Controller
}

type Body struct {
	Path string
	File string
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
	// fmt.Println("file:", body.File, "path:", body.Path, "path0:'", body.Path[0], body.Path[0] == '/')

	if body.File == "" || body.Path == "" || body.Path[0] != '/' {
		c.Data["json"] = "File or path can not be null.  And path should be start with '/'."
		c.ServeJSON()
		return
	}

	dynamincs[body.Path] = body.File
	c.Data["json"] = "success"
	c.ServeJSON()
	return

}
