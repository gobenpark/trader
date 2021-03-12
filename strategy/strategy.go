package strategy

import (
	"context"
	"fmt"

	"github.com/gobenpark/trader/event"
)

type StrategyEngine struct {
	E chan event.Event
}

func (s *StrategyEngine) eventListener() {
	for e := range s.E {
		fmt.Println(e)
	}
}

func (*StrategyEngine) Start(ctx context.Context) {

}