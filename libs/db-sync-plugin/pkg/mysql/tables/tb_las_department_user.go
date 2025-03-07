package tables

type TbLasDepartmentUser struct {
	TaskID         string `gorm:"column:task_id;type:varchar(20);not null" json:"task_id"`                   // 任务id
	ThirdCompanyID string `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"` // 三方租户id
	PlatformID     string `gorm:"column:platform_id;type:varchar(60);not null" json:"platform_id"`           // 平台id
	UID            string `gorm:"column:uid;type:varchar(255);not null" json:"uid"`                          // 三方用户id
	DID            string `gorm:"column:did;type:varchar(255);not null" json:"did"`                          // 三方部门id
	Order          int    `gorm:"column:order;type:int;default:0" json:"order"`                              // 用户在部门下的排序
	Main           int    `gorm:"column:main;type:int;default:0" json:"main"`                                // 是否是主部门
}

func (TbLasDepartmentUser) TableName() string {
	return "tb_las_department_user"
}
