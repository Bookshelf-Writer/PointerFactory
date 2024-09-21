package PointerFactory

import (
	"context"
	"regexp"
	"time"
)

////////////////////////////////////

type chObj struct {
	group     rune
	retOffset chan uint32
}

type GlobalObj struct {
	startPoint time.Time
	cluster    uint16

	minute    uint32
	groups    map[rune]uint32
	groupsBuf []rune
	ch        chan *chObj

	ctx       context.Context
	ctxCancel context.CancelFunc
}

////

func New(groups []rune, cluster uint16, startPoint time.Time) (*GlobalObj, error) {
	if len(groups) == 0 {
		return nil, ErrEmpyGroups
	}
	if startPoint.Unix() >= time.Now().Unix() {
		return nil, ErrInvalidStartPoint
	}

	obj := GlobalObj{}
	obj.startPoint = startPoint
	obj.cluster = cluster

	obj.groups = make(map[rune]uint32)
	obj.ctx, obj.ctxCancel = context.WithCancel(context.Background())

	for _, group := range groups {
		match, _ := regexp.MatchString("[a-z0-9]", string(group))
		if !match {
			return nil, ErrInvalidGroupElement
		}

		obj.groups[group] = 0
		obj.groupsBuf = append(obj.groupsBuf, group)
	}

	duration := obj.timeNow().Sub(obj.startPoint)
	obj.minute = uint32(duration.Minutes())

	go obj.loop()

	return &obj, nil
}

func (obj *GlobalObj) Close() {
	obj.ctxCancel()
}
