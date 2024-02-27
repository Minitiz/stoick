package urlshortener

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type URLShortenerI interface {
	Generate(ctx *gin.Context, req *GenerateReq, rsp *GenerateRsp, log *zap.Logger)
	Redirect(ctx *gin.Context, req *RedirectReq, rsp *RedirectRsp, log *zap.Logger)
}

type URLShortenerStruct struct{}

func New() URLShortenerI {
	return &URLShortenerStruct{}
}
