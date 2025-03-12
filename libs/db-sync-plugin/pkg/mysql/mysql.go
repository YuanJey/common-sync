package mysql

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"pl.ghgame.cn/gitea/yuanjie/db-sync-plugin/pkg/aksk"
	"strconv"
	"time"
)

var Database mysqlDB

type mysqlDB struct {
	db *gorm.DB
}
type Writer struct{}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func initDatabaseDsn(dsn string) {
	var db *gorm.DB
	var err1 error
	db, err := gorm.Open(mysql.Open(dsn), nil)
	if err != nil {
		time.Sleep(time.Duration(30) * time.Second)
		db, err1 = gorm.Open(mysql.Open(dsn), nil)
		if err1 != nil {
			panic(err1.Error() + " open failed " + dsn)
		}
	}
	newLogger := logger.New(
		Writer{},
		logger.Config{
			SlowThreshold:             time.Duration(500) * time.Millisecond, // Slow SQL threshold
			LogLevel:                  logger.LogLevel(1),                    // Log level
			IgnoreRecordNotFoundError: true,                                  // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                                  // Disable color
		},
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err.Error() + " Open failed " + dsn)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err.Error() + " db.DB() failed ")
	}

	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(5))
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	//err = db.AutoMigrate(&tables.TempUser{}, &tables.TempDept{})
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	Database.db = db
}

func init() {
	dbKey := os.Getenv("ECIS_ECISACCOUNTSYNC_DB")
	sk := ""
	if os.Getenv("SK_ENC") == "true" {
		ak, news := aksk.GetAkSK(os.Getenv("SK_DEC_KEY"))
		sk = news
		fmt.Println("ak sk is ", ak, sk)
	} else {
		sk = os.Getenv("SK")
	}
	dsn, err := decryptByAes(dbKey, sk)
	if err != nil {
		panic(err.Error())
	}
	initDatabaseDsn(dsn)
	//initTbCompanyCfg()
	cbs, _ := strconv.Atoi(os.Getenv("C_B_S"))
	if cbs > 0 {
		Database.db.Config.CreateBatchSize = cbs
	} else {
		Database.db.Config.CreateBatchSize = 1000
	}
}
func decryptByAes(enData, key string) (string, error) {
	// 应用SK
	//key := config.APIConfig.Env.SK
	h := md5.New()
	h.Write([]byte(key))
	akey := hex.EncodeToString(h.Sum(nil))

	enDataFromBase64, err := base64.StdEncoding.DecodeString(enData)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(akey))
	if err != nil {
		return "", err
	}
	iv := []byte(akey)[:aes.BlockSize]
	decrypter := cipher.NewCBCDecrypter(block, iv)

	dst := make([]byte, len(enDataFromBase64))
	decrypter.CryptBlocks(dst, enDataFromBase64)

	length := len(dst)
	unpadding := int(dst[length-1])
	if length-unpadding >= 0 {
		dst = dst[:(length - unpadding)]
	}
	plain := string(dst)
	return plain, nil
}
