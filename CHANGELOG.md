# Changelog

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
