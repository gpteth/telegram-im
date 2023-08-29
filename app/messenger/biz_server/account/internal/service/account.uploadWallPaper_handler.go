package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountUploadWallPaper(ctx context.Context, request *mtproto.TLAccountUploadWallPaper) (*mtproto.WallPaper, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.uploadWallPaper - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	err := mtproto.ErrMethodNotImpl
	return nil, err
}
