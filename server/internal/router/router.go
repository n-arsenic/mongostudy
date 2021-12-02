package router

import "github.com/julienschmidt/httprouter"

const (
	ReadSamplePath   = "/sample"
	CreateSamplePath = "/sample/new"
	UpdateSamplePath = "/sample/update"
	DeleteSamplePath = "/sample/delete"

	SampleCategoryTree = "/sample/tree"
)

func New() *httprouter.Router {
	router := httprouter.New()

	router.GET(ReadSamplePath, handleReadSample)
	router.GET(ReadSamplePath+":id", handleReadSample)
	router.POST(CreateSamplePath, handleCreateSample)
	router.PUT(UpdateSamplePath+":id", handleUpdateOrCreateSample)
	router.PATCH(UpdateSamplePath+":id", handleUpdateSample)

	router.GET(SampleCategoryTree, handleSampleCategoryTree)

	return router
}
