package main

import "sync"

var cacher *Cacher

type Cacher struct {
	ConditionMtx   sync.RWMutex
	LastConditions map[string]IsuCondition
}

func init() {
	CacheClear()
}

func CacheClear() {
	cacher = &Cacher{
		ConditionMtx:   sync.RWMutex{},
		LastConditions: map[string]IsuCondition{},
	}
}

func (ch *Cacher) GetLastCondition(jiaIsuUUID string) (IsuCondition, bool) {
	ch.ConditionMtx.RLock()
	defer ch.ConditionMtx.RUnlock()

	c, ok := ch.LastConditions[jiaIsuUUID]
	return c, ok
}

func (ch *Cacher) SetLastCondition(jiaIsuUUID string, c IsuCondition) {
	ch.ConditionMtx.Lock()
	defer ch.ConditionMtx.Unlock()

	ch.LastConditions[jiaIsuUUID] = c
}
