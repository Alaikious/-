package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CoachSubmit struct {
	Score  []Score `json:"scores"`
	Target string  `json:"target"` // 目标学生学号
	Year   int     `json:"year"`
}

type Score struct {
	Value float64 `json:"value"`
	Index string  `json:"index"` // 分数
}

// SubmitScores todo
func SubmitScores(ctx *gin.Context) {
	var coachSubmit CoachSubmit

	if err := ctx.ShouldBindJSON(&coachSubmit); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	var user model.AccountEntity
	utils.Db.Where("countid=?", coachSubmit.Target).Find(&user)
	if user.Username == "" {
		utils.Fail(ctx, "没有此学生学号")
		return
	}
	var entity []model.CountEntity
	for i := 0; i < len(coachSubmit.Score); i++ {
		var modelCount model.CountModelEntity
		utils.Db.Where("`index`=?", coachSubmit.Score[i].Index).Find(&modelCount)
		entityI := model.CountEntity{
			Index:    coachSubmit.Score[i].Index,
			Label:    modelCount.Label,
			Top:      modelCount.Top,
			Unique:   modelCount.Unique,
			Value:    coachSubmit.Score[i].Value,
			Countid:  coachSubmit.Target,
			Teamyear: coachSubmit.Year,
		}
		entity = append(entity, entityI)
		utils.Db.Where("index=? AND countid=?", entityI.Index, entityI.Countid).Delete(entityI)
		/*
			err := utils.Db.Updates(&entity).Error
			if err != nil {
				utils.Fail(ctx, "save failed")
				return
			}
		*/
	}
	utils.Db.Create(&entity)
	utils.Success1(ctx)
	return
}

func GetRejectReasons(ctx *gin.Context) {
	var reason []model.Reasons
	utils.Db.Find(&reason)
	var reasons model.Tag
	for i := 0; i < len(reason); i++ {
		reasons = append(reasons, reason[i].Content)
	}
	utils.Success(ctx, gin.H{
		"list": reasons,
	})
}

type ReasonsDto struct {
	Reasons model.Tag `json:"reasons"`
}

func SubmitRejectReasons(ctx *gin.Context) {
	var reason ReasonsDto
	var reasons []model.Reasons
	utils.Db.Find(&reasons)
	utils.Db.Where("1 = 1").Delete(&reasons)
	if err := ctx.ShouldBindJSON(&reason); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	var reasonss []model.Reasons
	for i := 0; i < len(reason.Reasons); i++ {
		reasonsss := model.Reasons{
			Content: reason.Reasons[i],
		}
		reasonss = append(reasonss, reasonsss)
	}
	err := utils.Db.Create(&reasonss).Error
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

	utils.Db.Where("id>=?", sizeI*(pageI-1)+1).Limit(sizeI).Find(&recordEntity)

	utils.Success(ctx, gin.H{
		"count": len(total),
		"list":  recordEntity,
	})
}

// Structure 完成勿动
func Structure(ctx *gin.Context) {
	var countModelEntity []model.CountModelEntity
	utils.Db.Find(&countModelEntity)
	utils.Success(ctx, gin.H{
		"list": countModelEntity,
	})
	/*
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

	*/
	return
}

