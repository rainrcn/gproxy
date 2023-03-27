# gproxy

golang implementation of http request proxy jump service

## How to use 
Modify the configuration file in the config directory

`app.yml`

```yaml

# Indicates the http protocol configuration that the service listens on
http:
  enabled: true
  listen-addr: ":80"

## Indicates the https protocol configuration for service listening
# https:
#   enabled: false
#   listen-addr: ":443"
#   cert-file:
#   key-file

# http request proxy mapping configuration
maps:
  test.bd1: https://baidu.com
  tww.bd2: https://www.baidu.com
```

#### maps:
Modify the above configuration and start. When the request address: `http://localhost:80/test.bd1/s? ie=UTF-8&wd=golang` ', the proxy service will request: `https://baidu.com/s?ie=UTF-8&wd=golang`.
The request host and `test.bd1` section will be replaced with `https://baidu.com`

#### global maps:
`app.yml`

```yaml

#... Same as above

# http request proxy mapping configuration
maps:
  /: https://www.baidu.com
```

When the address: `http://localhost:80` is requested, the proxy service will request `https://www.baidu.com`