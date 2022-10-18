package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        int32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	// db
	ListUser(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, id int32) (*User, error)
	CreateUser(ctx context.Context, User *User) error
	UpdateUser(ctx context.Context, User *User) error
	DeleteUser(ctx context.Context, User *User) error
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) List(ctx context.Context) (ps []*User, err error) {
	ps, err = uc.repo.ListUser(ctx)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) Get(ctx context.Context, id int32) (p *User, err error) {
	fmt.Println("biz-user")
	uc.log.WithContext(ctx).Infof("biz.Get: %d", id)
	p, err = uc.repo.GetUser(ctx, id)
	if err != nil {
		return
	}
	return
}

func (uc *UserUsecase) Create(ctx context.Context, User *User) error {
	return uc.repo.CreateUser(ctx, User)
}

func (uc *UserUsecase) Update(ctx context.Context, User *User) error {
	return uc.repo.UpdateUser(ctx, User)
}

func (uc *UserUsecase) Delete(ctx context.Context, User *User) error {
	return uc.repo.DeleteUser(ctx, User)
}
