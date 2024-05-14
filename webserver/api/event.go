package api

import (
	"context"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"net/http"
	"sportshot/utils/db"
	"sportshot/utils/models/event"
	pb "sportshot/utils/proto"
)

// UnaryInterceptor for debug
func UnaryInterceptor(
	ctx context.Context,
	method string,
	req, reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	zap.S().Infof("Making request to: %s", cc.Target())
	return invoker(ctx, method, req, reply, cc, opts...)
}

func GetEvents(ctx *gin.Context) {

	var f event.SearchEventsForm
	if err := ctx.ShouldBindQuery(&f); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.S().Errorf("GetEvents err: %v", err)
		return
	}

	// etcd
	zap.S().Info("connecting to etcd")
	cli, err := db.GetEtcdClient()
	if err != nil {
		zap.S().Errorf("fail to GetEtcdClient : %v", err)
	}
	defer func(cli *clientv3.Client) {
		err = cli.Close()
		if err != nil {
			zap.S().Errorf("error closing etcd: %v", err)
		}
	}(cli)

	// register to grpc
	resolver.Register(db.NewEtcdResolver(cli))

	// connect to grpc
	zap.S().Info("connecting to grpc")
	conn, err := grpc.Dial("etcd:///event", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(UnaryInterceptor))

	if err != nil {
		zap.S().Errorf("fail to connect to EventServer: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err = conn.Close()
		if err != nil {
			zap.S().Errorf("fail to close grpc connect: %v", err)
		}
	}(conn)

	c := pb.NewEventServiceClient(conn)
	zap.S().Info("grpc 1")

	res, err := c.SearchEvents(context.Background(), &pb.SearchEventsRequest{
		LeagueName: f.LeagueName,
		SportType:  f.SportType,
		StartDate:  f.StartDate,
		EndDate:    f.EndDate,
	})
	zap.S().Info("grpc 2")

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		zap.S().Errorf("search events error: %v", err)
		return
	}
	zap.S().Info("grpc 3")

	ctx.JSON(200, gin.H{
		"data": res,
		"code": 200,
	})
	zap.S().Info("grpc 4")

}
