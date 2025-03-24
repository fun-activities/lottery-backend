package admin

import (
	"context"

	"github.com/fun-activities/lottery-backend/internal/repo"
)

type LotteryAdminService struct {
	repo repo.LotteryAdmin
}

func NewLotteryAdminService(repo repo.LotteryAdmin) *LotteryAdminService {
	return &LotteryAdminService{
		repo: repo,
	}
}

func (l *LotteryAdminService) Create(ctx context.Context) {

}

func (l *LotteryAdminService) Get(ctx context.Context) {
	l.repo.Get(ctx)
}
