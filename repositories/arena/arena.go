package arena

type Arena struct {
	Id       int64   `db:"id"`
	Fighter1 string  `db:"fighter1"`
	Fighter2 string  `db:"fighter2"`
	Health1  int32   `db:"health1"`
	Health2  int32   `db:"health2"`
	X1       float32 `db:"x1"`
	Y1       float32 `db:"y1"`
	X2       float32 `db:"x2"`
	Y2       float32 `db:"y2"`
	Winner   string  `db:"winner"`
}
type Repository interface {
	Create(arena *Arena) (int64, error)
	Get(id string) (*Arena, error)
	Update(arena *Arena) error
	Delete(id string) error
}
