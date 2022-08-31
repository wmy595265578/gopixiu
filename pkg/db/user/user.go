package user

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/caoyingjunz/gopixiu/pkg/db/model"
)

type UserInterface interface {
	Get(ctx context.Context, uid int64) (*model.User, error)
	List(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, modelUser *model.User) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Delete(ctx context.Context, uid int64) error
	Update(ctx context.Context, modelUser *model.User) error
}

type user struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserInterface {
	return &user{db}
}

func (u *user) Get(ctx context.Context, uid int64) (*model.User, error) {
	var modelUser model.User
	if err := u.db.Where("id = ?", uid).First(&modelUser).Error; err != nil {
		return nil, err
	}
	return &modelUser, nil
}

func (u *user) List(ctx context.Context) ([]model.User, error) {
	var users []model.User
	if tx := u.db.Find(&users); tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func (u *user) Create(ctx context.Context, modelUser *model.User) (*model.User, error) {
	now := time.Now()
	modelUser.GmtCreate = now
	modelUser.GmtModified = now
	if err := u.db.Create(modelUser).Error; err != nil {
		return nil, err
	}
	return modelUser, nil
}

func (u *user) GetByName(ctx context.Context, name string) (*model.User, error) {
	var modelUser model.User
	if err := u.db.Where("name = ?", name).First(&modelUser).Error; err != nil {
		return nil, err
	}
	return &modelUser, nil
}

func (u *user) Delete(ctx context.Context, uid int64) error {
	return u.db.Delete(model.User{}, uid).Error
}

func (u *user) Update(ctx context.Context, modelUser *model.User) error {
	return u.db.Updates(*modelUser).Error
}