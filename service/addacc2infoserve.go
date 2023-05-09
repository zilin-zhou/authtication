package service

import "github.com/hyperledger/fabric-sdk-go/pkg/client/channel"

type AccumulatorInfo2 struct {
	Acc2     []byte `json:"acc2"`
	Witness1 []byte `json:"witness1"`
	Witness2 []byte `json:"witness2"`
	Witness3 []byte `json:"witness3"`
}

//使用sdk操作区块链 向区块添加信息
func (s *ServiceSetup) AddAcc2InfoServe(args [][]byte) (channel.Response, error) {
	eventID := "EventAddAccInfo"
	//注册链码事件
	reg, notifier := regitserEvent(s.Client, s.ChaincodeID, eventID)
	defer s.Client.UnregisterChaincodeEvent(reg)
	args = append(args, []byte(eventID))
	//调用链码
	req := channel.Request{ChaincodeID: s.ChaincodeID, Fcn: "addAccInfo2", Args: args}
	resp, err := s.Client.Execute(req)
	if err != nil {
		return channel.Response{}, err
	}
	//调用链码事件结果
	err = eventResult(notifier, eventID)
	if err != nil {
		return channel.Response{}, err
	}
	return resp, nil
}
