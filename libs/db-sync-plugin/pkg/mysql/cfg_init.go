package mysql

import (
	"fmt"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"
)

func InstallTbCompanyCfg(ThirdCompanyID, PlatformIDs, CompanyID string) {
	cfg := tables.TbCompanyCfg{ThirdCompanyID: ThirdCompanyID, PlatformIDs: PlatformIDs, CompanyID: CompanyID, Status: 1}
	tableName := cfg.TableName()
	err := Database.db.Table(tableName).Create(cfg).Error
	if err != nil {
		fmt.Println("tb_company_cfg create failed", err.Error())
	}
}
