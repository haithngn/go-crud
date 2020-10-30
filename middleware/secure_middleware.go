package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SecureMiddleware struct {
}

func (middleware *SecureMiddleware) EnsureOTP() gin.HandlerFunc {
	return func(context *gin.Context) {
		otp := context.GetHeader("OTP")
		if len(otp) == 0 {
			context.AbortWithStatus(http.StatusForbidden)
			return
		}

		context.Next()
	}
}
