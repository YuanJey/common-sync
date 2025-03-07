package tables

type TbLasUserIncrement struct {
	ThirdCompanyID   string `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"`                    // 三方租户id
	PlatformID       string `gorm:"column:platform_id;type:varchar(60);not null;uniqueIndex:idx_platform_uid" json:"platform_id"` // 平台id，用来区分多种数据源
	UID              string `gorm:"column:uid;type:varchar(255);not null;uniqueIndex:idx_platform_uid" json:"uid"`                // 三方用户id
	DefDID           string `gorm:"column:def_did;type:varchar(255);default:null" json:"def_did"`                                 // 默认三方部门id
	DefDIDOrder      int    `gorm:"column:def_did_order;type:int;default:0" json:"def_did_order"`                                 // 用户在默认部门下的排序
	Account          string `gorm:"column:account;type:varchar(255);default:null" json:"account"`                                 // 登录名
	NickName         string `gorm:"column:nick_name;type:varchar(255);default:null" json:"nick_name"`                             // 用户昵称
	Password         string `gorm:"column:password;type:varchar(255);default:null" json:"password"`                               // 密码
	Avatar           string `gorm:"column:avatar;type:varchar(255);default:null" json:"avatar"`                                   // 头像
	Email            string `gorm:"column:email;type:varchar(80);default:null" json:"email"`                                      // 邮箱
	Gender           string `gorm:"column:gender;type:varchar(60);default:null" json:"gender"`                                    // 性别，枚举: secrecy,male,female
	Title            string `gorm:"column:title;type:varchar(255);default:null" json:"title"`                                     // 职称
	WorkPlace        string `gorm:"column:work_place;type:varchar(255);default:null" json:"work_place"`                           // 办公地点
	Leader           string `gorm:"column:leader;type:varchar(255);default:null" json:"leader"`                                   // 上级主管ID
	Employer         string `gorm:"column:employer;type:varchar(60);default:null" json:"employer"`                                // 员工工号
	EmploymentStatus string `gorm:"column:employment_status;type:varchar(60);default:null" json:"employment_status"`              // 就职状态[active, notactive, disabled]
	EmploymentType   string `gorm:"column:employment_type;type:varchar(60);default:null" json:"employment_type"`                  // 就职类型，枚举: unknow, permanent, intern
	Phone            string `gorm:"column:phone;type:varchar(60);default:null" json:"phone"`                                      // 手机号
	Telephone        string `gorm:"column:telephone;type:varchar(60);default:null" json:"telephone"`                              // 座机号
	Source           string `gorm:"column:source;type:varchar(20);default:'sync'" json:"source"`                                  // 来源，buildin/sync
	CustomFields     string `gorm:"column:custom_fields;type:varchar(5000);default:null" json:"custom_fields"`                    // 自定义字段，json数组
	UpdateType       string `gorm:"column:update_type;type:varchar(20);not null" json:"update_type"`                              // 修改类型, user_del/user_update/user_add
	//CreationTime     time.Time `gorm:"column:ctime;type:timestamp" json:"ctime"`                                                     // 创建时间
	//ModificationTime time.Time `gorm:"column:mtime;type:timestamp" json:"mtime"`                                                     // 修改时间
}

func (TbLasUserIncrement) TableName() string {
	return "tb_las_user_increment"
}
