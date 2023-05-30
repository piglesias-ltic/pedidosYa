package routes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createReviewReq struct {
	OrderId     string  `json:"order_id"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
}

func reviewHandler(c *gin.Context) {
	r := createReviewReq{}
	req, err := io.ReadAll(c.Request.Body)
	if err := json.Unmarshal(req, &r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//res, err := injector.CreateReviewUsecase.Create(nil)
	returnResponseOrFail(c, res, err)
}
