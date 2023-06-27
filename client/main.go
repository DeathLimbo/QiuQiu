package main

import (
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/engi"
	"korok.io/korok/game"
	"korok.io/korok/gfx"
	"korok.io/korok/math/f32"
)

type StartScene struct {
	bg     engi.Entity
	ground engi.Entity
	bird   engi.Entity
}

func (sn *StartScene) OnEnter(g *game.Game) {
	//  get textue
	at, _ := asset.Texture.Atlas("images/bird.png")
	bird1, _ := at.GetByName("bird1.png")
	bird2, _ := at.GetByName("bird2.png")
	bird3, _ := at.GetByName("bird3.png")
	frames := []gfx.Tex2D{bird1, bird2, bird3}
	g.AnimationSystem.SpriteEngine.NewAnimation("flying", frames, true)

	// setup bird
	bird := korok.Entity.New()
	spr := korok.Sprite.NewCompX(bird, bird1)
	spr.SetSize(48, 32)
	spr.SetZOrder(100)
	xf := korok.Transform.NewComp(bird)
	xf.SetPosition(f32.Vec2{160, 240})
	sn.bird = bird

	anim := korok.Flipbook.NewComp(sn.bird)
	anim.SetRate(.1)
	anim.Play("flying")

	bg, _ := at.GetByName("background.png")
	//tt, _ := at.GetByName("game_name.png")
	ground, _ := at.GetByName("ground.png")

	// setup bg
	{
		entity := korok.Entity.New()
		spr := korok.Sprite.NewCompX(entity, bg)
		spr.SetSize(320, 480)
		xf := korok.Transform.NewComp(entity)
		xf.SetPosition(f32.Vec2{160, 240})
		sn.bg = entity
	}

	// setup ground {840 281}
	{
		entity := korok.Entity.New()
		spr := korok.Sprite.NewCompX(entity, ground)
		spr.SetSize(420, 140)
		spr.SetGravity(0, 1)
		spr.SetZOrder(1)
		xf := korok.Transform.NewComp(entity)
		xf.SetPosition(f32.Vec2{0, 100})
		sn.ground = entity
	}

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
