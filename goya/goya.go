package goyawp

import (
	"github.com/sagayosa/goya"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/utils"
)

func Get[T any](addr string, uri string, opt any) T {
	var result T
	resp := goya.Get[*baseresp.BaseResp](utils.StringPlus("http://", addr, uri), opt)
	if resp == nil {
		return result
	}
	if resp.Data == nil {
		return result
	}

	utils.UnMarshalUnknownAny(resp.Data, &result)
	return result
}

func Post[T any](addr string, uri string, opt any) T {
	var result T
	resp := goya.Post[*baseresp.BaseResp](utils.StringPlus("http://", addr, uri), opt)
	if resp == nil {
		return result
	}
	if resp.Data == nil {
		return result
	}

	utils.UnMarshalUnknownAny(resp.Data, &result)
	return result
}

func Put[T any](addr string, uri string, opt any) T {
	var result T
	resp := goya.Put[*baseresp.BaseResp](utils.StringPlus("http://", addr, uri), opt)
	if resp == nil {
		return result
	}
	if resp.Data == nil {
		return result
	}

	utils.UnMarshalUnknownAny(resp.Data, &result)
	return result
}

func Delete[T any](addr string, uri string, opt any) T {
	var result T
	resp := goya.Delete[*baseresp.BaseResp](utils.StringPlus("http://", addr, uri), opt)
	if resp == nil {
		return result
	}
	if resp.Data == nil {
		return result
	}

	utils.UnMarshalUnknownAny(resp.Data, &result)
	return result
}
