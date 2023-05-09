module application/chaincode

go 1.15

require (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20220920210243-7bc6fa0dd58b
	github.com/hyperledger/fabric-protos-go v0.0.0-20220827195505-ce4c067a561d
)

replace (
	github.com/hyperledger/fabric-chaincode-go v0.0.0-20220920210243-7bc6fa0dd58b => github.com/hyperledger/fabric-chaincode-go v0.0.0-20201119163726-f8ef75b17719
	github.com/hyperledger/fabric-protos-go v0.0.0-20220827195505-ce4c067a561d => github.com/hyperledger/fabric-protos-go v0.0.0-20210311171918-e08edaab0493
)
