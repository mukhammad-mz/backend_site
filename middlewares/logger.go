package middlewares

import (
	"bytes"
	"io"
	"io/ioutil"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

//Logger ...dfgh
func Logger() gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		defer c.Request.Body.Close()

		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		clientUserAgent := c.Request.UserAgent()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		var userID string
		uID, ok := c.Get("userUID")
		if ok {
			userID = uID.(string)
		}
		
		entry := log.Fields{
			"_hostname":     hostname,
			"_statusCode":   statusCode,
			"_latency":      latency, // time to process
			"_clientIP":     clientIP,
			"_method":       c.Request.Method,
			"_path":         path,
			"_dataLength":   dataLength,
			"_userAgent":    clientUserAgent,
			"_userID":       userID,
			"_requestBody":  string(body),
			"_responseBody": blw.body.String(),
		}
		
		if len(c.Errors) > 0 {
			log.WithFields(entry).Errorf("telemetry error: %v", c.Errors)
		} else {
			if statusCode > 499 {
				log.WithFields(entry).Error("telemetry")
			} else if statusCode > 399 {
				log.WithFields(entry).Warn("telemetry")
			} else {
				log.WithFields(entry).Info("telemetry")
			}
		}
	}
}
