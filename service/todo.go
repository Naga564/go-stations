package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/TechBowl-japan/go-stations/model"
)

// A TODOService implements CRUD of TODO entities.
type TODOService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewTODOService(db *sql.DB) *TODOService {
	return &TODOService{
		db: db,
	}
}

// CreateTODO creates a TODO on DB.
func (s *TODOService) CreateTODO(ctx context.Context, subject, description string) (*model.TODO, error) {
	const (
		insert  = `INSERT INTO todos(subject, description) VALUES(?, ?)`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	//戻り値構造体を有効化
	response := model.CreateTODOResponse{}

	if subject == "" {
		log.Println("Subject empty")

		//エラーコード400を返す
		// return
	}

	//descriptionが適切かどうか判定
	// if description = ""{

	//エラーコード400を返す
	// return
	// }

	//ステートメント(準備)の定義

	stmt, err := s.db.PrepareContext(ctx, insert)
	if err != nil {
		//エラーハンドリング
		log.Println("err")
	}

	//実行する
	result, err := s.db.ExecContext(ctx, insert, subject, description)
	//result, err := stmt.ExecContext(ctx, stmt)
	if err != nil {
		//エラーハンドリング
		log.Println("err")
	}
	log.Println("stmt:::")
	log.Println(stmt)
	log.Println("ID:::")
	//log.Println(result.LastInsertId())
	log.Println(result)

	//確認する(引数にidが必要)
	row := s.db.QueryRowContext(ctx, confirm)

	//値を返す
	response.TODO = row
	err := json.NewEncoder(w).Encode(response)
	// return nil, nil
}

// ReadTODO reads TODOs on DB.
func (s *TODOService) ReadTODO(ctx context.Context, prevID, size int64) ([]*model.TODO, error) {
	const (
		read       = `SELECT id, subject, description, created_at, updated_at FROM todos ORDER BY id DESC LIMIT ?`
		readWithID = `SELECT id, subject, description, created_at, updated_at FROM todos WHERE id < ? ORDER BY id DESC LIMIT ?`
	)

	return nil, nil
}

// UpdateTODO updates the TODO on DB.
func (s *TODOService) UpdateTODO(ctx context.Context, id int64, subject, description string) (*model.TODO, error) {
	const (
		update  = `UPDATE todos SET subject = ?, description = ? WHERE id = ?`
		confirm = `SELECT subject, description, created_at, updated_at FROM todos WHERE id = ?`
	)

	return nil, nil
}

// DeleteTODO deletes TODOs on DB by ids.
func (s *TODOService) DeleteTODO(ctx context.Context, ids []int64) error {
	const deleteFmt = `DELETE FROM todos WHERE id IN (?%s)`

	return nil
}
