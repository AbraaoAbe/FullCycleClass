package entity

import "errors"

type ActionTableInterface interface {
	Init()
	GetScore(action string) (int, error)
}

type ActionTable struct {
	Actions map[string]int
}

func (a *ActionTable) Init() {
	a.Actions = make(map[string]int)
	a.Actions["goal"] = 5
	a.Actions["yellow_card"] = -1
	a.Actions["red_card"] = -3
}

func (a *ActionTable) GetScore(action string) (int, error) {
	score, ok := a.Actions[action]
	if !ok {
		return 0, errors.New("action not found")
	}
	return score, nil
}
