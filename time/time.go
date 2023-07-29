/**
* @program: lemon
*
* @description:
*
* @author: lemon
*
* @create: 2019-11-01 20:53
**/

package time

import (
	"time"
)

const Ymd = "2006-01-02"
const Hms = "15:04:05"
const YmdHms = "2006-01-02 15:04:05"

const (
	Year = iota + 1
	Month
	Day
	Hour
	Minute
	Second
)

type Date struct {
	err  error
	time time.Time
}

type TickerInfo struct {
	duration time.Duration
	fn       func()
	ticker   *time.Ticker
}

type DateInfo struct {
	flag int
	time time.Time
}

func (d Date) Format(format string) string {
	return d.time.Format(format)
}

func (d Date) Time() time.Time {
	return d.time
}

func (d Date) Second() DateInfo {
	return DateInfo{time: d.time, flag: Second}
}

func (d Date) Minute() DateInfo {
	return DateInfo{time: d.time, flag: Minute}
}

func (d Date) Hour() DateInfo {
	return DateInfo{time: d.time, flag: Hour}
}

func (d Date) Day() DateInfo {
	return DateInfo{time: d.time, flag: Day}
}

func (d Date) Month() DateInfo {
	return DateInfo{time: d.time, flag: Month}
}

func (d Date) Year() DateInfo {
	return DateInfo{time: d.time, flag: Year}
}

func (d Date) Error() error {
	return d.err
}

func (t DateInfo) Get() int {
	switch t.flag {
	case Second:
		return t.time.Second()
	case Minute:
		return t.time.Minute()
	case Hour:
		return t.time.Hour()
	case Day:
		return t.time.Day()
	case Month:
		return int(t.time.Month())
	case Year:
		return t.time.Year()
	default:
		return 0
	}
}

func (t DateInfo) Begin() int64 {
	switch t.flag {
	case Second:
		return t.time.Unix()
	case Minute:
		return t.time.Unix() - int64(t.time.Second())
	case Hour:
		return t.time.Unix() - int64(t.time.Second()+t.time.Minute()*60)
	case Day:
		return t.time.Unix() - int64(t.time.Second()+t.time.Minute()*60+t.time.Hour()*60*60)
	case Month:
		return time.Date(t.time.Year(), t.time.Month(), 1, 0, 0, 0, 0, t.time.Location()).Unix()
	case Year:
		return time.Date(t.time.Year(), 1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	default:
		return 0
	}
}

func (t DateInfo) End() int64 {
	switch t.flag {
	case Second:
		return t.time.Unix()
	case Minute:
		return t.Begin() + 60
	case Hour:
		return t.Begin() + 3600
	case Day:
		return t.Begin() + 86400
	case Month:
		return time.Date(t.time.Year(), t.time.Month()+1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	case Year:
		return time.Date(t.time.Year()+1, 1, 1, 0, 0, 0, 0, t.time.Location()).Unix()
	default:
		return 0
	}
}

func New() Date {
	return Date{time: time.Now()}
}

func Time(t time.Time) Date {
	return Date{time: t}
}

func Timestamp(timestamp int64) Date {
	return Date{time: time.Unix(timestamp, 0)}
}

func Format(format, dateString string) Date {
	var t, err = time.ParseInLocation(format, dateString, time.Local)
	return Date{time: t, err: err}
}

func YMDHMSString(dateString string) Date {
	var t, err = time.ParseInLocation(YmdHms, dateString, time.Local)
	return Date{time: t, err: err}
}

func YMDString(dateString string) Date {
	var t, err = time.ParseInLocation(Ymd, dateString, time.Local)
	return Date{time: t, err: err}
}

func HMSString(dateString string) Date {
	var t, err = time.ParseInLocation(Hms, dateString, time.Local)
	return Date{time: t, err: err}
}

func Ticker(duration time.Duration, fn func()) *TickerInfo {
	return &TickerInfo{fn: fn, duration: duration}
}

func (ticker *TickerInfo) Start() {
	ticker.ticker = time.NewTicker(ticker.duration)
	go func() {
		for range ticker.ticker.C {
			ticker.fn()
		}
	}()
}

func (ticker *TickerInfo) Stop() {
	ticker.ticker.Stop()
}
