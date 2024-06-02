/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-28 02:12:57
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-02 13:27:01
 * @FilePath     : /v2/go-common-v2-dh-oauth2-server/oauth2_test.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhoauth2server

import (
	"testing"

	dhlog "github.com/lepingbeta/go-common-v2-dh-log"
	dhredis "github.com/lepingbeta/go-common-v2-dh-redis"
)

func TestMakeAuthCode(t *testing.T) {
	var jsonStr = `{
		"host": "dev.cezhiqiu.cn",
		"port": 6379,
		"pass": "dahe_redis_888",
		"network": "tcp",
		"max-idle": 3,
		"idle-timeout": 240,
		"max-active": 0,
		"db": 6,
		"prefix": "redis-test"
	}`
	dhredis.InitRedis(jsonStr)
	SetTokenExpires(180, 180, 180)
	MakeAuthCode("123hhh")
}

func TestMakeTwoToken(t *testing.T) {
	var jsonStr = `{
		"host": "dev.cezhiqiu.cn",
		"port": 6379,
		"pass": "dahe_redis_888",
		"network": "tcp",
		"max-idle": 3,
		"idle-timeout": 240,
		"max-active": 0,
		"db": 6,
		"prefix": "redis-test"
	}`
	dhredis.InitRedis(jsonStr)
	SetTokenExpires(180, 180, 180)
	at, rt, _ := MakeTwoToken("ARvGXIgXsyllhJm83FByoQ==")
	dhlog.DebugAny(at)
	dhlog.DebugAny(rt)
}

func TestRefreshToken(t *testing.T) {
	userId := "test_user_id"
	var jsonStr = `{
		"host": "dev.cezhiqiu.cn",
		"port": 6379,
		"pass": "dahe_redis_888",
		"network": "tcp",
		"max-idle": 3,
		"idle-timeout": 240,
		"max-active": 0,
		"db": 6,
		"prefix": "redis-test"
	}`
	dhredis.InitRedis(jsonStr)
	SetTokenExpires(180, 180, 180)
	code, _ := MakeAuthCode(userId)
	_, rt, _ := MakeTwoToken(code)
	at, rt, _ := RefreshToken(rt)
	dhlog.DebugAny(at)
	dhlog.DebugAny(rt)
}
