package example

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerFunc(t *testing.T) {
	client := htest.NewClient().ToFunc(NameHandler)
	body := client.Get("").Send().OK().JSON()
	body.String("name", "hexi")
}

func TestNameHandler(t *testing.T) {
	client := htest.NewClient().To(Mux)
	body := client.Get("/name").Send().OK().JSON()
	body.String("name", "hexi")
}

func TestNameHandlerEcho(t *testing.T) {
	client := htest.NewClient().To(server)
	client.Get("/name").Send().OK().JSON().String("name", "hexi")
	client.Get("/stuid").Send().OK().JSON().String("stuid", "3160100001")
}