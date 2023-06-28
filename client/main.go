package main

import (
	"client/src/screen"
	"korok.io/korok"
)

func main() {
	opt := korok.Options{
		Title:  "球球大作战",
		Width:  320,
		Height: 480,
	}
	korok.Run(&opt, &screen.StartScene{})
}
