// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"workspace/internal/svc"
	"workspace/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertLogic {
	return &ConvertLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
// Convert 转链：输入一个长链接 ---> 短链接
func (l *ConvertLogic) Convert(req *types.ConvertRequest) (resp *types.ConvertResponse, err error) {
	// 1. 校验输入的数据
	// 1.1 数据不能为空，使用validator校验库

	// 1.2 输入的长链接能ping通，是有效的网址
	// 1.3 判断数据库是否已经转链过（数据库中是否已经存在该长链）
	// 1.3.1 给长链接生成md5值
	// 1.3.2 拿md5去数据库中查是否存在,如果不存在，继续进行转链
	// 1.4 输入的不能是一个(数据库中已存在的)短链接（避免循环转链）
	// 2. 取号
	// 3. 号码转链
	// 4. 存储长短链接映射关系
	// 5. 返回响应
	return
}
