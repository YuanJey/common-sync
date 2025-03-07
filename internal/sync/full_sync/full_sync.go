package full_sync

import (
	"common-sync/pkg/config"
	"common-sync/pkg/third_data"
	"common-sync/pkg/utils"
	"github.com/YuanJey/go-log/pkg/log"
	utils2 "github.com/YuanJey/goutils2/pkg/utils"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
)

type FullSync struct {
	thirdData third_data.ThirdData
}

const (
	Source_Buildin = "buildin"
	Source_Sync    = "sync"
)

func (f *FullSync) Order(order int) int {
	if config.ServerConfig.SyncOptions.Sort == 0 {
		return 100000 - order
	}
	return order
}
func (f *FullSync) SyncDept(operationID, taskID string) error {
	list, err := f.thirdData.GetThirdDeptList()
	if err != nil {
		log.Error(operationID, "failed to get third dept list", err)
		return err
	}
	var lastDepartments []tables.TbLasDepartment
	for i := range list {
		dbData := utils.HandleDeptData(list[i])
		if dbData.DID != "" {
			dbData.TaskID = taskID
			dbData.PlatformID = config.ServerConfig.WPS.PlatformId
			dbData.ThirdCompanyID = config.ServerConfig.WPS.CompanyId
			dbData.Order = f.Order(i)
			dbData.Source = Source_Buildin
			lastDepartments = append(lastDepartments, dbData)
			continue
		}
		log.Error(operationID, "failed to handle dept data", utils2.StructToJsonString(list[i]))
	}
	for j := 0; j < len(lastDepartments); j += 100 {
		end := j + 100
		if end > len(lastDepartments) {
			end = len(lastDepartments)
		}
		temp := lastDepartments[j:end] // correct way to slice
		err = mysql.BatchInsertTbLasDepartment(temp)
		if err != nil {
			log.Error(operationID, "batch insert department data error ", err.Error())
			for i2 := range temp {
				err2 := mysql.BatchInsertTbLasDepartment([]tables.TbLasDepartment{temp[i2]})
				if err2 != nil {
					log.Error(operationID, "data is ", utils2.StructToJsonString(temp[i2]))
					return err2
				}
			}
			return err
		}
	}
	return nil
}
func (f *FullSync) SyncUser(operationID, taskID string) error {
	list, err := f.thirdData.GetThirdUserList()
	if err != nil {
		log.Error(operationID, "failed to get third user list", err.Error())
		return err
	}
	var lastDepartmentUsers []tables.TbLasDepartmentUser
	var lastUsers []tables.TbLasUser
	for i := range list {
		user, departmentUser := utils.HandleUserData(list[i])
		if user.UID != "" && departmentUser.UID != "" {
			user.TaskID = taskID
			user.PlatformID = config.ServerConfig.WPS.PlatformId
			user.ThirdCompanyID = config.ServerConfig.WPS.CompanyId
			user.DefDIDOrder = f.Order(i)
			user.Source = Source_Buildin

			departmentUser.TaskID = taskID
			departmentUser.PlatformID = config.ServerConfig.WPS.PlatformId
			departmentUser.ThirdCompanyID = config.ServerConfig.WPS.CompanyId
			departmentUser.Order = f.Order(i)
			departmentUser.Main = 1

			lastUsers = append(lastUsers, user)
			lastDepartmentUsers = append(lastDepartmentUsers, departmentUser)
			continue
		}
		log.Error(operationID, "failed to handle user data", utils2.StructToJsonString(list[i]))
	}
	for j := 0; j < len(lastUsers); j += 100 {
		end := j + 100
		if end > len(lastUsers) {
			end = len(lastUsers)
		}
		temp := lastUsers[j:end] // correct way to slice
		err = mysql.BatchInsertTbLasUser(temp)
		if err != nil {
			log.Error(operationID, "batch insert lastUsers data error ", err.Error())
			for i2 := range temp {
				err2 := mysql.BatchInsertTbLasUser([]tables.TbLasUser{temp[i2]})
				if err2 != nil {
					log.Error(operationID, "data is ", utils2.StructToJsonString(temp[i2]))
					return err2
				}
			}
			return err
		}
	}
	for j := 0; j < len(lastDepartmentUsers); j += 100 {
		end := j + 100
		if end > len(lastDepartmentUsers) {
			end = len(lastDepartmentUsers)
		}
		temp := lastDepartmentUsers[j:end] // correct way to slice
		err = mysql.BatchInsertTbLasDepartmentUser(temp)
		if err != nil {
			log.Error(operationID, "batch insert lastDepartmentUsers data error ", err.Error())
			for i2 := range temp {
				err2 := mysql.BatchInsertTbLasDepartmentUser([]tables.TbLasDepartmentUser{temp[i2]})
				if err2 != nil {
					log.Error(operationID, "data is ", utils2.StructToJsonString(temp[i2]))
					return err2
				}
			}
			return err
		}
	}
	return nil
}
