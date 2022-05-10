package crypto

import (
	"fmt"
	"github.com/blockchain-jd-com/framework-go/crypto/adv"
	"github.com/blockchain-jd-com/framework-go/crypto/classic"
	"github.com/blockchain-jd-com/framework-go/crypto/framework"
	"github.com/blockchain-jd-com/framework-go/crypto/sm"
)

/**
 * @Author: imuge
 * @Date: 2020/4/29 10:22 上午
 */

var (
	functions  = make(map[int16]framework.CryptoFunction)
	algorithms = make(map[int16]framework.CryptoAlgorithm)
	names      = make(map[string]int16)
)

func init() {
	classicArray := classic.NewClassicCryptoService().GetFunctions()
	for _, function := range classicArray {
		functions[function.GetAlgorithm().Code] = function
		algorithms[function.GetAlgorithm().Code] = function.GetAlgorithm()
		names[function.GetAlgorithm().Name] = function.GetAlgorithm().Code
	}
	smArray := sm.NewSMCryptoService().GetFunctions()
	for _, function := range smArray {
		functions[function.GetAlgorithm().Code] = function
		algorithms[function.GetAlgorithm().Code] = function.GetAlgorithm()
		names[function.GetAlgorithm().Name] = function.GetAlgorithm().Code
	}
	advArray := adv.NewAdvCryptoService().GetFunctions()
	for _, function := range advArray {
		functions[function.GetAlgorithm().Code] = function
		algorithms[function.GetAlgorithm().Code] = function.GetAlgorithm()
		names[function.GetAlgorithm().Name] = function.GetAlgorithm().Code
	}

}

// 获取算法定义
func GetAlgorithmByCode(code int16) framework.CryptoAlgorithm {
	algorithm, exists := algorithms[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return algorithm
}

func GetAlgorithmByName(name string) framework.CryptoAlgorithm {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetAlgorithmByCode(code)
}

// 随机算法实现
func GetRandomFunctionByCode(code int16) framework.RandomFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.RandomFunction)
}

func GetRandomFunctionByName(name string) framework.RandomFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetRandomFunctionByCode(code)
}

func GetRandomFunction(algorithm framework.CryptoAlgorithm) framework.RandomFunction {
	return GetRandomFunctionByCode(algorithm.Code)
}

// Hash算法实现
func GetHashFunctionByCode(code int16) framework.HashFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.HashFunction)
}

func GetHashFunctionByName(name string) framework.HashFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetHashFunctionByCode(code)
}

func GetHashFunction(algorithm framework.CryptoAlgorithm) framework.HashFunction {
	return GetHashFunctionByCode(algorithm.Code)
}

// 非对称加密算法实现
func GetAsymmetricEncryptionFunctionByCode(code int16) framework.AsymmetricEncryptionFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.AsymmetricEncryptionFunction)
}

func GetAsymmetricEncryptionFunctionByName(name string) framework.AsymmetricEncryptionFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetAsymmetricEncryptionFunctionByCode(code)
}

func GetAsymmetricEncryptionFunction(algorithm framework.CryptoAlgorithm) framework.AsymmetricEncryptionFunction {
	return GetAsymmetricEncryptionFunctionByCode(algorithm.Code)
}

// 签名算法实现
func GetSignatureFunctionByCode(code int16) framework.SignatureFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.SignatureFunction)
}

func GetSignatureFunctionByName(name string) framework.SignatureFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetSignatureFunctionByCode(code)
}

func GetSignatureFunction(algorithm framework.CryptoAlgorithm) framework.SignatureFunction {
	return GetSignatureFunctionByCode(algorithm.Code)
}

// 对称加密算法
func GetSymmetricEncryptionFunctionByCode(code int16) framework.SymmetricEncryptionFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function.(framework.SymmetricEncryptionFunction)
}

func GetSymmetricEncryptionFunctionByName(name string) framework.SymmetricEncryptionFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetSymmetricEncryptionFunctionByCode(code)
}

func GetSymmetricEncryptionFunction(algorithm framework.CryptoAlgorithm) framework.SymmetricEncryptionFunction {
	return GetSymmetricEncryptionFunctionByCode(algorithm.Code)
}

// 密码算法实现
func GetCryptoFunctionByCode(code int16) framework.CryptoFunction {
	function, exists := functions[code]
	if !exists {
		panic(fmt.Sprintf("Algorithm [code:%d] has no service provider!", code))
	}

	return function
}

func GetCryptoFunctionByName(name string) framework.CryptoFunction {
	code, exists := names[name]
	if !exists {
		panic(fmt.Sprintf("Algorithm [name:%s] has no service provider!", name))
	}
	return GetCryptoFunctionByCode(code)
}

func GetCryptoFunction(algorithm framework.CryptoAlgorithm) framework.CryptoFunction {
	return GetCryptoFunctionByCode(algorithm.Code)
}
