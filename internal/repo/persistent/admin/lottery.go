package admin

import (
	"context"

	"log"

	"github.com/jmoiron/sqlx"
)

type LotteryAdminRepo struct {
	db *sqlx.DB
}

func NewLotteryAdminRepo(db *sqlx.DB) *LotteryAdminRepo {
	return &LotteryAdminRepo{
		db: db,
	}
}

func (l LotteryAdminRepo) Create(ctx context.Context) {

}

func (l LotteryAdminRepo) Update(ctx context.Context) {

}

func (l LotteryAdminRepo) Delete(ctx context.Context) {

}

type User struct {
	ID   int64 `db:"id"` // 注意 db 标签
	Name int   `db:"name"`
	Age  int   `db:"age"`
}

func (l LotteryAdminRepo) Get(ctx context.Context) {
	user := &User{}
	err := l.db.Get(user, "select * from test limit 1")
	log.Println(err, user.ID)
}
