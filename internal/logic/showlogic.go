// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"database/sql"
	"errors"

	"workspace/internal/svc"
	"workspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShowLogic {
	return &ShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShowLogic) Show(req *types.ShowRequest) (resp *types.ShowResponse, err error) {
	// 查看短链接，输入q1mi.cn/lucky ---> 重定向到真实的链接
	// req.ShortUrl = lucky
	// 1. 根据短链接查询原始的长链接
	// 1.1 查询数据库之前增加缓存层
	u, err := l.svcCtx.ShortUrlMapModel.FindOneBySurl(l.ctx, sql.NullString{Valid: true, String: req.ShortUrl})
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("404")
		}
		logx.Errorw("ShortUrlMapModel.FindOneBySurl error", logx.LogField{Value: err.Error(), Key: "error"})
		return nil, err
	}
	// 2. 返回查询到的长链接，在handler层返回重定向响应

	return &types.ShowResponse{LongUrl: u.Lurl.String}, nil
}
