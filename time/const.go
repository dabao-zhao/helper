// Copyright (c) 2020 Qi Yin <qiyin@thinkeridea.com>

package xtime

import "time"

const (
	Nanosecond time.Duration = 1

	Microsecond = 1000 * Nanosecond

	Millisecond = 1000 * Microsecond

	Second = 1000 * Millisecond

	Minute = 60 * Second

	Hour = 60 * Minute

	Day = 24 * Hour

	Week = 7 * Day
)
