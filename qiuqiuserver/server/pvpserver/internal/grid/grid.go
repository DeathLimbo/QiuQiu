package grid

import (
	"sync"
)

type Grid struct {
	pos      Pos      //格子ID
	Entities sync.Map //当前格子内的实体
}

type GridManger struct {
	grids         map[Pos]*Grid
	width, height int32
	rate          int32
}

func NewGridManager(width, height, rate int32) *GridManger {
	ret := &GridManger{
		grids:  make(map[Pos]*Grid),
		width:  width,
		height: height,
		rate:   rate,
	}
	x := RateGrid(width, rate)
	y := RateGrid(height, rate)
	//生成 格子
	for i := int32(0); i < x; i++ {
		for j := int32(0); j < y; j++ {
			grid := &Grid{
				pos: Pos{
					x: i,
					y: j,
				},
			}
			ret.grids[grid.pos] = grid
		}
	}

	return ret
}

func (g *GridManger) findEntity(pos Pos, id uint32) Entity {
	grid, ok := g.grids[pos]
	if !ok {
		return nil
	}
	var ret Entity
	grid.Entities.Range(func(key, value any) bool {
		ole, ok := value.(Entity)
		if !ok {
			return true
		}
		if ole.GetId() == id {
			ret = ole
			return false
		}
		return true
	})
	return ret
}

func (g *GridManger) deleteEntity(e Entity) { //通知周围格子
	gpos := g.GetGridPos(e.GetPos())
	grid, ok := g.grids[gpos]
	if !ok {
		return
	}
	grid.Entities.Delete(e.GetId())
}

func (g *Grid) broad(e Entity) {
	g.Entities.Range(func(key, value any) bool {
		target, ok := value.(Entity)
		if !ok {
			return true
		}
		target.Process(e)
		return true
	})
}

func (g *GridManger) addEntity(e Entity) {
	gpos := g.GetGridPos(e.GetPos())
	grid, ok := g.grids[gpos]
	if !ok {
		return
	}
	grid.Entities.Store(e.GetId(), e)
}

// 广播给周围格子
func (g *GridManger) broadRandGrid(e Entity) {
	pos := g.GetGridPos(e.GetPos())
	for i := pos.x - 1; i < pos.x+1; i++ {
		for j := pos.y - 1; j < pos.y+1; j++ {
			pos := Pos{
				x: i,
				y: j,
			}
			grid, ok := g.grids[pos]
			if !ok {
				continue
			}
			grid.broad(e)
		}
	}
}

func (g *GridManger) Moveto(to Pos, e Entity) {
	gPos := g.GetGridPos(e.GetPos())
	gTo := g.GetGridPos(to)

	ole := g.findEntity(gPos, e.GetId())
	defer func() {
		//通知周围的格子
		g.broadRandGrid(e)
	}()
	if ole != nil { //代表以前就有数据
		if gTo.x == gPos.x && gTo.y == gPos.y { //代表格子没换
			return
		}
		//删除当前格子的 entity
		g.deleteEntity(e)
	}

	e.SetPos(to)

	//添加到新格子
	g.addEntity(e)

}

func RateGrid(length int32, rate int32) int32 {
	x := length / rate
	return x
}

func (g *GridManger) GetGridPos(client Pos) Pos {
	x := RateGrid(client.x, g.rate)
	y := RateGrid(client.y, g.rate)
	return Pos{
		x: x,
		y: y,
	}
}
