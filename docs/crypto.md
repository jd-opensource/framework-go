## Crypto

### 核心类型

![function](crypto-function.png)

`CryptoFunction`:
```go
type CryptoFunction interface {
	GetAlgorithm() CryptoAlgorithm
}
```

`SymmetricKeyGenerator`， 对称秘钥生成接口:
```go
type SymmetricKeyGenerator interface {
	GenerateSymmetricKey() SymmetricKey
}
```

`AsymmetricKeypairGenerator`，非对称公私钥生成接口:
```go
type AsymmetricKeypairGenerator interface {
	GenerateKeypair() AsymmetricKeypair
}
```

`SymmetricEncryptionFunction`，对称加密算法实现:
```go
type SymmetricEncryptionFunction interface {
	SymmetricKeyGenerator
	CryptoFunction

	// 加密
	Encrypt(key SymmetricKey, data []byte) Ciphertext

	// 解密
	Decrypt(key SymmetricKey, ciphertext Ciphertext) []byte

	// 校验对称密钥格式是否满足要求
	SupportSymmetricKey(symmetricKeyBytes []byte) bool

	// 将字节数组形式的密钥转换成SymmetricKey格式
	ResolveSymmetricKey(symmetricKeyBytes []byte) SymmetricKey

	// 校验密文格式是否满足要求
	SupportCiphertext(ciphertextBytes []byte) bool

	// 将字节数组形式的密文转换成SymmetricCiphertext格式
	ParseCiphertext(ciphertextBytes []byte) SymmetricCiphertext
}
```

`HashFunction`，哈希算法实现:
```go
type HashFunction interface {
	CryptoFunction

	// 计算指定数据的 hash
	Hash(data []byte) HashDigest

	// 校验 hash 摘要与指定的数据是否匹配
	Verify(digest HashDigest, data []byte) bool

	// 校验字节数组形式的hash摘要的格式是否满足要求
	SupportHashDigest(digestBytes []byte) bool

	// 将字节数组形式的hash摘要转换成HashDigest格式
	ParseHashDigest(digestBytes []byte) HashDigest
}
```

`RandomFunction`，随机算法实现:
```go
type RandomFunction interface {
	CryptoFunction

	Generate(seed []byte) RandomGenerator
}
```

`SignatureFunction`，签名算法实现:
```go
type SignatureFunction interface {
	AsymmetricKeypairGenerator
	CryptoFunction

	// 计算指定数据的 hash
	Sign(privKey PrivKey, data []byte) SignatureDigest

	// 校验签名摘要和数据是否一致
	Verify(digest SignatureDigest, pubKey PubKey, data byte) bool

	// 使用私钥恢复公钥
	RetrievePubKey(privKey PrivKey) PubKey

	// 校验私钥格式是否满足要求
	SupportPrivKey(privKeyBytes []byte) bool

	// 将字节数组形式的私钥转换成PrivKey格式
	ParsePrivKey(privKeyBytes []byte) PrivKey

	// 校验公钥格式是否满足要求
	SupportPubKey(pubKeyBytes []byte) bool

	// 将字节数组形式的密钥转换成PubKey格式
	ParsePubKey(pubKeyBytes []byte) PubKey

	// 校验字节数组形式的签名摘要的格式是否满足要求
	SupportDigest(digestBytes []byte)

	// 将字节数组形式的签名摘要转换成SignatureDigest格式
	ParseDigest(digestBytes []byte) SignatureDigest
}
```

`AsymmetricEncryptionFunction`，非对称加密算法实现:
```go
type AsymmetricEncryptionFunction interface {
	AsymmetricKeypairGenerator
	CryptoFunction
}
```

`CryptoService`，密码算法服务:
```go
type CryptoService interface {
	GetFunctions() []CryptoFunction
}
```

`ClassicCryptoService`，传统密码算法:
```go
var (
	AES       = NewAESEncryptionFunction()
	ED25519   = NewED25519SignatureFunction()
	RIPEMD160 = NewRIPEMD160HashFunction()
	SHA256    = NewSHA256HashFunction()
	GO_RANDOM = NewGoRandomFunction()
	ECDSA     = NewECDSASignatureFunction()
	RSA       = NewRSACryptoFunction()
)

var _ framework.CryptoService = (*ClassicCryptoService)(nil)

type ClassicCryptoService struct {
	functions []framework.CryptoFunction
}

func NewClassicCryptoService() ClassicCryptoService {
	return ClassicCryptoService{
		[]framework.CryptoFunction{AES, ED25519, RIPEMD160, SHA256, GO_RANDOM, ECDSA, RSA},
	}
}

func (c ClassicCryptoService) GetFunctions() []framework.CryptoFunction {
	return c.functions
}
```

`SMCryptoService`，`SM`相关算法:
```go
var (
	SM2 = NewSM2CryptoFunction()
	SM3 = NewSM3HashFunction()
	SM4 = NewSM4EncryptionFunction()
)

var _ framework.CryptoService = (*SMCryptoService)(nil)

type SMCryptoService struct {
	functions []framework.CryptoFunction
}

func NewClassicCryptoService() SMCryptoService {
	return SMCryptoService{
		[]framework.CryptoFunction{SM2, SM3, SM4},
	}
}

func (c SMCryptoService) GetFunctions() []framework.CryptoFunction {
	return c.functions
}
```

