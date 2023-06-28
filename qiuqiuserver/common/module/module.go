package module

import (
	"sync"
)

type Module interface {
	OnInit()
	OnDestroy()
	Run(closeSig chan bool)
}

type module struct {
	mi       Module
	closeSig chan bool
	wg       sync.WaitGroup
}

var mods []*module

func register(mi Module) *module {
	m := new(module)
	m.mi = mi
	m.closeSig = make(chan bool, 1)

	mods = append(mods, m)

	return m
}

func Destroy() {
	for i := len(mods) - 1; i >= 0; i-- {
		m := mods[i]
		m.closeSig <- true
		m.wg.Wait()
		destroy(m)
	}
}

func destroy(m *module) {
	defer func() {
		if r := recover(); r != nil {
			//if conf.LenStackBuf > 0 {
			//	buf := make([]byte, conf.LenStackBuf)
			//	l := runtime.Stack(buf, false)
			//	log.Error("%v: %s", r, buf[:l])
			//} else {
			//	log.Error("%v", r)
			//}
		}
	}()

	m.mi.OnDestroy()
}

func Run(mod Module) {
	m := register(mod)
	mod.OnInit()

	go func() {
		m.wg.Add(1)
		m.mi.Run(m.closeSig)
		m.wg.Done()
	}()

}

func RunModule(mod Module) {
	if mod == nil {
		panic("mod is nil")
		return
	}

	Run(mod)
}
