package internal

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewGridManager(t *testing.T) {
	rand.Seed(time.Now().Unix())
	m := NewGridManager(100, 100, 100)
	wait := sync.WaitGroup{}
	for i := uint32(1); i <= 2; i++ {
		player := &Player{
			BaseEntity{id: i},
		}
		go func() {
			for {
				pos := Pos{x: GetRandomWithAll(0, 100), y: GetRandomWithAll(0, 100)}
				//fmt.Println(fmt.Sprintf("%v:开始移动x:%v->%v y:%v->%v,", player.id, player.GetPos().x, pos.x, player.GetPos().y, pos.y))
				m.Moveto(pos, player)
				time.Sleep(1 * time.Second)
			}
			wait.Done()
		}()
		wait.Add(1)
	}
	wait.Wait()
}

func GetRandomWithAll(min, max int) int32 {
	if min > max {
		min, max = max, min
	}
	return int32(rand.Intn(max-min+1) + min)
}

func TestNewGridManager2(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(GetRandomWithAll(0, 10))
	}
}
