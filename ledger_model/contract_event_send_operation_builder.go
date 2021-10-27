package ledger_model

import (
	"errors"
	"github.com/blockchain-jd-com/framework-go/utils/bytes"
	"reflect"
)

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:25
 */

type ContractEventSendOperationBuilder struct {
	address []byte
	factory *BlockchainOperationFactory
}

func NewContractEventSendOperationBuilder(address []byte, factory *BlockchainOperationFactory) *ContractEventSendOperationBuilder {
	return &ContractEventSendOperationBuilder{address: address, factory: factory}
}

/*
	address 合约地址
	event   合约版本
	event   合约方法
	args	参数列表， 仅支持 bool/int16/int32/int64/string/[]byte
*/
func (cesob *ContractEventSendOperationBuilder) Send(version int64, event string, args []interface{}) error {
	params := []BytesValue{}
	for i := 0; i < len(args); i++ {
		v := reflect.ValueOf(args[i])
		if v.Type() == reflect.TypeOf([]byte{}) {
			params = append(params, BytesValue{
				BYTES,
				v.Bytes(),
			})
		} else {
			switch v.Kind() {
			case reflect.Bool:
				params = append(params, BytesValue{
					BOOLEAN,
					[]byte{bytes.BoolToBytes(v.Bool())},
				})
			case reflect.Int16:
				params = append(params, BytesValue{
					INT16,
					bytes.Int16ToBytes(int16(v.Int())),
				})
			case reflect.Int32:
				params = append(params, BytesValue{
					INT32,
					bytes.Int32ToBytes(int32(v.Int())),
				})
			case reflect.Int64:
				params = append(params, BytesValue{
					INT64,
					bytes.Int64ToBytes(v.Int()),
				})
			case reflect.String:
				params = append(params, BytesValue{
					TEXT,
					[]byte(v.String()),
				})
			default:
				return errors.New("only bool/int16/int32/int64/string/[]byte supported in args")
			}
		}
	}
	operation := ContractEventSendOperation{
		ContractAddress: cesob.address,
		Event:           event,
		Args:            BytesValueList{params},
		Version:         version,
	}
	if cesob.factory != nil {
		cesob.factory.addOperation(operation)
	}

	return nil
}
