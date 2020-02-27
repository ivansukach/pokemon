package usersProperties

type Properties struct {
	IdOfPokemon string `db:"idofpokemon"`
	Login       string `db:"login"`
	Name        string `db:"name"`
}
type Repository interface {
	Create(property *Properties) error
	Delete(id string) error
	GetAll() ([]Properties, error)
}
