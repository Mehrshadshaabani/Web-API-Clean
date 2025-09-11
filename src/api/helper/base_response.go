package helper

import "github.com/Mehrshadshaabani/Web-API-Clean/api/validations"

type BaseHttpResponse struct {
	Result           any                           `json:"result"`
	Success          bool                          `json:"success"`
	ResultCode       int                           `json:"result_code"`
	ValidationErrors []validations.ValidationError `json:"validationerrors"`
	Error            any                           `json:"error"`
}

func GenerateBaseHttpResponse(result any, success bool, resultCode int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
	}
}
func GenerateBaseHttpResponseWithError(result any, success bool, resultCode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:     result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}
}
func GenerateBaseHttpResponseWithValidationErrors(result any, success bool, resultcode int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		ResultCode:       resultcode,
		ValidationErrors: *validations.GetValidationErrors(err),
	}
}
