package admin

import (
	"common-sync/pkg/config"
	"github.com/YuanJey/goutils2/pkg/utils"
	"github.com/gin-gonic/gin"
	"reflect"
)

type UserReq struct {
	UID              string `json:"uid" binding:"required"`
	NickName         string `json:"name" binding:"required"`
	DefDID           string `json:"def_did"  binding:"required"`
	DefDIDOrder      string `json:"def_did_order"`
	Account          string `json:"account" binding:"required"`
	Password         string `json:"password"`
	Avatar           string `json:"avatar"`
	Email            string `json:"email"`
	Gender           string `json:"gender"`
	Title            string `json:"title"`
	WorkPlace        string `json:"work_place"`
	Leader           string `json:"leader"`
	Employer         string `json:"employer"`
	EmploymentStatus string `json:"employment_status"`
	EmploymentType   string `json:"employment_type"`
	Phone            string `json:"phone"`
	Telephone        string `json:"telephone"`
}
type DeptReq struct {
	DID   string `json:"did" binding:"required"`
	PID   string `json:"pid" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Order string `json:"order"`
}

func DeptMap(c *gin.Context) {
	operationID := utils.OperationIDGenerator()
	req := DeptReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"operationID": operationID, "msg": "参数错误", "err": err.Error()})
		return
	}
	config.ServerConfig.SetDeptFields(req.DID, "DID")
	config.ServerConfig.SetDeptFields(req.PID, "PID")
	config.ServerConfig.SetDeptFields(req.Name, "Name")
	if req.Order != "" {
		config.ServerConfig.SetDeptFields(req.Order, "Order")
	}
	c.JSON(200, config.ServerConfig.DeptFields)
}
func UserMap(c *gin.Context) {
	operationID := utils.OperationIDGenerator()
	req := UserReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"operationID": operationID, "msg": "参数错误", "err": err.Error()})
		return
	}
	value := reflect.ValueOf(req)
	for i := 0; i < value.NumField(); i++ {
		if value.Field(i).String() != "" {
			config.ServerConfig.SetUserFields(value.Field(i).String(), value.Type().Field(i).Name)
		}
	}
	c.JSON(200, config.ServerConfig.UserFields)
}
