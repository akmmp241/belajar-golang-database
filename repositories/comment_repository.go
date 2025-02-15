package repositories

import (
	"belajar-golang-database/models"
	"context"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment models.Comment) (models.Comment, error)
	FindById(ctx context.Context, id int32) (models.Comment, error)
	FindAll(ctx context.Context) ([]models.Comment, error)
}
