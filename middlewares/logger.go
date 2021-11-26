package middlewares

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"site_backend/helper"
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
			"_hostname":   hostname,
			"_statusCode": statusCode,
			"_latency":    latency, // time to process
			"_clientIP":   clientIP,
			"_method":     c.Request.Method,
			"_path":       path,
			"_dataLength": dataLength,
			"_userAgent":  clientUserAgent,
			"_userID":     userID,
		}
		Logger1(entry)
	}
}

func Logger1(logfile interface{}) {
	namefile := helper.GetDate()
	str := fmt.Sprintf("%v", logfile)
	file, err := os.OpenFile("./logs/"+namefile+".log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	if _, err := file.WriteString(helper.GetDateTime() + " " + str + "\n"); err != nil {
		log.Println(err)
		fmt.Println("Log Error: ", err)
	}
}
