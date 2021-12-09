package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/n-arsenic/mongostudy/server/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	store *storage.SampleStore
}

func NewHandler(store *storage.SampleStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) handleReadSample(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	samples := h.store.All()
	data, err := json.MarshalIndent(samples, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) handleCreateSample(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	sample := &storage.Sample{}
	err := json.NewDecoder(r.Body).Decode(sample)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resultID, err := h.store.Insert(sample)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	resp, _ := json.Marshal(struct{ SampleID string }{SampleID: resultID})

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) handleUpdateSample(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var options storage.UpdateOptions
	err := json.NewDecoder(r.Body).Decode(&options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.store.Update(&options)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) handleDeleteSample(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}

func (h *Handler) handleUpdateCategory(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
