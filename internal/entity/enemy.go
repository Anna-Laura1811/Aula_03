package entity

import "github.com/google/uuid"

type Enemy struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
}

func NewEnemy(nickname string, life, attack int) *Enemy {
	return &Enemy{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		Attack:   attack,
	}
}