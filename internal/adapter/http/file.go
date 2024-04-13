package http // import "go.prajeen.com/objekt/internal/adapter/http"

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.prajeen.com/objekt/internal/core/domain"
	"go.prajeen.com/objekt/internal/core/port"
)

type FileHandler struct {
	router *httprouter.Router
	svc    port.FileService
}

func NewFileHandler(router *httprouter.Router, svc port.FileService) *FileHandler {
	return &FileHandler{
		router: router,
		svc:    svc,
	}
}

type createFileRequest struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	BucketName string `json:"bucket_name"`
	MimeType   string `json:"mime_type"`
}

func (h *FileHandler) AddRoutes() {
	h.router.POST("/files", h.CreateFile)
	h.router.DELETE("/files/:id", h.DeleteFile)
	h.router.GET("/files/:id", h.GetFile)
	h.router.GET("/buckets/:bucket_id/files", h.GetFilesByBucketName)
}

func (h *FileHandler) CreateFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var requestBody createFileRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	file := &domain.File{
		Name:       requestBody.Name,
		Size:       requestBody.Size,
		BucketName: requestBody.BucketName,
		MimeType:   requestBody.MimeType,
	}

	file, err = h.svc.CreateFile(r.Context(), file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *FileHandler) DeleteFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	err := h.svc.DeleteFile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *FileHandler) GetFile(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	file, err := h.svc.GetFile(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *FileHandler) GetFilesByBucketName(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bucketID := p.ByName("bucket_id")
	files, err := h.svc.GetFilesByBucketID(r.Context(), bucketID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(files); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}