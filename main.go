package main

import (
	"fmt"
	"github.com/ivansukach/pokemon-auth/repositories/users"
	"github.com/ivansukach/pokemon/repositories/abilities"
	"github.com/ivansukach/pokemon/repositories/pokemonDescription"
	"github.com/ivansukach/pokemon/repositories/usersProperties"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello, world")
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=books sslmode=disable")
	if err != nil {
		log.Error(err)
		return
	}
	_ = abilities.New(db)
	_ = usersProperties.New(db)
	_ = users.New(db)
	_ = pokemonDescription.New(db)
}