type SubmitElementDto struct {
	Element []SubmitElement `json:"applications"`
}

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
	ctx.ShouldBindJSON(&submitElementDto)
	for j := 0; j < len(submitElementDto.Element); j++ {
		var recordEntity model.RecordEntity
		val, _ := strconv.Atoi(submitElementDto.Element[j].Target)
		utils.Db.Where("id=?", val).Find(&recordEntity)
		if submitElementDto.Element[j].State == "approved" {
			var count model.CountEntity
			yearI, _ := strconv.Atoi(recordEntity.Year)
			utils.Db.Where("`index`=? AND countid=? AND teamyear=?", recordEntity.Index, recordEntity.Countid, yearI).Find(&count)
			count.Value = float64(recordEntity.Value)
			var modelCount model.CountModelEntity
			utils.Db.Where("`index`=?", recordEntity.Index).Find(&modelCount)
			countCopy := model.CountEntity{
				Index:    recordEntity.Index,
				Label:    modelCount.Label,
				Top:      modelCount.Top,
				Unique:   modelCount.Unique,
				Value:    float64(recordEntity.Value),
				Countid:  recordEntity.Countid,
				Teamyear: yearI,
			}
			utils.Db.Where("`index`=? AND countid=? AND teamyear=?", count.Index, count.Countid, yearI).Delete(&count)
			utils.Db.Create(&countCopy)
		}
		if submitElementDto.Element[j].State == "withdraw" {
			var count model.CountEntity
			if recordEntity.State == "approved" {
				yearI, _ := strconv.Atoi(recordEntity.Year)
				utils.Db.Where("`index`=? AND countid=? AND teamyear=?", recordEntity.Index, recordEntity.Countid, yearI).Find(&count).Delete(&count)
			}
			utils.Db.Where("id=?", val).Delete(&recordEntity)
			utils.Success1(ctx)
			return
		}
		recordEntity.State = submitElementDto.Element[j].State
		recordEntity.Reason = submitElementDto.Element[j].Reason
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
	var countEntityI []model.CountEntity
	utils.Db.Where("countid=?", id).Find(&countEntityI)
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
		"score": countEntityI,
	})
	return
}

type List struct {
	Id       int                 `json:"id"`
	Username string              `json:"username"`
	Name     string              `json:"name"`
	Scores   []model.CountEntity `json:"scores"`
	Year     string              `json:"year"`
}

func QueryStudentScores(ctx *gin.Context) {
	println("fuck")
	page := ctx.Query("page")
	size := ctx.Query("size")
	pageI, _ := strconv.Atoi(page)
	sizeI, _ := strconv.Atoi(size)
	username := ctx.Query("username")
	if username != "" {
		var countEntityLike []model.CountEntity
		utils.Db.Where("countid LIKE ?", username+"%").Find(&countEntityLike)
		var user []model.AccountEntity
		var had []string
		had = append(had, "ok")
		for i := 0; i < len(countEntityLike); i++ {
			var users model.AccountEntity
			var ok = false
			for j := 0; j < len(had); j++ {
				if had[j] == countEntityLike[i].Countid {
					ok = true
				}
			}
			if ok == true {
				continue
			}
			utils.Db.Where("countid=?", countEntityLike[i].Countid).Find(&users)
			if users.Id != 0 {
				user = append(user, users)
			}
			had = append(had, countEntityLike[i].Countid)
		}
		var list []List
		for i := 0; i < len(user); i++ {
			var listOne List
			listOne.Id = int(user[i].Id)
			listOne.Name = user[i].Username
			listOne.Username = user[i].Countid
			listOne.Year = "2023"
			for j := 0; j < len(countEntityLike); j++ {
				if countEntityLike[j].Countid == user[i].Countid {
					listOne.Scores = append(listOne.Scores, countEntityLike[j])
				}
			}
			list = append(list, listOne)
		}

		utils.Success(ctx, gin.H{
			"total": len(user),
			"list":  list,
		})
		return
	}

	var user []model.AccountEntity
	var countEntityAll []model.CountEntity
	utils.Db.Find(&countEntityAll)
	for i := pageI*sizeI - sizeI + 1; i < pageI*sizeI+1; i++ {
		var users model.AccountEntity
		utils.Db.Where("role=?", "student").Where("id=?", i).Find(&users)
		if users.Id != 0 {
			user = append(user, users)
		}
	}
	var list []List
	for i := 0; i < len(user); i++ {
		var listOne List
		listOne.Id = int(user[i].Id)
		listOne.Name = user[i].Username
		listOne.Username = user[i].Countid
		listOne.Year = "2023"
		for j := 0; j < len(countEntityAll); j++ {
			if countEntityAll[j].Countid == user[i].Countid {
				listOne.Scores = append(listOne.Scores, countEntityAll[j])
			}
		}
		list = append(list, listOne)
	}

	utils.Success(ctx, gin.H{
		"total": len(user),
		"list":  list,
	})
	/*
		var son [100]model.CountEntity
		for i := 0; i < len(countEntityAll); i++ {
			countEntityII := make([]model.CountEntity, 100)
			for j := 0; j < len(countEntityAll); j++ {
				if countEntityAll[j].Realindex-countEntityAll[i].Realindex > 0 && countEntityAll[j].Realindex-countEntityAll[i].Realindex < 10 {
					countEntityII = append(countEntityII, countEntityAll[j])
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
		for i := 0; i < len(countEntityAll); i++ {
			countEntityAll[i].List = son[i].List
		}

		for i := 0; i < len(countEntityAll); i++ { //去除错误切片
			if countEntityAll[i].Realindex%10 != 0 || countEntityAll[i].Realindex == 0 {
				countEntityAll = append(countEntityAll[:i], countEntityAll[i+1:]...)
				i--
			}
		}

	*/
	return

}

