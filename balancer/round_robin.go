package balancer

import (
	"github.com/rxrw/go-grpc-zookeeper/common"

	"math/rand"
	"sync"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
)

const RoundRobin = "round_robin_x"

// newRoundRobinBuilder creates a new roundrobin balancer builder.
func newRoundRobinBuilder() balancer.Builder {
	return base.NewBalancerBuilder(RoundRobin, &roundRobinPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newRoundRobinBuilder())
}

type roundRobinPickerBuilder struct{}

func (*roundRobinPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	grpclog.Infof("roundrobinPicker: newPicker called with readySCs: %v", info.ReadySCs)

	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var scs []balancer.SubConn
	for addr, sc := range info.ReadySCs {
		weight := common.GetWeight(sc.Address)
		for i := 0; i < weight; i++ {
			scs = append(scs, addr)
		}
	}

	return &roundRobinPicker{
		subConns: scs,
		next:     rand.Intn(len(scs)),
	}
}

type roundRobinPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
	next     int
}

func (p *roundRobinPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	sc := p.subConns[p.next]
	p.next = (p.next + 1) % len(p.subConns)
	p.mu.Unlock()
	return balancer.PickResult{SubConn: sc}, nil
}
