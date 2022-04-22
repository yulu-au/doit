package resolve

import (
	"context"
	"grpc-client/etcd"
	"strings"
	"time"

	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

func NewBuilder() resolver.Builder {
	return &echoBuilder{etcdClient: etcd.EClient, store: make(map[string]map[string]struct{})}
}

type echoBuilder struct {
	store      map[string]map[string]struct{}
	etcdClient *clientv3.Client
}

func (e *echoBuilder) Scheme() string {
	return "etcd"
}

func (e *echoBuilder) Build(target resolver.Target,
	cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	e.store[target.Endpoint] = make(map[string]struct{})

	r := &echoResolver{
		client: e.etcdClient,
		target: target,
		cc:     cc,
		store:  e.store[target.Endpoint],
		stopCh: make(chan struct{}, 1),
		rn:     make(chan struct{}, 1),
		t:      time.NewTicker(time.Second * 2),
	}
	// 开启后台更新 goroutine
	go r.start(context.Background())
	// 全量更新服务地址
	r.ResolveNow(resolver.ResolveNowOptions{})

	return r, nil
}

type echoResolver struct {
	client *clientv3.Client
	target resolver.Target
	cc     resolver.ClientConn //这里是resolver与conn的桥梁
	store  map[string]struct{}
	stopCh chan struct{}
	// rn channel is used by ResolveNow() to force an immediate resolution of the target.
	rn chan struct{}
	t  *time.Ticker
}

func (r *echoResolver) start(ctx context.Context) {
	target := r.target.Endpoint

	w := clientv3.NewWatcher(r.client)
	rch := w.Watch(ctx, target+"/", clientv3.WithPrefix())
	for {
		select {
		case <-r.rn:
			r.resolveNow()
		// case <-r.t.C:
		// r.ResolveNow(resolver.ResolveNowOptions{})
		case <-r.stopCh:
			w.Close()
			return
		case wresp := <-rch:
			for _, ev := range wresp.Events {
				switch ev.Type {
				case mvccpb.PUT:
					r.store[string(ev.Kv.Value)] = struct{}{}
				case mvccpb.DELETE:
					delete(r.store, strings.Replace(string(ev.Kv.Key), target+"/", "", 1))
				}
			}
			r.updateTargetState()
		}
	}
}

func (r *echoResolver) resolveNow() {
	target := r.target.Endpoint
	resp, err := r.client.Get(context.Background(), target+"/", clientv3.WithPrefix())
	if err != nil {
		r.cc.ReportError(err)
		return
	}

	for _, kv := range resp.Kvs {
		r.store[string(kv.Value)] = struct{}{}
	}

	r.updateTargetState()
}

func (r *echoResolver) updateTargetState() {
	addrs := make([]resolver.Address, len(r.store))
	i := 0
	for k := range r.store {
		addrs[i] = resolver.Address{Addr: k}
		i++
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}

func (r *echoResolver) ResolveNow(o resolver.ResolveNowOptions) {
	select {
	case r.rn <- struct{}{}:
	default:

	}
}

func (r *echoResolver) Close() {
	r.t.Stop()
	close(r.stopCh)
}
