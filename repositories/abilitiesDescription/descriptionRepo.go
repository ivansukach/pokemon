package abilitiesDescription

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type descriptionRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &descriptionRepository{db: db}
}

func (dr *descriptionRepository) Create(description *Description) error {
	_, err := dr.db.NamedExec("INSERT INTO abilitiesDescription VALUES (:name, :animation, :damage, :mana, :attackrange)", description)
	return err
}
func (dr *descriptionRepository) Get(name string) (*Description, error) {
	d := Description{}
	err := dr.db.QueryRowx("SELECT * FROM abilitiesDescription WHERE Name=$1", name).StructScan(&d)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &d, err
}
func (dr *descriptionRepository) Update(description *Description) error {
	_, err := dr.db.NamedExec("UPDATE abilitiesDescription SET (Name=:name, Animation=:animation, Damage=:damage,"+
		" Mana=:mana, AttackRange=:attackrange) WHERE Name=:name", description)
	return err
}
func (dr *descriptionRepository) Delete(name string) error {
	_, err := dr.db.Exec("DELETE FROM abilitiesDescription WHERE login=$1", name)
	return err
}
