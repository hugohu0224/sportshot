package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"sportshot/utils/global"
	"sportshot/utils/models/event"
	pb "sportshot/utils/proto"
)

func GetEvents(ctx *gin.Context) {
	var f event.SearchEventsForm
	if err := ctx.ShouldBindQuery(&f); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.S().Errorf("GetEvents err: %v", err)
		return
	}

	// get eventserver ip from etcd
	// dial and close

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to connect to EventServer: %v", err)
	}

	res, err := global.EventServerClient.SearchEvents(context.Background(), &pb.SearchEventsRequest{
		LeagueName: f.LeagueName,
		SportType:  f.SportType,
		StartDate:  f.StartDate,
		EndDate:    f.EndDate,
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
