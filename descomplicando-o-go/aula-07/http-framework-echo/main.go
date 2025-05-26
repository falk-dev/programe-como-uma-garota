package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GithubWebhookRequestBody struct {
	Action     string
	Number     int64
	Repository string
	Sender     string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func echoHandler(c echo.Context) error {
	req := c.Request()
	body, err := io.ReadAll(req.Body)
	checkErr(err)

	obj := new(GithubWebhookRequestBody)
	err = json.Unmarshal(body, obj)
	checkErr(err)

	fmt.Println("Action", obj.Action)
	fmt.Println("Number", obj.Number)
	fmt.Println("Sender", obj.Sender)
	fmt.Println("Repository", obj.Repository)

	return c.String(http.StatusOK, "Ok")
}

func main() {
	server := echo.New()

	server.POST("/", echoHandler)

	server.Logger.Fatal(server.Start(":8080"))
}
