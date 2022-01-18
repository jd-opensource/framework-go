package ledger_model

import (
	"fmt"
	"testing"
)

func TestLoadProperties(t *testing.T) {
	properties := LoadProperties("/home/jodad/JD/nodes/peer0/config/init/bftsmart.config")
	for _, v := range properties {
		fmt.Println(v.Name + " = " + v.Value)
	}
}
