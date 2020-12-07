package random

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/blockchain-jd-com/framework-go/utils/sha"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/12/1 下午6:01
 */

func TestHashSecureRandom_Read(t *testing.T) {
	hrr := NewHashSecureRandom([]byte("a"), sha.Sha256)
	p := make([]byte, 32)
	hrr.Read(p)
	fmt.Println(base58.Encode(p))
	p = make([]byte, 32)
	hrr.Read(p)
	fmt.Println(base58.Encode(p))

}
