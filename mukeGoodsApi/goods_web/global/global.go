package global

import (
	"webApi/goods_web/config"
	proto "webApi/goods_web/proto/v1"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig   *config.ServerConfig = &config.ServerConfig{}
	Trans          ut.Translator
	GoodsSrvClient proto.GoodsClient
)
