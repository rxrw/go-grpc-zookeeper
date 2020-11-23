package balancer

import (
	"fmt"
	"grpc-worker/common"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
)

const ConsistentHash = "consistent_hash_x"

var DefaultConsistentHashKey = "consistent-hash"

func InitConsistentHashBuilder(consistanceHashKey string) {
	balancer.Register(newConsistentHashBuilder(consistanceHashKey))
}

// newConsistanceHashBuilder creates a new ConsistanceHash balancer builder.
func newConsistentHashBuilder(consistentHashKey string) balancer.Builder {
	return base.NewBalancerBuilder(ConsistentHash, &consistentHashPickerBuilder{consistentHashKey}, base.Config{HealthCheck: true})
}

type consistentHashPickerBuilder struct {
	consistentHashKey string
}

func (b *consistentHashPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	grpclog.Infof("consistentHashPicker: newPicker called with readySCs: %v", info.ReadySCs)
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	picker := &consistentHashPicker{
		subConns:          make(map[string]balancer.SubConn),
		hash:              NewKetama(10, nil),
		consistentHashKey: b.consistentHashKey,
	}

	for sc, addr := range info.ReadySCs {
		weight := common.GetWeight(addr.Address)
		for i := 0; i < weight; i++ {
			node := wrapAddr(addr.Address.Addr, i)
			picker.hash.Add(node)
			picker.subConns[node] = sc
		}
	}
	return picker
}

type consistentHashPicker struct {
	subConns          map[string]balancer.SubConn
	hash              *Ketama
	consistentHashKey string
}

func (p *consistentHashPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	var sc balancer.SubConn
	key := p.consistentHashKey

	targetAddr, _ := p.hash.Get(key)

	sc = p.subConns[targetAddr]
	return balancer.PickResult{
		SubConn: sc,
	}, nil
}

func wrapAddr(addr string, idx int) string {
	return fmt.Sprintf("%s-%d", addr, idx)
}
