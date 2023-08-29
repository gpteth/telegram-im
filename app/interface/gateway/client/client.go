package gateway_client

import (
	"context"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/interface/gateway/egatepb"
	"google.golang.org/grpc"
)

func NewClient(target string, cfg *warden.ClientConfig, opts ...grpc.DialOption) (egatepb.EGateClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), target, warden.WithDialLogFlag(warden.LogFlagDisableArgs))
	if err != nil {
		return nil, err
	}

	return egatepb.NewEGateClient(cc), nil
}
