package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetMyScores(ctx *gin.Context) {
	var user = GetUsername(ctx)
	countEntityI := make([]model.CountEntity, 100)
	utils.Db.Where("countid=?", user.Countid).Find(&countEntityI)
	var son [100]model.CountEntity
	for i := 0; i < len(countEntityI); i++ {
		countEntityII := make([]model.CountEntity, 100)
		for j := 0; j < len(countEntityI); j++ {
			if countEntityI[j].Realindex-countEntityI[i].Realindex > 0 && countEntityI[j].Realindex-countEntityI[i].Realindex < 10 {
				countEntityII = append(countEntityII, countEntityI[j])
			}
		}
		for j := 0; j < len(countEntityII); j++ { //去除错误切片
			if countEntityII[j].Realindex == 0 {
				countEntityII = append(countEntityII[:j], countEntityII[j+1:]...)
				j--
			}
		}
		son[i].List = countEntityII
	}
	for i := 0; i < len(countEntityI); i++ {
		countEntityI[i].List = son[i].List
	}

	for i := 0; i < len(countEntityI); i++ { //去除错误切片
		if countEntityI[i].Realindex%10 != 0 || countEntityI[i].Realindex == 0 {
			countEntityI = append(countEntityI[:i], countEntityI[i+1:]...)
			i--
		}
	}
	utils.Success(ctx, gin.H{
		"score": countEntityI,
	})
	return
}

func QueryOthersScore(ctx *gin.Context) {
	var id = ctx.Query("target")
	countEntityI := make([]model.CountEntity, 10)
	utils.Db.Where("countid=?", id).Find(&countEntityI)
	println(countEntityI[0].Realindex)
	var all float64
	all = 0
	for i := 0; i < len(countEntityI); i++ {
		all += countEntityI[i].Value
		println("单项分+1:且数据为:", countEntityI[i].Value)
	}

	utils.Success(ctx, gin.H{
		"score": all,
	})
	return
}

type Coach struct {
	Coachmame string `json:"mame"`
	Coachid   string `json:"id"`
}

// GetMyCoaches 显示为无 todo
func GetMyCoaches(ctx *gin.Context) {
	var user = GetUsername(ctx)
	println("用户为", user.Username)
	var coachI []Coach
	println("学号为", user.Countid)
	var coach [5]model.AccountEntity
	utils.Db.Where("coachid=?", user.Countid).Find(&coach)

	for i := 0; i < 5; i++ {
		var ccc Coach
		ccc.Coachid = coach[i].Countid
		ccc.Coachmame = coach[i].Username
		coachI = append(coachI, ccc)
	}
	for i := 0; i < len(coachI); i++ { //去除错误切片
		if coachI[i].Coachmame == "" {
			coachI = append(coachI[:i], coachI[i+1:]...)
			i--
		}
	}
	utils.Success(ctx, gin.H{
		"list": coachI,
	})
	return
}

func GetMyApplyHistory(ctx *gin.Context) {
	var user = GetUsername(ctx)
	var history []model.RecordEntity
	utils.Db.Where("countid=?", user.Countid).Find(&history)
	utils.Success(ctx, gin.H{
		"list": history,
	})
	return
}

func GetApplyItemList(ctx *gin.Context) {
	var apply []model.ApplyEntity
	utils.Db.Where("state=?", "allowed").Find(&apply) //allowed即同意，disallowed不同意
	utils.Success(ctx, gin.H{
		"list": apply,
	})
	return
}

type Account struct {
	Files   model.Tag `json:"files"` // 图片地址数组
	Value   string    `json:"value"` // 分数
	Index   string    `json:"name"`
	Content string    `json:"content"`
}

// ApplyScore todo 修改模型格式

func ApplyScore(ctx *gin.Context) {
	var account Account

	var user = GetUsername(ctx)
	if err := ctx.ShouldBindJSON(&account); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}

	val, _ := strconv.Atoi(account.Value)
	neww := model.RecordEntity{
		Value:   int64(val),
		Content: account.Content,
		Index:   account.Index,
		Files:   account.Files,
	}
	neww.Time = times()
	var bee model.RecordEntity
	utils.Db.Where("index=?", neww.Index).First(&bee)
	neww.Label = bee.Label
	neww.State = "pending" //todo 确认状态
	neww.CountID = user.Countid
	utils.Db.Create(&neww)
	var newnew model.RecordEntity
	utils.Db.Where("time=?", neww.Time).First(&newnew)
	utils.Success(ctx, gin.H{
		"id": newnew.Id,
	})
	return
}

type ReAccount struct {
	Id      string `json:"id"` // 图片地址数组
	Content string `json:"content"`
}

func SubmitComplain(ctx *gin.Context) {
	var reAccount ReAccount
	if err := ctx.ShouldBindJSON(&reAccount); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	var newnew model.RecordEntity
	val, _ := strconv.Atoi(reAccount.Id)
	utils.Db.Where("id=?", val).First(&newnew)
	newnew.Content = reAccount.Content
	newnew.Time = times()
	newnew.State = "complain"
	utils.Db.Where("id=?", val).Updates(&newnew)
	utils.Success1(ctx)
	return
}

type SuggestDto struct {
	Content string `json:"content"` // 建议内容
	IsAnon  string `json:"isAnon"`  // 是否匿名
	Target  string `json:"target"`  // 辅导员工号
}

func SubmitAdvice(ctx *gin.Context) {
	var suggest SuggestDto
	var user = GetUsername(ctx)
	if err := ctx.ShouldBindJSON(&suggest); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	if suggest.IsAnon == "true" {
		abc := model.SuggestEntity{
			Content: suggest.Content,
			Target:  suggest.Target,
			Countid: user.Countid,
		}
		utils.Db.Create(abc)
		utils.Success1(ctx)
		return
	}
	abc := model.SuggestEntity{
		Content: suggest.Content,
		Target:  suggest.Target,
	}
	utils.Db.Create(abc)
	utils.Success1(ctx)
	return
}
