/*
 * Copyright (c) 2021. Lorem ipsum dolor sit amet, consectetur adipiscing elit.
 * Morbi non lorem porttitor neque feugiat blandit. Ut vitae ipsum eget quam lacinia accumsan.
 * Etiam sed turpis ac ipsum condimentum fringilla. Maecenas magna.
 * Proin dapibus sapien vel ante. Aliquam erat volutpat. Pellentesque sagittis ligula eget metus.
 * Vestibulum commodo. Ut rhoncus gravida arcu.
 */
package container

import (
	"sync"
	"time"
)

type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []Candle
	Add(candle Candle)
	Code() string
	Level() time.Duration
}

type SaveMode int

const (
	InMemory = iota
	External
)

type Info struct {
	Code             string
	CompressionLevel time.Duration
}

//TODO: inmemory or external storage
type DataContainer struct {
	mu         sync.RWMutex
	CandleData []Candle
	Info
}

func NewDataContainer(info Info) *DataContainer {
	return &DataContainer{
		CandleData: []Candle{},
		Info:       info,
	}
}

func (t *DataContainer) Empty() bool {
	return len(t.CandleData) == 0
}

func (t *DataContainer) Size() int {
	l := 0
	t.mu.RLock()
	l = len(t.CandleData)
	t.mu.RUnlock()
	return l
}

func (t *DataContainer) Clear() {
	t.CandleData = []Candle{}
}

func (t *DataContainer) Values() []Candle {
	t.mu.Lock()
	d := make([]Candle, len(t.CandleData))
	copy(d, t.CandleData)
	t.mu.Unlock()
	return d
}

// Add foreword append container candle data
// current candle [0] index
func (t *DataContainer) Add(candle Candle) {
	if len(t.CandleData) != 0 {
		for _, i := range t.CandleData {
			if i.Date.Equal(candle.Date) {
				return
			}
		}
	}
	t.mu.Lock()
	t.CandleData = append([]Candle{candle}, t.CandleData...)
	t.mu.Unlock()
}

func (t *DataContainer) Code() string {
	return t.Info.Code
}

func (t *DataContainer) Level() time.Duration {
	return t.Info.CompressionLevel
}
