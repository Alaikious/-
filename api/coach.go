package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CoachSubmit struct {
	Score  []Score `json:"score"`
	Target string  `json:"target"` // 目标学生学号
}

type Score struct {
	Content []string `json:"content"` // 具体内容数组，累计分数的项目列表，若`unique`为真，则无此项
	Value   float64  `json:"value"`
	Index   string   `json:"index"` // 分数
}

// SubmitScores todo
func SubmitScores(ctx *gin.Context) {
	var coachSubmit CoachSubmit

	if err := ctx.ShouldBindJSON(&coachSubmit); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}

	for i := 1; i < len(coachSubmit.Score); i++ {
		var modelCount model.CountModelEntity
		utils.Db.Where("index=?", coachSubmit.Score[i].Index).Find(&modelCount)
		entity := model.CountEntity{
			ApplyTime: times(),
			Content:   coachSubmit.Score[i].Content,
			Index:     modelCount.Index,
			Label:     modelCount.Label,
			Top:       modelCount.Top,
			Unique:    modelCount.Unique,
			Value:     coachSubmit.Score[i].Value,
			Countid:   coachSubmit.Target,
			Realindex: modelCount.Realindex,
		}
		err := utils.Db.Updates(&entity).Error
		if err != nil {
			utils.Fail(ctx, "save failed")
			return
		}
	}

	utils.Success1(ctx)
}

func GetRejectReasons(ctx *gin.Context) {
	var reason []model.Reasons
	utils.Db.Find(&reason)
	utils.Success(ctx, reason)
}

func SubmitRejectReasons(ctx *gin.Context) {
	var reason []string

	if err := ctx.ShouldBindJSON(&reason); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	toReason := model.Reasons{
		Reason: reason,
	}
	err := utils.Db.Create(&toReason).Error
	if err != nil {
		utils.Fail(ctx, "save failed")
		return
	}
	utils.Success1(ctx)
	return
}

func GetApplyList(ctx *gin.Context) {
	recordEntity := make([]model.RecordEntity, 0)
	var total []model.RecordEntity
	utils.Db.Find(&total)
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageI, _ := strconv.Atoi(page)
	sizeI, _ := strconv.Atoi(size)
	for i := pageI*sizeI - sizeI + 1; i < pageI*sizeI+1; i++ {
		var recordEntityI model.RecordEntity
		utils.Db.Where("id=?", i).Find(&recordEntityI)
		recordEntity = append(recordEntity, recordEntityI)
	}

	for i := 0; i < len(recordEntity); i++ { //去除空切片
		if recordEntity[i].Id == 0 {
			recordEntity = append(recordEntity[:i], recordEntity[i+1:]...)
			i--
		}
	}
	utils.Success(ctx, gin.H{
		"count": len(total),
		"list":  recordEntity,
	})
}

// Structure 完成勿动
func Structure(ctx *gin.Context) {

	countEntityI := make([]model.CountModelEntity, 0, 100)
	countEntityII := make([]model.CountModelEntity, 0, 100)
	var countEntityCount []model.CountModelEntity
	utils.Db.Find(&countEntityCount)
	for i := 1; i < len(countEntityCount); i++ {
		var countEntityIII model.CountModelEntity
		utils.Db.Where("id=?", i).Find(&countEntityIII)
		println(countEntityIII.Realindex)
		countEntityI = append(countEntityI, countEntityIII)
		println("realindex:", countEntityI[0].Realindex)
		var countEntityIIII model.CountModelEntity
		utils.Db.Where("id=?", i).Find(&countEntityIIII)
		println(countEntityIIII.Realindex)
		countEntityII = append(countEntityII, countEntityIIII)
	}

	for i := 0; i < len(countEntityI); i++ {
		for j := 0; j < len(countEntityII); j++ {
			if countEntityI[i].Realindex%10 == 0 && countEntityII[j].Realindex-countEntityI[i].Realindex > 0 && countEntityII[j].Realindex-countEntityI[i].Realindex < 10 {
				countEntityI[i].List = append(countEntityI[i].List, countEntityII[j])
				println("+1次放入")
			}
		}
	}

	for i := 0; i < len(countEntityI); i++ { //去除错误切片
		if countEntityI[i].Realindex%10 != 0 || countEntityI[i].Realindex == 0 {
			countEntityI = append(countEntityI[:i], countEntityI[i+1:]...)
			i--
		}
	}

	utils.Success(ctx, gin.H{
		"top":   100,
		"list":  countEntityI,
		"index": "v-all",
		"label": "all counts",
	})
	return
}

