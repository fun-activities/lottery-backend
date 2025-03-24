package admin

import (
	"context"

	"github.com/fun-activities/lottery-backend/internal/entity"
	"github.com/fun-activities/lottery-backend/internal/repo"
	"github.com/fun-activities/lottery-backend/pkg/uuid"
)

const PREFIX = "p-"

type PrizeAdminService struct {
	repo repo.PrizeAdmin
}

func NewPrizeAdminService(repo repo.PrizeAdmin) *PrizeAdminService {
	return &PrizeAdminService{
		repo: repo,
	}
}

func (p *PrizeAdminService) Create(ctx context.Context, prizeEntity entity.Prize) (int64, error) {
	prizeEntity.PrizeId = PREFIX + uuid.Gen()
	return p.repo.Create(ctx, prizeEntity)
}

func (p *PrizeAdminService) Update(ctx context.Context, prizeEntity entity.Prize) (int64, error) {
	return p.repo.Update(ctx, prizeEntity)
}

func (p *PrizeAdminService) UpdateStatus(ctx context.Context, prizeID string, status entity.Status) (int64, error) {
	return p.repo.UpdateStatus(ctx, prizeID, status)
}

func (p *PrizeAdminService) Get(ctx context.Context, query entity.PrizeQuery) ([]entity.Prize, int,error) {
	return p.repo.Get(ctx, query)
}
