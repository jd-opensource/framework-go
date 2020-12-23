package random

import (
	"github.com/blockchain-jd-com/framework-go/utils/base58"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
 * Author: imuge
 * Date: 2020/12/1 下午6:01
 */

func TestHashSecureRandom_Read(t *testing.T) {
	hrr := NewHashSecureRandom([]byte("a"), 32, Sha256)
	p1 := make([]byte, 32)
	hrr.Read(p1)
	p2 := make([]byte, 32)
	hrr = NewHashSecureRandom([]byte("a"), 32, Sha256)
	hrr.Read(p2)
	assert.EqualValues(t, p1, p2)

	// generated from jd chain
	assert.Equal(t, "41PnrB7Q6qAivS5iQPiatYkbh7br7xQhg24atm5XLBiF", base58.Encode(p1))

}
