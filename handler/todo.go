package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TechBowl-japan/go-stations/model"
	"github.com/TechBowl-japan/go-stations/service"
)

// A TODOHandler implements handling REST endpoints.
type TODOHandler struct {
	svc *service.TODOService
}

// NewTODOHandler returns TODOHandler based http.Handler.
func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

// Create handles the endpoint that creates the TODO.
func (h *TODOHandler) Create(ctx context.Context, req *model.CreateTODORequest) (*model.CreateTODOResponse, error) {
	_, _ = h.svc.CreateTODO(ctx, "", "")
	return &model.CreateTODOResponse{}, nil
}

// Read handles the endpoint that reads the TODOs.
func (h *TODOHandler) Read(ctx context.Context, req *model.ReadTODORequest) (*model.ReadTODOResponse, error) {
	_, _ = h.svc.ReadTODO(ctx, 0, 0)
	return &model.ReadTODOResponse{}, nil
}

// Update handles the endpoint that updates the TODO.
func (h *TODOHandler) Update(ctx context.Context, req *model.UpdateTODORequest) (*model.UpdateTODOResponse, error) {
	_, _ = h.svc.UpdateTODO(ctx, 0, "", "")
	return &model.UpdateTODOResponse{}, nil
}

// Delete handles the endpoint that deletes the TODOs.
func (h *TODOHandler) Delete(ctx context.Context, req *model.DeleteTODORequest) (*model.DeleteTODOResponse, error) {
	_ = h.svc.DeleteTODO(ctx, nil)
	return &model.DeleteTODOResponse{}, nil
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	request := &model.CreateTODORequest{}
	responce := &model.CreateTODOResponse{}

	//POSTかどうかの判定
	if r.Method == http.MethodPost {
		//JSON DECODE
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "error", http.StatusBadRequest)
		}

		//subjectが空文字かどうかを判定
		if request.Subject == "" {
			//空の場合
			//400 BadRequestを返却
			http.Error(w, "error", http.StatusBadRequest)
		} else {
			//空ではない場合
			//(3)をやる(途中)
			r.Context()
			//APIを呼び出す
			//responce =

			//戻ってきたものをエンコード
			err := json.NewEncoder(w).Encode(&responce)
			//エラーハンドリング
			if err != nil {
				log.Println(err)
			}
			//正常終了
			w.WriteHeader(http.StatusOK)
		}
	} else {
		//GETだったとき
	}

	// err := json.NewEncoder(w).Encode(responce)

	// //エラーハンドリング
	// if err != nil {
	// 	log.Println(err)
	// }

}
