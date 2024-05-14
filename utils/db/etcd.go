package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

type etcdResolver struct {
	client *clientv3.Client    // etcd客戶端
	target resolver.Target     // 解析目標
	cc     resolver.ClientConn // 解析器客戶端連接
	ctx    context.Context     // 上下文
	cancel context.CancelFunc  // 取消函數
}

func (r *etcdResolver) start() {
	go r.watch() // 啟動監視
}

// 監視服務地址變化
func (r *etcdResolver) watch() {
	// 啟動監視器，監視目標的前綴
	watchChan := r.client.Watch(r.ctx, r.target.Endpoint(), clientv3.WithPrefix())
	for {
		select {
		case <-r.ctx.Done():
			return
		case resp := <-watchChan:
			if resp.Err() != nil {
				log.Printf("watch error: %v", resp.Err())
				continue
			}
			// 處理變更事件
			var addrs []resolver.Address
			for _, ev := range resp.Events {
				if ev.Type == clientv3.EventTypePut {
					addrs = append(addrs, resolver.Address{Addr: string(ev.Kv.Value)})
				}
			}
			r.cc.UpdateState(resolver.State{Addresses: addrs})
		}
	}
}

func (r *etcdResolver) ResolveNow(options resolver.ResolveNowOptions) {
	// 立即解析（此處留空，因為我們使用watch動態更新）
}

func (r *etcdResolver) Close() {
	r.cancel() // 取消監視
}

type etcdResolverBuilder struct {
	client *clientv3.Client // etcd客戶端
}

func (b *etcdResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())
	r := &etcdResolver{
		client: b.client,
		target: target,
		cc:     cc,
		ctx:    ctx,
		cancel: cancel,
	}
	r.start() // 啟動解析器並開始監視
	return r, nil
}

func (b *etcdResolverBuilder) Scheme() string {
	return "etcd"
}

func NewEtcdResolver(client *clientv3.Client) resolver.Builder {
	return &etcdResolverBuilder{client: client}
}

func GetEtcdClient() (*clientv3.Client, error) {
	// single etcd
	endpoints := []string{fmt.Sprintf("%s:%s", viper.GetString("etcd.host"), viper.GetString("etcd.port"))}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		zap.S().Errorf("fail connect to etcd error: %v, endpoints %v", err, endpoints)
		return nil, err
	}
	return cli, nil
}

func RegisterToEtcd(etcdClient *clientv3.Client, serviceKey string, serviceValue string) error {
	// grant
	lease, err := etcdClient.Grant(context.Background(), 5)
	if err != nil {
		log.Fatalf("failed to create lease: %v", err) // 处理错误
	}

	// register to etcd
	_, err = etcdClient.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(lease.ID))
	if err != nil {
		zap.S().Errorf("fail to register service: %v", err)
		return err
	}

	// keep
	ch, err := etcdClient.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		zap.S().Errorf("etcd keep alive failed: %v", err)
		return err
	}

	// goroutine the keep
	go func() {
		for {
			<-ch
		}
	}()
	return nil
}
