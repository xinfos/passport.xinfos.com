package v1

import "github.com/gin-gonic/gin"

//LoginRequest 登陆请求
type LoginRequest struct {
	LoginID     string `json:"login_id"`
	Password    string `json:"password"`
	AppName     string `json:"app_name"`
	AppEntrance string `json:"app_entrance"`
	CsrfToken   string `json:"csrf_token"`
	IsMobile    string `json:"is_mobile"`
	ReturnURL   string `json:"return_url"`
}

//LoginByPwd 密码登陆
func LoginByPwd(c *gin.Context) {
	// var req LoginRequest
}
