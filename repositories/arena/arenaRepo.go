package arena

import (
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type arenaRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &arenaRepository{db: db}
}

func (ar *arenaRepository) Create(arena *Arena) (int64, error) {
	res, err := ar.db.NamedExec("INSERT INTO arena(Fighter1, Fighter2, Health1, Health2, X1, Y1, X2, Y2) VALUES"+
		" (:fighter1, :fighter2, :health1, :health2, 0, 0, 0, 0) RETURNING Id", arena)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
func (ar *arenaRepository) Get(id string) (*Arena, error) {
	a := Arena{}
	err := ar.db.QueryRowx("SELECT * FROM arena WHERE Id=$1", id).StructScan(&a)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &a, err
}
func (ar *arenaRepository) Update(arena *Arena) error {
	_, err := ar.db.NamedExec("UPDATE arena SET (:health1, :health2, :x1, :y1, :x2, :y2, :winner) WHERE Id=:id", arena) //Health1=:health1
	return err
}
func (ar *arenaRepository) Delete(id string) error {
	_, err := ar.db.Exec("DELETE FROM arena WHERE id=$1", id) //id!=Id???
	return err
}
