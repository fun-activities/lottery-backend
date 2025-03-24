package repo

import (
	"context"

	"github.com/fun-activities/lottery-backend/internal/entity"
	"github.com/fun-activities/lottery-backend/internal/repo/persistent/admin"
	"github.com/jmoiron/sqlx"
)

type LotteryAdmin interface {
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	Get(ctx context.Context)
}

type Lottery interface {
	DoLottery(ctx context.Context)
	Get(ctx context.Context)
}

type PrizeAdmin interface {
	Create(ctx context.Context, prize entity.Prize) (int64, error)
	Update(ctx context.Context, prize entity.Prize) (int64, error)
	UpdateStatus(ctx context.Context, prizeId string, status entity.Status) (int64, error)
	Get(ctx context.Context, query entity.PrizeQuery) ([]entity.Prize, int,error)
}

type Repo struct {
	LotteryAdmin LotteryAdmin
	PrizeAdmin   PrizeAdmin
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		LotteryAdmin: admin.NewLotteryAdminRepo(db),
		PrizeAdmin:   admin.NewPrizeAdminRepo(db),
	}
}
