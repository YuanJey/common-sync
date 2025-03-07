package tables

type TbLasDepartmentUserIncrement struct {
	ThirdCompanyID string `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"`                        // 三方租户id
	PlatformID     string `gorm:"column:platform_id;type:varchar(60);not null;uniqueIndex:idx_platform_did_uid" json:"platform_id"` // 平台id，用来区分多种数据源，platform_id + did 唯一
	UID            string `gorm:"column:uid;type:varchar(255);not null" json:"uid"`                                                 // 三方用户id
	DID            string `gorm:"column:did;type:varchar(255);not null" json:"did"`                                                 // 三方部门id
	DIDs           string `gorm:"column:did;type:varchar(2000);default:null" json:"dids"`                                           // JSONArray, 形如：[{"did": 1, "order": 1}] 仅当update_type=user_dept_move时生效
	Order          int    `gorm:"column:order;type:int;default:0" json:"order"`                                                     // 用户在部门下的排序
	Main           int    `gorm:"column:main;type:int;default:0" json:"main"`                                                       // 是否是主部门，1：是，0：不是
	UpdateType     string `gorm:"column:update_type;type:varchar(20);not null" json:"update_type"`                                  // 修改类型, user_dept_add/user_dept_del/user_dept_update/user_dept_move
	//CreationTime     time.Time `gorm:"column:ctime;type:timestamp" json:"ctime"`                                                         // 创建时间
	//ModificationTime time.Time `gorm:"column:mtime;type:timestamp" json:"mtime"`                                                         // 修改时间
}

func (TbLasDepartmentUserIncrement) TableName() string {
	return "tb_las_department_user_increment"
}
