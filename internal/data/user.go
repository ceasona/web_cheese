package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"web_cheese/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	fmt.Println("NewUserRepo")
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ar *userRepo) ListUser(ctx context.Context) ([]*biz.User, error) {
	var users []biz.User
	rv := make([]*biz.User, 0)
	result := ar.data.db.Find(&users)
	rows, err := ar.data.db.Model(&biz.User{}).Rows()
	for rows.Next() {
		user := &biz.User{}
		err = ar.data.db.ScanRows(rows, user)
		rv = append(rv, user)
		if err != nil {
			return nil, err
		}
	}
	return rv, result.Error
}

func (ar *userRepo) GetUser(ctx context.Context, id int32) (*biz.User, error) {
	fmt.Println("data-user")
	var u biz.User
	ar.data.db.Where("id = ?", id).First(&u)
	ar.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (ar *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	result := ar.data.db.Create(&user)
	return result.Error
}

func (ar *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	id := user.ID
	var u biz.User
	ar.data.db.Where("id = ?", id).First(&u)
	if u == (biz.User{}) {
		var errdelete = errors.New("no this row")
		return errdelete
	}
	result := ar.data.db.Save(&user)
	return result.Error
}

func (ar *userRepo) DeleteUser(ctx context.Context, user *biz.User) error {
	result := ar.data.db.Delete(&user)
	if result.RowsAffected == 0 {
		var errdelete = errors.New("no rows")
		return errdelete
	}
	return result.Error
}
