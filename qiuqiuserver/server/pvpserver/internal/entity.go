package internal

import "fmt"

type Pos struct {
	x, y int32
}

type EntityTyp int

const (
	ETYP_PLAYER EntityTyp = 1 << iota
	ETYP_CYCLE
)

type Entity interface {
	GetName() string
	GetId() uint32
	GetPos() Pos
	SetPos(pos Pos)
	GetDistance(p *Pos) int32
	GetTyp() EntityTyp
	OnLeave()
	OnJoin()
	Process(e Entity)
}

type BaseEntity struct {
	id   uint32
	name string
	pos  Pos
	typ  EntityTyp
}

func (p *BaseEntity) GetName() string {
	return p.name
}

func (p *BaseEntity) GetId() uint32 {
	return p.id
}

func (p *BaseEntity) GetTyp() EntityTyp {
	return p.typ
}

func (p *BaseEntity) GetPos() Pos {
	return p.pos
}

func (p *BaseEntity) SetPos(pos Pos) {
	p.pos = pos
}

func (p *BaseEntity) OnLeave() {
}

func (p *BaseEntity) OnJoin() {
}

func (p *BaseEntity) Process(entity Entity) {
	fmt.Println(fmt.Sprintf("玩家:%v接收到通知:%v pos %v", p.GetId(), entity.GetId(), entity.GetPos()))
}

func (p *BaseEntity) GetDistance(pos *Pos) int32 {
	if pos == nil {
		return 0
	}
	distance := (pos.x - p.pos.x) ^ 2 + (pos.y - p.pos.y) ^ 2
	return distance
}

type Player struct {
	BaseEntity
}

type Cycle struct {
	BaseEntity
}
