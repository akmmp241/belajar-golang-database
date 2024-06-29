package repositories

import (
	"belajar-golang-database/models"
	"context"
	"database/sql"
	"errors"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repository *CommentRepositoryImpl) Insert(ctx context.Context, comment models.Comment) (models.Comment, error) {
	command := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := repository.DB.ExecContext(ctx, command, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (models.Comment, error) {
	command := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, command, id)
	comment := models.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}

		return comment, nil
	} else {
		return comment, errors.New("Comment Not Found")
	}
}

func (repository *CommentRepositoryImpl) FindAll(ctx context.Context) ([]models.Comment, error) {
	command := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, command)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		comment := models.Comment{}
		err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		if err != nil {
			return nil, err
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
