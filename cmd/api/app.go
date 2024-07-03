package api

import "github.com/gin-gonic/gin"

type CTX = *gin.Context
type Map = gin.H

var App *gin.Engine = gin.New()
