
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import cmd "github.com/hyperledger/fabric/discovery/cmd"
import common "github.com/hyperledger/fabric/cmd/common"
import discovery "github.com/hyperledger/fabric/discovery/client"
import mock "github.com/stretchr/testify/mock"

//存根是为存根类型自动生成的模拟类型
type Stub struct {
	mock.Mock
}

//发送提供给定字段的模拟函数：服务器、CONF、Req
func (_m *Stub) Send(server string, conf common.Config, req *discovery.Request) (cmd.ServiceResponse, error) {
	ret := _m.Called(server, conf, req)

	var r0 cmd.ServiceResponse
	if rf, ok := ret.Get(0).(func(string, common.Config, *discovery.Request) cmd.ServiceResponse); ok {
		r0 = rf(server, conf, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cmd.ServiceResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, common.Config, *discovery.Request) error); ok {
		r1 = rf(server, conf, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
