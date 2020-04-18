package v1

import (
	"net/http"
	"passport.xinfos.com/api"
	"passport.xinfos.com/internal/model"
	"passport.xinfos.com/internal/service"
	"passport.xinfos.com/pkg/logger"
	"passport.xinfos.com/utils/errs"
	"passport.xinfos.com/utils/identity"

	"github.com/gin-gonic/gin"
)

//GetUserInfoByIDRequest - request
type GetUserInfoByIDRequest struct {
	Request string `json:"request_id"`
	UserID  uint64 `json:"user_id"`
}

//GetUserInfoByID - Get user info by user_id
func GetUserInfoByID(c *gin.Context) {
	var req GetUserInfoByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		api.JSON(c, http.StatusOK, nil, nil)
		return
	}
	user, _ := service.NewUserService().GetUserInfoByID(req.UserID)
	api.JSON(c, http.StatusOK, user, nil)
	return
}

//GetAllUsers - request
func GetAllUsers(c *gin.Context) {

}

type CreateUserRequest struct {
	Request string `json:"request_id"`
	Name    string `json:"name"  binding:"required"`
	Phone   string `json:"phone"  binding:"required"`
	IDCard  string `json:"id_card" binding:"required"`
}

//CreateUser - create user request
func CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Error(err.Error())
		api.JSON(c, errs.ErrParamInvalid, nil, nil)
		return
	}

	//身份证校验
	idCard, err := identity.IsValidCitizenNo18(req.IDCard)
	if err != nil {
		api.JSON(c, errs.ErrParamInvalid, nil, err)
		return
	}

	userId, err := service.NewUserService().Create(&model.User{
		Name:     req.Name,
		Phone:    req.Phone,
		IDCard:   req.IDCard,
		Birthday: idCard.Birthday,
		Gender:   idCard.Gender,
		Age:      idCard.Age,
	})
	if err != nil {
		logger.Error(err.Error())
		api.JSON(c, errs.ErrParamInvalid, nil, nil)
		return
	}

	api.JSON(c, http.StatusOK, map[string]uint64{"user_id": userId}, nil)
	return
}
