package abilitiesDescription

type Description struct {
	Name        string `db:"name"`
	Animation   string `db:"animation"`
	Damage      int32  `db:"damage"`
	Mana        int32  `db:"mana"`
	AttackRange int32  `db:"attackrange"`
}
type Repository interface {
	Create(description *Description) error
	Get(name string) (*Description, error)
	Update(description *Description) error
	Delete(name string) error
}
