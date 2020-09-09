package core

import "github.com/jinzhu/gorm"

type Dong struct {
	gorm.Model
	Dong     string
	Category string
}

func (d *Dong) Stringify() string {
	return d.Dong
}
