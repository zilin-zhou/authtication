module application

go 1.15

require (
	gitee.com/frankyu365/gocrypto v1.0.7-alpha
	github.com/gin-gonic/gin v1.8.1
	github.com/hyperledger/fabric-protos-go v0.0.0-20220827195505-ce4c067a561d
	github.com/hyperledger/fabric-sdk-go v1.0.0
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/wealdtech/go-merkletree v1.0.0
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064
	golang.org/x/sys v0.0.0-20220319134239-a9b59b0215f8 // indirect
	vuvuzela.io/concurrency v0.0.0-20190327123758-e608f351e310

)

replace github.com/hyperledger/fabric-protos-go v0.0.0-20220827195505-ce4c067a561d => github.com/hyperledger/fabric-protos-go v0.0.0-20210311171918-e08edaab0493
