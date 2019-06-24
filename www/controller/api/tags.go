package api

import (
	"github.com/gin-gonic/gin"
)

func GetTags(c *gin.Context){
	c.JSON(200, gin.H{"msg":"Test get targets!"})
	//return
}

func PostTags(c *gin.Context){

}
