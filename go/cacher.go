package main

import "sync"

var cacher *Cacher

type Cacher struct {
	ConditionMtx       sync.RWMutex
	LastConditions     map[string]IsuCondition
	IsuMtx             sync.RWMutex
	IsuListByCharacter map[string][]Isu
	Isus               map[string]Isu
}

func init() {
	CacheClear()
}

func CacheClear() {
	cacher = &Cacher{
		ConditionMtx:       sync.RWMutex{},
		LastConditions:     map[string]IsuCondition{},
		IsuMtx:             sync.RWMutex{},
		IsuListByCharacter: map[string][]Isu{},
		Isus:               map[string]Isu{},
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

	if _c, ok := ch.LastConditions[jiaIsuUUID]; ok {
		if _c.Timestamp.After(c.Timestamp) {
			return
		}
	}

	ch.LastConditions[jiaIsuUUID] = c
}

func (ch *Cacher) GetIsu(jiaIsuUUID string) (Isu, bool) {
	ch.IsuMtx.RLock()
	defer ch.IsuMtx.RUnlock()

	isu, ok := ch.Isus[jiaIsuUUID]
	return isu, ok
}

func (ch *Cacher) GetIsuListByCharacter() map[string][]Isu {
	ch.IsuMtx.RLock()
	defer ch.IsuMtx.RUnlock()

	return ch.IsuListByCharacter
}

func (ch *Cacher) AddIsu(isu Isu) {
	ch.IsuMtx.Lock()
	defer ch.IsuMtx.Unlock()

	ch.Isus[isu.JIAIsuUUID] = isu
	ch.IsuListByCharacter[isu.Character] = append(ch.IsuListByCharacter[isu.Character], isu)
}
