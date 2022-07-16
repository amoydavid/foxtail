package forms

type PasswordLoginForm struct {
	// 密码  binding:"required"为必填字段,长度大于3小于20
	Password string `form:"password" json:"password" binding:"required,min=3"`
	//用户名
	Email string `form:"email" json:"email" binding:"required"`
}
