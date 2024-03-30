package sayorpc

import (
	"encoding/json"
	"net/http"

	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	"github.com/sagayosa/sayo_utils/utils"
)

func DirectGet(addr string, uri string, data map[string]interface{}) (result interface{}, err error) {
	code, body, err := Get(utils.StringPlus("http://", addr, uri), data)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrDirectHttpFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrDirectHttpFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data, nil
}

func DirectPost(addr string, uri string, data interface{}) (result interface{}, err error) {
	code, body, err := Post(utils.StringPlus("http://", addr, uri), data)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrDirectHttpFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrDirectHttpFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data, nil
}

func DirectPut(addr string, uri string, data interface{}) (result interface{}, err error) {
	code, body, err := Put(utils.StringPlus("http://", addr, uri), data)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrDirectHttpFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrDirectHttpFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data, nil
}

func DirectDelete(addr string, uri string, data interface{}) (result interface{}, err error) {
	code, body, err := Delete(utils.StringPlus("http://", addr, uri), data)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return "", sayoerror.ErrorInStatusCode(sayoerror.ErrDirectHttpFailed, code)
	}

	resp := &baseresp.BaseResp{}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.Code != sayoerror.SuccessCode {
		return "", sayoerror.ErrorInMsgCode(sayoerror.ErrDirectHttpFailed, int(resp.Code), resp.Msg)
	}

	return resp.Data, nil
}
