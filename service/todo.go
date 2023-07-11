package service

import (
	"context"
	"database/sql"
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

	//todoを保存する流れ

	//ステートメント(準備)の定義
	todo := TODOService{}
	responce := model.CreateTODOResponse{}

	//引数チェック
	if subject == "" {
		//エラーコード400を返す
		// return nil, http.StatusBadRequest
		log.Println("error")
	}
	defer log.Println("error")

	stmt, err := todo.db.PrepareContext(ctx, insert)
	if err != nil {
		//エラーハンドリング
		log.Println("err")
	

	//実行する
	//result, err := todo.db.ExecContext(ctx, insert)
	result, err := stmt.ExecContext(ctx, subject,description)
	if err != nil {
		//エラーハンドリング
		log.Println("err")
	}
	// log.Println(result)
	//IDをExecContextの結果から取得
	id, err := result.LastInsertId()

	//確認する
	todo.db.QueryRowContext(ctx, confirm,id)
	row.Scan(&responce)

	//値を返す

	return &responce.TODO, nil
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
