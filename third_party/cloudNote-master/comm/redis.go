package comm

import (
	"fmt"
	"strconv"

	"channel"
	"comm/mgodb"
	"comm/myredis"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

var (
	bm myredis.Cache
)

func init() {
	var err error
	redisIp := beego.AppConfig.String("redisIp")
	redisPort := beego.AppConfig.String("redisPort")
	connectInfo := redisIp + redisPort
	redis := `{"conn":"%[1]v","dbNum":"1"}`
	redisInfo := fmt.Sprintf(redis, connectInfo)
	fmt.Println(redisInfo)
	err = bm.StartAndGC(redisInfo)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for msg := range channel.Sendc {
			fmt.Println("Send msg:", msg)
			var operatelog mgodb.OperateLog
			operatelog.UserId = msg["userId"].(string)
			operatelog.RedisCmd = msg["cmd"].(string)
			operatelog.TimeStamp = msg["time"].(int)
			result := mgodb.AddRecord("cloudNote", "operate_log", operatelog)
			fmt.Println(result)
		}
	}()
}

//PutDataRedis put message to redis
func PutDataToRedis(name, key, userId string, val map[string]interface{}) error {
	var slice []interface{}
	slice = append(slice, name)
	for k, v := range val {
		switch v.(type) {
		case int:
			v = strconv.Itoa(v.(int))
		default:
		}
		if v.(string) != "" {
			slice = append(slice, k, v.(string))
		}

	}

	beego.Debug(slice)
	err := bm.Put(key, userId, slice)
	beego.Debug(name, err)
	return err
}

//GetDataAllFromRedis get value from redis
func GetDataAllFromRedis(key, userId string) interface{} {
	data := bm.GetAll(key, userId)
	return data
}

//GetDataKeyFromRedis get key from redis
func GetDataKeyFromRedis(key, userId string) []string {
	data := bm.GetKey(key, userId)
	result, err := redis.Strings(data, nil)
	if err != nil {
		return nil
	}
	return result
}
