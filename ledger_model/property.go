package ledger_model

import "github.com/blockchain-jd-com/framework-go/utils/bytes"

type Property struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p *Property) ToBytes() []byte {
	nbs := bytes.StringToBytes(p.Name)
	bs := bytes.NUMBERMASK_NORMAL.WriteMask(int64(len(nbs)))
	bs = append(bs, nbs...)
	vbs := bytes.StringToBytes(p.Value)
	bs = append(bs, bytes.NUMBERMASK_NORMAL.WriteMask(int64(len(vbs)))...)
	return append(bs, vbs...)
}
