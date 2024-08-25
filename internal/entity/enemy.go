package entity

import "github.com/google/uuid"

type Enemy struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
	Defense  int
	Heal     int // Novo atributo para cura
}

func NewEnemy(nickname string, life, attack, defense, heal int) *Enemy {
	return &Enemy{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		Attack:   attack,
		Defense:  defense,
		Heal:     heal, // Inicializando a cura
	}
}

// Função para aplicar cura ao inimigo
func (e *Enemy) ApplyHeal() {
	e.Life += e.Heal
}