type SubmitElementDto []SubmitElement

type SubmitElement struct {
	Reason string `json:"reason"` // 理由，若`state`为`rejected`则有此项
	State  string `json:"state"`  // 设置申报状态
	Target string `json:"target"` // 目标申报id
}

// SubmitApproves 完成勿-s
//
//	SubmitApproves todo 加一个批量s
func SubmitApproves(ctx *gin.Context) {
	var submitElementDto SubmitElementDto
	if err := ctx.ShouldBindJSON(&submitElementDto); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}

	for j := 0; j < len(submitElementDto); j++ {
		var recordEntity model.RecordEntity
		val, _ := strconv.Atoi(submitElementDto[j].Target)
		utils.Db.Where("id=?", val).Find(&recordEntity)
		recordEntity.State = submitElementDto[j].State
		recordEntity.Reason = submitElementDto[j].Reason

		err := utils.Db.Where("id=?", recordEntity.Id).Updates(&recordEntity).Error
		if err != nil {
			utils.Fail(ctx, "save failed")
			return
		}
	}
	utils.Success1(ctx)
	return
}

func Query(ctx *gin.Context) {
	var id = ctx.Query("username")
	countEntityI := make([]model.CountEntity, 100)
	utils.Db.Where("countid=?", id).Find(&countEntityI)
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

func QueryStudent(ctx *gin.Context) {
	countEntityI := make([]model.CountEntity, 100)
	utils.Db.Find(&countEntityI)
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

func Line(count []model.CountModelEntity) {
	for i := 0; i < len(count); i++ {
		if count[i].father == nil {
			continue
		}
		var index = -1
		for j := 0; j < len(count); j++ {
			if count[i].father == count[j].Index {
				index = j
				break
			}
		}
		/*
			var son = count[i]
			count[index].List = append(count[index].List, son)
			count = append(count[:i], count[i+1:]...)
		*/
		count[index].List = append(count[index].List, count[i])
	}
	var roots []model.CountModelEntity
	for i := 0; i < len(count); i++ {
		if count[i].father == nil {
			roots = append(roots, count[i])
			println(count[i].Index)
		}
	}
}

/*
func Line(count []model.CountModelEntity, val int) []model.CountModelEntity {

	if val == 0 {
		for i := 0; i < len(count); i++ {
			var had string
			if strings.Contains(had, count[i].Index[0:1]) || len(count[i].Index) != val+1 {
				continue
			}

			for j := 0; j < len(count); j++ {
				if len(count[j].Index) == val+3 && count[i].Index[val:val+1] == count[j].Index[val:val+1] {
					var son model.CountModelEntity
					son = count[j]
					count[i].List = append(count[i].List, son)
					count = append(count[:j], count[j+1:]...)
				}
			}
			Line(count[i].List, val+2)
		}
	}
	if val!=0{
		for i := 0; i < len(count); i++ {
			var had string
			if strings.Contains(had, count[i].Index[0:1]) || len(count[i].Index) != val+1 {
				continue
			}
			for j := 0; j < len(count); j++ {
				if len(count[j].Index) == val+3 && count[i].Index[val:val+1] == count[j].Index[val:val+1] {
					var son model.CountModelEntity
					son = count[j]
					count[i].List = append(count[i].List, son)
					count = append(count[:j], count[j+1:]...)
				}
			}
			Line(count[i].List, val+2)
		}
	}
	return
}
*/
