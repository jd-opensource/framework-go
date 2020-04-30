package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:49 下午
 */

type AsymmetricEncryptionFunction interface {
	AsymmetricKeypairGenerator
	CryptoFunction

	/**
	 * 加密；
	 *
	 * @param data
	 * @return
	 */
	Encrypt(pubKey PubKey, data []byte) AsymmetricCiphertext

	/**
	 * 解密；
	 *
	 * @param privKey
	 * @param ciphertext
	 * @return
	 */
	Decrypt(privKey PrivKey, ciphertext AsymmetricCiphertext) []byte

	/**
	 * 使用私钥恢复公钥；
	 *
	 * @param privKey PrivKey形式的私钥信息
	 * @return PubKey形式的公钥信息
	 */
	RetrievePubKey(privKey PrivKey) PubKey

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
	 * @param privKeyBytes 包含算法标识和私钥的字节数组
	 * @return PrivKey形式的私钥
	 */
	ParsePrivKey(privKeyBytes []byte) PrivKey

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
	 * @param pubKeyBytes 包含算法标识和公钥的字节数组
	 * @return PubKey形式的公钥
	 */
	ParsePubKey(pubKeyBytes []byte) PubKey

	/**
	 * 校验密文格式是否满足要求；
	 *
	 * @param ciphertextBytes 包含算法标识和密文的字节数组
	 * @return 是否满足指定算法的密文格式
	 */
	SupportCiphertext(ciphertextBytes []byte) bool

	/**
	 * 将字节数组形式的密文转换成AsymmetricCiphertext格式；
	 *
	 * @param ciphertextBytes 包含算法标识和密文的字节数组
	 * @return AsymmetricCiphertext形式的签名摘要
	 */
	ParseCiphertext(ciphertextBytes []byte) AsymmetricCiphertext
}
