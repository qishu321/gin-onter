package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

type Level int

const (
	INFO Level = iota
	WARNING
	ERROR
	FATAL
)

var (
	file *os.File
	e    error
)
type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 日志中间件。
func Logger() gin.HandlerFunc {
	// Create or open the log file
	file, e := os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
	if e != nil {
		panic(e)
	}
	//defer file.Close() // Ensure the file is closed when the handler exits

	gin.DefaultWriter = io.MultiWriter(file, os.Stdout) // Output to both file and console

	g := func(c *gin.Context) {
		// Log Gin's debug messages
		gin.DefaultWriter.Write([]byte(fmt.Sprintf(c.Errors.ByType(gin.ErrorTypePrivate).String())))

		// 使用自定义 ResponseWriter
		crw := &CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = crw
		c.Next()

		// 记录回包内容和处理时间
		respBody := string(crw.body.Bytes())

		// Open the log file again for writing
		file, err := os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Create a logger with the log file
		logger := log.New(file, "[INFO]: ", log.LstdFlags)
		logger.Printf("%s %s (%v)\n", c.Request.Method, c.Request.RequestURI, respBody)
	}

	return g
}



//func CustomLogger() gin.HandlerFunc {
//	gin.DisableConsoleColor()
//	file, e = os.OpenFile("logs/gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0664)
//	if e != nil {
//		panic(e)
//	}
//	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
//
//	//// 创建一个临时的 gin.Context，以便在日志中记录响应内容
//	//tmpContext := &gin.Context{}
//	//crw := &CustomResponseWriter{
//	//	body:           bytes.NewBufferString(""),
//	//	ResponseWriter: tmpContext.Writer,
//	//}
//	//tmpContext.Writer = crw
//	//tmpContext.Next()
//	//
//	//respBody := string(crw.body.Bytes())
//
//	g := gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
//		levelFlags := []string{"INFO", "WARN", "ERROR", "FATAL"}
//		var level string
//		status := param.StatusCode
//
//		switch {
//		case status > 499:
//			level = levelFlags[FATAL]
//		case status > 399:
//			level = levelFlags[ERROR]
//		case status > 299:
//			level = levelFlags[WARNING]
//		default:
//			level = levelFlags[INFO]
//		}
//
//		return fmt.Sprintf("[%s] - %s - [%s] \"%s %s %s %d %s \"%s\" %s\" %s\" %s\"\n",
//			level,
//			param.ClientIP,
//			param.TimeStamp.Format(conf.TimeFormat),
//			param.Method,
//			param.Path,
//			param.Request.Proto,
//			status,
//			param.Latency,
//			param.Request.UserAgent(),
//			param.ErrorMessage,
//			param.Request.RequestURI,
//		)
//	})
//
//	return g
//}