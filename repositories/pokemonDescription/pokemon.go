package pokemonDescription

type Pokemon struct {
	Name          string  `db:"name"`
	Image         string  `db:"image"`
	Price         int32   `db:"price"`
	HealthPoints  int32   `db:"healthpoints"`
	HealthRegen   int32   `db:"healthregen"`
	Mana          int32   `db:"mana"`
	ManaRegen     int32   `db:"manaregen"`
	Armor         float32 `db:"armor"`
	Damage        int32   `db:"damage"`
	MovementSpeed int32   `db:"movementspeed"`
	AttackSpeed   int32   `db:"attackspeed"`
	AttackRange   int32   `db:"attackrange"`
}
type Repository interface {
	Create(pokemon *Pokemon) error
	Get(name string) (*Pokemon, error)
	Update(pokemon *Pokemon) error
	Delete(name string) error
	Listing() ([]Pokemon, error)
}
