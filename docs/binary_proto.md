## Binary proto

### 核心方法

1. [`Codec`](../binary-proto/codec.go)
序列化：
```go
func (c *Codec) Encode(obj interface{}) ([]byte, error) {
    ...
}
```

反序列化：
```go
func (c *Codec) Decode(data []byte, obj interface{}) {
    ...
}
```

### 实现

1. 协议文件

使用`JD Chain`相关接口获取`JD Chain`中定义的契约结构。

2. Go类型生成

根据契约结构，生成Go版本契约`struct`/`interface`。

3. 类型注册

实现[`DataContract`](../binary-proto/data_contract.go)接口注册契约类型

实现[`EnumContract`](../binary-proto/enum_contract.go)接口注册枚举契约

实现[`GenericContract`](../binary-proto/generic_contract.go)接口注册泛型接口和实现

4. 序列化

根据`Tag`定义区分固定长度类型字段和动态长度字段，具体编码规则参照`JD Chain`实现

5. 反序列化

参照序列化编码规则，利用`Codec`中注册的契约信息使用反射生成对应结构对象。