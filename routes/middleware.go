package routes

import (
	"github.com/JuanTorr/project/controller/analytics"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

//Analitycs monitoring module
func Analitycs(r *gin.Engine) {
	r.Use(analytics.AnalizeRequest)
}

//RedirectHTTPS redirects from http to https protocol
func RedirectHTTPS(r *gin.Engine) {
	secureFunc := func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:3000",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
	r.Use(secureFunc)
}