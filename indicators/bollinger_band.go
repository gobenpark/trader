package indicators

import (
	"fmt"
	"math"

	"github.com/gobenpark/trader/container"
)

type BollingerBand struct {
	period int
	Top    []Indicate
	Mid    []Indicate
	Bottom []Indicate
}

func NewBollingerBand(period int) *BollingerBand {
	return &BollingerBand{period: period}
}

func (b *BollingerBand) mean(data []container.Candle) float64 {
	total := 0.0
	for _, i := range data {
		total += i.Close
	}

	return total / float64(len(data))
}

func (b *BollingerBand) standardDeviation(mean float64, data []container.Candle) float64 {
	total := 0.0
	for _, i := range data {
		da := i.Close - mean
		total += math.Pow(da, 2)
	}
	return math.Sqrt(total / float64(len(data)))
}

func (b *BollingerBand) Calculate(container container.Container) {
	c := container.Values()
	if len(c) < b.period {
		return
	}

	slice := len(c) - b.period
	for i := slice - 1; i >= 0; i-- {
		fmt.Println(len(c[i : i+b.period]))
		mean := b.mean(c[i : i+b.period])
		sd := b.standardDeviation(mean, c[i:i+b.period])

		b.Mid = append([]Indicate{{
			Data: mean,
			Date: c[i].Date,
		}}, b.Mid...)

		b.Top = append([]Indicate{{
			Data: mean + (sd * 2),
			Date: c[i].Date,
		}}, b.Top...)

		b.Bottom = append([]Indicate{{
			Data: mean - (sd * 2),
			Date: c[i].Date,
		}}, b.Bottom...)
	}

}

func (b *BollingerBand) Get() []Indicate {
	panic("implement me")
}
