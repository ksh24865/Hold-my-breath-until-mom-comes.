package model

type Signal struct {
	ID          int     `json:"id" gorm:"primary_key"`
	Temperature float32 `json:"temp" gorm:"type:double(10,2)"`
	Humidity    float32 `json:"humid" gorm:"type:double(10,2)"`
	Weight      float32 `json:"wegiht" gorm:"type:double(10,2)"`
	
}
