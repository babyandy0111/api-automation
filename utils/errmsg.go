package utils

var message map[int]string

func init() {
	message = make(map[int]string)
	message[OK] = "SUCCESS"
	message[BAD_REUQEST_ERROR] = "服務器繁忙,請稍後再試"
	message[REUQES_PARAM_ERROR] = "參數錯誤"
	message[USER_NOT_FOUND] = "用戶不存在"
}

func MapErrMsg(errcode int) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return message[BAD_REUQEST_ERROR]
	}
}
