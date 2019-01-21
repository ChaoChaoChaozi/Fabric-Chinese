
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//Code generated by mockery v1.0.0. 不要编辑。
package mocks

import mock "github.com/stretchr/testify/mock"
import txvalidator "github.com/hyperledger/fabric/core/committer/txvalidator"
import validation "github.com/hyperledger/fabric/core/handlers/validation/api"

//pluginmapper是为pluginmapper类型自动生成的模拟类型
type PluginMapper struct {
	mock.Mock
}

//PlugInfactoryByName提供了一个具有给定字段的模拟函数：name
func (_m *PluginMapper) PluginFactoryByName(name txvalidator.PluginName) validation.PluginFactory {
	ret := _m.Called(name)

	var r0 validation.PluginFactory
	if rf, ok := ret.Get(0).(func(txvalidator.PluginName) validation.PluginFactory); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.PluginFactory)
		}
	}

	return r0
}