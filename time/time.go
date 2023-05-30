package time

import "time"

func Midnight() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
}

// ParseInLocal
// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>
func ParseInLocal(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, time.Local)
}
