package models

import "time"

type Item struct {
    ID          uint      `gorm:"primaryKey"`
    Ticker      string    `gorm:"uniqueIndex"`
    Company     string
    Brokerage   string
    Action      string
    RatingFrom  string
    RatingTo    string
    TargetFrom  float64
    TargetTo    float64
    Time        time.Time
}
