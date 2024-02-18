package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitLogger(c *fiber.Ctx, res any, code int) ILogger {
	log := &logger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		Path:       c.Path(),
		StatusCode: code,
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *logger) Print() ILogger {
	return nil
}

func (l *logger) Save() {

}

func (l *logger) SetQuery(c *fiber.Ctx) {

}

func (l *logger) SetBody(c *fiber.Ctx) {

}

func (l *logger) SetResponse(res any) {
}
