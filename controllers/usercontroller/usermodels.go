package usercontroller

type createReq struct {
	HereditaryTitle   string `json:"hereditary_title"`
	EducationalTitle  string `json:"educational_title"`
	ProfesionalTitle  string `json:"professional_title"`
	FirstName         string `json:"first_name" validate:"required"`
	MiddleName        string `json:"middle_name"`
	LastName          string `json:"last_name" validate:"required"`
	PopularName       string `json:"popular_name" `
	NickName          string `json:"nick_name" `
	Gender            string `json:"gender" validate:"required"`
	DOBActual         string `json:"dob_actual" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DOBVirtual        string `json:"dob_virtual" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	MobileCountryCode string `json:"mobile_country_code"`
	Mobile            string `json:"mobile"`
	Address           string `json:"address"`
	Address1          string `json:"address1"`
	Address2          string `json:"address2"`
	Village           string `json:"village"`
	City              string `json:"city"`
	District          string `json:"district"`
	State             string `json:"state"`
	Province          string `json:"province"`
	Country           string `json:"Country"`
	Planet            string `json:"planet"`
}
