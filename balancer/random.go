package balancer

import (
	"grpc-worker/common"
	"math/rand"
	"sync"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
)

const Random = "random_x"

// newRandomBuilder creates a new random balancer builder.
func newRandomBuilder() balancer.Builder {
	return base.NewBalancerBuilder(Random, &randomPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newRandomBuilder())
}

type randomPickerBuilder struct{}

func (*randomPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	grpclog.Infof("randomPicker: newPicker called with readySCs: %v", info.ReadySCs)
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
	return &randomPicker{
		subConns: scs,
		rand:     rand.New(rand.NewSource(time.Now().Unix())),
	}
}

type randomPicker struct {
	subConns []balancer.SubConn
	mu       sync.Mutex
	rand     *rand.Rand
}

func (p *randomPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	p.mu.Lock()
	sc := p.subConns[p.rand.Intn(len(p.subConns))]
	p.mu.Unlock()

	return balancer.PickResult{SubConn: sc}, nil
}
