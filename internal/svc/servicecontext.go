// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"workspace/internal/config"
	"workspace/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	ShortUrlMapModel model.ShortUrlMapModel // 接口类型。 代表了 short_url_map表
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.ShortUrlDB.DSN)
	return &ServiceContext{
		Config:           c,
		ShortUrlMapModel: model.NewShortUrlMapModel(conn),
	}
}