`Crypto`:
```go
// 获取算法定义
func GetAlgorithmByCode(code int16) framework.CryptoAlgorithm {
	algorithm, exists := algorithms[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return algorithm
}

func GetAlgorithmByName(name string) framework.CryptoAlgorithm {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetAlgorithmByCode(code)
}

// 随机算法实现
func GetRandomFunctionByCode(code int16) framework.RandomFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.RandomFunction)
}

func GetRandomFunctionByName(name string) framework.RandomFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetRandomFunctionByCode(code)
}

func GetRandomFunction(algorithm framework.CryptoAlgorithm) framework.RandomFunction {
	return GetRandomFunctionByCode(algorithm.Code)
}

// Hash算法实现
func GetHashFunctionByCode(code int16) framework.HashFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.HashFunction)
}

func GetHashFunctionByName(name string) framework.HashFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetHashFunctionByCode(code)
}

func GetHashFunction(algorithm framework.CryptoAlgorithm) framework.HashFunction {
	return GetHashFunctionByCode(algorithm.Code)
}

// 非对称加密算法实现
func GetAsymmetricEncryptionFunctionByCode(code int16) framework.AsymmetricEncryptionFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.AsymmetricEncryptionFunction)
}

func GetAsymmetricEncryptionFunctionByName(name string) framework.AsymmetricEncryptionFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetAsymmetricEncryptionFunctionByCode(code)
}

func GetAsymmetricEncryptionFunction(algorithm framework.CryptoAlgorithm) framework.AsymmetricEncryptionFunction {
	return GetAsymmetricEncryptionFunctionByCode(algorithm.Code)
}

// 签名算法实现
func GetSignatureFunctionByCode(code int16) framework.SignatureFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.SignatureFunction)
}

func GetSignatureFunctionByName(name string) framework.SignatureFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetSignatureFunctionByCode(code)
}

func GetSignatureFunction(algorithm framework.CryptoAlgorithm) framework.SignatureFunction {
	return GetSignatureFunctionByCode(algorithm.Code)
}

// 对称加密算法
func GetSymmetricEncryptionFunctionByCode(code int16) framework.SymmetricEncryptionFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.SymmetricEncryptionFunction)
}

func GetSymmetricEncryptionFunctionByName(name string) framework.SymmetricEncryptionFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetSymmetricEncryptionFunctionByCode(code)
}

func GetSymmetricEncryptionFunction(algorithm framework.CryptoAlgorithm) framework.SymmetricEncryptionFunction {
	return GetSymmetricEncryptionFunctionByCode(algorithm.Code)
}

// 密码算法实现
func GetCryptoFunctionByCode(code int16) framework.CryptoFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function
}

func GetCryptoFunctionByName(name string) framework.CryptoFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetCryptoFunctionByCode(code)
}

func GetCryptoFunction(algorithm framework.CryptoAlgorithm) framework.CryptoFunction {
	return GetCryptoFunctionByCode(algorithm.Code)
}

```
### 示例

#### RandomFunction

```go
// 获取随机算法实现
rf := framework.GetRandomFunctionByName("GO_RANDOM")
// 生成10个随机字符
bytes := rf.NextBytes(10)
```

#### HashFunction
```go
// 获取哈希算法实现
hf := framework.GetHashFuction("SHA256")
// 计算哈希
hashBytes := []byte("bytes to hash")
hash := hf.hash(hashBytes)
// 校验哈希
ok := hf.verify(hash, hashBytes)
```

#### AsymmetricEncryptionFunction

```go
// 获取非对称加密算法实现
aef := framework.GetAsymmetricEncryptionFunction("RSA")
// 生成公私钥对
keypair := aef.GenerateKeypair()
// 加密
encryptBytes := []byte("bytes to encrypt")
ciphertext := aef.Encrypt(keypair.PubKey, encryptBytes)
// 解密
decryptBytes :=aef.Decrypt(keypair.PrivKey, ciphertext)
```

#### SymmetricEncryptionFunction
```go
// 获取对称加密算法实现
sef := framework.GetSymmetricEncryptionFunction("AES")
// 生成秘钥
key := sef.GenerateSymmetricKey()
// 加密
encryptBytes := []byte("bytes to encrypt")
ciphertext := sef.Encrypt(key, encryptBytes)
// 解密
decryptBytes :=sef.Decrypt(key, ciphertext)
```

#### SignatureFunction
```go
// 获取哈希算法实现
sf := framework.GetSignatureFunction("ED25519")
// 生成公私钥对
keypair := sf.GenerateKeypair()
// 计算签名
signBytes := []byte("bytes to sign")
sign := sf.Sign(keypair.PrivKey, hashBytes)
// 校验签名
ok := hf.verify(sign, keypair.PubKey, signBytes)
```

### KeyGen

生成/解析与JD Chain Java版本保存格式一致的的公私钥工具类

例如，使用`ed25519`生成/解析与`JD Chain` `Java`版本互通的公私钥对：
```go
function := crypto.GetCryptoFunctionByName(classic.ED25519_ALGORITHM.Name)
f1 := function.(framework.AsymmetricKeypairGenerator)
keypair := f1.GenerateKeypair()

// Encode
base58PubKey := crypto.EncodePubKey(keypair.PubKey)
pwd := []byte("abc")
base58PrivKey := crypto.EncodePrivKeyWithRawPwd(keypair.PrivKey, pwd)
base58PrivKey := crypto.EncodePrivKey(keypair.PrivKey, sha.Sha256(pwd))

// Decode
decPrivKey := crypto.DecodePrivKeyWithRawPwd(base58PrivKey, pwd)
decPrivKey := crypto.DecodePrivKey(base58PrivKey, sha.Sha256(pwd))
```