package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
)

func GetApplyScoreRules(ctx *gin.Context) {
	entity := model.TipEntity{}
	utils.Success(ctx, gin.H{
		"object": entity,
	})
	return
}
