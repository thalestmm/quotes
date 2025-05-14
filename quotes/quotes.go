package quotes

import "gorm.io/gorm"

type Quote struct {
	gorm.Model
	Text   string `json:"text"`
	Author string `json:"author"`
}
