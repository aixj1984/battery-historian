package myredis

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"channel"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

// Cache is Redis cache adapter.
type Cache struct {
	p        *redis.Pool // redis connection pool
	conninfo string
	dbNum    int
	key      string
	password string
}

var Cmd string

//JointCmd joint to cmd
func JointCmd(cmd string, key ...interface{}) string {
	for i := 0; i < len(key); i++ {
		cmd = cmd + " " + key[i].(string)
	}
	beego.Debug(cmd)
	return cmd
}

// actually do the redis cmds
func (rc *Cache) do(commandName, userId string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()
	Cmd = JointCmd(commandName, args...)
	channel.SendToChannel(userId, Cmd)
	return c.Do(commandName, args...)
}

// Get cache from redis.
func (rc *Cache) GetAll(key, userId string) interface{} {
	if v, err := rc.do("HGETALL", userId, key); err == nil {

		return v
	}
	return nil
}
func (rc *Cache) GetKey(key, userId string) interface{} {
	if v, err := rc.do("HKEYS", userId, key); err == nil {
		return v
	}
	return nil
}

// Put put cache to redis.
func (rc *Cache) Put(key, userId string, val interface{}) error {
	var err error
	value := val.([]interface{})
	if len(value) > 3 {
		if _, err = rc.do("HMSET", userId, value...); err != nil {
			return err
		}
	} else {
		if _, err = rc.do("HSET", userId, value...); err != nil {
			return err
		}
	}
	if _, err = rc.do("HSET", userId, key, value[0], userId); err != nil {
		return err
	}
	return err
}

// Delete delete cache in redis.
//func (rc *Cache) Delete(key string) error {
//	var err error
//	if _, err = rc.do("DEL", key); err != nil {
//		return err
//	}
//	_, err = rc.do("HDEL", rc.key, key)
//	return err
//}

// StartAndGC start redis cache adapter.
// config is like {"key":"collection key","conn":"connection info","dbNum":"0"}
// the cache item in redis are stored forever,
// so no gc operation.
func (rc *Cache) StartAndGC(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = "myredis"
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}
	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	rc.key = cf["key"]
	rc.conninfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (rc *Cache) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.conninfo)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selecterr := c.Do("SELECT", rc.dbNum)
		if selecterr != nil {
			c.Close()
			return nil, selecterr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}
