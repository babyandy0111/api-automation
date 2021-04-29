package utils

type (
	NullJson struct{}

	ResponseSuccessBean struct {
		Code int         `json:"code"`
		Msg  string      `json:"message"`
		Data interface{} `json:"data"`
	}
)

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, "OK", data}
}

type ResponseErrorBean struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

func Error(errCode int, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
