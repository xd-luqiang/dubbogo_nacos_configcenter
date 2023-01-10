package main

import (
	m "dubbogo_nacos_configcenter/model_service"

	dubbo_config "dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

type Provider struct {
	m.UnimplementedServiceServer
}

func main() {
	dubbo_config.SetProviderService(&Provider{})
	rootConfig := dubbo_config.NewRootConfigBuilder().
		SetConfigCenter(dubbo_config.NewConfigCenterConfigBuilder().
			SetProtocol("nacos").
			SetAddress("mse-31a78442-nacos-ans.mse.aliyuncs.com:8848"). // 根据配置结构，设置配置中心
			SetDataID("dubbo").                                         // 设置配置ID
			SetGroup("test").
			Build()).
		// SetLogger(LoggerConfig).
		Build()

	if err := rootConfig.Init(); err != nil { // 框架启动
		panic(err)
	}
	select {}
}
