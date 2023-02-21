只有一个请求，可以打印出所在机器的ip
```
# curl 127.0.0.1:8081/index
{
    "Msg": "hello~",
    "IP": "172.17.0.5"
}
```

dockerfile已准备好
镜像可以在https://hub.docker.com/r/betaberry/go-helloworld 访问

