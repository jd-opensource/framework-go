package ledger_model

type DataAccountOperationBuilder struct {
	kvBuilder         *DataAccountKVSetOperationBuilder
	permissionBuilder *AccountPermissionSetOperationBuilder
}

func NewDataAccountOperationBuilder(address []byte, factory *BlockchainOperationFactory) *DataAccountOperationBuilder {
	return &DataAccountOperationBuilder{
		kvBuilder:         NewDataAccountKVSetOperationBuilder(address, factory),
		permissionBuilder: NewAccountPermissionSetOperationBuilder(DATA, address, factory),
	}
}

func (daob *DataAccountOperationBuilder) SetBytes(key string, value []byte, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetBytes(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) SetImage(key string, value []byte, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetImage(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) SetText(key, value string, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetText(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) SetJSON(key, value string, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetJSON(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) SetInt64(key string, value int64, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetInt64(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) SetTimestamp(key string, value int64, expVersion int64) *DataAccountKVSetOperationBuilder {
	daob.kvBuilder.SetTimestamp(key, value, expVersion)

	return daob.kvBuilder
}

func (daob *DataAccountOperationBuilder) Permission() *AccountPermissionSetOperationBuilder {
	return daob.permissionBuilder
}
