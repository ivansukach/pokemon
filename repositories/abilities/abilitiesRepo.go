package abilities

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type abilitiesRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &abilitiesRepository{db: db}
}

func (pr *abilitiesRepository) Create(ability *Ability) error {
	_, err := pr.db.NamedExec("INSERT INTO abilities VALUES (:name,:electroshock,:flamethrower,:waterflow)", ability)
	return err
}

func (pr *abilitiesRepository) Delete(name string) error {
	_, err := pr.db.Exec("DELETE FROM abilities WHERE Name=$1", name)
	return err
}
func (pr *abilitiesRepository) Get(name string) (*Ability, error) {
	ability := Ability{}
	err := pr.db.QueryRowx("SELECT * FROM abilities WHERE Name=$1", name).StructScan(&ability)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &ability, err
}
