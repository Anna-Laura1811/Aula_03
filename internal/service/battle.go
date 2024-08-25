package service

import (
	"errors"
	"math/rand"
	"strconv"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/repository"
)

type BattleService struct {
	PlayerRepository repository.PlayerRepository
	EnemyRepository  repository.EnemyRepository
	BattleRepository repository.BattleRepository
}

func NewBattleService(playerRepo repository.PlayerRepository, enemyRepo repository.EnemyRepository, battleRepo repository.BattleRepository) *BattleService {
	return &BattleService{
		PlayerRepository: playerRepo,
		EnemyRepository:  enemyRepo,
		BattleRepository: battleRepo,
	}
}

func (bs *BattleService) CreateBattle(playerNickname, enemyNickname string) (*entity.Battle, string, error) {
	player, err := bs.PlayerRepository.LoadPlayerByNickname(playerNickname)
	if err != nil || player == nil {
		return nil, "", errors.New("jogador não encontrado")
	}

	enemy, err := bs.EnemyRepository.LoadEnemyByNickname(enemyNickname)
	if err != nil || enemy == nil {
		return nil, "", errors.New("inimigo não encontrado")
	}

	if player.Life <= 0 || enemy.Life <= 0 {
		return nil, "", errors.New("tanto o jogador quanto o inimigo devem ter vida > 0 para batalhar")
	}

	// Determinar a ação do jogador (ataque ou cura)
	action := ""
	if player.Life <= 2 {
		action = "heal"
	} else {
		action = "attack"
	}

	// Chamar NewBattle com todos os argumentos necessários
	battle := entity.NewBattle(player.ID, enemy.ID, player.Nickname, enemy.Nickname, action, player.Life, enemy.Life)

	dice := battle.DiceThrown
	var result string

	if dice <= 2 { // Inimigo ataca
		damage := enemy.Attack - player.Defense
		if damage < 0 {
			damage = 0
		}
		player.Life -= damage
		if player.Life < 0 {
			player.Life = 0
		}
		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do jogador")
		}

		result = "Inimigo atacou. Dano causado: " + strconv.Itoa(damage) +
			" | Vida do Jogador: " + strconv.Itoa(player.Life) +
			" | Vida do Inimigo: " + strconv.Itoa(enemy.Life) +
			" | Ataque do Jogador: " + strconv.Itoa(player.Attack) +
			" | Defesa do Jogador: " + strconv.Itoa(player.Defense) +
			" | Ataque do Inimigo: " + strconv.Itoa(enemy.Attack) +
			" | Defesa do Inimigo: " + strconv.Itoa(enemy.Defense)
	} else if dice <= 4 { // Jogador ataca
		damage := player.Attack - enemy.Defense
		if damage < 0 {
			damage = 0
		}
		enemy.Life -= damage
		if enemy.Life < 0 {
			enemy.Life = 0
		}
		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do inimigo")
		}

		result = "Jogador atacou. Dano causado: " + strconv.Itoa(damage) +
			" | Vida do Jogador: " + strconv.Itoa(player.Life) +
			" | Vida do Inimigo: " + strconv.Itoa(enemy.Life) +
			" | Ataque do Jogador: " + strconv.Itoa(player.Attack) +
			" | Defesa do Jogador: " + strconv.Itoa(player.Defense) +
			" | Ataque do Inimigo: " + strconv.Itoa(enemy.Attack) +
			" | Defesa do Inimigo: " + strconv.Itoa(enemy.Defense)
	} else if dice == 5 { // Jogador se cura
		heal := rand.Intn(10) + 1 // Valor de cura aleatório entre 1 e 10
		player.Life += heal
		if err := bs.PlayerRepository.SavePlayer(player.ID, player); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do jogador")
		}

		result = "Jogador se curou. Vida recuperada: " + strconv.Itoa(heal) +
			" | Vida atual do Jogador: " + strconv.Itoa(player.Life) +
			" | Vida do Inimigo: " + strconv.Itoa(enemy.Life) +
			" | Ataque do Jogador: " + strconv.Itoa(player.Attack) +
			" | Defesa do Jogador: " + strconv.Itoa(player.Defense) +
			" | Ataque do Inimigo: " + strconv.Itoa(enemy.Attack) +
			" | Defesa do Inimigo: " + strconv.Itoa(enemy.Defense)
	} else { // Inimigo se cura
		heal := rand.Intn(10) + 1 // Valor de cura aleatório entre 1 e 10
		enemy.Life += heal
		if err := bs.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
			return nil, "", errors.New("falha ao atualizar a vida do inimigo")
		}

		result = "Inimigo se curou. Vida recuperada: " + strconv.Itoa(heal) +
			" | Vida atual do Inimigo: " + strconv.Itoa(enemy.Life) +
			" | Vida do Jogador: " + strconv.Itoa(player.Life) +
			" | Ataque do Jogador: " + strconv.Itoa(player.Attack) +
			" | Defesa do Jogador: " + strconv.Itoa(player.Defense) +
			" | Ataque do Inimigo: " + strconv.Itoa(enemy.Attack) +
			" | Defesa do Inimigo: " + strconv.Itoa(enemy.Defense)
	}

	if player.Life == 0 {
		battle.Result = "Inimigo venceu"
		result += " | Inimigo venceu a batalha"
	} else if enemy.Life == 0 {
		battle.Result = "Jogador venceu"
		result += " | Jogador venceu a batalha"
	} else {
		battle.Result = "A batalha continua"
		result += " | A batalha continua"
	}

	if _, err := bs.BattleRepository.AddBattle(battle); err != nil {
		return nil, "", err
	}

	return battle, result, nil
}
