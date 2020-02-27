package pokemonDescription

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type pokemonRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &pokemonRepository{db: db}
}
func (pr *pokemonRepository) Create(pokemon *Pokemon) error {
	_, err := pr.db.NamedExec("INSERT INTO pokemonDescription VALUES (:name, :image, :price, :healthpoints, :healthregen, "+
		":mana, :manaregen, :armor, :damage, :movementspeed, :attackspeed, , :attackrange)", pokemon)
	return err
}
func (pr *pokemonRepository) Get(name string) (*Pokemon, error) {
	p := Pokemon{}
	err := pr.db.QueryRowx("SELECT * FROM pokemonDescription WHERE name=$1", name).StructScan(&p)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &p, err
}
func (pr *pokemonRepository) Update(pokemon *Pokemon) error {
	_, err := pr.db.NamedExec("UPDATE pokemonDescription SET (Name=:name, Image=:image, Price=:price, HealthPoints=:healthpoints,"+
		" HealthRegen=:healthregen, Mana=:mana, ManaRegen=:manaregen, Armor=:armor, Damage=:damage, MovementSpeed=:movementspeed,"+
		" AttackSpeed=:attackspeed, AttackRange=:attackrange) WHERE Name=:name", pokemon)
	return err
}
func (pr *pokemonRepository) Delete(name string) error {
	_, err := pr.db.Exec("DELETE FROM pokemonDescription WHERE name=$1", name)
	return err
}
func (pr *pokemonRepository) Listing() ([]Pokemon, error) {
	rows, err := pr.db.Queryx("SELECT * FROM pokemonDescription")
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	p := make([]Pokemon, 0)
	for rows.Next() {
		pokemon := Pokemon{}
		err = rows.StructScan(&pokemon)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		p = append(p, pokemon)
	}
	return p, err
}
