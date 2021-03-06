# hostnic-cni

[English](README.md)|中文

**hostnic-cni** 是一个 [Container Network Interface](https://github.com/containernetworking/cni) 插件。 本插件会直接调用 IaaS 的接口去创建网卡并且关联到容器的网络命名空间。当前支持的 IaaS：[QingCloud](http://qingcloud.com)，未来会支持更多。

### 使用说明

1. 从 [CNI release 页面](https://github.com/containernetworking/cni/releases)  下载 CNI 包，解压到 /opt/cni/bin 下。
2. 从 [release 页面](https://github.com/yunify/hostnic-cni/releases) 下载 hostnic 放置到 /opt/cni/bin/ 路径下。
3. 增加 cni 的配置

```bash

cat >/etc/cni/net.d/10-hostnic.conf <<EOF
{
    "cniVersion": "0.3.0",
    "name": "hostnic",
    "type": "hostnic",
    "provider": "qingcloud",
    "args": {
      "providerConfigFile":"/etc/qingcloud/client.yaml",
      "vxNets":["vxnet-xxxxx","vxnet-xxxx"]
    },    
    "ipam":{
          "routes":[{"dst":"kubernetes service cidr","gw":"hostip or 0.0.0.0"}]
  　},
    "isGateway": true
}
EOF

cat >/etc/cni/net.d/99-loopback.conf <<EOF
{
	"cniVersion": "0.2.0",
	"type": "loopback"
}
EOF
```
3. 增加 IaaS 的 sdk 配置文件

```bash
cat >/etc/qingcloud/client.yaml <<EOF
qy_access_key_id: "Your access key id"
qy_secret_access_key: "Your secret access key"
# your instance zone
zone: "pek3a"
EOF
```
### 配置说明
* **provider** IaaS 的提供方，当前只支持 qingcloud，未来会支持更多。
* **providerConfigFile** IaaS 提供方的配置文件
* **vxNets** nic 所在的私有网络，数组格式，支持多个，多个私有网络必须在同一个 vpc 下。
* **ipam** 给nic设置路由条目。（可选）

### kubernetes用户的说明
kubernetes管理集群时会给每一个服务分配一个集群ip地址。kube-proxy会负责服务负载均衡，由于默认的方案所有pod的网络请求通过主机的网卡转发后才会进入路由器。但给pod单独分配网卡后，流量不再经过kube-proxy，所有调用服务的请求会失败。这里我们提供设置路由的功能来解决这个问题。pod分配网卡后，用户可以指定将集群ip网段的请求转发给虚拟机，交由kube-proxy去处理。由于网关必须满足一定条件，用户可以指定网关为0.0.0.0，插件会自动寻找可用网管，如果都不满足条件，会自动分配一个。
