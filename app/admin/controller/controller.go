package controller

import (
	"cdnFiber/infrastructure/valid"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func BindReq(ctx *fiber.Ctx, v interface{}) (err error) {
	err = ctx.BodyParser(v)
	if err != nil {
		return err
	}

	if err = ctx.QueryParser(v); err != nil {
		return err
	}

	if err = valid.Validate(v); err != nil {
		return err
	}

	return nil
}

func SuccessRes(c *fiber.Ctx, v interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fmt.Sprintf("操作成功"),
		"data":    v,
	})
}

func ErrorRes(c *fiber.Ctx, msg string, Code int) error {
	code := fiber.StatusBadRequest
	if Code != 0 {
		code = Code
	}
	return c.Status(code).JSON(fiber.Map{
		"message": fmt.Sprintf("操作失败: %s", msg),
	})
}
