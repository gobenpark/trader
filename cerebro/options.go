package cerebro

import (
	"time"

	"github.com/gobenpark/trader/broker"
	"github.com/gobenpark/trader/store"
	"github.com/gobenpark/trader/strategy"
	"github.com/rs/zerolog"
)

type Option func(*Cerebro)

func WithBroker(b broker.Broker) Option {
	return func(c *Cerebro) {
		c.broker = b
		c.broker.SetEventBroadCaster(c.eventEngine)
	}
}

func WithStrategy(s ...strategy.Strategy) Option {
	return func(c *Cerebro) {
		c.strategies = s
	}
}

func WithStore(s store.Store, codes ...string) Option {
	return func(c *Cerebro) {
		c.storengine.Stores[s.Uid()] = s
		c.storengine.Mapper[s.Uid()] = append(c.storengine.Mapper[s.Uid()], codes...)
	}
}

func WithLogLevel(level zerolog.Level) Option {
	return func(c *Cerebro) {
		c.log = c.log.Level(level)
	}
}

func WithLive(isLive bool) Option {
	return func(c *Cerebro) {
		c.isLive = isLive
	}
}

func WithResample(code string, level time.Duration, leftEdge bool) Option {
	return func(c *Cerebro) {
		c.compress[code] = append(c.compress[code], CompressInfo{level: level, LeftEdge: leftEdge})
	}
}

func WithPreload(b bool) Option {
	return func(c *Cerebro) {
		c.preload = b
	}
}
