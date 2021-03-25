package utils

import "time"

func GetCurrentTimeString() string{
	now := time.Now()
	DateCreated := now.Format("02-01-2006T15:04:05Z")
	return DateCreated
}