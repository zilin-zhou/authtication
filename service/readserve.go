package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (s *ServiceSetup) ReadInfo(args [][]byte) ([]byte, error) {

	//SDK调用链码
	req := channel.Request{ChaincodeID: s.ChaincodeID, Fcn: "readInfo", Args: args}
	resp, err := s.Client.Query(req)
	if err != nil {
		if err != nil {
			return []byte{0x00}, err
		}
	}
	return resp.Payload, nil
}
