package helpers

import "gin-api-blog/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	ResultCode       ResultCode                     `json:"resultCode"`
	ValidationErrors *[]validations.ValidationError `json:"validationErrors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, resultCode ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success:    success,
		ResultCode: resultCode,
		Result:     result,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
		Result:     result,
	}
}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode ResultCode, err any) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
		Result:     result,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationError(err),
		Result:           result,
	}
}
