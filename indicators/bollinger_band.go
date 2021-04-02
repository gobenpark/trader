/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package indicators

import (
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

func (b *BollingerBand) Calculate(c container.Container) {
	con := c.Values()
	if len(con) < b.period {
		return
	}

	slice := len(con) - b.period
	for i := slice - 1; i >= 0; i-- {
		mean := b.mean(con[i : i+b.period])
		sd := b.standardDeviation(mean, con[i:i+b.period])

		b.Mid = append([]Indicate{{
			Data: mean,
			Date: con[i].Date,
		}}, b.Mid...)

		b.Top = append([]Indicate{{
			Data: mean + (sd * 2),
			Date: con[i].Date,
		}}, b.Top...)

		b.Bottom = append([]Indicate{{
			Data: mean - (sd * 2),
			Date: con[i].Date,
		}}, b.Bottom...)
	}

}

func (b *BollingerBand) Get() []Indicate {
	panic("implement me")
}
