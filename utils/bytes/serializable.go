package bytes

/**
 * @Author: imuge
 * @Date: 2020/4/28 5:39 下午
 */

type BytesSerializable interface {

	/**
	 * 以字节数组形式获取字节块的副本；
	 * @return byte[]
	 */
	ToBytes() []byte
}
