package main

import (
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/game"
	"korok.io/korok/math/f32"
)

type StartScene struct {
}

func (sn *StartScene) OnEnter(g *game.Game) {
	//  get textue
	at, _ := asset.Texture.Atlas("images/bird.png")
	bird1, _ := at.GetByName("bird1.png")

	// setup bird
	bird := korok.Entity.New()
	spr := korok.Sprite.NewCompX(bird, bird1)
	spr.SetSize(48, 32)
	xf := korok.Transform.NewComp(bird)
	xf.SetPosition(f32.Vec2{160, 240})
}
func (sn *StartScene) Update(dt float32) {
	// draw something
}
func (sn *StartScene) OnExit() {
}

func (sn *StartScene) Load() {
	asset.Texture.LoadAtlas("images/bird.png", "images/bird.json")
}

func main() {
	opt := korok.Options{
		Title:  "球球大作战",
		Width:  320,
		Height: 480,
	}
	korok.Run(&opt, &StartScene{})
}
