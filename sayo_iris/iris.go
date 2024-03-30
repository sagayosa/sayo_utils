package sayoiris

import (
	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	sayolog "github.com/sagayosa/sayo_utils/sayo_log"
)

type HandlerFunc func(iris.Context)

func IrisCtxJSONWrap(f func(ctx iris.Context) (*baseresp.BaseResp, error)) HandlerFunc {
	return func(ctx iris.Context) {
		resp, err := f(ctx)
		if err != nil {
			sayolog.Err(err).Error()
		}
		ctx.JSON(resp)
	}
}
