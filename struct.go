package PointerFactory

import (
	"context"
	"regexp"
	"sync/atomic"
	"time"
)

////////////////////////////////////

type GlobalObj struct {
	startPoint time.Time
	cluster    uint16
	base       int32

	isActive bool
	minute   atomic.Uint32
	groups   map[rune]*atomic.Uint32

	ctx       context.Context
	ctxCancel context.CancelFunc
}

////

func New(groups []rune, cluster uint16, base int32, startPoint time.Time) (*GlobalObj, error) {
	if len(groups) == 0 {
		return nil, ErrEmpyGroups
	}
	if base < 2 || base > 36 {
		return nil, ErrInvalidBase
	}
	if startPoint.Unix() >= time.Now().Unix() {
		return nil, ErrInvalidStartPoint
	}

	obj := GlobalObj{}
	obj.startPoint = startPoint
	obj.cluster = cluster
	obj.base = base

	obj.groups = make(map[rune]*atomic.Uint32)

	for _, group := range groups {
		match, _ := regexp.MatchString("[a-z0-9]", string(group))
		if !match {
			return nil, ErrInvalidGroupElement
		}

		obj.groups[group] = &atomic.Uint32{}
	}

	obj.ctx, obj.ctxCancel = context.WithCancel(context.Background())
	duration := obj.timeNow().Sub(obj.startPoint)
	obj.minute.Store(uint32(duration.Minutes()))

	go obj.loop()
	for !obj.isActive {
		time.Sleep(10 * time.Millisecond)
	}

	return &obj, nil
}

func (obj *GlobalObj) Close() {
	obj.ctxCancel()
}

//

func (obj *GlobalObj) IsActive() bool {
	return obj.isActive
}
