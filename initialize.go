package gg

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Initialize[T Entities](r gin.IRouter, db *gorm.DB) error {
	repo := NewPGStore[T](db)
	srv := NewService[T](repo)

	MountRoutes[T](r, srv)

	return nil
}
