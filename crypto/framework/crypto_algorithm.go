package framework

import (
	"fmt"
	binary_proto "github.com/blockchain-jd-com/framework-go/binary-proto"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
)

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:13 下午
 */

var (
	// 随机数算法标识
	RANDOM_ALGORITHM = 0x1000

	// 哈希数算法标识
	HASH_ALGORITHM = 0x2000

	// 签名算法标识
	SIGNATURE_ALGORITHM = 0x4000

	// 加密算法标识
	ENCRYPTION_ALGORITHM = 0x8000

	/**
	 * 扩展密码算法标识；
	 * 表示除了
	 * {@link #RANDOM_ALGORITHM}、{@link #HASH_ALGORITHM}、{@link #SIGNATURE_ALGORITHM}、{@link #ENCRYPTION_ALGORITHM}
	 * 之外的其它非标准分类的密码算法，诸如加法同态算法、多方求和算法等；
	 */
	EXT_ALGORITHM = 0x0000

	// 非对称密钥标识
	ASYMMETRIC_KEY = 0x0100

	// 对称密钥标识
	SYMMETRIC_KEY = 0x0200

	// 算法编码的字节长度；等同于 {@link #getCodeBytes(CryptoAlgorithm)} 返回的字节数组的长度
	CODE_SIZE = 2
)

var _ binary_proto.DataContract = (*CryptoAlgorithm)(nil)

func init() {
	binary_proto.Cdc.RegisterContract(CryptoAlgorithm{})
}

type CryptoAlgorithm struct {
	Code int16  `primitiveType:"INT16"`
	Name string `primitiveType:"TEXT"`
}

func (ca CryptoAlgorithm) ContractCode() int32 {
	return binary_proto.CRYPTO_ALGORITHM
}

func (ca CryptoAlgorithm) ContractName() string {
	return "CryptoAlgorithm"
}

func (ca CryptoAlgorithm) Description() string {
	return ""
}

func (ca CryptoAlgorithm) ToString() string {
	return fmt.Sprintf("%s[%d]", ca.Name, ca.Code&-1)
}

func (ca CryptoAlgorithm) getCodeBytes() []byte {
	return bytes.Int16ToBytes(ca.Code)
}

func (ca CryptoAlgorithm) Match(algorithmBytes []byte, offset int) bool {
	return ca.Code == bytes.ToInt16(algorithmBytes[offset:2])
}

func (ca CryptoAlgorithm) IsRandomAlgorithm() bool {
	return RANDOM_ALGORITHM == (int(ca.Code) & RANDOM_ALGORITHM)
}

func (ca CryptoAlgorithm) IsHashAlgorithm() bool {
	return HASH_ALGORITHM == (int(ca.Code) & HASH_ALGORITHM)
}

func (ca CryptoAlgorithm) IsSignatureAlgorithm() bool {
	return SIGNATURE_ALGORITHM == (int(ca.Code) & SIGNATURE_ALGORITHM)
}

func (ca CryptoAlgorithm) IsEncryptionAlgorithm() bool {
	return ENCRYPTION_ALGORITHM == (int(ca.Code) & ENCRYPTION_ALGORITHM)
}

func (ca CryptoAlgorithm) IsExtAlgorithm() bool {
	return EXT_ALGORITHM == (int(ca.Code) & 0xF000)
}

func (ca CryptoAlgorithm) HasAsymmetricKey() bool {
	return ASYMMETRIC_KEY == (int(ca.Code) & ASYMMETRIC_KEY)
}

func (ca CryptoAlgorithm) HasSymmetricKey() bool {
	return SYMMETRIC_KEY == (int(ca.Code) & SYMMETRIC_KEY)
}

func (ca CryptoAlgorithm) IsSymmetricEncryptionAlgorithm() bool {
	return ca.IsEncryptionAlgorithm() && ca.HasSymmetricKey()
}

func (ca CryptoAlgorithm) IsAsymmetricEncryptionAlgorithm() bool {
	return ca.IsEncryptionAlgorithm() && ca.HasAsymmetricKey()
}

func (ca CryptoAlgorithm) Equals(algorithm CryptoAlgorithm) bool {
	return ca.Code == algorithm.Code
}

/**
 * 声明一项哈希算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefineHash(name string, uid byte) CryptoAlgorithm {
	code := int16(HASH_ALGORITHM | (int(uid) & 0x00FF))
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}

/**
 * 声明一项非对称密码算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefineSignature(name string, encryptable bool, uid byte) CryptoAlgorithm {
	var code int16
	if encryptable {
		code = int16(SIGNATURE_ALGORITHM | ENCRYPTION_ALGORITHM | ASYMMETRIC_KEY | (int(uid) & 0x00FF))
	} else {
		code = int16(SIGNATURE_ALGORITHM | ASYMMETRIC_KEY | (int(uid) & 0x00FF))
	}
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}

/**
 * 声明一项非对称加密算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefineAsymmetricEncryption(name string, uid byte) CryptoAlgorithm {
	code := int16(ENCRYPTION_ALGORITHM | ASYMMETRIC_KEY | (int(uid) & 0x00FF))
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}

/**
 * 声明一项对称密码算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefineSymmetricEncryption(name string, uid byte) CryptoAlgorithm {
	code := int16(ENCRYPTION_ALGORITHM | SYMMETRIC_KEY | (int(uid) & 0x00FF))
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}

/**
 * 声明一项随机数算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefineRandom(name string, uid byte) CryptoAlgorithm {
	code := int16(RANDOM_ALGORITHM | (int(uid) & 0x00FF))
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}

/**
 * 声明一项扩展的密码算法；
 *
 * @param name 算法名称；
 * @param uid  算法ID；需要在同类算法中保持唯一性；
 * @return
 */
func DefinExt(name string, uid byte) CryptoAlgorithm {
	code := int16(EXT_ALGORITHM | (int(uid) & 0x00FF))
	return CryptoAlgorithm{
		Code: code,
		Name: name,
	}
}
