package tracerotel

import (
	"context"
	"strconv"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
)

const (
	typeStr         = "tracerotel"
	defaultInterval = 1
)

func CreateDefaultConfig() component.ReceiverConfig {

	return &Config{
		ReceiverSettings: config.NewReceiverSettings(component.NewID(typeStr)),
		Interval:         strconv.Itoa(defaultInterval),
	}
}

// func CreateTracesReceiver(_ context.Context, params component.ReceiverCreateSettings, bcfg Config,
//
//	nextConsumer Component.consumer.Traces) (component.TracesReceiver, error){
//		return nil, nil
//	}
func createTracesReceiver(s context.Context, p component.ReceiverCreateSettings, baseCfg component.ReceiverConfig, consumer consumer.Traces) (component.TracesReceiver, error) {
	if consumer == nil{
		return nil, component.ErrNilNextConsumer
	}
	logger := p.Logger
	tracerotelCfg := baseCfg.(*Config)

	traceRcvr := &tracerotelReceiver{
		logger: logger,
		nextConsumer: consumer,
		config: tracerotelCfg,
	}

	return traceRcvr, nil
}

func NewFactory() component.ReceiverFactory {
	return component.NewReceiverFactory(
		typeStr,
		CreateDefaultConfig,
		component.WithTracesReceiver(createTracesReceiver, component.StabilityLevelStable),
	)

}