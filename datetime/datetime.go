// https://github.com/jinzhu/now
package datetime

import (
	"time"
)

// WeekStartDay set week start day, default is sunday
var WeekStartDay = time.Sunday

type DateTime struct {
	time.Time
}

func new(t time.Time) *DateTime {
	return &DateTime{t}
}

func (t *DateTime) beginningOfDay() time.Time {
	y, m, d := t.Date()

	return time.Date(y, m, d, 0, 0, 0, 0, t.Time.Location())
}

func BeginningOfWeek() time.Time {
	return new(time.Now()).beginningOfWeek()
}

func (d *DateTime) beginningOfWeek() time.Time {
	t := d.beginningOfDay()
	weekday := int(t.Weekday())

	if WeekStartDay != time.Sunday {
		weekStartDayInt := int(WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}

	return t.AddDate(0, 0, -weekday)
}
