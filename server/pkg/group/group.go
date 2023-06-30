package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BatchGroup struct {
	groups []*gin.RouterGroup
}

func NewBatchGroup(groups []*gin.RouterGroup) BatchGroup {
	return BatchGroup{
		groups,
	}
}

func (batchGroup *BatchGroup) Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) {
	for _, group := range batchGroup.groups {
		group.Handle(httpMethod, relativePath, handlers...)
	}
}

func (group *BatchGroup) POST(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodPost, relativePath, handlers...)
}

func (group *BatchGroup) GET(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodGet, relativePath, handlers...)
}

func (group *BatchGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodDelete, relativePath, handlers...)
}

func (group *BatchGroup) PATCH(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodPatch, relativePath, handlers...)
}

func (group *BatchGroup) PUT(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodPut, relativePath, handlers...)
}

func (group *BatchGroup) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodOptions, relativePath, handlers...)
}

func (group *BatchGroup) HEAD(relativePath string, handlers ...gin.HandlerFunc) {
	group.Handle(http.MethodHead, relativePath, handlers...)
}
