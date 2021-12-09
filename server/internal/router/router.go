package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/n-arsenic/mongostudy/server/internal/storage"
)

const SamplePath = "/sample"

func NewSampleRouter(db *storage.MongoDB) *httprouter.Router {
	router := httprouter.New()
	store := storage.NewSampleStore(db)
	h := NewHandler(store)

	router.GET(SamplePath, h.handleReadSample)
	router.GET(SamplePath+":id", h.handleReadSample)
	router.POST(SamplePath, h.handleCreateSample)
	router.PUT(SamplePath, h.handleUpdateSample)

	return router
}
