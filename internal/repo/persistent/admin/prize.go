package admin

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/fun-activities/lottery-backend/internal/entity"
	"github.com/jmoiron/sqlx"
)

type PrizeAdminRepo struct {
	db *sqlx.DB
}

func NewPrizeAdminRepo(db *sqlx.DB) *PrizeAdminRepo {
	return &PrizeAdminRepo{
		db: db,
	}
}

func (p PrizeAdminRepo) Create(ctx context.Context, prize entity.Prize) (int64, error) {
	log.Printf("%+v", prize)
	query := "INSERT INTO prizes (prize_id,name,type,quantity,image_url,price,total_stock,ext)" +
		"VALUES (?,?,?,?,?,?,?,?)"
	result, err := p.db.ExecContext(ctx, query,
		prize.PrizeId, prize.Name,
		prize.Type, prize.Quantity,
		prize.ImageUrl,
		prize.Price, prize.TotalStock, prize.Ext,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (p PrizeAdminRepo) Update(ctx context.Context, prize entity.Prize) (int64, error) {
	query := "UPDATE prizes SET name = ?, type = ? ,status = ? ,image_url=?" +
		",price=?,total_stock=?,ext=? WHERE prize_id = ?"
	result, err := p.db.ExecContext(ctx, query,
		prize.Name,
		prize.Type, prize.Status,
		prize.ImageUrl,
		prize.Price, prize.TotalStock, prize.Ext, prize.PrizeId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()

}

func (p PrizeAdminRepo) UpdateStatus(ctx context.Context, prizeId string, status entity.Status) (int64, error) {
	query := "UPDATE prizes SET status = ? WHERE prize_id = ?"
	result, err := p.db.ExecContext(ctx, query, status, prizeId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (p PrizeAdminRepo) Get(ctx context.Context, query entity.PrizeQuery) ([]entity.Prize, int, error) {
	var filterValues []interface{}

	var where = []string{"1=1"}
	if len(query.Search) > 0 {
		filterValues = append(filterValues, "%"+query.Search+"%")
		where = append(where, "name like ?")
	}

	if len(query.PrizeId) > 0 {
		filterValues = append(filterValues, query.PrizeId)
		where = append(where, "prize_id = ?")
	}

	if query.Status > 0 {
		filterValues = append(filterValues, query.Status)
		where = append(where, "status = ?")
	}

	sql_total := "select count(*) from prizes where " + strings.Join(where, " and ")
	var total int
	err := p.db.GetContext(ctx, &total, sql_total, filterValues...)
	if err != nil {
		return nil, 0, err
	}
	if total == 0 {
		return nil, 0, fmt.Errorf("empty data")
	}

	sql := "select * from prizes where " + strings.Join(where, " and ")

	page := entity.Page
	pageSize := entity.PageSize

	if query.Page > 0 {
		page = query.Page
	}
	if query.PageSize > 0 {
		pageSize = query.PageSize
	}
	sql += " order by id desc limit ? offset ?"
	filterValues = append(filterValues, pageSize, pageSize*(page-1))

	var prizes []entity.Prize
	err = p.db.SelectContext(ctx, &prizes, sql, filterValues...)
	return prizes, total, err
}
