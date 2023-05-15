package response

import "github.com/go-playground/validator/v10"

//HTTPResponse represents response body of this API
type HTTPResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

//ErrorResponse to req
type ErrorResponse struct {
	FailedField string `json:"failed_fields"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

//ValidateStruct to validate
func ValidateStruct(input interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
