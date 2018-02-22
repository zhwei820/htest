package htest

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"testing"
	"time"
)

const (
	WrongXMLData = `
<?xml version="1.0" encoding="UTF-8"?>
<user>
	<id>1</id>
	<name>hexi</name>
`

	JSONAssertData = `
{
	"number": 1,
	"time": "2018-2-22 00:00:00",
	"ok": true,
	"no": false,
}
`
	JSONAssertDataTimeStr = "2018-02-22 00:00:00"
	StandardTime          = "2006-01-02 15:04:05"
)

var (
	JSONAssertDataTime, _ = time.Parse(StandardTime, JSONAssertDataTimeStr)
)

func TestJSON_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/name").Send().StatusOK().JSON().Exist("name").NotExist("stuid")
}

func TestJSON_String(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).String("time", JSONAssertDataTimeStr)
}

func TestJSON_Int(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Int("number", int64(1))
}

func TestJSON_True(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).True("ok")
}

func TestJSON_False(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).False("no")
}

func TestJSON_Uint(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Uint("number", uint64(1))
}

func TestJSON_Time(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Time("time", JSONAssertDataTime)
	t.Logf("Time: %s\n", gjson.Get(`{"time": "2018-02-22 00:00:00"}`, "time").Time())
}

func TestJSON_Float(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Float("number", float64(1))
}

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Send().StatusOK().JSON().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestJSON_NotEmpty(t *testing.T) {
	NewXML([]byte(UserDataXML), t).JSON.NotEmpty()
}

func TestXML_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Send().StatusOK().XML().Exist("user.name").NotExist("user.stuid")
}

func TestXML_String(t *testing.T) {
	NewXML([]byte(UserDataXML), t).String("user.name", "hexi")
}

func TestXML_Empty(t *testing.T) {
	NewXML([]byte(""), t).Empty()
}

func TestXML_NotEmpty(t *testing.T) {
	NewXML([]byte(WrongXMLData), t).NotEmpty()
}

func TestWrongXML_JSON_Empty(t *testing.T) {
	NewXML([]byte(WrongXMLData), t).JSON.Empty()
}

func TestXML_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Send().StatusOK().XML().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func UserDataHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserData)
}

func UserDataXMLHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserDataXML)
}
