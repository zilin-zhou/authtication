package tools

import (
	"application/sdkInit"
	"application/service"
	"fmt"
	"os"
)

var FabricSetup *service.ServiceSetup

const (
	cc_name    = "zero_gnark"
	cc_version = "1.0.0"
)

//读取配置文件创建sdk服务实例
func SdkSetup() {

	var orgs = []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    1,
			OrgAnchorFile: os.Getenv("GOPATH") + "/src/application/fixtures/channel-artifacts/Org1MSPanchors.tx",
		},
	}

	// init sdk env info
	var info = sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    os.Getenv("GOPATH") + "/src/application/fixtures/channel-artifacts/channel.tx",
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    os.Getenv("GOPATH") + "/src/application/chaincode/",
		ChaincodeVersion: cc_version,
	}
	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Printf(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> 通过链码外部服务设置链码状态......")

	serviceSetup, err := service.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk)
	if err != nil {
		fmt.Println()
		os.Exit(-1)
	}

	FabricSetup = serviceSetup
}
