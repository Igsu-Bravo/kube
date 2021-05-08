package health2

import (
	"github.com/astaxie/beego"
)

type HealthController struct {
	beego.Controller
}

func (c *HealthController) Get() {
	c.Ctx.WriteString("Ok!")
}

func Greet() string {
	return "Hello!"
}
