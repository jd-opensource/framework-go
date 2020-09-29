package framework

import "github.com/blockchain-jd-com/framework-go/utils/bytes"

/**
 * @Author: imuge
 * @Date: 2020/4/29 9:06 上午
 */

// 算法标识符的长度
var ALGORYTHM_CODE_SIZE = CODE_SIZE

type CryptoBytes interface {
	bytes.BytesSerializable

	/**
	 * 算法；
	 *
	 * @return
	 */
	GetAlgorithm() int16
}
