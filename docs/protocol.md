# protocol

## 详解主要字段

```sh
# 设备端发送
REGISTER sip:34020000002000000001@3402000000 SIP/2.0
Via: SIP/2.0/UDP 192.168.0.64:5060;rport;branch=z9hG4bK1945388331
From:<sip:34020000001320000002@3402000000>;tag=33226951
To: <sip:34020000001320000002@3402000000>
Call-ID: 1455222403
CSeq: 1 REGISTER
Contact: <sip:34020000001320000002@192.168.0.64:5060>
Max-Forwards: 70
User-Agent: IP Camera
Expires: 3600
Content-Length: 0

# 服务端应答
SIP/2.0 200 OK
Via: SIP/2.0/UDP 192.168.0.64:5060;rport;branch=z9hG4bK1494593151
From: <sip:34020000001320000002@3402000000>;tag=33226951
To: <sip:34020000001320000002@3402000000>
Call-ID: 1455222403
CSeq: 2 REGISTER
User-Agent: DuHua
Date: 2025-04-22T10:48:38.234
Content-Length: 0
```

- 消息头: SIP消息的方法`Method`(`REGISTER`), 接收方URI(`sip:34020000002000000001@3402000000`), SIP协议版本.

- `Via`头: 包含了发送请求方的相关信息, 后续需要使用这些信息进行回复.
  - `SIP/2.0/UDP`: 表示使用的是2.0版本的SIP协议, 使用的传输协议是UDP, 也可以使用TCP协议, 大部分设备默认是UDP协议, GB2016以前只支持UDP协议.
  - `192.168.0.64:5060`: 为请求发送方的IP地址和端口号.
  - `branch`: 具体值是一个在整个SIP通信过程中不重复的数值. `branch`是一个事务ID(Transaction ID), 用于区分同一个UA所发起的不同Transaction, 它不会对未来的`request`或者是`response`造成影响, 对于遵循IETF RFC3261规范的实现, 这个`branch`参数的值必须用`z9hG4bK`字符串打头. 其它部分是对`To`, `From`, `Call-ID`头域和`Request-URI`按一定的算法加密后得到, 也可以是随机数或者UUID, 目前没发现有什么用途.
  - `rport`: 表示使用rport机制路由响应, 即发送的响应时, 按照rport中的端口发送SIP响应. 就是说IP和端口均完全遵照从哪里来的, 发回哪里去的原则. 如果没有rport字段时, 服务端的策略是IP使用UDP包中的地址, 即从哪里来回哪里去, 但是端口使用的是via中的端口, 详情见IETF RFC35818.

- `From`头, 包含了**请求发送方**的逻辑标识. 在GB28181协议中是发送请求的设备国标ID和域国标ID信息. `tag`参数是为了身份认证的, 值为随机数字字符.

- `To`头, 标明**请求接收方**的逻辑标识的. 在GB28181协议中填写的是发送请求的设备国标ID和域国标ID信息.

- `Call-ID`头: Call-ID头是全局唯一的, 在同一个session中保持一致, 在不同session中不同.

- `CSeq`头, CSeq头又叫Command Seqence(命令队列), 用于标识命令顺序. 值为序号+Method, 序号部分为无符号整数, 最大值为2^31. 序号起始值是随机的, 后续在同一个session中依次递增. 对于ACK和CANCEL中的CSeq与INVITE中的Cseq保持一致.

- `Contact`头, 包含**源**的URI信息, 用来给响应消息直接和源建立连接用. 在GB28181协议中为SIP设备编码@源IP地址端口.

- `Max-Forwards`头,用于设置包最大中转次数, 默认是70.
3
- `User-Agent`头, 用于设置关于UA的信息, 用户可以自定义.

- `Expires`头, 表示超时时间. `0`表示注销.

- `Content-Length`头, 表示消息体的长度, 因为`REGISTER`消息不需要消息体, 所以为0. 如果携带了`xml`或者`sdp`等消息体, 则>0.
