package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Anna-Laura1811/RPG_GO_BD/internal/handler"
	repository 	"github.com/Anna-Laura1811/RPG_GO_BD/internal/repositoy"
	"github.com/Anna-Laura1811/RPG_GO_BD/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	// "postgresql://<username>:<password>@<database_ip>/todos?sslmode=disable"
	dsn := "postgresql://postgres:root@localhost/postgres?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}

	playerRepository := repository.NewPlayerRepository(db)
	playerService := service.NewPlayerService(*playerRepository)
	playerHandler := handler.NewPlayerHandler(playerService)
	enemyRepository := repository.NewEnemyRepository(db)
	enemyService := service.NewEnemyService(*enemyRepository)
	enemyHandler := handler.NewEnemyHandler(enemyService)
     
	battleRepository := repository.NewBattleRepository(db)
	battleService := service.NewBattleService(*playerRepository, *enemyRepository, *battleRepository)
	battleHandler := handler.NewBattleHandler(battleService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /player", playerHandler.AddPlayer)
	mux.HandleFunc("GET /player", playerHandler.LoadPlayers)
	mux.HandleFunc("DELETE /player/{id}", playerHandler.DeletePlayer)
	mux.HandleFunc("GET /player/{id}", playerHandler.LoadPlayer)
	mux.HandleFunc("PUT /player/{id}", playerHandler.SavePlayer)
	mux.HandleFunc("POST /enemy", enemyHandler.AddEnemy)
	mux.HandleFunc("GET /enemy", enemyHandler.LoadEnemies)
	mux.HandleFunc("DELETE /enemy/{id}", enemyHandler.DeleteEnemy)
	mux.HandleFunc("GET /enemy/{id}", enemyHandler.LoadEnemy)
	mux.HandleFunc("PUT /enemy/{id}", enemyHandler.SaveEnemy)
	mux.HandleFunc("POST /battle", battleHandler.AddBattle)
	mux.HandleFunc("GET /battle", battleHandler.LoadBatlle)
	mux.HandleFunc("DELETE /battle/{id}", battleHandler.DeleteBattle)
	mux.HandleFunc("GET /battle/{id}", battleHandler.LoadBattle)
	mux.HandleFunc("PUT /battle/{id}", battleHandler.SaveBattle)


	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}