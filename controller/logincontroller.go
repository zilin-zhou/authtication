package controller

import (
	"application/service"
	"application/tools"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginControll struct {
}
type LoginRequest struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"` //QKD
	ID       string   `json:"ID"`    //指纹
	UserID   string   `json:"UserID"`
	Witness  *big.Int `json:"witness"`
}

func (lc LoginControll) Login(ctx *gin.Context) {
	var count = 0
	request := new(LoginRequest)
	//查看响应时间

	err := ctx.ShouldBind(&request)
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("Parameter binding failed:%s", err))
		return
	}
	var requestbytes [][]byte
	//逻辑电路
	var LCCSstr []string
	var Lccs [][]byte
	//判断
	switch {
	case request.Username == "" && request.Password == "":
		if request.Phone == "" && request.ID != " " {
			_, str := MyHash([]byte(request.ID))
			LCCSstr = append(LCCSstr, str)
			Lccs = Lsscbytes(LCCSstr)
		}
		if request.Phone != "" && request.ID != "" {
			requestbytes = append(requestbytes, []byte(request.Phone), []byte(request.ID))
			for i := 0; i < len(requestbytes); i++ {
				_, str := MyHash(requestbytes[i])
				LCCSstr = append(LCCSstr, str)
			}
			Lccs = Lsscbytes(LCCSstr)
		}
	case request.Username != "" && request.Password != "":
		if request.Phone != "" && request.ID != "" {
			requestbytes = append(requestbytes, []byte(request.Username), []byte(request.Password), []byte(request.Phone), []byte(request.ID))
			for i := 0; i < len(requestbytes); i++ {
				_, str := MyHash(requestbytes[i])
				LCCSstr = append(LCCSstr, str)
			}
			strtemp := StrByXOR(LCCSstr[0], LCCSstr[1])
			//删除前两个元素
			LCCSstr = LCCSstr[2:]
			LCCSstr = append(LCCSstr, strtemp)
			Lccs = Lsscbytes(LCCSstr)
		}
		if request.Phone != "" && request.ID == "" {
			requestbytes = append(requestbytes, []byte(request.Username), []byte(request.Password), []byte(request.Phone))
			for i := 0; i < len(requestbytes); i++ {
				_, str := MyHash(requestbytes[i])
				LCCSstr = append(LCCSstr, str)
			}
			strtemp := StrByXOR(LCCSstr[0], LCCSstr[1])
			//删除前两个元素
			LCCSstr = LCCSstr[2:]
			LCCSstr = append(LCCSstr, strtemp)
			Lccs = Lsscbytes(LCCSstr)
		}
	}

	//上链查找累加器的值
	if request.Witness == nil {
		tools.Failed(ctx, fmt.Sprintln("Witness is nil"))
		return
	}
	witnessbyte, _ := json.Marshal(request.Witness)

	_, witnessstr := MyHash(witnessbyte)

	//上链查找
	accresult, err := tools.FabricSetup.ReadInfo([][]byte{[]byte(witnessstr)})
	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("The query acc failed:%s\n", err))
		return
	}
	accinfo := service.AccumulatorInfo{}
	if err = json.Unmarshal(accresult, &accinfo); err != nil {
		log.Println("Failed to query accumulator information! No login!Please enter the correct verification information!")
		return
	}

	var acc *big.Int
	if err = json.Unmarshal(accinfo.Acc, &acc); err != nil {
		tools.Failed(ctx, fmt.Sprintf("acc Deserialization failed:%s\n", err))
		return
	}

	time3 := time.Now().UnixNano()
	bl := Public1.Verify(acc, request.Witness, []byte(request.UserID))
	time4 := time.Now().UnixNano()
	log.Printf("累加器1验证时间:%d", time4-time3)

	if !bl {
		tools.Failed(ctx, fmt.Sprintf("witness failed:%s\n", err))
		return
	}
	_, acc1str := MyHash(accinfo.Acc)

	acc2result, err := tools.FabricSetup.ReadInfo([][]byte{[]byte(acc1str)})

	if err != nil {
		tools.Failed(ctx, fmt.Sprintf("The query acc failed:%s\n", err))
		return
	}
	acclccs := ACCLCSS{}
	acc2info := service.AccumulatorInfo2{}
	err = json.Unmarshal(acc2result, &acc2info)
	if err != nil {
		log.Println(err)
	}
	var acc2 *big.Int
	json.Unmarshal(acc2info.Acc2, &acc2)
	acclccs.ACC = acc2

	var witness1 *big.Int
	json.Unmarshal(acc2info.Witness1, &witness1)
	acclccs.Witness = append(acclccs.Witness, witness1)

	var witness2 *big.Int
	json.Unmarshal(acc2info.Witness2, &witness2)
	acclccs.Witness = append(acclccs.Witness, witness2)

	var witness3 *big.Int
	json.Unmarshal(acc2info.Witness3, &witness3)
	acclccs.Witness = append(acclccs.Witness, witness3)
	var timetotal int64
	timetotal = 0
	for i := 0; i < len(acclccs.Witness); i++ {
		for j := 0; j < len(Lccs); j++ {
			public2 := UserPublic2[request.UserID]
			time1 := time.Now().UnixNano()
			bl = public2.Verify(acclccs.ACC, acclccs.Witness[i], Lccs[j])
			time2 := time.Now().UnixNano()
			if bl {
				timetotal += time2 - time1
				count++
			}
			time1 = 0
			time2 = 0
		}
	}

	log.Printf("累加器2验证时间:%d", timetotal)
	if count == len(Lccs) {
		tools.Success(ctx, "验证成功!")
	} else {
		tools.Failed(ctx, "失败")
	}
}
