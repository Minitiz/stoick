package urlshortener

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minitiz/url-shortener/pkg/db"
	"go.uber.org/zap"
)

type RedirectReq struct {
	Short string `json:"short"`
}

type RedirectRsp struct {
	Expiry bool `json:"expiry"`
}

func (s *URLShortenerStruct) Redirect(ctx *gin.Context, req *RedirectReq, rsp *RedirectRsp, log *zap.Logger) {
	Status := 201
	defer ctx.PureJSON(Status, rsp)

	log.Info("[Start]")
	defer log.Info("[end]")
	log.Info("Redirect", zap.Any("req", req))
	data, err := db.ReadDataAndIncrementAccess(req.Short)
	if err != nil {
		log.Error("Error reading data", zap.Error(err))
		Status = 500
		return
	}
	log.Info("Data", zap.Any("data", data))
	if data.URL == "" || time.Now().Unix() > data.Expiry {
		Status = http.StatusGone
		rsp.Expiry = true
		return
	}
	ctx.Redirect(301, data.URL)
}
