package dtmsvr

import (
	"github.com/yedf/dtm/common"
	"github.com/yedf/dtm/dtmcli"
)

type dtmsvrConfig struct {
	TransCronInterval int64             `yaml:"TransCronInterval"` // 单位秒 当事务等待这个时间之后，还没有变化，则进行一轮处理，包括prepared中的任务和committed的任务
	DB                map[string]string `yaml:"DB"`
}

var config = &dtmsvrConfig{
	TransCronInterval: 10,
}

var dbName = "dtm"

func init() {
	common.InitConfig(dtmcli.GetProjectDir(), &config)
	config.DB["database"] = ""
}
