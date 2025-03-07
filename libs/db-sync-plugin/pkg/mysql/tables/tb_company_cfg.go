package tables

type TbCompanyCfg struct {
	ThirdCompanyID string `gorm:"column:third_company_id;type:varchar(20);not null;unique" json:"third_company_id"` // 三方租户id，要求唯一且不变
	PlatformIDs    string `gorm:"column:platform_ids;type:varchar(100);not null;unique" json:"platform_ids"`        // 客户数据源id，要求唯一且不变
	CompanyID      string `gorm:"column:company_id;type:varchar(20);not null" json:"company_id"`                    // 云文档租户id
	Status         int    `gorm:"column:status;type:tinyint(4);default:1" json:"status"`                            // 0-禁用 1-启用
}

func (TbCompanyCfg) TableName() string {
	return "tb_company_cfg"
}
