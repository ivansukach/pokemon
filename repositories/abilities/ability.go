package abilities

type Ability struct {
	Name         string `db:"name"`
	Electroshock bool   `db:"electroshock"`
	Flamethrower bool   `db:"flamethrower"`
	Waterflow    bool   `db:"waterflow"`
}
type Repository interface {
	Create(ability *Ability) error
	Get(name string) (*Ability, error)
	Delete(login string) error
}
