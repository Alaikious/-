package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetMyScores(ctx *gin.Context) {
	var user = GetUsername(ctx)
	var countEntityI []model.CountEntity
	year := ctx.Query("year")
	utils.Db.Where("countid=? AND teamyear=?", user.Countid, year).Find(&countEntityI)
	/*
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

	*/
	utils.Success(ctx, gin.H{
		"list": countEntityI,
	})
	return
}

func QueryOthersScore(ctx *gin.Context) {
	var id = ctx.Query("target")
	var user model.AccountEntity
	countEntityI := make([]model.CountEntity, 10)
	utils.Db.Where("countid=?", id).Find(&countEntityI)
	utils.Db.Where("countid=?", id).Find(&user)
	var all float64
	all = 0
	for i := 0; i < len(countEntityI); i++ {
		all += countEntityI[i].Value
		println("单项分+1:且数据为:", countEntityI[i].Value)
	}

	utils.Success(ctx, gin.H{
		"score": all,
		"name":  user.Username,
	})
	return
}

type Coach struct {
	Coachmame string `json:"name"`
	Coachid   string `json:"id"`
}

// GetMyCoaches 显示为无 todo
func GetMyCoaches(ctx *gin.Context) {
	var user = GetUsername(ctx)
	println("用户为", user.Username)
	var coachI []Coach
	println("学号为", user.Countid)
	var coach [5]model.AccountEntity
	utils.Db.Where("countid=?", user.Coachid).Find(&coach)

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

type RecordSendEntity struct {
	Id       int64     `json:"id"`
	Complain *string   `json:"complain"`     // 申诉内容
	Content  model.Tag `json:"content"`      // 文字描述
	Files    model.Tag `json:"files"`        // 图片地址数组
	Countid  string    `json:"username"`     // 申报id
	Index    string    `json:"index"`        // 申报项目代号，通过层级组合生成; `v-2-2`表示**德育素质-纪实加减分-社会责任记实分**; `v-1`表示**德育素质-基本评定分**
	Label    string    `json:"label"`        // 中文名
	Reason   *string   `json:"rejectReason"` // 驳回理由
	State    string    `json:"state"`        // 状态
	Time     string    `json:"time"`         // 申报时间
	Value    float32   `json:"value"`        // 分数
	Target   string    `json:"target"`
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
	Value   float32   `json:"value"` // 分数
	Index   string    `json:"index"`
	Content model.Tag `json:"contents"`
	Target  string    `json:"target"`
	Year    string    `json:"year"`
}

// ApplyScore todo 修改模型格式

/*
	type Form struct {
		Value map[string][]string
		File  map[string][]*FileHeader
	}

	type FileHeader struct {
		Filename string
		Header   textproto.MIMEHeader
		Size     int64
		content []byte
		tmpfile string
	}
*/
func ApplyScore(ctx *gin.Context) {
	var account Account

	var user = GetUsername(ctx)
	if err := ctx.ShouldBindJSON(&account); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}

	neww := model.RecordEntity{
		Value:   account.Value,
		Content: account.Content,
		Index:   account.Index,
		Files:   account.Files,
		Target:  account.Target,
		Year:    account.Year,
	}
	neww.Time = times()
	for i := 0; i < (len(neww.Index)+1)/2; i++ {
		var bee model.CountModelEntity
		utils.Db.Where("`index`=?", neww.Index[0:i*2+1]).Find(&bee)
		if i == 0 {
			neww.Label += bee.Label
			println(neww.Label)
		}
		if i != 0 {
			neww.Label += "-" + bee.Label
			println(neww.Label)
		}
	}
	neww.State = "pending" //todo 确认状态
	neww.Countid = user.Countid
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
	newnew.Complain = reAccount.Content
	newnew.Time = times()
	newnew.State = "complain"
	utils.Db.Where("id=?", val).Updates(&newnew)
	utils.Success1(ctx)
	return
}

type SuggestDto struct {
	Content string `json:"content"` // 建议内容
	IsAnon  bool   `json:"isAnon"`  // 是否匿名
	Target  string `json:"target"`  // 辅导员工号
}

func SubmitAdvice(ctx *gin.Context) {
	var suggest SuggestDto
	var user = GetUsername(ctx)
	if err := ctx.ShouldBindJSON(&suggest); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	if user.Coachid != suggest.Target {
		utils.Fail(ctx, "你没有这个辅导员")
		return
	}
	if suggest.IsAnon != true {
		abc := model.SuggestEntity{
			Content: suggest.Content,
			Target:  suggest.Target,
			Countid: user.Countid,
		}
		utils.Db.Create(&abc)
		utils.Success1(ctx)
		return
	}
	abc := model.SuggestEntity{
		Content: suggest.Content,
		Target:  suggest.Target,
	}
	utils.Db.Create(&abc)
	utils.Success1(ctx)
	return
}
