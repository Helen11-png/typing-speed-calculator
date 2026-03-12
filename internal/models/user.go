package models

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string // в реальности - хеш
	CreatedAt time.Time
}

type Profile struct {
	Username     string
	TestsCount   int
	AverageSpeed int
	BestSpeed    int
	JoinDate     string
}
