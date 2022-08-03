package goods

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"webApi/goods_web/global"

	proto "webApi/goods_web/proto/v1"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func List(ctx *gin.Context) {
	// 获取传参
	page := ctx.DefaultQuery("page", "0")
	pn, _ := strconv.Atoi(page)

	Size := ctx.DefaultQuery("size", "0")
	PSize, _ := strconv.Atoi(Size)
	rsp, err := global.GoodsSrvClient.BrandList(context.Background(), &proto.BrandFilterRequest{
		Pages:       int32(pn),
		PagePerNums: int32(PSize),
	})
	if err != nil {
		zap.S().Errorw("[Brand List] get list error:%v", err.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	// result := make([]interface{}, 0)
	// for _, value := range rsp.Data {

	// 	result = append(result, user)
	// }
	ctx.JSON(http.StatusOK, rsp)
	zap.S().Debug("Get Brand list")
}
