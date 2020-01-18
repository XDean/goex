package xapp

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

type (
	Runner struct {
		Config ConfigRegistry
		inits  []InitTask
		mains  []MainTask
	}
	InitTask struct {
		Ready func() bool
		Init  func()
	}
	MainTask func()
)

func (a *Runner) RegisterInitFunc(f func()) {
	a.inits = append(a.inits, InitTask{
		Ready: func() bool {
			return true
		},
		Init: f,
	})
}

func (a *Runner) RegisterInit(f InitTask) {
	a.inits = append(a.inits, f)
}

func (a *Runner) RegisterRun(f func()) {
	a.mains = append(a.mains, f)
}

func (a *Runner) RegisterConfigPath(path string) {
	a.RegisterInitFunc(func() {
		err := a.Config.Load(path)
		if err != nil {
			log.WithError(err).Fatal("Fail to load config file")
		}
	})
}

func (a *Runner) Run() {
	a.runInits()
	a.runMains()
}

func (a *Runner) runInits() {
	tasks := make(map[int]InitTask, len(a.inits))
	for k, v := range a.inits {
		tasks[k] = v
	}
	for len(tasks) > 0 {
		toRemove := make([]int, 0)
		for k, v := range tasks {
			if v.Ready() {
				v.Init()
				toRemove = append(toRemove, k)
			}
		}
		for _, v := range toRemove {
			delete(tasks, v)
		}
	}
}

func (a *Runner) runMains() {
	wg := sync.WaitGroup{}
	for _, v := range a.mains {
		wg.Add(1)
		go func() {
			v()
			wg.Done()
		}()
	}
	wg.Wait()
}
