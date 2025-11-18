package repository

import (
	"context"
	"fmt"
	"log"
	"testing"
	_ "github.com/lib/pq"
	belajargolangdatabase "github.com/MargaTzi/belajar-golang-database"
	"github.com/MargaTzi/belajar-golang-database/entity"
)

func TestCommentInsert(t *testing.T) {
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email: "Ditotest@mail.com",
		Comment: "Ini coba",
	}

	data, err := CommentRepository.Insert(ctx, comment)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(data)
}

func TestFindById(t *testing.T){
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comment, err := CommentRepository.FindById(ctx, 43)
	if err != nil {
		log.Fatalln(err)
	}	
	fmt.Println(comment)
}

func TestFindAll(t *testing.T){
	CommentRepository := NewCommentRepository(belajargolangdatabase.GetConnection())

	ctx := context.Background()
	comments, err := CommentRepository.FindAll(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, comment := range comments{
		fmt.Println(comment)
	}
}