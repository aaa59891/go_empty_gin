package controllers

import (
	"net/http"

	"github.com/aaa59891/go_empty_gin/constants"
	"github.com/gin-gonic/gin"
)

func ErrorHandler404(c *gin.Context) {
	GoToErrorResponse(http.StatusNotFound, c, constants.Err404NotFound)
}
