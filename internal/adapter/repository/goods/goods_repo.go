package goods

import "database/sql"

type goodsRepository struct {
	db *sql.DB
}

func NewGoodsRepository(db *sql.DB) *goodsRepository {
	return &goodsRepository{
		db: db,
	}
}
