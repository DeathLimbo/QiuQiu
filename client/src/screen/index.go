package screen

import (
	"korok.io/korok"
	"korok.io/korok/asset"
	"korok.io/korok/engi"
	"korok.io/korok/game"
	"korok.io/korok/gfx"
	"korok.io/korok/gui"
	"korok.io/korok/hid/input"
	"korok.io/korok/math/f32"
)

type StateEnum int

const (
	Ready StateEnum = iota
	Running
	Over
)

// 定义重力加速度和冲量
const (
	Gravity    = 600
	TapImpulse = 280
)

// game.go
type GameScene struct {
	bg     engi.Entity
	ground engi.Entity
	ready  struct {
		gfx.Tex2D
		gui.Rect
	}
	tap struct {
		gfx.Tex2D
		gui.Rect
	}
	state StateEnum
	bird  struct {
		engi.Entity
		f32.Vec2
		vy   float32
		w, h float32
	}
}

func (sn *GameScene) OnEnter(g *game.Game) {
	// 重新调整鸟的位置
	korok.Transform.Comp(sn.bird.Entity).SetPosition(f32.Vec2{80, 240})

	at, _ := asset.Texture.Atlas("images/bird.png")

	// ready and tap image
	sn.ready.Tex2D, _ = at.GetByName("getready.png")
	sn.ready.Rect = gui.Rect{
		X: (320 - 233) / 2,
		Y: 70,
		W: 233,
		H: 70,
	}
	sn.tap.Tex2D, _ = at.GetByName("tap.png")
	sn.tap.Rect = gui.Rect{
		X: (320 - 143) / 2,
		Y: 200,
		W: 143, // 286
		H: 123, // 246
	}
}
func (sn *GameScene) Update(dt float32) {
	if st := sn.state; st == Ready {
		sn.showReady(dt)
		return
	} else if st == Over {
		sn.showOver(dt)
		return
	}

	// 检测屏幕点击，每次点击给鸟施加一次冲量
	if input.PointerButton(0).JustPressed() {
		sn.bird.vy = TapImpulse
	}
	// 模拟物理加速
	sn.bird.vy -= Gravity * dt
	sn.bird.Vec2[1] += sn.bird.vy * dt
	sn.bird.Vec2[0] = 80

	if sn.bird.Vec2[1] <= 120 {
		sn.bird.Vec2[1] = 120
	} else if sn.bird.Vec2[1] >= 450 {
		sn.bird.Vec2[1] = 450
	}
	// update bird position
	b := korok.Transform.Comp(sn.bird.Entity)
	b.SetPosition(sn.bird.Vec2)

}
func (sn *GameScene) OnExit() {
}

// 在当前帧添加方法和属性
func (sn *GameScene) borrow(bird, bg, ground engi.Entity) {
	sn.bird.Entity, sn.bg, sn.ground = bird, bg, ground
}

func (sn *GameScene) showReady(dt float32) {
	// show ready
	gui.Image(1, sn.ready.Rect, sn.ready.Tex2D, nil)

	// show tap hint
	gui.Image(2, sn.tap.Rect, sn.tap.Tex2D, nil)

	// check any click
	if input.PointerButton(0).JustPressed() {
		sn.state = Running
	}
}

func (sn *GameScene) showOver(dt float32) {

}
