package controller

import (
	"application/accumulator"
	"application/service"
	"application/tools"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"gitee.com/frankyu365/gocrypto/ecc"
	"github.com/gin-gonic/gin"
)

var Acc1 *big.Int
var Public1 *accumulator.PublicKey
var Accdata [][]byte
var UserPublic2 map[string]*accumulator.PublicKey

//接受传来的属性结构体
type RegisterController struct {
}
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"` //QKD
	ID       string `json:"ID"`    //指纹
}
type ACCLCSS struct {
	ACC     *big.Int
	Witness []*big.Int
}

type UserInfomation struct {
	Lccs   [][]byte
	UserID string
}

func (r RegisterController) Register(ctx *gin.Context) {
	//UserPublic2 = make(map[string]*accumulator.PublicKey)
	registerbody := new(RegisterRequest)

	err := ctx.ShouldBind(&registerbody)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("Parameter binding failed:%s", err))
		return
	}

	//产生属性的       merkle树
	var attributedata [][]byte
	attributedata = append(attributedata, []byte(registerbody.Username))
	attributedata = append(attributedata, []byte(registerbody.Password))
	attributedata = append(attributedata, []byte(registerbody.Phone))
	attributedata = append(attributedata, []byte(registerbody.ID))

	//生成用户的公私钥
	bl, _ := PathExists("/home/zilin/go/src/application/kgc/user/" + registerbody.Username + "eccPublic.pem")
	if bl {
		log.Println("The user's private key already exists, determine whether the user ID exists")
	} else {
		ecc.GenerateECCKey(256, "./kgc/user/"+registerbody.Username)
	}
	userprivate, _ := GereratePub("/home/zilin/go/src/application/kgc/user/" + registerbody.Username + "eccPublic.pem")
	//生成用户标识和逻辑控制电路信息
	userID, lccs := GenerateUserID(attributedata, userprivate)

	userinfo := UserInfomation{
		Lccs:   lccs,
		UserID: userID,
	}

	//信息序列化
	userinfobyte, _ := json.Marshal(userinfo)

	//对userinfo签名
	ecctime1 := time.Now().UnixNano()
	rText, sText, _ := ecc.ECCSign(userinfobyte, "/home/zilin/go/src/application/kgc/user/"+registerbody.Username+"eccPrivate.pem")
	ecctime2 := time.Now().UnixNano()
	log.Printf("签名时间%d", ecctime2-ecctime1)
	//使用register的pk进行加密
	ecctime1 = time.Now().UnixNano()
	cipherText, _ := ecc.EccEncrypt(userinfobyte, "/home/zilin/go/src/application/kgc/register/eccPublic.pem")
	ecctime2 = time.Now().UnixNano()
	log.Printf("加密时间%d", ecctime2-ecctime1)
	//使用注册服务器的公钥 进行加密上链
	registerpublic, _ := GereratePub("/home/zilin/go/src/application/kgc/register/eccPublic.pem")
	var bodyBytes [][]byte

	//rpk公钥进行hash
	_, rpkshashtr := MyHash(registerpublic)
	bodyBytes = append(bodyBytes, []byte(rpkshashtr))
	bodyBytes = append(bodyBytes, rText)
	bodyBytes = append(bodyBytes, sText)
	bodyBytes = append(bodyBytes, cipherText)

	tools.FabricSetup.AddServe(bodyBytes)

	//进行注册
	//注册服务器查找信息    验证身份
	result, err := tools.FabricSetup.ReadInfo([][]byte{[]byte(rpkshashtr)})
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("The query failed:%s\n", err))
		return
	}

	var queryresult service.Information
	err = json.Unmarshal(result, &queryresult)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("sign Deserialization failed:%s\n", err))
		return
	}
	//解密
	eccDtime1 := time.Now().UnixNano()
	plainText, err := ecc.EccDecrypt(queryresult.EncryptInfo, "/home/zilin/go/src/application/kgc/register/eccPrivate.pem")
	eccDtime2 := time.Now().UnixNano()
	log.Printf("解密时间为%d", eccDtime2-eccDtime1)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("userinfo Decryption failed:%s\n", err))
		return
	}

	//验签
	eccDtime1 = time.Now().UnixNano()
	bl, err = ecc.ECCVerify(plainText, queryresult.RSign, queryresult.SSign, "/home/zilin/go/src/application/kgc/user/"+registerbody.Username+"eccPublic.pem")
	eccDtime2 = time.Now().UnixNano()
	log.Printf("验证签名时间为%d", eccDtime2-eccDtime1)
	if !bl || err != nil {
		tools.Failed(ctx, fmt.Sprintf("Signature verification failed:%s\n", err))
		return
	}

	userinfo1 := UserInfomation{}

	err = json.Unmarshal(plainText, &userinfo1)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("userinfo Deserialization failed:%s\n", err))
		return
	}

	//创建累加器1   将UserID 放入累加器
	if Acc1 == nil && Public1 == nil {
		Public1, _, _ = accumulator.GenerateKey(rand.Reader)
	}
	Accdata = append(Accdata, []byte(userinfo1.UserID))
	time1 := time.Now().UnixNano()
	acc1, witness1 := Public1.Accumulate(Accdata...)
	time2 := time.Now().UnixNano()

	log.Printf("累加器1时间:%d", time2-time1)
	Acc1 = acc1

	//累加器的值上链 并且返回用户witness

	witnessbyte, _ := json.Marshal(witness1[len(witness1)-1])

	_, witnessstr1 := MyHash(witnessbyte)

	Accbyte1, err := json.Marshal(Acc1)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("Accumulator serialization failed:%s\n", err))
		return
	}
	_, err = tools.FabricSetup.AddAccInfoServe([][]byte{[]byte(witnessstr1), Accbyte1})
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("The accumulator failed to write to the ledger:%s\n", err))
		return
	}

	//创建用户累加器2
	Public2, _, _ := accumulator.GenerateKey(rand.Reader)

	UserPublic2 = map[string]*accumulator.PublicKey{
		userID: Public2,
	}

	time3 := time.Now().UnixNano()
	acc2, witness2 := Public2.Accumulate(lccs...)
	time4 := time.Now().UnixNano()
	log.Printf("累加器2时间:%d", time4-time3)

	//序列化
	_, acc1str := MyHash(Accbyte1)
	var Acc2infobytes [][]byte
	Acc2infobytes = append(Acc2infobytes, []byte(acc1str))
	acc2byte, _ := json.Marshal(acc2)
	Acc2infobytes = append(Acc2infobytes, acc2byte)
	for i := 0; i < len(witness2); i++ {
		witness2byte, _ := json.Marshal(witness2[i])
		Acc2infobytes = append(Acc2infobytes, witness2byte)
	}

	_, err = tools.FabricSetup.AddAcc2InfoServe(Acc2infobytes)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("The accumulator failed to write to the ledger:%s\n", err))
		return
	}

	//得到的witness 加密返回给用户
	data := &struct {
		Msg     string
		Witness *big.Int
		ID      string
	}{
		Msg:     "register successful!",
		Witness: witness1[len(witness1)-1],
		ID:      userID,
	}
	tools.Success(ctx, data)
}

