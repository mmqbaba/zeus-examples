## 配置入口
```json
	// 默认路径 ./conf/zeus.json
	{
		"engine_type": "etcd",
		"config_path": "/zeus/zeus-examples/sample", // 服务应用的配置路径
		"config_format": "json",     // 配置格式
		"endpoints": ["127.0.0.1:2379"],
		"username": "root",
		"password": "123456"
	}
```

## 应用服务配置
```json
	// 路径 /zeus/zeus-examples/sample
	{
		"redis": {
		"host": "127.0.0.1:6379",
			"pwd": ""
	},
		"go_micro": {
			"service_name": "sample",
			"registry_plugin_type": "etcd",
			"registry_addrs": ["127.0.0.1:2379"],
			"registry_authuser": "root",
			"registry_authpwd": "123456"
	}
	}
```

## gen-proto
```bash
	./build-proto.sh
```

## run
```bash
	go run ./cmd/app
```
