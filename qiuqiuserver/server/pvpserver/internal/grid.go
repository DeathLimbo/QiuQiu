package internal

import (
	"fmt"
	"sync"
)

type Grid struct {
	pos      Pos      //格子ID
	Entities sync.Map //当前格子内的实体
}

type GridManger struct {
	grids map[Pos]*Grid
}

func NewGridManager(x, y int32) *GridManger {
	ret := &GridManger{
		grids: make(map[Pos]*Grid),
	}
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

func (g *GridManger) deleteEntity(e Entity) {
	grid, ok := g.grids[e.GetPos()]
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
	grid, ok := g.grids[e.GetPos()]
	if !ok {
		return
	}
	grid.Entities.Store(e.GetId(), e)
}

// 广播给周围格子
func (g *GridManger) broadRandGrid(e Entity) {
	for i := e.GetPos().x - 1; i < e.GetPos().x+1; i++ {
		for j := e.GetPos().y - 1; j < e.GetPos().y+1; j++ {
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
	ole := g.findEntity(e.GetPos(), e.GetId())
	if ole != nil { //代表以前就有数据
		if ole.GetPos().x == e.GetPos().x && ole.GetPos().y == e.GetPos().y { //代表格子没换
			fmt.Println("在这里？？", e.GetId(), ole.GetId(), ole.GetPos().x, e.GetPos().x, ole.GetPos().y, e.GetPos().y)
			return
		}
		//删除当前格子的 entity
		g.deleteEntity(e)
	}
	e.SetPos(to)

	//添加到新格子
	g.addEntity(e)

	//通知周围的格子
	g.broadRandGrid(e)
}
