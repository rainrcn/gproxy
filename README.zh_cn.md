# gproxy

> GO语言实现的HTTP请求代理跳板服务

## 如何使用
### 修改config目录中的配置文件 
`app.yml`
```yaml

# 表示服务监听的http协议配置
http:
  enabled: true
  listen-addr: ":80"

## 表示服务监听的https协议配置
# https:
#   enabled: false
#   listen-addr: ":443"
#   cert-file:
#   key-file 

# http请求代理映射配置
maps:
  test.bd1: https://baidu.com
  tww.bd2: https://www.baidu.com
```
#### maps：
修改以上配置并启动。当请求地址： `http://localhost:80/test.bd1/s?ie=UTF-8&wd=golang` ，代理服务将会请求： `https://baidu.com/s?ie=UTF-8&wd=golang` 。 
请求主机与`test.bd1` 段部分将会被替换为 `https://baidu.com` 

#### 全局maps：
`app.yml`
```yaml

# ...和上面相同

# http请求代理映射配置
maps:
  /: https://www.baidu.com
```

当请求地址： `http://localhost:80` ，代理服务将会请求 `https://www.baidu.com`