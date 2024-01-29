package sayoerror

func ErrorInStatusCode(err error, statusCode int) error {
	return Msg(err, "status code = %v", statusCode)
}

func ErrorInMsgCode(err error, code int, msg string) error {
	return Msg(err, "code = %v msg = %v", code, msg)
}
