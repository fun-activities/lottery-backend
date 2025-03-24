package admin

import (
	"context"
	"log"
	"github.com/jmoiron/sqlx"
)

type UserAdminRepo struct {
	db *sqlx.DB
}

func (u UserAdminRepo) Create(ctx context.Context) {

}

func (u UserAdminRepo) Update(ctx context.Context) {

}

func (u UserAdminRepo) Delete(ctx context.Context) {

}

func (u UserAdminRepo) Get(ctx context.Context) {
	log.Println("sss")
}
