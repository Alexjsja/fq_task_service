package types

import (
	"github.com/StanDenisov/fq_utils/users"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Author users.User `json:"author,omitempty" gorm:"many2many:user_comments"`
	Text   string     `json:"text"`
}
