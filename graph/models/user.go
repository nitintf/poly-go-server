package models

import (
	"context"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            string     `json:"id"`
	FullName      string     `json:"full_name"`
	Email         string     `json:"email"`
	Password      string     `json:"password"`
	Active        bool       `json:"active"`
	EmailVerified bool       `json:"email_verified"`
	Role          UserRole   `json:"role"`
	AvatarURL     string     `json:"avatar_url"`
	Groups        []Groups   `json:"groups" pg:"many2many:user_groups"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"-" pg:",soft_delete"`
}

var _ pg.BeforeInsertHook = (*User)(nil)

func (u *User) BeforeInsert(ctx context.Context) (context.Context, error) {
	bytePassword := []byte(u.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to hash password")
		return ctx, err
	}

	u.Password = string(passwordHash)

	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

var _ pg.BeforeUpdateHook = (*User)(nil)

func (u *User) BeforeUpdate(ctx context.Context) (context.Context, error) {
	u.UpdatedAt = time.Now()
	return ctx, nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
