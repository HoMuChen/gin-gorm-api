package models

import (
  "time"
  "github.com/lib/pq"
)

type Media struct {
  Id            string    `gorm:"primary_key"`
  Description   string
  Comment_cout  int
  Image_url     string
  Like_count    int
  Location_id   string
  Location_name string
  Owner_id      string
  Tags          pq.StringArray
  Tm            time.Time
  Updated_at    time.Time
  Url           string
}
