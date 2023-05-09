package domain

import (
	"time"
)

type DateRange struct {
	StartTime time.Time
	EndTime   time.Time
}

func NewDateRange(startTime time.Time, endTime time.Time) DateRange {
	return DateRange{
		StartTime: startTime,
		EndTime:   endTime,
	}
}

func (d DateRange) OverlapsWith(dateRange DateRange) bool {
	return (d.StartTime.Before(dateRange.StartTime) && d.EndTime.After(dateRange.StartTime)) ||
		(d.StartTime.Before(dateRange.EndTime) && d.EndTime.After(dateRange.EndTime)) ||
		(d.StartTime.After(dateRange.StartTime) && d.EndTime.Before(dateRange.EndTime))
}

func (d DateRange) IsLesserThan(time time.Time) bool {
	return d.EndTime.Before(time)
}
