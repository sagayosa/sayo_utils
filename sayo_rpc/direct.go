package sayorpc

import "github.com/grteen/sayo_utils/utils"

func DirectGet(addr string, uri string, data map[string]interface{}) (code int, body []byte, err error) {
	return Get(utils.StringPlus("http://", addr, uri), data)
}

func DirectPost(addr string, uri string, data interface{}) (code int, body []byte, err error) {
	return Post(utils.StringPlus("http://", addr, uri), data)
}

func DirectPut(addr string, uri string, data interface{}) (code int, body []byte, err error) {
	return Put(utils.StringPlus("http://", addr, uri), data)
}

func DirectDelete(addr string, uri string, data interface{}) (code int, body []byte, err error) {
	return Delete(utils.StringPlus("http://", addr, uri), data)
}
