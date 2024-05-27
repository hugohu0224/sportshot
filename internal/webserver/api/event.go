package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.etcd.io/etcd/client/v3/naming/resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"sportshot/pkg/utils/global"
	"sportshot/pkg/utils/models/event"
	"sportshot/pkg/utils/proto"
)

// UnaryInterceptor checking the target
func UnaryInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	zap.S().Infof("making request to: %s", cc.Target())
	return invoker(ctx, method, req, reply, cc, opts...)
}

func GetEvents(ctx *gin.Context) {
	var f event.SearchEventsForm
	// empty query params
	if len(ctx.Request.URL.Query()) == 0 {
		ctx.JSON(200, gin.H{
			"message": "Empty query params.",
			"code":    200,
		})
		return
	}

	// validate params
	if err := ctx.ShouldBindQuery(&f); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.S().Errorf("GetEvents err: %v", err)
		return
	}

	// resolver
	etcdResolver, err := resolver.NewBuilder(global.EtcdClient)

	// dial
	conn, err := grpc.Dial("etcd:///event",
		grpc.WithResolvers(etcdResolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(UnaryInterceptor),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)

	// start to search
	c := proto.NewEventServiceClient(conn)
	res, err := c.SearchEvents(context.Background(), &proto.SearchEventsRequest{
		LeagueName: f.LeagueName,
		HomeName:   f.HomeName,
		AwayName:   f.AwayName,
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

	// response
	ctx.JSON(200, gin.H{
		"data": res,
		"code": 200,
	})
}
