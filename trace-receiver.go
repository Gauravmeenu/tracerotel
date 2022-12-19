package tracerotel

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.uber.org/zap"
)

type tracerotelReceiver struct{
	host component.Host
	cancel context.CancelFunc
	logger *zap.Logger
	nextConsumer consumer.Traces
	config *Config

}

func (tracerotelRcvr *tracerotelReceiver)Start(ctx context.Context, host component.Host)error{
	tracerotelRcvr.host=host
	ctx=context.Background()
	ctx, tracerotelRcvr.cancel=context.WithCancel(ctx)
	interval, _ := time.ParseDuration(tracerotelRcvr.config.Interval)
	go func(){
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for{
			select{
			case <-ticker.C:
				tracerotelRcvr.logger.Info("I should start processing now!!")
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (tracerotelRcvr *tracerotelReceiver)Shutdown(ctx context.Context)error{
	tracerotelRcvr.cancel()
	return nil
}