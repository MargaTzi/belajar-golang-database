package repository

import (
	"context"
	"database/sql"

	"github.com/MargaTzi/belajar-golang-database/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment)(entity.Comment, error){
	query := "insert into komen(email, comment) values($1, $2) returning id"

	err := repository.DB.QueryRowContext(ctx, query, comment.Email, comment.Comment).Scan(&comment.Id)
	if err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int)(entity.Comment, error){
	query := "select id, email, comment from komen where id = $1 limit 1"

	var comment entity.Comment

	err := repository.DB.QueryRowContext(ctx, query, id).Scan(&comment.Id, &comment.Email, &comment.Comment)
	if err != nil {
		return entity.Comment{}, err
	} 

	return comment, nil
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context)([]entity.Comment, error){
	query := "select id, email, comment from komen"

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	var comments []entity.Comment
	
	for rows.Next(){
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}