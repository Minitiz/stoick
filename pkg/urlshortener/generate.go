package urlshortener

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minitiz/url-shortener/pkg/db"
	"go.uber.org/zap"
)

type GenerateReq struct {
	URL string `json:"url"`
}

type GenerateRsp struct {
	Short string `json:"short"`
}

func (s *URLShortenerStruct) Generate(ctx *gin.Context, req *GenerateReq, rsp *GenerateRsp, log *zap.Logger) {
	Status := 200
	defer ctx.PureJSON(Status, rsp)
	Data := db.Data{
		URL:    req.URL,
		Short:  Base62Encode(stringToInt64(req.URL)),
		Access: 0,
		Expiry: time.Now().Unix() + 30,
	}
	if err := Data.InsertURL(); err != nil {
		log.Error("Error inserting URL", zap.Error(err))
		Status = 500
	}
	rsp.Short = Data.Short
}
