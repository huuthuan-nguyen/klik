package transformer

import (
	"github.com/huuthuan-nguyen/klik-dokter/app/model"
	"time"
)

type UserTransformer struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	IsActive  bool      `json:"is_active"`
	CreateAt  time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Transform /**
func (user *UserTransformer) Transform(e any) any {
	userModel, ok := e.(model.User)
	if !ok {
		return e
	}

	user.ID = userModel.ID
	user.Email = userModel.Email
	user.IsActive = userModel.IsActive
	user.CreateAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt
	return *user
}
