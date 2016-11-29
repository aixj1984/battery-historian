package channel

import (
	"time"

	"github.com/astaxie/beego"
)

var channelNum, _ = beego.AppConfig.Int("channelNum")
var Sendc = make(chan map[string]interface{}, channelNum)

//SendToChanner send message to channel
func SendToChannel(userId, cmd string) {
	tm := int(time.Now().Unix())
	data := map[string]interface{}{
		"userId": userId,
		"cmd":    cmd,
		"time":   tm,
	}
	Sendc <- data
}
