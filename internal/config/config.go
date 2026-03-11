// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	ShortUrlDB ShortUrlDB
	Sequence   struct { // 匿名结构体
		DSN string
	}
	BaseString        string   // base62指定的基础字符串
	ShortUrlBlackList []string // 黑名单
	ShortDomain       string   // 短链接域名
}

type ShortUrlDB struct { // 命名结构体
	DSN string
}
