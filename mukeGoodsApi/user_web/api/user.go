package api

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"webApi/user_web/forms"
	"webApi/user_web/global"
	"webApi/user_web/global/response"
	"webApi/user_web/middlewares"
	"webApi/user_web/models"
	userpb "webApi/user_web/proto/v1"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "error: cannot Dial user serve ",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "其他错误",
				})
			}
			return
		}
	}

}

func GetUserList(ctx *gin.Context) {

	// 获取传参
	page := ctx.DefaultQuery("page", "0")
	pn, _ := strconv.Atoi(page)

	Size := ctx.DefaultQuery("size", "0")
	PSize, _ := strconv.Atoi(Size)
	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &userpb.PageInfo{
		Pn:    uint32(pn),
		PSize: uint32(PSize),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] get list error:%v", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Userinfo {
		user := response.UserRespone{
			ID:       value.Id,
			NickName: value.NickName,
			// BirthDay: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("006-01-02 15:04:05"),
			BirthDay: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Mobile:   value.Mobile,
			Gender:   value.Gender,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	zap.S().Debug("Get user list")
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}

	for filed, val := range fileds {
		rsp[filed[strings.Index(filed, ".")+1:]] = val
	}
	return rsp
}
func HandleValidatorErr(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}

// login
func PasswordLogin(c *gin.Context) {
	// 表单验证
	passwordForm := forms.PassWordForm{}
	if err := c.ShouldBind(&passwordForm); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	// 处理登录逻辑
	// 首先查看用户是否存在
	// 存在则直接返回 jwt token
	// 不存在 返回错误
	ctx := context.Background()
	rsp, err := global.UserSrvClient.GetUserByMobile(ctx, &userpb.MobileRequest{
		Mobile: passwordForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "GetUserByMobile 失败",
				})
			}
		}
		return
	}

	// 获取用户 进行check
	pas, pasErr := global.UserSrvClient.CheckAPassword(ctx, &userpb.CheckPasswordRequest{
		Password:          passwordForm.PassWord,
		EncruptedPassword: passwordForm.RePassWord,
	})
	zap.S().Infof("check pas success:", pas.Success)
	if pasErr != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"password": "check 错误",
		})
		return
	}
	if !pas.Success {
		c.JSON(http.StatusBadRequest, map[string]string{
			"password": "登录失败",
		})
		return
	}

	// 验证通过后，生成jwt token
	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(rsp.Id),
		NickName:    rsp.NickName,
		AuthorityId: uint(rsp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            //签名生效时间
			ExpiresAt: time.Now().Unix() + 86400*10, //有效期
			Issuer:    "webApi",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "token签名失败",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "ok",
		"data": map[string]interface{}{
			"token":      token,
			"nickname":   rsp.NickName,
			"expired_at": claims.StandardClaims.ExpiresAt * 1000,
		},
	})

}
