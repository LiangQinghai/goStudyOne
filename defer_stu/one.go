package defer_stu

import "sync"

var m sync.Mutex

func NormalCall()  {
	m.Lock()
	m.Unlock()
}

func DeferCall()  {
	defer m.Unlock()
	m.Lock()
}
