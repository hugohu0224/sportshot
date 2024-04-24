package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	pb "sportshot/proto"
	"sportshot/webservice/global"
)

func GetEvents(ctx *gin.Context) {
	res, err := global.EnevtServerClient.SearchEvents(context.Background(), &pb.SearchEventsRequest{
		Name: "",
		Type: "",
		Date: "",
	})
	if err != nil {
		zap.S()
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": res,
		"code": 200,
	})
}
