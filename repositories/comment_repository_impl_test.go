package repositories

import (
	belajargolangdatabase "belajar-golang-database"
	"belajar-golang-database/models"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestCommentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	comment := models.Comment{
		Email:   "kamal@gmail.com",
		Comment: "Mantap",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentRepositoryImpl_FindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()

	comment := models.Comment{
		Email:   "kamal@gmail.com",
		Comment: "Uwu",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	commentResult, err := commentRepository.FindById(ctx, result.Id)
	if err != nil {
		panic(err)
	}

	fmt.Println(commentResult)
}

func TestCommentRepositoryImpl_FindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
