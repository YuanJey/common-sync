package tables

type TbLasUser struct {
	TaskID           string  `gorm:"column:task_id;type:varchar(20);not null" json:"task_id"`
	ThirdCompanyID   string  `gorm:"column:third_company_id;type:varchar(20);not null" json:"third_company_id"`
	PlatformID       string  `gorm:"column:platform_id;type:varchar(60);not null" json:"platform_id"`
	UID              string  `gorm:"column:uid;type:varchar(255);not null" json:"uid"`
	DefDID           *string `gorm:"column:def_did;type:varchar(255);default:null" json:"def_did,omitempty"`
	DefDIDOrder      int     `gorm:"column:def_did_order;type:int;default:0" json:"def_did_order"`
	Account          string  `gorm:"column:account;type:varchar(255);not null" json:"account"`
	NickName         string  `gorm:"column:nick_name;type:varchar(255);not null" json:"nick_name"`
	Password         *string `gorm:"column:password;type:varchar(255);default:null" json:"password,omitempty"`
	Avatar           *string `gorm:"column:avatar;type:varchar(255);default:null" json:"avatar,omitempty"`
	Email            *string `gorm:"column:email;type:varchar(80);default:null" json:"email,omitempty"`
	Gender           *string `gorm:"column:gender;type:varchar(60);default:null" json:"gender,omitempty"` // possible values: secrecy, male, female
	Title            *string `gorm:"column:title;type:varchar(255);default:null" json:"title,omitempty"`
	WorkPlace        *string `gorm:"column:work_place;type:varchar(255);default:null" json:"work_place,omitempty"`
	Leader           *string `gorm:"column:leader;type:varchar(255);default:null" json:"leader,omitempty"`
	Employer         *string `gorm:"column:employer;type:varchar(255);default:null" json:"employer,omitempty"`
	EmploymentStatus string  `gorm:"column:employment_status;type:varchar(60);default:'notactive'" json:"employment_status"` // possible values: active, notactive, disabled
	EmploymentType   *string `gorm:"column:employment_type;type:varchar(60);default:null" json:"employment_type,omitempty"`  // possible values: unknown, permanent, intern
	Phone            *string `gorm:"column:phone;type:varchar(60);default:null" json:"phone,omitempty"`
	Telephone        *string `gorm:"column:telephone;type:varchar(60);default:null" json:"telephone,omitempty"`
	Source           string  `gorm:"column:source;type:varchar(20);default:'sync'" json:"source"`                         // possible values: buildin, sync
	CustomFields     *string `gorm:"column:custom_fields;type:varchar(5000);default:null" json:"custom_fields,omitempty"` // JSON array
}

func (TbLasUser) TableName() string {
	return "tb_las_user"
}
