package res

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	c              *gin.Context
	httpStatusCode int
	response       ApiResponse
}

func (j *JsonResponse) Send() {
	j.c.JSON(j.httpStatusCode, j.response)
}

func NewSuccessJsonResponse(c *gin.Context, httpCode int, code string, msg string, data interface{}) AppHttpResponse {
	httpStatusCode, res := NewSuccessMessage(httpCode, code, msg, data)
	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}

func NewErrorJsonResponse(c *gin.Context, httpCode int, code string, err error) AppHttpResponse {
	httpStatusCode, res := NewFailedMessage(httpCode, code, err)
	return &JsonResponse{
		c,
		httpStatusCode,
		res,
	}
}
