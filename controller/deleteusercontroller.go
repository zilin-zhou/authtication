package controller

// import (
// 	"application/tools"
// 	"encoding/json"
// 	"log"

// 	"github.com/gin-gonic/gin"
// )

// type DeleteUserController struct{}
// type DeleteUserRequest struct {
// 	UserID string `json:"userid"`
// }

// func (dc DeleteUserController) DeleteUser(ctx *gin.Context) {
// 	deleteuser := DeleteUserRequest{}

// 	err := ctx.Bind(&deleteuser)
// 	if err != nil {
// 		log.Printf("The binding parameter is incorrect!")
// 		return
// 	}
// 	userInfo := UsersMap[deleteuser.UserID]
// 	// fmt.Println(UsersMap)
// 	delete(UsersMap, deleteuser.UserID)
// 	log.Println("Register the user information in the user list maintained by the module!")
// 	log.Println("Delete the corresponding accumulator information in the ledger!")
// 	// fmt.Println(UsersMap)
// 	keybyte, _ := json.Marshal(userInfo)
// 	_, key := MyHash(keybyte)
// 	delete(TreeMap, key)
// 	log.Println("Delete the corresponding tree MerkleTree information!")
// 	resp, err := tools.FabricSetup.DeleteServe([][]byte{[]byte(key)})
// 	if err != nil {
// 		log.Println("Failed to delete information in ledger!")
// 		return
// 	}
// 	log.Println("The user was deleted successfully!")
// 	tools.Success(ctx, string(resp.Payload))
// }
