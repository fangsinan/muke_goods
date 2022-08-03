package global

import (
	"webApi/user_web/config"
	userpb "webApi/user_web/proto/v1"

	ut "github.com/go-playground/universal-translator"
)

var (
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	Trans         ut.Translator
	UserSrvClient userpb.UserClient
)
