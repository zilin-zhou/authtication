//上链的链码
package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

//定义信息

type Information struct {
	RSign       []byte `json:"rsign"`
	SSign       []byte `json:"ssign"`
	EncryptInfo []byte `json:"info"`
}
type AccumulatorInfo struct {
	Acc []byte `json:"acc"`
}
type InfoChaincode struct {
}
type AccumulatorInfo2 struct {
	Acc2     []byte `json:"acc2"`
	Witness1 []byte `json:witness1`
	Witness2 []byte `json:witness2`
	Witness3 []byte `json:witness3`
}

//链码初始方法
func (ic *InfoChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (ic *InfoChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	//获取用户操作
	fn, args := stub.GetFunctionAndParameters()
	if fn == "addInfo" {
		return ic.addInfo(stub, args)
	} else if fn == "readInfo" {
		return ic.readInfo(stub, args)
	} else if fn == "addAccInfo" {
		return ic.addAccInfo(stub, args)
	} else if fn == "deleteinfo" {
		return ic.deleteinfo(stub, args)
	} else if fn == "addAccInfo2" {
		return ic.addAccInfo2(stub, args)
	}
	return shim.Error("指定的函数名称错误")
}

/*
	把信息添加到区块中
	参数一：验证公钥
	参数二：加密的proof
	参数三：链码事件
*/
func (ic *InfoChaincode) addInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 5 {
		shim.Error("The number of parameters is incorrect!")
	}

	info := Information{
		RSign:       []byte(args[1]),
		SSign:       []byte(args[2]),
		EncryptInfo: []byte(args[3]),
	}
	infobyte, err := json.Marshal(info)
	if err != nil {
		return shim.Error("Parameter parsing failed!")
	}
	//写入账本
	err = stub.PutState(args[0], infobyte)
	if err != nil {
		return shim.Error("Write to ledger failed!")
	}
	//设置链码事件
	err = stub.SetEvent(args[4], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("The information was added successfully!"))
}
func (ic *InfoChaincode) addAccInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 3 {
		shim.Error("The number of parameters is incorrect!")
	}

	info := AccumulatorInfo{
		Acc: []byte(args[1]),
	}
	infobyte, err := json.Marshal(info)
	if err != nil {
		return shim.Error("Parameter parsing failed!")
	}
	//写入账本
	err = stub.PutState(args[0], infobyte)
	if err != nil {
		return shim.Error("Write to ledger failed!")
	}
	//设置链码事件
	err = stub.SetEvent(args[2], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("The information was added successfully!"))
}

////从区块中读出信息
//参数一：zpk
func (ic *InfoChaincode) readInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//根据zpk值查找加密的信息
	proofbyte, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("An error occurred while taking value from the ledger!")
	}

	return shim.Success(proofbyte)
}

func (ic *InfoChaincode) deleteinfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("参数个数错误!")
	}

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}
	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("信息删除成功"))
}
func (ic *InfoChaincode) addAccInfo2(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	info := AccumulatorInfo2{
		Acc2:     []byte(args[1]),
		Witness1: []byte(args[2]),
		Witness2: []byte(args[3]),
		Witness3: []byte(args[4]),
	}
	infobyte, err := json.Marshal(info)
	if err != nil {
		return shim.Error("Parameter parsing failed!")
	}
	//写入账本
	err = stub.PutState(args[0], infobyte)
	if err != nil {
		return shim.Error("Write to ledger failed!")
	}
	//设置链码事件
	err = stub.SetEvent(args[5], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte("The information was added successfully!"))
}
func main() {
	err := shim.Start(new(InfoChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
