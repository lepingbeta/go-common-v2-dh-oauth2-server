/*
 * @Author       : Symphony zhangleping@cezhiqiu.com
 * @Date         : 2024-05-28 02:12:57
 * @LastEditors  : Symphony zhangleping@cezhiqiu.com
 * @LastEditTime : 2024-06-02 13:22:47
 * @FilePath     : /inner-user-center-api/data/mycode/dahe/go-common/v2/go-common-v2-dh-oauth2-server/oauth2.go
 * @Description  :
 *
 * Copyright (c) 2024 by 大合前研, All Rights Reserved.
 */
package dhoauth2server

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	dhredis "github.com/lepingbeta/go-common-v2-dh-redis"
)

// TokenData 存储令牌的详细信息，包括三种不同的过期时间（以整数形式表示秒数）。
var oauth2Config struct {
	codeExpire         int
	accessTokenExpire  int
	refreshTokenExpire int
}

// SetTokenExpires 设置 TokenData 结构体中的过期时间字段。
// 参数 expireSeconds 表示过期时间，单位为秒。
func SetTokenExpires(codeExpire, accessTokenExpire, refreshTokenExpire int) {
	// 将当前时间加上过期秒数，转换为整数形式的过期时间
	oauth2Config.codeExpire = codeExpire
	oauth2Config.accessTokenExpire = accessTokenExpire
	oauth2Config.refreshTokenExpire = refreshTokenExpire
}

// GenerateAuthCode 生成一个随机的授权码
func GenerateAuthCode() (string, error) {
	const codeLength = 16
	bytes := make([]byte, codeLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func MakeAuthCode(userId string) (string, error) {
	code, err := GenerateAuthCode()
	if err != nil {
		return "", err
	}

	dhredis.Set(fmt.Sprintf("oauth2:code:%s", code), userId, oauth2Config.codeExpire)

	return code, nil
}

func MakeTwoToken(code string) (string, string, error) {
	userId := dhredis.Get(fmt.Sprintf("oauth2:code:%s", code))
	if len(userId) == 0 {
		return "", "", fmt.Errorf("code 没有匹配的userId")
	}

	at, err := GenerateAccessToken()
	if err != nil {
		return "", "", err
	}

	rt, err := GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}

	dhredis.Set(fmt.Sprintf("oauth2:access_token:%s", at), userId, oauth2Config.accessTokenExpire)
	dhredis.Set(fmt.Sprintf("oauth2:refresh_token:%s", rt), userId, oauth2Config.refreshTokenExpire)

	return at, rt, nil
}

// GenerateAccessToken 生成一个随机的令牌
func GenerateAccessToken() (string, error) {
	const tokenLength = 32
	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

// GenerateRefreshToken 生成一个随机的令牌
func GenerateRefreshToken() (string, error) {
	const tokenLength = 32
	bytes := make([]byte, tokenLength)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

func RefreshToken(rt string) (string, string, error) {
	userId := dhredis.Get(fmt.Sprintf("oauth2:refresh_token:%s", rt))
	if len(userId) == 0 {
		return "", "", fmt.Errorf("refresh_token 没有匹配的userId")
	}
	at, err := GenerateAccessToken()
	if err != nil {
		return "", "", err
	}

	rt, err = GenerateRefreshToken()
	if err != nil {
		return "", "", err
	}
	dhredis.Set(fmt.Sprintf("oauth2:access_token:%s", at), userId, oauth2Config.accessTokenExpire)
	dhredis.Set(fmt.Sprintf("oauth2:refresh_token:%s", rt), userId, oauth2Config.refreshTokenExpire)
	return at, rt, nil
}

func GetUserId(at string) (string, error) {
	userId := dhredis.Get(fmt.Sprintf("oauth2:access_token:%s", at))
	if len(userId) == 0 {
		return "", fmt.Errorf("access token 没有匹配的userId")
	}

	return userId, nil
}
