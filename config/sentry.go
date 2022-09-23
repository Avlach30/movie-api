package config

import (
	"fmt"
	"movie-api/helper"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func SentryConfig(router *gin.Engine) gin.HandlerFunc {
	SENTRY_DSN := helper.GetEnvValue("SENTRY_DSN")

	err := sentry.Init(sentry.ClientOptions{
		Dsn: SENTRY_DSN,
		TracesSampleRate: 1.0,
	})
	if (err != nil) {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	sentryInit := sentrygin.New(sentrygin.Options{})

	return sentryInit
}