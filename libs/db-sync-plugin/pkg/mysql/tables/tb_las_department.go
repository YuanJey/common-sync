package tables

type TbLasDepartment struct {
	DID            string `gorm:"column:did;type:varchar(255);not null;uniqueIndex:idx_platform_did;comment:三方部门id" json:"did"`                        // 三方部门id
	TaskID         string `gorm:"column:task_id;type:varchar(20);not null" json:"task_id"`                                                             // 任务id
	ThirdCompanyID string `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"`                                           // 三方租户id
	PlatformID     string `gorm:"column:platform_id;type:varchar(60);not null;uniqueIndex:idx_platform_did;comment:平台id，用来区分多种数据源" json:"platform_id"` // 平台id
	PID            string `gorm:"column:pid;type:varchar(255);not null;comment:三方父部门id" json:"pid"`                                                    // 三方父部门id
	Name           string `gorm:"column:name;type:varchar(255);not null;comment:部门名称" json:"name"`                                                     // 部门名称
	Source         string `gorm:"column:source;type:varchar(20);default:'sync';comment:来源，buildin/sync" json:"source"`                                 // 来源
	Order          int    `gorm:"column:order;type:int;default:0;comment:排序" json:"order"`                                                             // 排序
}

func (*TbLasDepartment) TableName() string {
	return "tb_las_department"
}

//func (d *TbLasDepartment) BeforeSave(tx *gorm.DB) (err error) {
//	return err
//}
