package proc

import (
	"context"
	"fmt"
	"go-arch/util/config/yml"
	"go-arch/util/err"
	"go.uber.org/zap"
	"time"
)

func Main(ctx context.Context, config *yml.Config) error {
	for {
		select {
		case <-ctx.Done():
			return err.MainProcessContextTerminated
		default:
			zap.L().Info("hi",
				zap.String("config", fmt.Sprintf("%+v", config)))
			time.Sleep(1 * time.Second)
		}
	}
}
