package example

import (
	"io"
	"net/http"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/labstack/echo"
)

var (
	Mux         *http.ServeMux
	ghttpServer *ghttp.Server
	server      *echo.Echo
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/name", NameHandler)
	server = echo.New()
	server.GET("/name", NameHandlerEcho)

	ghttpServer = g.Server()
	ghttpServer.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write(g.Map{"name": "hexi"})
	})
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}

func NameHandlerEcho(c echo.Context) error {
	return c.String(http.StatusOK, `{"name": "hexi"}`)
}
