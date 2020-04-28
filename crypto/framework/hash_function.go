package framework

/**
 * @Author: imuge
 * @Date: 2020/4/28 2:37 下午
 */

type HashFunction interface {
	CryptoFunction

	/**
	 * 计算指定数据的 hash；
	 *
	 * @param data
	 * @return
	 */
	Hash(data []byte) HashDigest

	/**
	 * 校验 hash 摘要与指定的数据是否匹配；
	 *
	 * @param digest
	 * @param data
	 * @return
	 */
	Verify(digest HashDigest, data []byte) bool

	/**
	 * 校验字节数组形式的hash摘要的格式是否满足要求；
	 *
	 * @param digestBytes 包含算法标识和hash摘要的字节数组
	 * @return 是否满足指定算法的hash摘要格式
	 */
	SupportHashDigest(digestBytes []byte) bool

	/**
	 * 将字节数组形式的hash摘要转换成HashDigest格式；
	 *
	 * @param digestBytes 包含算法标识和hash摘要的字节数组
	 * @return HashDigest形式的hash摘要
	 */
	ResolveHashDigest(digestBytes []byte) HashDigest
}
