package tables

type TbLasDepartmentIncrement struct {
	ThirdCompanyID string `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"`                    // 三方租户id
	PlatformID     string `gorm:"column:platform_id;type:varchar(60);not null;uniqueIndex:idx_platform_did" json:"platform_id"` // 平台id，用来区分多种数据源，platform_id + did 唯一
	DID            string `gorm:"column:did;type:varchar(255);not null;uniqueIndex:idx_platform_did" json:"did"`                // 三方部门id
	PID            string `gorm:"column:pid;type:varchar(255);default:null" json:"pid"`                                         // 三方父部门id
	Name           string `gorm:"column:name;type:varchar(255);default:null" json:"name"`                                       // 部门名称
	Order          int    `gorm:"column:order;type:int(11);default:0" json:"order"`                                             // 排序
	Source         string `gorm:"column:source;type:varchar(20);default:'sync'" json:"source"`                                  // 来源，buildin/sync
	UpdateType     string `gorm:"column:update_type;type:varchar(20);not null" json:"update_type"`                              // 修改类型, dept_del/dept_update/dept_add/dept_move
	//CreationTime     time.Time `gorm:"column:ctime;type:timestamp" json:"ctime"`                                                     // 创建时间
	//ModificationTime time.Time `gorm:"column:mtime;type:timestamp" json:"mtime"`                                                     // 修改时间
}

func (TbLasDepartmentIncrement) TableName() string {
	return "tb_las_department_increment"
}

//func (t *TbLasDepartmentIncrement) BeforeUpdate(tx *gorm.DB) (err error) {
//	action := LatestDeptActions{
//		DID:    t.DID,
//		Action: t.UpdateType,
//	}
//	return tx.Table(action.TableName()).Save(action).Error
//}
