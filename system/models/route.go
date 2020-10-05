package models

import "github.com/gin-gonic/gin"

type Route struct {
	//	Name        string
	Method      string
	Pattern     string
	MiddleWare  gin.HandlersChain
	HandlerFunc gin.HandlerFunc
}
