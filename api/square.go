package api

import (
	"fmt"
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type TopicDto struct {
	Content  string `json:"content"`  // 正文内容
	Title    string `json:"title"`    // 正文标题
	Type     string `json:"type"`     // 帖子类型
	UserRole string `json:"userRole"` // 用户身份
}

type CommentDto struct {
	Content string `json:"content"` // 评论内容
	Target  string `json:"target"`  // 目标帖子id
}

func Topic(ctx *gin.Context) {
	var topicDto TopicDto
	if err := ctx.ShouldBindJSON(&topicDto); err != nil {
		utils.Fail(ctx, "请求参数错误")
		println("错误日志:", err)
		return
	}

	topicEntity := model.TopicEntity{
		Content: topicDto.Content,
		Title:   topicDto.Title,
		Type:    topicDto.Type,
		Role:    topicDto.UserRole,
		Time:    times(),
	}
	err := utils.Db.Create(&topicEntity).Error
	if err != nil {
		utils.Fail(ctx, "save failed")
		return
	}
	utils.Success(ctx, gin.H{
		"id": topicEntity.Id,
	})
	return
}

func Comment(ctx *gin.Context) {
	var commentDto CommentDto
	if err := ctx.ShouldBindJSON(&commentDto); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}

	commentEntity := model.CommentEntity{
		Content: commentDto.Content,
		Target:  commentDto.Target,
		Time:    times(),
	}
	err := utils.Db.Create(&commentEntity).Error
	if err != nil {
		utils.Fail(ctx, "save failed")
		return
	}
	utils.Success(ctx, gin.H{
		"id": string(commentEntity.Id),
	})
	return
}

func TopicGet(ctx *gin.Context) {

	//可以改进...
	//commentEntity := make([][]model.CommentEntity, 200)
	topicGetEntity := make([]model.TopicGetEntity, 0)

	/*问题很大这块码
	for i := 1; i <= 198; i++ {
		utils.Db.Where("id=?", i).Find(&topicGetEntity[i])
		entity := &topicGetEntity[i]
		for j := 1; j <= 198; j++ {
			utils.Db.Where("target=?", string(i)).Find(&commentEntity[i][j])
		}
		_ = &entity.Comment
		entities := &commentEntity[i]
		_ = entities
	}
	*/

	page := ctx.Query("page")
	size := ctx.Query("size")
	pageI, _ := strconv.Atoi(page)
	sizeI, _ := strconv.Atoi(size)

	var topicGetEntityAll []model.TopicGetEntity
	utils.Db.Find(&topicGetEntityAll)

	for i := pageI*sizeI - sizeI + 1; i < pageI*sizeI+1; i++ {
		var topicGetEntityI model.TopicGetEntity
		utils.Db.Where("id=?", i).Find(&topicGetEntityI)
		topicGetEntity = append(topicGetEntity, topicGetEntityI)
	}
	for i := 0; i < len(topicGetEntity); i++ { //去除空切片
		if topicGetEntity[i].Id == 0 {
			topicGetEntity = append(topicGetEntity[:i], topicGetEntity[i+1:]...)
			i--
		}
	}

	fmt.Println(topicGetEntity)
	/*
		buffer := make([]byte, 2048)
		length, _ := ctx.Request.Body.Read(buffer)
		if err := json.Unmarshal(buffer[:length], &topicGetEntity); err != nil {
			panic(err)
		}
	*/
	utils.Success(ctx, gin.H{
		"count": len(topicGetEntityAll),
		"list":  topicGetEntity,
	})

}

func TopicAllGet(ctx *gin.Context) {
	aimId := ctx.Param("id")
	aimIdInt, _ := strconv.Atoi(aimId)
	var topicGetEntityI model.TopicGetEntity
	commentEntity := make([]model.CommentEntity, 0, 100)
	for j := 0; j < 10; j++ {
		var commentEntityI []model.CommentEntity
		utils.Db.Where("target=?", aimId).Find(&commentEntityI)
		commentEntity = commentEntityI
	}

	utils.Db.Where("id=?", aimIdInt).Find(&topicGetEntityI)
	topicGetEntityI.Comment = commentEntity
	utils.Success(ctx, gin.H{
		"topic": topicGetEntityI,
	})
	return
}

func times() string {
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	return timeStr
}
