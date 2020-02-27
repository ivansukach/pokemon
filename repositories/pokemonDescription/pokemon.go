package pokemonDescription

type Pokemon struct {
	name          string  `db:"name"`
	image         string  `db:"image"`
	price         int32   `db:"price"`
	healthPoints  int32   `db:"healthpoints"`
	healthRegen   int32   `db:"healthregen"`
	mana          int32   `db:"mana"`
	manaRegen     int32   `db:"manaregen"`
	armor         float32 `db:"armor"`
	damage        int32   `db:"damage"`
	movementSpeed int32   `db:"movementspeed"`
	attackSpeed   int32   `db:"attackspeed"`
	attackRange   int32   `db:"attackrange"`
}
type Repository interface {
	Create(pokemon *Pokemon) error
	Get(name string) (*Pokemon, error)
	Update(pokemon *Pokemon) error
	Delete(name string) error
	Listing() ([]Pokemon, error)
}
