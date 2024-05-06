package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"sportshot/utils/models/event"
	pb "sportshot/utils/proto"
	"sportshot/webserver/global"
)

func GetEvents(ctx *gin.Context) {
	var form event.SearchEventsForm
	if err := ctx.ShouldBindQuery(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.S().Errorf("GetEvents err: %v", err)
		return
	}

	res, err := global.EventServerClient.SearchEvents(context.Background(), &pb.SearchEventsRequest{
		LeagueName: form.LeagueName,
		Type:       form.Type,
		Date:       form.Date,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		zap.S().Errorf("search events error: %v", err)
		return
	}

	ctx.JSON(200, gin.H{
		"data": res,
		"code": 200,
	})
}
