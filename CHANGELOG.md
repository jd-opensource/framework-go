# Changelog

## v1.3.4
2022.12.14

**Bug Fixes**
* 修复合约语言解析异常
* 合约调用空字符参数处理
* 修复国密TLS异常

## v1.3.3
2022.5.11

**FEATURES**
* 支持JD Chain 1.6.4，支持国密非国密连接

## v1.3.2
2021.12.08

**FEATURES**
* 提供非panic方式创建网关连接

## v1.3.1
2021.11.10

**FEATURES**
* 创建网关连接可以不传Keypair参数
* 根据seed生成公私钥对

## v1.3.0
2021.11.01

适配`JD Chain` [1.6.0](https://github.com/blockchain-jd-com/jdchain/releases/tag/1.6.0)

**FEATURES**
* 添加[证书](https://github.com/blockchain-jd-com/jdchain/wiki/CA)支持基础工具类
* 增加用户证书和生命周期相关操作
* 增加合约生命周期相关操作
* 适配`1.6.0`最新查询接口

## v1.2.0
2021.09.03
**FEATURES**
* 去除冗余`value`字段
**Bug Fixes**
* 解决空字节数组序列化与`java`版本不一致

## v1.1.2
2021.07.02
**Bug Fixes**
* [JD Chain 1.4.0以上版本激活参与方缺少shutdown参数](https://github.com/blockchain-jd-com/framework-go/issues/9)
* [交易详情缺少对合约调用操作的解析](https://github.com/blockchain-jd-com/framework-go/issues/10)
* TransactionState#GetValue默认值

## v1.1.1
2021.03.22
**Bug Fixes**
* [NewTransaction传入自解析账本异常](https://github.com/blockchain-jd-com/framework-go/issues/5)
* JD Chain 1.4.0 公钥兼容

## v1.1.0
2021.01.08
对应`JD Chain` `v1.4.0` 

**ENHANCEMENTS**
* 序列化框架并优化

**FEATURES**
* 根据`seed`生成公私钥对，目前Go版本仅`ED25519`支持

## v1.0.1
2021.01.08

**Bug Fixes**
* [频繁发送交易失败](https://github.com/blockchain-jd-com/framework-go/issues/3)

## v1.0.0
2020.11.27

**FEATURES**
* [crypto] 提供与`JD Chain`完全互通的密码算法相关实现
* [binary proto] 与`JD Chain`完全互通的自研序列化/反序列化框架
* [sdk] 提供与`JD Chain` `Java`版本`SDK`使用方式基本一致的实现，支持交易发送以及区块链相关数据查询
