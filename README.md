# 快速开始
## 下载可执行文件
在 [Releases](https://github.com/rocketk/revproxy/releases) 中下载最新的可执行文件，请选择你所需要的平台。

以`revproxy-darwin-amd64`为例，将其改名为`revproxy`
```bash
mv revproxy-darwin-amd64 revproxy
```

## 编写配置文件
在`revproxy`同级目录下新建一个配置文件`revproxy.yaml`
```bash
touch revproxy.yaml
vim revproxy.yaml
```
将以下内容放入配置文件中
```yaml
target:
  base-url: http://example.com/
  prefix-to-trim: 
server:
  port: 9080
```
请注意将`target.base-url`修改为你的目标服务的域名  
`server.port`是`revproxy`的端口
## 启动反向代理程序
```bash
./revproxy
```
建议使用`nohup`后台运行，例：
```bash
nohup ./revproxy > revproxy-logs.log 2>1&
```