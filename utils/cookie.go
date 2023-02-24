package utils

/*
// CookieSetService 设置cookie
func CookieSetService(c *gin.Context) {
	cookieName := &http.Cookie{
		Name:     "spring",
		Value:    "nihao",
		Path:     "/",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookieName)
}

// CookieGetService 获取cookie
func CookieGetService(c *gin.Context) {
	cookieKey := "spring"
	if cookie, err := c.Request.Cookie(cookieKey); err == nil {
		fmt.Fprintln(c.Writer, cookie)
	} else {
		fmt.Fprintln(c.Writer, "")
	}
}

// CookieClearService 删除cookie
func CookieClearService(c *gin.Context) {
	clearCookieName := &http.Cookie{
		Name:     "spring",
		Value:    "",
		Path:     "/",
		MaxAge:   -1, //<0 意味着删除cookie
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, clearCookieName)
}
*/
