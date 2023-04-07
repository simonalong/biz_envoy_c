### biz_config_b


记得将Dockerfile文件内容修改为对应的 Dockerfile-none-envoy 和 Dockerfile-have-envoy
```shell
# 有envoy
docker build -t biz-c-service:1.0.0-have-envoy .
# 有envoy【本地调试版本】
docker build -t biz-c-service:1.0.0-have-envoy-debug .
# 有envoy【链路数据】
docker build -t biz-c-service:1.0.0-have-envoy-trace .
# 没有envoy
docker build -t biz-c-service:1.0.0-none-envoy .
```

```shell
# 有envoy
docker run -d --name biz-c-service -p 18002:18002 biz-c-service:1.0.0-have-envoy
# 没有envoy
docker run -d --name biz-c-service -p 18002:18002 biz-c-service:1.0.0-none-envoy
```

```shell
# 测试
curl http://localhost:18012/api/c/front/cf/true/false
```

```shell
# 打镜像
docker build -t biz-c-service:1.0.0-none-envoy .
# 镜像保存
docker save biz-c-service:1.0.0-none-envoy -o biz-c-service.tar
# 镜像文件上传
scp -v biz-c-service.tar root@10.30.30.78:/root/zhouzy/biz-c-service.tar
# 开发环境中载入
docker load -i /root/zhouzy/biz-c-service.tar
```
