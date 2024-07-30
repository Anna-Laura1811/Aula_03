package service

import (
	"errors"
	"fmt"

	"github.com/Anna-Laura1811/RPG_GO_BD/internal/entity"
	repository "github.com/Anna-Laura1811/RPG_GO_BD/internal/service"
)

type BattleService struct {
	PlayerRepository repository.PlayerRepository
	EnemyRepository  repository.EnemyRepository
}

func NewBattleService(playerRepo repository.PlayerRepository, enemyRepo repository.EnemyRepository) *BattleService {
	return &BattleService{
		PlayerRepository: playerRepo,
		EnemyRepository:  enemyRepo,
	}
}

func (bs *BattleService) Battle(playerID, enemyID string, diceThrown int) (*entity.BattleResult, error) {
	player, err := bs.PlayerRepository.LoadPlayerById(playerID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if player == nil {
		return nil, errors.New("player id not found")
	}

	enemy, err := bs.EnemyRepository.LoadEnemyById(enemyID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}

	playerAttack := player.Attack + diceThrown
	enemyAttack := enemy.Attack + diceThrown

	player.Life -= enemyAttack
	enemy.Life -= playerAttack

	var winner string
	if player.Life > 0 && enemy.Life <= 0 {
		winner = "player"
	} else if enemy.Life > 0 && player.Life <= 0 {
		winner = "enemy"
	} else {
		winner = "draw"
	}

	result := &entity.BattleResult{
		PlayerID:    player.ID,
		EnemyID:     enemy.ID,
		Winner:      winner,
		PlayerLife:  player.Life,
		EnemyLife:   enemy.Life,
		PlayerAttack: playerAttack,
		EnemyAttack:  enemyAttack,
	}

	return result, nil
}
