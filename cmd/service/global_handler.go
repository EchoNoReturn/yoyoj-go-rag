package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func logAllRequests() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// 请求进入
		log.Infof("[Request] %s %s", c.Method(), c.OriginalURL())

		// 执行下一步（继续路由处理）
		err := c.Next()

		// 请求返回
		log.Infof("[Response] %s %s -> %d", c.Method(), c.OriginalURL(), c.Response().StatusCode())

		return err
	}
}

func globalErrorHandler() fiber.ErrorHandler {
	// 全局错误处理函数
	// 捕获所有未处理的错误并进行统一处理
	return func(c *fiber.Ctx, err error) error {
		if err != nil {
			log.Errorf("Global error handler caught an error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal Server Error",
			})
		}
		return nil
	}
}
