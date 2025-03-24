package service

import (
	"context"

	"github.com/fun-activities/lottery-backend/internal/entity"
	"github.com/fun-activities/lottery-backend/internal/repo"
	a "github.com/fun-activities/lottery-backend/internal/service/admin"
)

type lotteryAdmin interface {
	Create(ctx context.Context)
	Get(ctx context.Context)
}

type prizeAdmin interface {
	Get(ctx context.Context, query entity.PrizeQuery) ([]entity.Prize, int,error)
	Create(ctx context.Context, prizeEntity entity.Prize) (int64, error)
	Update(ctx context.Context, prizeEntity entity.Prize) (int64, error)
	UpdateStatus(ctx context.Context, prizeID string, status entity.Status) (int64, error)
}

type Service struct {
	LotteryAdmin lotteryAdmin
	PrizeAdmin   prizeAdmin
}

type Dependent struct {
	Repo *repo.Repo
}

func New(dep Dependent) Service {
	return Service{
		LotteryAdmin: a.NewLotteryAdminService(dep.Repo.LotteryAdmin),
		PrizeAdmin:   a.NewPrizeAdminService(dep.Repo.PrizeAdmin),
	}
}
