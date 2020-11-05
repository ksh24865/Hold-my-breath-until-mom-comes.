package handler

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (h *GinHandler) ShowGraphPopup(c *gin.Context) {
	// Check if the article ID is valid
	if weightID, err := strconv.Atoi(c.Param("weight_id")); err == nil {
		// Check if the article exists
		if weight, err := h.mwuc.GetWeightByID(weightID); err == nil {
			// Call the render function with the title, article and the name of the
			// template
			render(c, gin.H{
				"title":   "weightgraph",
				"payload": weight}, "graph.html")
		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}