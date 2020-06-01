package ledger_model

/*
 * Author: imuge
 * Date: 2020/6/1 下午7:23
 */

type ContractOperator interface {

	// 部署合约
	Contracts() *ContractCodeDeployOperationBuilder
}