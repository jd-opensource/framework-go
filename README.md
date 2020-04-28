## JD Chain Crypto in Go

### AES

`package`：`golang.org/x/crypto/aes`

`JD Chain`: `ECB PKCS7`

`Confirmed`: `true`



### Base58

`package`: [shengdoushi/base58](#https://github.com/shengdoushi/base58)

`JD Chain`: `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`

`Confirmed`: `true`



### SHA256

`package`：`golang.org/x/crypto/sha256`

`JD Chain`: `SHA128=SHA256[:16]`

`Confirmed`: `true`



### ED25519

`package`：`golang.org/x/crypto/ed25519`

`Confirmed`: `true`

> `JD Chain PrivKey` `32`位，`Go`中`64`位，低位为`PubKey`，签名和验签的时候需要注意



### RSA

`package`：`golang.org/x/crypto/rsa`

`Confirmed`: `true`

> `PKCS1`, `SHA256`



### ECDSA

`package`：[ThePiachu/Golang-Koblitz-elliptic-curve-DSA-library](#https://github.com/ThePiachu/Golang-Koblitz-elliptic-curve-DSA-library)

`Confirmed`: `true`

> `JD Chain` `PublicKey`的`getRawBytes`是`65`字节，较`Go`版本多了`0x04`



### RIPEMD160

`package`：`golang.org/x/crypto/ripemd160`

`Confirmed`: `true`

> out piut size 20



### SM2

`packahge`: [ZZMarquis/gm](#https://github.com/ZZMarquis/gm)/[tjfoc/gmsm](#https://github.com/tjfoc/gmsm)

`Confirmed`: `true`

> `JD Chain` `public key`会多一个`0x04`前缀
> `sign`/`verify`注意`uid`的问题




### SM3

`packahge`: [ZZMarquis/gm](#https://github.com/ZZMarquis/gm)/[tjfoc/gmsm](#https://github.com/tjfoc/gmsm)

`Confirmed`: `true`




### SM4

`packahge`: [ZZMarquis/gm](#https://github.com/ZZMarquis/gm)/[tjfoc/gmsm](#https://github.com/tjfoc/gmsm)

`Confirmed`: `true`

> `go`:`SM4/ECB/NoPadding`
> `JD Chain`:`SM4/CBC/PKCS7Padding`, [gmhelper](#https://github.com/ZZMarquis/gmhelper.git )

> [scloudrun/go-sm4](#https://github.com/scloudrun/go-sm4) 可作为解决方案：`JD Chain`加密结果会加上`16`字节的`iv`作为前缀，`Go`版本解密后去除前`16`个字节；`Go`版本的加密结果需要加上`iv`后`JD Chain`才能正确解密。