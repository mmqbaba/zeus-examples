## 配置入口
```json
// 默认路径/etc/tif/zeus.json
{
    "engine_type": "etcd",
    "config_path": "/zeus/dzqz", // 服务应用的配置路径
    "config_format": "json",     // 配置格式
	"endpoints": ["127.0.0.1:2379"],
	"username": "root",
	"password": "123456"
}
```

## gen-proto
```bash
./build-proto.sh
```

## run
```bash
go run ./cmd
```