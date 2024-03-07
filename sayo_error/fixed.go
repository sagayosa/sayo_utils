package sayoerror

import "fmt"

func ErrorInStatusCode(err error, statusCode int) error {
	return ErrMsg(err, fmt.Sprintf("status code = %v", statusCode))
}

func ErrorInMsgCode(err error, code int, msg string) error {
	return ErrMsg(err, fmt.Sprintf("code = %v msg = %v", code, msg))
}
