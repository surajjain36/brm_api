package authcontroller

//LoginReq from happay
type LoginReq struct {
	MobileNumber string `json:"mobile_number"  validate:"required"`
	Password     string `json:"password" validate:"required"`
}

//RegisterReq from happay
type RegisterReq struct {
	MobileNumber    string `json:"mobile_number"  validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=ConfirmPassword"`
}

//ForgotPWReq from happay
type ForgotPWReq struct {
	MobileNumber string `json:"mobile_number"  validate:"required"`
}

//VerifyOTPReq from happay
type VerifyOTPReq struct {
	MobileNumber    string `json:"mobile_number"  validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=ConfirmPassword"`
	OTP             string `json:"otp" validate:"required"`
}
