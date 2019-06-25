package main

import (
	"log"
	"net/http"

	"gopkg.in/macaron.v1"

	"github.com/masudur-rahman/go-chat/clients"
)

func main() {
	log.Println("Websocket is running")
	hub := clients.NewHub()
	go hub.Run()

	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", ServeHome)
	m.Get("/ws", func(ctx *macaron.Context) {
		clients.ServeWs(hub, ctx)
	})
	m.Run()
}

func ServeHome(ctx *macaron.Context) {
	log.Println(ctx.Req.URL)
	ctx.HTML(http.StatusOK, "home")
}
