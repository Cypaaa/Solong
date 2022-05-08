package main

import (
	"GOLONG/module/Console"
	"GOLONG/module/Server"
	"fmt"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

func main() {
	Console.Clear()
	fmt.Println("You've launched GOLONG!")

	client := youtube.Client{}
	video, err := client.GetVideo("dQw4w9WgXcQ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(video.Title)

	http.HandleFunc("/", Server.Index)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":80", nil)
}
