package usersProperties

type Properties struct {
	IdOfPokemon int64  `db:"idofpokemon"`
	Login       string `db:"login"`
	Name        string `db:"name"`
}
type Repository interface {
	Create(property *Properties) error
	Delete(id int64) error
	GetAll(login string) ([]Properties, error)
	GetNameById(id int64) (string, error)
}
