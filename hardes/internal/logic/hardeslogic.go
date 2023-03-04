package logic

import (
	"context"

	"hardes/internal/svc"
	"hardes/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HardesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHardesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HardesLogic {
	return &HardesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HardesLogic) Hardes(req *types.SaveReq) (resp *types.SaveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
