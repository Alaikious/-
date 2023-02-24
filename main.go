package main

import (
	"gin-api/api"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/api/login", api.Login)
	router.POST("/api/square/topic", api.Topic)
	router.POST("/api/square/topic/comment", api.Comment)
	router.POST("/api/changePassword", api.ChangePassword)
	router.GET("/api/login", api.LoginCookie)
	router.GET("/api/getApplyScoreRules", api.GetApplyScoreRules)
	router.GET("/api/square/getTopics", api.TopicGet)
	router.GET("/api/square/topic/:id", api.TopicAllGet)
	router.GET("/api/coach/submitScores", api.SubmitScores)
	router.GET("/api/coach/getRejectReasons", api.GetRejectReasons)
	router.POST("/api/coach/submitRejectReasons", api.SubmitRejectReasons)
	router.GET("/api/coach/getApplyList", api.GetApplyList)
	router.GET("/api/coach/scores/structure", api.Structure)
	router.POST("/api/coach/submitApproves", api.SubmitApproves)
	router.GET("/api/student/getMyScores", api.GetMyScores)
	router.GET("/api/student/queryOthersScore", api.QueryOthersScore)
	router.GET("/api/student/getMyCoaches", api.GetMyCoaches)
	router.GET("/api/student/getMyApplyHistory", api.GetMyApplyHistory)
	router.GET("/api/student/getApplyItemList", api.GetApplyItemList)
	router.POST("/api/student/applyScore", api.ApplyScore)
	router.POST("/api/student/submitAdvice", api.SubmitAdvice)
	router.POST("/api/student/submitComplain", api.SubmitComplain)
	router.GET("/api/coach/scores/query", api.Query)
	router.GET("/api/coach/queryStudent", api.QueryStudent)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
