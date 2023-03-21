/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-01 20:53
**/

package utils

import (
	"time"
)

type ti int

const Time ti = iota

const ymd = "2006-01-02"
const hms = "15:04:05"
const full = "2006-01-02 15:04:05"

const year = 1
const month = 2
const day = 3
const hour = 4
const minute = 5
const second = 6

type date struct {
	err  error
	time time.Time
}

type tickerInfo struct {
	duration time.Duration
	fn       func()
	ticker   *time.Ticker
}

type timeInfo struct {
	flag int
	time time.Time
}

func (d date) Format(format string) string {
	return d.time.Format(format)
}

func (d date) Time() time.Time {
	return d.time
}

func (d date) Second() timeInfo {
	return timeInfo{time: d.time, flag: second}
}

func (d date) Minute() timeInfo {
	return timeInfo{time: d.time, flag: minute}
}

func (d date) Hour() timeInfo {
	return timeInfo{time: d.time, flag: hour}
}

func (d date) Day() timeInfo {
	return timeInfo{time: d.time, flag: day}
}

func (d date) Month() timeInfo {
	return timeInfo{time: d.time, flag: month}
}

func (d date) Year() timeInfo {
	return timeInfo{time: d.time, flag: year}
}

func (d date) Error() error {
	return d.err
}

func (t timeInfo) Get() int {
	switch t.flag {
	case second:
		return t.time.Second()
	case minute:
		return t.time.Minute()
	case hour:
		return t.time.Hour()
	case day:
		return t.time.Day()
	case month:
		return int(t.time.Month())
	case year:
		return t.time.Year()
	default:
		return 0
	}
}

func (t timeInfo) Begin() int64 {
	switch t.flag {
	case second:
		return t.time.Unix()
	case minute:
		return t.time.Unix() - int64(t.time.Second())
	case hour:
		return t.time.Unix() - int64(t.time.Second()+t.time.Minute()*60)
	case day:
		return t.time.Unix() - int64(t.time.Second()+t.time.Minute()*60+t.time.Hour()*60*60)
	case month:
		return time.Date(t.time.Year(), t.time.Month(), 1, 0, 0, 0, 0, t.time.Location()).Unix()
	case year:
		return time.Date(t.time.Year(), 1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	default:
		return 0
	}
}

func (t timeInfo) End() int64 {
	switch t.flag {
	case second:
		return t.time.Unix()
	case minute:
		return t.Begin() + 60
	case hour:
		return t.Begin() + 3600
	case day:
		return t.Begin() + 86400
	case month:
		return time.Date(t.time.Year(), t.time.Month()+1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	case year:
		return time.Date(t.time.Year()+1, 1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	default:
		return 0
	}
}

func (ti ti) New() date {
	return date{time: time.Now()}
}

func (ti ti) Time(t time.Time) date {
	return date{time: t}
}

func (ti ti) Timestamp(timestamp int64) date {
	return date{time: time.Unix(timestamp, 0)}
}

func (ti ti) Format(format, dateString string) date {
	var t, err = time.ParseInLocation(format, dateString, time.Local)
	return date{time: t, err: err}
}

func (ti ti) String(dateString string) date {
	var t, err = time.ParseInLocation(full, dateString, time.Local)
	return date{time: t, err: err}
}

func (ti ti) Ticker(duration time.Duration, fn func()) *tickerInfo {
	return &tickerInfo{fn: fn, duration: duration}
}

func (ticker *tickerInfo) Start() {
	ticker.ticker = time.NewTicker(ticker.duration)
	go func() {
		for range ticker.ticker.C {
			ticker.fn()
		}
	}()
}

func (ticker *tickerInfo) Stop() {
	ticker.ticker.Stop()
}
