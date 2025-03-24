package lottery

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type LotteryRepo struct {
	db *sqlx.DB
}

func NewLotteryRepo(db *sqlx.DB) *LotteryRepo {
	return &LotteryRepo{
		db: db,
	}
}

func (l *LotteryRepo) DoLottery(ctx context.Context) {
	l.db.Query("select * from test limit 1")
}

func (l *LotteryRepo) Get(ctx context.Context) {
	l.db.Query("select * from test limit 1")
}
