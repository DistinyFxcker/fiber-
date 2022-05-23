package controller

import (
	"go_web_example/cmd"
	//"cdnFiber/app/adminService/cmd"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"go_web_example/infrastructure/valid"
	"go_web_example/utility/errors"
	"net/http"
)

func BindReq(ctx *gin.Context, v interface{}) (err error) {
	err = ctx.ShouldBind(v)
	if err != nil {
		return err
	}

	if err = valid.Validate(v); err != nil {
		return err
	}

	return nil
}

func SuccessRes(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("操作成功"),
		"data":    v,
	})
	return
}

//func ErrorRes(c *fiber.Ctx, msg string, Code int) error {
//	code := fiber.StatusBadRequest
//	if Code != 0 {
//		code = Code
//	}
//	return c.Status(code).JSON(fiber.Map{
//		"message": fmt.Sprintf("操作失败: %s", msg),
//	})
//}

func ErrorRes(c *gin.Context, msg string, err error) {
	Info := fiber.Map{}
	if err != nil {
		errorType := errors.GetType(err)
		switch errorType {
		case errors.NoType:
			cmd.ZapLog.Error(err.Error())
			Info["message"] = fmt.Sprintf("操作失败: %s", msg)
		default:
			Info["message"] = fmt.Sprintf("操作失败: %s", msg)
		}
	} else {
		Info["message"] = fmt.Sprintf("操作失败: %s", msg)
	}

	c.JSON(http.StatusInternalServerError, Info)
	return
}
