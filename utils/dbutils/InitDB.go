package dbutils

import (
	"github.com/sunmfei/mfu/common/MFei"
	"github.com/sunmfei/mfu/common/sun"
	"github.com/sunmfei/mfu/utils/dbutils/redisUtil"
)

// GetDB 开放给外部获得db连接
func GetDB(typ string) {
	MFei.DB = setup(typ)
	sun.Redis = redisUtil.Setup()
}