func Line(count []model.CountModelEntity) {
	/*
		var indexLine []int
		var Lines []model.CountModelEntity
		for i := 0; i < len(count); i++ {
			indexLine = append(indexLine, count[i].Id)
		}
		sort.Ints(indexLine)
		for i := 0; i < len(indexLine); i++ {
			for j := 0; j < len(count); j++ {
				if indexLine[i] == count[j].Id {
					Lines = append(Lines, count[j])
				}
			}
		}
		return Lines

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
						var son = count[i]
						count[index].List = append(count[index].List, son)
						count = append(count[:i], count[i+1:]...)

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
	return
}

type Major struct {
	Major string `json:"major" gorm:"major"`
}

func (t *Major) TableName() string {
	return "major"
}

type Student struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Major    string `json:"major"`
	year     int    `json:"year"`
}

func QueryStudent(ctx *gin.Context) {
	aimId := ctx.Param("username")
	if len(aimId) != 12 {
		utils.Fail(ctx, "学号格式错误")
	}
	year, _ := strconv.Atoi(aimId[0:4])
	var major Major
	var user model.AccountEntity
	utils.Db.Where("countid=?", aimId).Find(&user)
	utils.Db.Where("countid=?", aimId[4:8]).First(&major)
	println(aimId[4:8])
	student := Student{
		Id:       int(user.Id),
		Username: user.Username,
		Major:    major.Major,
		year:     year,
	}
	utils.Success(ctx, gin.H{
		"student": student,
	})
	return
}

type SuggestSend struct {
	Content string `json:"content"` // 建议内容
	Countid string `json:"studentId"`
	Name    string `json:"studentName"`
	Target  string `json:"coachId"` // 辅导员工号
}

type SuggestSendAnno struct {
	Content string `json:"content"` // 建议内容
	Target  string `json:"coachId"` // 辅导员工号
}

func GetSuggest(ctx *gin.Context) {
	var coach = GetUsername(ctx)
	var suggest []model.SuggestEntity
	var suggestSend []interface{}
	utils.Db.Where("target=?", coach.Countid).Find(&suggest)
	for i := 0; i < len(suggest); i++ {
		if suggest[i].Countid != "" {
			var student model.AccountEntity
			utils.Db.Where("countid=?", suggest[i].Countid).Find(&student)
			send := SuggestSend{
				Content: suggest[i].Content,
				Countid: suggest[i].Countid,
				Name:    student.Username,
				Target:  suggest[i].Target,
			}
			suggestSend = append(suggestSend, send)
			continue
		}
		if suggest[i].Countid == "" {
			send := SuggestSendAnno{
				Content: suggest[i].Content,
				Target:  suggest[i].Target,
			}
			suggestSend = append(suggestSend, send)
		}

	}
	utils.Success(ctx, gin.H{
		"advice": suggestSend,
	})
}
