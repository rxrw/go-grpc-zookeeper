package balancer

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/grpclog"
)

const LeastConnection = "least_connection_x"

// newLeastConnectionBuilder creates a new leastConnection balancer builder.
func newLeastConnectionBuilder() balancer.Builder {
	return base.NewBalancerBuilder(LeastConnection, &leastConnectionPickerBuilder{}, base.Config{HealthCheck: true})
}

func init() {
	balancer.Register(newLeastConnectionBuilder())
}

type leastConnectionPickerBuilder struct{}

func (*leastConnectionPickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	grpclog.Infof("leastConnectionPicker: newPicker called with readySCs: %v", info.ReadySCs)
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}
	var nodes []*Node

	for scs := range info.ReadySCs {
		nodes = append(nodes, &Node{scs, 0})
	}

	return &leastConnectionPicker{
		nodes: nodes,
		rand:  rand.New(rand.NewSource(time.Now().Unix())),
	}
}

type Node struct {
	balancer.SubConn
	inflight int64
}

type leastConnectionPicker struct {
	nodes []*Node
	mu    sync.Mutex
	rand  *rand.Rand
}

func (p *leastConnectionPicker) Pick(opts balancer.PickInfo) (balancer.PickResult, error) {
	if len(p.nodes) == 0 {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	var node *Node
	if len(p.nodes) == 1 {
		node = p.nodes[0]
	} else {
		p.mu.Lock()
		a := p.rand.Intn(len(p.nodes))
		b := p.rand.Intn(len(p.nodes))
		p.mu.Unlock()
		if a == b {
			b = (b + 1) % len(p.nodes)
		}
		if p.nodes[a].inflight < p.nodes[b].inflight {
			node = p.nodes[a]
		} else {
			node = p.nodes[b]
		}
	}

	atomic.AddInt64(&node.inflight, 1)

	var pickResult = balancer.PickResult{
		SubConn: node,
		Done: func(info balancer.DoneInfo) {
			atomic.AddInt64(&node.inflight, -1)
		},
	}

	return pickResult, nil
}
