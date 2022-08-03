package forms

type PassWordForm struct {
	Mobile     string `form:"mobile" json:"mobile" binding:"required,mobileValidator"`
	PassWord   string `form:"password" json:"password" binding:"required,min=3,max=10"`
	RePassWord string `form:"re_password" json:"re_password" binding:"required,eqfield=PassWord"`
}
