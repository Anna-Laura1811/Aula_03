package entity

import "github.com/google/uuid"

type Player struct {
	ID       string
	Nickname string
	Life     int
	Attack   int
	Defense  int
	Heal     int // Atributo para cura
}

func NewPlayer(nickname string, life, attack, defense, heal int) *Player {
	return &Player{
		ID:       uuid.New().String(),
		Nickname: nickname,
		Life:     life,
		Attack:   attack,
		Defense:  defense,
		Heal:     heal, // Inicializando a cura
	}
}

// Função para aplicar cura ao jogador
func (p *Player) ApplyHeal() {
	p.Life += p.Heal
}

