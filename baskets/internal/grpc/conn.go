package grpc

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/stackus/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Dial(ctx context.Context, endpoint string) (conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(
		endpoint,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			if err = conn.Close(); err != nil {
				log.Err(errors.Wrap(err, "failed dial grpc conn"))
			}
		}

		go func() {
			<-ctx.Done()
			if err = conn.Close(); err != nil {
				log.Err(errors.Wrap(err, "failed dial grpc conn"))
			}
		}()
	}()

	return
}
