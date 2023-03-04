package hardes1

import (
	"context"

	"hardes/internal/svc"
	"hardes/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveLogic {
	return &SaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveLogic) Save(req *types.SaveReq) (resp *types.SaveResp, err error) {
	// todo: add your logic here and delete this line

	return
}
