package ledger_model

type EventAccountOperationBuilder struct {
	eventPublishBuilder *EventPublishOperationBuilder
	permissionBuilder   *AccountPermissionSetOperationBuilder
}

func NewEventAccountOperationBuilder(address []byte, factory *BlockchainOperationFactory) *EventAccountOperationBuilder {
	return &EventAccountOperationBuilder{
		eventPublishBuilder: NewEventPublishOperationBuilder(address, factory),
		permissionBuilder:   NewAccountPermissionSetOperationBuilder(EVENT, address, factory),
	}
}

func (eaob *EventAccountOperationBuilder) PublishBytes(key string, content []byte, sequence int64) *EventPublishOperationBuilder {
	eaob.eventPublishBuilder.PublishBytes(key, content, sequence)

	return eaob.eventPublishBuilder
}

func (eaob *EventAccountOperationBuilder) PublishInt64(key string, content int64, sequence int64) *EventPublishOperationBuilder {
	eaob.eventPublishBuilder.PublishInt64(key, content, sequence)

	return eaob.eventPublishBuilder
}

func (eaob *EventAccountOperationBuilder) PublishString(key string, content string, sequence int64) *EventPublishOperationBuilder {
	eaob.eventPublishBuilder.PublishString(key, content, sequence)

	return eaob.eventPublishBuilder
}

func (eaob *EventAccountOperationBuilder) Permission() *AccountPermissionSetOperationBuilder {
	return eaob.permissionBuilder
}
