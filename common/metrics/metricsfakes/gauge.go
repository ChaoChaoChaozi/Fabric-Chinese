
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package metricsfakes

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
)

type Gauge struct {
	WithStub        func(labelValues ...string) metrics.Gauge
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		labelValues []string
	}
	withReturns struct {
		result1 metrics.Gauge
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Gauge
	}
	AddStub        func(delta float64)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		delta float64
	}
	SetStub        func(value float64)
	setMutex       sync.RWMutex
	setArgsForCall []struct {
		value float64
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Gauge) With(labelValues ...string) metrics.Gauge {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		labelValues []string
	}{labelValues})
	fake.recordInvocation("With", []interface{}{labelValues})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(labelValues...)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.withReturns.result1
}

func (fake *Gauge) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *Gauge) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return fake.withArgsForCall[i].labelValues
}

func (fake *Gauge) WithReturns(result1 metrics.Gauge) {
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *Gauge) WithReturnsOnCall(i int, result1 metrics.Gauge) {
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Gauge
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Gauge
	}{result1}
}

func (fake *Gauge) Add(delta float64) {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		delta float64
	}{delta})
	fake.recordInvocation("Add", []interface{}{delta})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		fake.AddStub(delta)
	}
}

func (fake *Gauge) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Gauge) AddArgsForCall(i int) float64 {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].delta
}

func (fake *Gauge) Set(value float64) {
	fake.setMutex.Lock()
	fake.setArgsForCall = append(fake.setArgsForCall, struct {
		value float64
	}{value})
	fake.recordInvocation("Set", []interface{}{value})
	fake.setMutex.Unlock()
	if fake.SetStub != nil {
		fake.SetStub(value)
	}
}

func (fake *Gauge) SetCallCount() int {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return len(fake.setArgsForCall)
}

func (fake *Gauge) SetArgsForCall(i int) float64 {
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	return fake.setArgsForCall[i].value
}

func (fake *Gauge) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.setMutex.RLock()
	defer fake.setMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Gauge) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.Gauge = new(Gauge)
