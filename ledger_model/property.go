package ledger_model

import (
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"github.com/magiconair/properties"
)

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

type Properties []Property

func LoadProperties(configFile string) Properties {
	ps, err := properties.LoadFile(configFile, properties.UTF8)
	if err != nil {
		panic(err)
	}
	m := ps.Map()
	properties := Properties{}
	for k, v := range m {
		properties = append(properties, Property{
			Name:  k,
			Value: v,
		})
	}
	return properties
}
