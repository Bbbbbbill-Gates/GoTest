
# 总结

## 微服务设计

前端 -> (CDN、4\7层负载均衡ELB) ->API Gateway Envoy + BFF(Backend For Frontend) -> MicroService:
- API Gateway Envoy 用于跨端的逻辑处理，降低这些逻辑与面向不同端的服务之间的耦合
- BFF用于统一的协议出口，向外暴露友好、统一的接口
- MicroService 划分通过Business Capability业务职能或 Bunded Context限界上下文

服务发现：
- 客户端服务发现：Server 注册服务，Client 获取到服务连接池在本地做一个负载均衡逻辑，相对于服务端的服务发现不容易产生流量热点，遵守了去中心化治理的理念。
- 服务端服务发现：通过服务端的服务发现做负载均衡，转发客户端的请求到下游 API，容易产生流量热点，但好处是客户端可以做的更加简单
- Service Mash：连接 Proxy，Proxy 做一个负载均衡逻辑，且 Proxy 通过本地 IPC 和服务 API 做通信

多集群多租户：
- 多集群带来更好的性能和冗余能力
- 多租户染色发布，可以实现大规模发布，在真实环境中引流测试流量

## 异常处理

```go

if (err != nil)

```

错误产生就立即处理，不会产生出现错误情况下的分支。

## 并发编程

go routine：GO 语言支持的轻量级线程（线程），也叫做协程，其本质就是用户态线程。

chan: Go语言提倡使用通信的方法代替共享内存，当一个资源需要在 goroutine 之间共享时，通道在 goroutine 之间架起了一个管道，并提供了确保同步交换数据的机制。声明通道时，需要指定将要被共享的数据的类型。可以通过通道共享内置类型、命名类型、结构类型和引用类型的值或者指针。

Package sync：保证 happens before 数据同步问题必要的一些关键字

contxt: 类似于 TLS（线程局部存储），我们为一些 goroutine 的 context 保存一些相关的信息方便我们对超时、错误的 fan out 做相应处理

## 工程事件

目录结构：
- cmd: 项目主干
- pkg: 外部可重用代码
- internal: 隐藏代码，不希望公开
- kit project: 基础库，统一、像标准库一样布局、高度抽象、支持插件
- service application project layout: api、config、test
  - 4+1 类 APP 服务类型：interface（对外BFF）、service（对内微服务）、job（流式任务处理）、admin（面向运营）、task（定时任务）

API 设计：
