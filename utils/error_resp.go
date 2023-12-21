package utils

import "github.com/gin-gonic/gin"

type APIErrors struct {
	StatusCode          int
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

// ErrorMessage ...
func ErrorMessage(c *gin.Context, status int, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"responseCode":        "1111",
		"responseDescription": msg,
	})
	return c
}

// ErrorMessage ...
func NewErrorMessage(c *gin.Context, apiError *APIErrors) *gin.Context {
	c.Abort()
	c.JSON(apiError.StatusCode, gin.H{
		"responseCode":        apiError.ResponseCode,
		"responseDescription": apiError.ResponseDescription,
	})
	return c
}
