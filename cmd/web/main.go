package main

import "github.com/MyongJi-StudentVote/StudentVote-Admin/internal/server"

func main() {
	app := server.NewApplication()
	app.Run()
}
