package model

type Weight struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Weight      float32 `json:"wegiht" gorm:"type:float"`
	Date	    string  `json:"dates" gorm:"type:char(40)"`
	
}
