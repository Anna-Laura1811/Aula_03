package service

import (
	"errors"
	"fmt"
	"math/rand"

	"RPG_AULA03/internal/entity"
	"RPG_AULA03/internal/repository"
)

type EnemyService struct {
	EnemyRepository repository.EnemyRepository
}

func NewEnemyService(EnemyRepository repository.EnemyRepository) *EnemyService {
	return &EnemyService{EnemyRepository: EnemyRepository}
}

func (es *EnemyService) AddEnemy(nickname string, life, attack, defense, heal int) (*entity.Enemy, error) {
	// Verificações de validação
	if nickname == "" || life == 0 || attack == 0 || defense == 0 || heal == 0 {
		return nil, errors.New("enemy nickname, life, attack, defense and heal are required")
	}

	if len(nickname) > 255 {
		return nil, errors.New("enemy nickname cannot exceed 255 characters")
	}

	if defense > 10 || defense <= 0 {
		return nil, errors.New("enemy defense must be between 1 and 10")
	}

	if attack > 10 || attack <= 0 {
		return nil, errors.New("enemy attack must be between 1 and 10")
	}

	if life > 100 || life <= 0 {
		return nil, errors.New("enemy life must be between 1 and 100")
	}

	if heal > 10 || heal <= 0 {
		return nil, errors.New("enemy heal must be between 1 and 10")
	}

	// Verificar se o inimigo já existe pelo nickname
	enemy, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy != nil {
		return nil, errors.New("enemy nickname already exists")
	}

	// Criar e salvar o inimigo
	enemy = entity.NewEnemy(nickname, life, attack, defense, heal) // Inclua heal aqui
	if _, err := es.EnemyRepository.AddEnemy(enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}

func (es *EnemyService) LoadEnemies() ([]*entity.Enemy, error) {
	enemies, err := es.EnemyRepository.LoadEnemies()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}

	if enemies == nil {
		return []*entity.Enemy{}, nil
	}
	return enemies, nil
}

func (es *EnemyService) DeleteEnemy(id string) error {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	if enemy == nil {
		return errors.New("enemy id not found")
	}
	if err := es.EnemyRepository.DeleteEnemyById(id); err != nil {
		fmt.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func (es *EnemyService) LoadEnemy(id string) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}
	return enemy, nil
}

func (es *EnemyService) SaveEnemy(id, nickname string, life, attack, defense, heal int) (*entity.Enemy, error) {
	enemy, err := es.EnemyRepository.LoadEnemyById(id)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	if enemy == nil {
		return nil, errors.New("enemy id not found")
	}

	if nickname != "" && nickname != enemy.Nickname {
		hasNickname, err := es.EnemyRepository.LoadEnemyByNickname(nickname)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("internal server error")
		}
		if hasNickname != nil {
			return nil, errors.New("enemy nickname already exists")
		}
		if len(nickname) > 255 {
			return nil, errors.New("enemy nickname cannot exceed 255 characters")
		}
		enemy.Nickname = nickname
	}

	if attack != 0 && attack != enemy.Attack {
		if attack > 10 || attack <= 0 {
			return nil, errors.New("enemy attack must be between 1 and 10")
		}
		enemy.Attack = attack
	}
	if defense != 0 && defense != enemy.Defense {
		if defense > 10 || defense <= 0 {
			return nil, errors.New("enemy defense must be between 1 and 10")
		}
		enemy.Defense = defense
	}

	if life != 0 && life != enemy.Life {
		if life > 100 || life <= 0 {
			return nil, errors.New("enemy life must be between 1 and 100")
		}
		enemy.Life = life
	}

	if heal != 0 && heal != enemy.Heal {
		if heal > 5 || heal <= 0 {
			return nil, errors.New("enemy heal must be between 1 and 10")
		}
		enemy.Heal = heal
	}

	if err := es.EnemyRepository.SaveEnemy(id, enemy); err != nil {
		fmt.Println(err)
		return nil, errors.New("internal server error")
	}
	return enemy, nil
}

// Método para curar o inimigo
func (es *EnemyService) HealEnemy(id string) (*entity.Enemy, error) {
	enemy, err := es.LoadEnemy(id)
	if err != nil {
		return nil, err
	}

	// Determinar a quantidade de cura (por exemplo, um valor aleatório entre 1 e 10)
	healAmount := rand.Intn(10) + 1
	enemy.Life += healAmount

	// Garantir que a vida do inimigo não exceda o máximo de 100
	if enemy.Life > 100 {
		enemy.Life = 100
	}

	if err := es.EnemyRepository.SaveEnemy(enemy.ID, enemy); err != nil {
		return nil, errors.New("falha ao atualizar a vida do inimigo após a cura")
	}

	return enemy, nil
}
