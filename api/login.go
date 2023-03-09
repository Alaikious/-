package api

import (
	"gin-api/model"
	"gin-api/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // 身份
}

type RoleJs struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

// Login countid为学号
func Login(ctx *gin.Context) {
	var loginDto LoginDto
	if err := ctx.ShouldBindJSON(&loginDto); err != nil {
		utils.Fail(ctx, "请求参数错误")
		return
	}
	var accountEntity model.AccountEntity
	if result := utils.Db.Debug().Where("countid=?", loginDto.Username).First(&accountEntity).Error; result == gorm.ErrRecordNotFound {
		utils.Fail(ctx, "账号不存在")
		println("错误账号:", loginDto.Username)
		return
	}
	if accountEntity.Password != loginDto.Password {
		utils.Fail(ctx, "密码错误")
		return
	}
	if accountEntity.Role != loginDto.Role {
		utils.Fail(ctx, "无此用户")
		return
	}

	/*
			//  生产token返回给前端
			hmacUser := utils.HmacUser{
				Id:       int(accountEntity.Id),
				Username: accountEntity.Username,
			}


		if token, err := utils.GenerateToken(hmacUser); err == nil {
			utils.Success(ctx, gin.H{
				"name":  accountEntity.Username,
				"token": token,
			})
	*/

	var session model.Sessions
	utils.Db.Debug().Where("countid=?", loginDto.Username).First(&session)
	var user = RoleJs{
		Username: accountEntity.Countid,
		Name:     accountEntity.Username,
	}

	utils.Success(ctx, gin.H{
		"user":  user,
		"token": session.Session,
	})

	cookieToken := http.Cookie{Name: "sms-session", Value: session.Session, Path: "/", MaxAge: 86400}
	http.SetCookie(ctx.Writer, &cookieToken)
	/*这个不行
	ctx.SetCookie("sms-session", token, 24*60*60, "/", "localhost", false, true)

		req, err := ctx.Request, ctx.Errors
		if err != nil {
			fmt.Println("Error creating request:error", err)
			return
		}
		cookieIt := &http.Cookie{Name: "sms-session", Value: token}
		req.AddCookie(cookieIt)
	*/
	/*
			cookie := model.Cookie{
				Id:       hmacUser.Id,
				Token:    token,
				Username: loginDto.Username,
				Role:     loginDto.Role,
			}
			utils.Db.Debug().Create(cookie)
			return
		} else {
			utils.Fail(ctx, "账号或密码错误")
			return
		}
	*/
	return
}

func LoginCookie(ctx *gin.Context) {
	var session model.Sessions
	val := ctx.GetHeader("Authorization")
	utils.Db.Where("session=?", val).First(&session)
	if session.Username == "" {
		//	utils.Fail(ctx, "fail for wrong cookie")
		return
	}
	var users model.AccountEntity
	if utils.Db.Where("username=?", session.Username).First(&users).Error != nil {
		//	utils.Fail(ctx, "error adapted user")
		return
	}
	var user = RoleJs{
		Username: users.Countid,
		Name:     users.Username,
	}

	utils.Success(ctx, gin.H{
		"user": user,
		"role": users.Role,
	})
	return
}

func GetUsername(ctx *gin.Context) model.AccountEntity { //用session返回用户本体
	var session model.Sessions
	val := ctx.GetHeader("Authorization")
	println(val)
	if utils.Db.Where("session=?", val).First(&session).Error != nil {
	}
	var user model.AccountEntity
	utils.Db.Where("username=?", session.Username).First(&user)
	return user
}

/*测试
func Test(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	fmt.Println("当前用户id", userId)
	utils.Success(ctx, "成功token")
}
*/
