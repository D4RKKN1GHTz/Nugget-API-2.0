package user

import (
	"app/app/modules/base"
	"app/app/modules/user/dto"
	"time"

	// "app/app/provider/database"
	// "app/app/provider/redis"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserCTL struct {
	userService *UserService
}

func newCTL(userService *UserService) *UserCTL {
	return &UserCTL{
		userService: userService,
	}
}

func (u *UserCTL) Get(ctx *gin.Context) {
	// database.DB()
	// database.DB("db2")
	// redis.DB()
	// redis.DB("rdb2")
	base.Success(ctx, dto.UserDTOResponse{})
}

func (u *UserCTL) UserCreate(ctx *gin.Context) {
	req := dto.UserDTOResponse{
		FirstName: "TTTT",
		LastName:  "YYYY",
		Username:  "TYTY",
		Password:  "12345678",
	}

	data, err := u.userService.UserCreateSRV(ctx, &req)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
		return
	}

	resp := dto.UserDTOResponse{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
	}

	base.Success(ctx, resp)
}

func (u *UserCTL) UserGetAll(ctx *gin.Context) {
	resp, err := u.userService.GetAllUsersSRV(ctx)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
	}
	base.Success(ctx, resp)
}

func (u *UserCTL) UserGetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	resp, err := u.userService.GetUserByIDSRV(ctx, id)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
	}
	base.Success(ctx, resp)
}

func (u *UserCTL) UserUpdate(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		base.BadRequest(ctx, "invalid ID", nil)
		return
	}

	var req dto.UserDTOResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		base.BadRequest(ctx, "invalid data", nil)
		return
	}

	req.ID = userID

	data, err := u.userService.UpdateUserSRV(ctx, &req)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
		return
	}

	parsedID, err := uuid.Parse(data.ID)
	if err != nil {
		base.BadRequest(ctx, "error parsing user ID", nil)
		return
	}

	resp := dto.UserDTOResponse{
		ID:        parsedID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		UpdatedAt: time.Now(),
	}

	base.Success(ctx, resp)
}

func (u *UserCTL) UserDelete(ctx *gin.Context) {
	userID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		base.BadRequest(ctx, "invalid id", nil)
		return
	}

	err = u.userService.DeleteUserSRV(ctx, userID)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
		return
	}
	base.Success(ctx, nil)
}

func (u *UserCTL) ResetPassword(ctx *gin.Context) {
	req := dto.UserResetPasswordReq{
		Password:    "12345678",
		NewPassword: "87654321",
	}

	data, err := u.userService.ResetPasswordSRV(ctx, &req)
	if err != nil {
		base.BadRequest(ctx, err.Error(), nil)
		return
	}

	resp := dto.UserDTOResponse{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
	}

	base.Success(ctx, resp)
}
