package api

import (
	"gin-api/utils"
	"github.com/gin-gonic/gin"
)

type NewPassword struct {
	New string `json:"new"` // 新密码
	Old string `json:"old"` // 旧密码
}

func ChangePassword(ctx *gin.Context) {
	var newPassword NewPassword
	if err := ctx.ShouldBindJSON(&newPassword); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	var abc = GetUsername(ctx)
	abc.Password = newPassword.New
	utils.Db.Where("username=?", abc.Username).Updates(abc)
	if abc.Password != newPassword.Old {
		utils.Fail(ctx, "密码错误")
		return
	}
	/*
		println(newPassword.Old + " is old")
		var cookie model.Cookie
		var accountEntity model.AccountEntity
		val, _ := ctx.Cookie("sms-session")
		utils.Db.Debug().Where("token=?", val).First(&cookie)
		if result := utils.Db.Debug().Where("username=?", cookie.Username).First(&accountEntity).Error; result == gorm.ErrRecordNotFound {
			utils.Fail(ctx, "账号不存在")
			return
		}
		if accountEntity.Password != newPassword.Old {
			utils.Fail(ctx, "密码错误")
			return
		}
		println(accountEntity.Username + " is username")
		accountEntity.Password = newPassword.New
		println(accountEntity.Password + " is password")
		println(newPassword.New + " is new")
		err := utils.Db.Where("username=?", accountEntity.Username).Updates(&accountEntity).Error
		if err != nil {
			utils.Fail(ctx, "save failed")
			return
		}
	*/
	//utils.Db.Debug().Where("username=?", accountEntity.Username).Save(&accountEntity)
	utils.Success1(ctx)
	return
}

/* CookieGetChangePassword 获取cookie 已废弃
func CookieGetChangePassword(c *gin.Context) {
	cookieKey := "sms-session"
	if cookie, err := c.Request.Cookie(cookieKey); err == nil {
		fmt.Fprintln(c.Writer, cookie)
	} else {
		fmt.Fprintln(c.Writer, "")
	}
}
*/
