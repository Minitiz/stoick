package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minitiz/url-shortener/pkg/logger"
	"github.com/minitiz/url-shortener/pkg/urlshortener"
)

func (h *Handlers) urlShortener(r *gin.Engine) {
	pkg := urlshortener.New()
	r.POST("/generate", func(ctx *gin.Context) {
		log := logger.Set(ctx)
		log.Info("[Start]")
		defer log.Info("[end]")
		var req urlshortener.GenerateReq
		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			ctx.String(http.StatusOK, "Bad Request")
			return
		}
		pkg.Generate(ctx, &req, &urlshortener.GenerateRsp{}, log)
	})

	r.GET("/*short", func(ctx *gin.Context) {
		log := logger.Set(ctx)
		log.Info("[Start]")
		defer log.Info("[end]")
		pkg.Redirect(ctx, &urlshortener.RedirectReq{Short: strings.TrimLeft(ctx.Param("short"), "/")}, &urlshortener.RedirectRsp{}, log)
	})
}
