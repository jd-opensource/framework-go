package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:18 下午
 */

type SymmetricEncryptionFunction interface {
	SymmetricKeyGenerator
	CryptoFunction

	/**
	 * 加密；
	 *
	 * @param key 密钥；
	 * @param data 明文；
	 * @return
	 */
	Encrypt(key SymmetricKey, data []byte) SymmetricCiphertext

	/**
	 * 解密；
	 *
	 * @param key 密钥；
	 * @param ciphertext 密文；
	 * @return
	 */
	Decrypt(key SymmetricKey, ciphertext SymmetricCiphertext) []byte

	/**
	 * 校验对称密钥格式是否满足要求；
	 *
	 * @param symmetricKeyBytes 包含算法标识、密钥掩码和对称密钥的字节数组
	 * @return 是否满足指定算法的对称密钥格式
	 */
	SupportSymmetricKey(symmetricKeyBytes []byte) bool

	/**
	 * 将字节数组形式的密钥转换成SymmetricKey格式；
	 *
	 * @param symmetricKeyBytes 包含算法标识、密钥掩码和对称密钥的字节数组
	 * @return SymmetricKey形式的对称密钥
	 */
	ParseSymmetricKey(symmetricKeyBytes []byte) SymmetricKey

	/**
	 * 校验密文格式是否满足要求；
	 *
	 * @param ciphertextBytes 包含算法标识和密文的字节数组
	 * @return 是否满足指定算法的密文格式
	 */
	SupportCiphertext(ciphertextBytes []byte) bool

	/**
	 * 将字节数组形式的密文转换成SymmetricCiphertext格式；
	 *
	 * @param ciphertextBytes 包含算法标识和密文的字节数组
	 * @return SymmetricCiphertext形式的签名摘要
	 */

	ParseCiphertext(ciphertextBytes []byte) SymmetricCiphertext
}
