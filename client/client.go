package main

import (
	m "dubbogo_nacos_configcenter/model_service"

	dubbo_config "dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
)

var grpcGreeterImpl = new(m.ServiceClientImpl)

func main() {
	dubbo_config.SetConsumerService(grpcGreeterImpl)
	rootConfig := dubbo_config.NewRootConfigBuilder().
		SetConfigCenter(dubbo_config.NewConfigCenterConfigBuilder().
			SetProtocol("nacos").
			SetAddress("mse-31a78442-nacos-ans.mse.aliyuncs.com:8848"). // 根据配置结构，设置配置中心
			SetDataID("client").                                        // 设置配置ID
			SetGroup("test").
			Build()).
		// SetLogger(LoggerConfig).
		Build()

	if err := dubbo_config.Load(dubbo_config.WithRootConfig(rootConfig)); err != nil { // 框架启动
		panic(err)
	}
	// res, err := grpcGreeterImpl.Get(context.Background(), &m.Request{Name: "ocpq_model_1_int64_float64"})
	// if err != nil {
	// 	log.Print(res)
	// }
}