//hash函数
func MyHash(data []byte) ([]byte, string) {
	//创建hash对象
	MyHash := sha256.New()
	//写入数据
	MyHash.Write(data)
	//计算结果
	hash := MyHash.Sum(nil)
	hashstr := hex.EncodeToString(hash)
	return hash, hashstr
}

//得到公钥
func GereratePub(filename string) ([]byte, error) {
	//打开公钥文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	buf := make([]byte, fileInfo.Size())
	file.Read(buf)
	file.Close()

	block, _ := pem.Decode(buf)
	pubinterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := pubinterface.(*ecdsa.PublicKey)
	byte, err := json.Marshal(publicKey)
	return byte, err
}

//生成用户标识
func GenerateUserID(attrs [][]byte, private []byte) (string, [][]byte) {
	var hashstrs []string

	for _, attr := range attrs {
		_, hashstr := MyHash(attr)
		hashstrs = append(hashstrs, hashstr)
	}
	_, privatehashstr := MyHash(private)

	result := hashstrs[0]
	for i := 0; i < len(hashstrs)-1; i++ {
		result = StrByXOR(result, hashstrs[i+1])
	}
	var lccs []string
	lccsstr1 := StrByXOR(hashstrs[0], hashstrs[1])

	lccs = append(lccs, lccsstr1, hashstrs[2], hashstrs[3])
	lccsbytes := Lsscbytes(lccs)
	result = StrByXOR(result, privatehashstr)
	resultbyte, _ := json.Marshal(result)
	_, resultstr := MyHash(resultbyte)
	return resultstr, lccsbytes
}

func StrByXOR(str1 string, str2 string) string {
	strlen := len(str1)

	result := ""

	for i := 0; i < strlen; i++ {
		result += string(str1[i] ^ str2[i])
	}
	return result
}
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func HashString(input string) string {
	// Create a new SHA-256 hash object
	hash := sha256.New()

	// Write the input string to the hash object
	hash.Write([]byte(input))

	// Get the hash sum as a byte slice and convert to a hex string
	hashSum := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashSum)

	return hashStr
}

func Lsscbytes(lccs []string) [][]byte {
	var lccsbytes [][]byte
	for i := 0; i < len(lccs); i++ {
		lccsbyte, _ := json.Marshal(lccs[i])
		lccsbytes = append(lccsbytes, lccsbyte)
	}
	return lccsbytes
}
