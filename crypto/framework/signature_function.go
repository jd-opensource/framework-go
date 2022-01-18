package framework

import "github.com/blockchain-jd-com/framework-go/crypto/ca"

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:53 下午
 */

type SignatureFunction interface {
	AsymmetricKeypairGenerator
	CryptoFunction

	/**
	 * 计算指定数据的 hash；
	 *
	 * @param data 被签名消息
	 * @return SignatureDigest形式的签名摘要
	 */
	Sign(privKey *PrivKey, data []byte) (*SignatureDigest, error)

	/**
	 * 校验签名摘要和数据是否一致；
	 *
	 * @param digest 待验证的签名摘要
	 * @param data 被签名信息
	 * @return 是否验证通过
	 */
	Verify(pubKey *PubKey, data []byte, digest *SignatureDigest) bool

	/**
	 * 使用私钥恢复公钥；
	 *
	 * @param privKey PrivKey形式的私钥信息
	 * @return PubKey形式的公钥信息
	 */
	RetrievePubKey(privKey *PrivKey) (*PubKey, error)

	/**
	 * 校验私钥格式是否满足要求；
	 *
	 * @param privKeyBytes 包含算法标识、密钥掩码和私钥的字节数组
	 * @return 是否满足指定算法的私钥格式
	 */
	SupportPrivKey(privKeyBytes []byte) bool

	/**
	 * 将字节数组形式的私钥转换成PrivKey格式；
	 *
	 * @param privKeyBytes 包含算法标识、密钥掩码和私钥的字节数组
	 * @return PrivKey形式的私钥
	 */
	ParsePrivKey(privKeyBytes []byte) (*PrivKey, error)

	/**
	 * 校验公钥格式是否满足要求；
	 *
	 * @param pubKeyBytes 包含算法标识、密钥掩码和公钥的字节数组
	 * @return 是否满足指定算法的公钥格式
	 */
	SupportPubKey(pubKeyBytes []byte) bool

	/**
	 * 将字节数组形式的密钥转换成PubKey格式；
	 *
	 * @param pubKeyBytes 包含算法标识、密钥掩码和公钥的字节数组
	 * @return PubKey形式的公钥
	 */
	ParsePubKey(pubKeyBytes []byte) (*PubKey, error)

	/**
	 * 校验字节数组形式的签名摘要的格式是否满足要求；
	 *
	 * @param digestBytes 包含算法标识和签名摘要的字节数组
	 * @return 是否满足指定算法的签名摘要格式
	 */

	SupportDigest(digestBytes []byte) bool

	/**
	 * 将字节数组形式的签名摘要转换成SignatureDigest格式；
	 *
	 * @param digestBytes 包含算法标识和签名摘要的字节数组
	 * @return SignatureDigest形式的签名摘要
	 */
	ParseDigest(digestBytes []byte) (*SignatureDigest, error)

	// 从证书解析公钥
	RetrievePubKeyFromCA(cert *ca.Certificate) (*PubKey, error)

	// 从私钥文件解析私钥
	RetrievePrivKey(privateKey string) (*PrivKey, error)

	// 从私钥文件解析加密私钥
	RetrieveEncrypedPrivKey(privateKey string, pwd []byte) (*PrivKey, error)
}
