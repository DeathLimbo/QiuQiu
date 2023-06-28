package internal

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewGridManager(t *testing.T) {
	m := NewGridManager(10, 10)
	wait := sync.WaitGroup{}
	for i := uint32(1); i <= 2; i++ {
		player := &Player{
			BaseEntity{id: i},
		}
		go func() {
			for {
				m.Moveto(Pos{x: GetRandomWithAll(0, 10), y: GetRandomWithAll(0, 10)}, player)
				time.Sleep(1 * time.Second)
			}
			wait.Done()
		}()
		wait.Add(1)
	}
	wait.Wait()
}

func GetRandomWithAll(min, max int) int32 {
	rand.Seed(time.Now().Unix())
	if min > max {
		min, max = max, min
	}
	return int32(rand.Intn(max-min+1) + min)
}
