package PointerFactory

import "time"

////////////////////////////////////

func GetTimeMin(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		0,
		0,
		0,
		t.Location(),
	)
}

//

func (obj *GlobalObj) timeNow() time.Time {
	return GetTimeMin(time.Now())
}
