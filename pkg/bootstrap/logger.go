package bootstrap

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

func RequestResponseLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		today := time.Now().Format("2006-01-02")
		logFile := filepath.Join("storage", "logs", today+"_request.log")
		f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("Could not open log file: %v", err))
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)

		writer := &bodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		c.Next()

		log := fmt.Sprintf(
			"[%s] %s %s %d %s \nRequestBody: %s\nResponseBody: %s\nIP=%s\n\n",
			time.Now().Format("15:04:05"),
			c.Request.Method,
			c.Request.URL.String(),
			c.Writer.Status(),
			time.Since(start),
			string(bodyBytes),
			writer.body.String(),
			c.ClientIP(),
		)
		_, err = f.WriteString(log)
		if err != nil {
			return
		}
	}
}

type bodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
