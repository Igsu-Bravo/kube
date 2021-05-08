package main

import (
	_ "kube/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	/* This would match routes like the following:
	/sum/3/5
	/product/6/23
	...
	*/
	// beego.Router("/:operation/:num1:int/:num2:int", &mainController{})
	web.Run()
}

/* type healthController struct {
	beego.Controller
}
*/
/* type mainController struct {
	beego.Controller
} */

/*
func (c *healthController) Get() {
	c.Ctx.WriteString("Ok!")
} */

/* func (c *mainController) Get() {

	//Obtain the values of the route parameters defined in the route above
	operation := c.Ctx.Input.Param(":operation")
	num1, _ := strconv.Atoi(c.Ctx.Input.Param(":num1"))
	num2, _ := strconv.Atoi(c.Ctx.Input.Param(":num2"))

	// Perform the calculation depending on the 'operation' route parameter
	switch operation {
	case "sum":
		c.Ctx.WriteString(strconv.Itoa(add(num1, num2)))
	case "product":
		c.Ctx.WriteString(strconv.Itoa(multiply(num1, num2)))
	default:
		c.Ctx.WriteString("What?")
	}
} */
