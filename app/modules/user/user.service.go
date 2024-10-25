package user

import (
	"app/app/modules/user/dto"
	"app/app/modules/user/entity"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type UserService struct {
	db *bun.DB
}

func newService(db *bun.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (service *UserService) User() *dto.UserDTOResponse {
	return &dto.UserDTOResponse{
		ID:        [16]byte{},
		FirstName: "FirstName",
		LastName:  "LastName",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (service *UserService) UserCreateSRV(ctx context.Context, req *dto.UserDTOResponse) (*entity.UserEntity, error) {
	user := &entity.UserEntity{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Username:  req.Username,
		Password:  req.Password,
	}
	_, err := service.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) GetAllUsersSRV(ctx context.Context) ([]*entity.UserEntity, error) {
	var users []*entity.UserEntity
	err := service.db.NewSelect().Model(&users).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *UserService) GetUserByIDSRV(ctx context.Context, id string) (*entity.UserEntity, error) {
	user := new(entity.UserEntity)
	err := service.db.NewSelect().
		Model(user).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) UpdateUserSRV(ctx context.Context, req *dto.UserDTOResponse) (*entity.UserEntity, error) {
	user := new(entity.UserEntity)
	err := service.db.NewSelect().
		Model(user).
		Where("id = ?", req.ID).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Username = req.Username
	user.UpdatedAt = time.Now().Unix()

	_, err = service.db.NewUpdate().
		Model(user).
		Where("id = ?", req.ID).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (service *UserService) DeleteUserSRV(ctx context.Context, userID uuid.UUID) error {
	user := new(entity.UserEntity)
	err := service.db.NewSelect().
		Model(user).
		Where("id = ?", userID).
		Scan(ctx)

	if err != nil {
		return err
	}

	_, err = service.db.NewUpdate().
		Model(user).
		Set("deleted_at = ?", time.Now()).
		Where("id = ?", userID).
		Exec(ctx)

	return err

}

func (service *UserService) ResetPasswordSRV(ctx context.Context, req *dto.UserResetPasswordReq) (*entity.UserEntity, error) {
	user := &entity.UserEntity{}
	err := service.db.NewSelect().Model(user).Where("id = ?", req.ID).Scan(ctx)
	if err != nil {
		return nil, err
	}

	if user.Password != req.Password {
		return nil, fmt.Errorf("รหัสเก่าไม่ถูกต้อง")
	}

	if user.Password == req.Password {
		return nil, fmt.Errorf("รหัสไม่สามารถตรงกันได้")
	}

	_, err = service.db.NewUpdate().Model(user).Where("id = ?", req.ID).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil

}
