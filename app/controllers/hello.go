package controllers
import (
	"github.com/revel/revel"
	"errors"
	"time"
    "net/http"
)

type Hello struct {
	*revel.Controller
}

func (c Hello) Index() revel.Result {
	greeting := "Hello World!"
	message := "これはStep2の実装です。"
	return c.Render(greeting, message)
}

func (c Hello) Greet(greeting string) revel.Result {
	message := "これはStep2の実装です。(その2)"
	c.RenderArgs["greeting"] = greeting
	c.RenderArgs["message"] = message
	return c.RenderTemplate("Hello/Index.html")
}

func (c Hello) HelloServerError() revel.Result {
	err := errors.New("Hello ServerError!")
	return c.RenderError(err)
}

func (c Hello) HelloNotFound() revel.Result {
	time := time.Now()
	msg := "Error Message"
	return c.NotFound("Hello NotFound!, %s, [%s]", msg, time.String())
}

func (c Hello) HelloForbidden() revel.Result {
	return c.Forbidden("Hello Forbidden!")
}

func (c Hello) HelloCustomError() revel.Result {
	c.Response.Status = http.StatusBadRequest
	c.RenderArgs["custom"] = "カスタムエラー"
	return c.RenderTemplate("errors/custom.html")
}
