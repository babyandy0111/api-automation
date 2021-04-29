package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/status"
	"net/http"
)

func PasswordEncrypt(salt, password string) string {
	dk, _ := scrypt.Key([]byte(password), []byte(salt), 32768, 8, 1, 32)
	return fmt.Sprintf("%x", string(dk))
}

func GetJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

// http方法
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		// 成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		// 错误返回
		errcode := BAD_REUQEST_ERROR
		errmsg := message[BAD_REUQEST_ERROR]
		if e, ok := err.(*CodeError); ok {
			// 自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			originErr := errors.Cause(err) // err类型
			if gstatus, ok := status.FromError(originErr); ok {
				// grpc err错误
				errmsg = gstatus.Message()
			}
		}
		logx.WithContext(r.Context()).Error("【GATEWAY-SRV-ERR】 : %+v ", err)

		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
	}
}

// http参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", MapErrMsg(REUQES_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(REUQES_PARAM_ERROR, errMsg))
}
