package hardes1

import (
	"context"

	"hardes/internal/svc"
	"hardes/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfLogic {
	return &GetConfLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfLogic) GetConf(req *types.GetConfReq) (resp *types.GetConfResp, err error) {
	// todo: add your logic here and delete this line

	return
}
