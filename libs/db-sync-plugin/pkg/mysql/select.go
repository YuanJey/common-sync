package mysql

import "pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/mysql/tables"

func FindTbLasUserIncrementDeduplication(currentPage int, pageSize int) ([]tables.TbLasUserIncrement, error) {
	var list []tables.TbLasUserIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		Select("DISTINCT uid, *"). // 根据 uid 去重
		//Distinct("uid").
		Order("ctime ASC"). // 先按 uid 排序，然后按 ctime 排序
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error
	return list, err
}
func FindTbLasUserIncrement(currentPage int, pageSize int) ([]tables.TbLasUserIncrement, error) {
	var list []tables.TbLasUserIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		//Where("status = ? or status = ?", 1, 2).
		Limit(pageSize).
		Offset(offset).
		Order("ctime ASC").
		Find(&list).Error
	return list, err
}

func FirstTbLasUserIncrement(uid string) (count int64, err error) {
	err = Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		Where("(update_type =? or update_type = ? ) and status in(?,?) and uid = ?", "user_add", "user_update", 1, 0, uid).
		Count(&count).Error
	return count, err
}

func FirstTbLasUserIncrementUpdateType(uid string) (updateType string, err error) {
	err = Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		Select("update_type").
		Where("uid = ?", uid).
		Order("ctime desc").
		Limit(1).
		Row().
		Scan(&updateType)
	return updateType, err
}

func FindTbLasDepartmentUserIncrement(currentPage int, pageSize int) ([]tables.TbLasDepartmentUserIncrement, error) {
	var list []tables.TbLasDepartmentUserIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasDepartmentUserIncrement{}.TableName()).
		//Where("status = ? or status = ?", 1, 2).
		Limit(pageSize).
		Offset(offset).
		Order("ctime ASC").
		Find(&list).Error
	return list, err
}

func FindTbLasDepartmentIncrement(currentPage int, pageSize int) ([]tables.TbLasDepartmentIncrement, error) {
	var list []tables.TbLasDepartmentIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		//Where("status = ? or status = ?", 1, 2).
		Limit(pageSize).
		Offset(offset).
		Order("ctime ASC").
		Find(&list).Error
	return list, err
}
func FindTbLasDepartmentIncrementDeduplication(currentPage int, pageSize int) ([]tables.TbLasDepartmentIncrement, error) {
	var list []tables.TbLasDepartmentIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		//Distinct("did").
		Select("DISTINCT did, *"). // 根据 uid 去重
		Order("ctime ASC").        // 先按 uid 排序，然后按 ctime 排序
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error
	return list, err
}
func FindTbLasUserIncrementByUidAndUpdateType(currentPage, pageSize, status int, uid, updateType string) ([]tables.TbLasUserIncrement, error) {
	var list []tables.TbLasUserIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		Where("uid = ? and update_type = ? and status = ?", uid, updateType, status).
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error
	return list, err
}
func FindTbLasDepartmentIncrementByDidAndUpdateType(currentPage, pageSize, status int, did, updateType string) ([]tables.TbLasDepartmentIncrement, error) {
	var list []tables.TbLasDepartmentIncrement
	offset := (currentPage - 1) * pageSize
	err := Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		Where("did = ? and update_type = ? and status = ?", did, updateType, status).
		Limit(pageSize).
		Offset(offset).
		Find(&list).Error
	return list, err
}
func FirstTbLasDeptIncrement(did string) (count int64, err error) {
	err = Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		Where("(update_type =? or update_type = ? or update_type =?) and status in(?,?) and did = ?", "dept_add", "dept_update", "dept_move", 0, 1, did).
		Count(&count).Error
	return count, err
}
func FirstTbLasDeptIncrementUpdateType(did string) (updateType string, err error) {
	err = Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		Select("update_type").
		Where("did = ?", did).
		Order("ctime desc").
		Limit(1).
		Row().
		Scan(&updateType)
	return updateType, err
}
func FindLatestTbLasUserIncrement(uid string) (tables.TbLasUserIncrement, error) {
	var action tables.TbLasUserIncrement
	err := Database.db.Table(tables.TbLasUserIncrement{}.TableName()).
		Where("uid = ?", uid).
		Order("ctime desc").
		Limit(1).
		Find(&action).Error
	return action, err
}
func FindLatestTbLasDepartmentIncrement(did string) (tables.TbLasDepartmentIncrement, error) {
	var action tables.TbLasDepartmentIncrement
	err := Database.db.Table(tables.TbLasDepartmentIncrement{}.TableName()).
		Where("did = ?", did).
		Order("ctime desc").
		Limit(1).
		Find(&action).Error
	return action, err
}
func FindTbLasDepartmentUserIncrementByUid(uid string) ([]tables.TbLasDepartmentUserIncrement, error) {
	var list []tables.TbLasDepartmentUserIncrement
	err := Database.db.Table(tables.TbLasDepartmentUserIncrement{}.TableName()).
		Where("uid = ?", uid).
		Find(&list).Error
	return list, err
}
