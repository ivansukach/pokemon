package usersProperties

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type propertiesRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &propertiesRepository{db: db}
}

func (pr *propertiesRepository) Create(property *Properties) error {
	_, err := pr.db.NamedExec("INSERT INTO userProperties VALUES (:login, :name)", property) //??
	return err
}

func (pr *propertiesRepository) Delete(id string) error {
	_, err := pr.db.Exec("DELETE FROM userProperties WHERE IdOfPokemon=$1", id)
	return err
}
func (pr *propertiesRepository) GetAll(login string) ([]Properties, error) {
	rows, err := pr.db.Queryx("SELECT * FROM userProperties WHERE Login=$1", login)
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	p := make([]Properties, 0)
	for rows.Next() {
		property := Properties{}
		err = rows.StructScan(&property)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		p = append(p, property)
	}
	return p, err
}
