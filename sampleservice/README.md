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

## 应用服务配置
```json
// 路径/zeus/dzqz
{
	"redis": {
		"host": "127.0.0.1:6379",
		"pwd": "",
		"enable": false
	},
	"go_micro": {
		"server_name": "zeus",
		"registry_plugin_type": "etcd",
		"registry_addrs": ["127.0.0.1:2379"],
		"registry_authuser": "root",
		"registry_authpwd": "123456"
	}
}
```
```bash
# 设值
ETCDCTL_API=3 etcdctl put /zeus/dzqz {\"go_micro\":{\"server_name\":\"zeus\",\"registry_plugin_type\":\"etcd\",\"registry_addrs\":[\"127.0.0.1:2379\"],\"registry_authuser\":\"root\",\"registry_authpwd\":\"123456\"},\"mongodb\":{\"host\":\"127.0.0.1:27017\",\"user\":\"root\",\"pwd\":\"123456\"},\"mysql_source\":{\"e_seal\":{\"datasourcename\":\"root:123456@tcp(localhost:3306)/e_seal\",\"maxidleconns\":30,\"maxopenconns\":1024}},\"redis\":{\"host\":\"127.0.0.1:6379\",\"pwd\":\"\",\"enable\":false},\"log_conf\":{\"log\":\"console\",\"level\":\"debug\",\"format\":\"text\",\"rotation_time\":\"hour\",\"log_dir\":\"./\"},\"ext\":{\"httphandler_pathprefix\":\"\",\"grpcgateway_pathprefix\":\"\"}}
```

## gen-proto
```bash
./build-proto.sh
```

## run
```bash
go run ./cmd/app
```