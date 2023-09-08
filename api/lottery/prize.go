package lottery

import (
	"gin-onter/models/models_lottery"
	"gin-onter/service/lottery"
	"gin-onter/utils/msg"
	"gin-onter/utils/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Addprize(c *gin.Context) {
	var data models_lottery.Prize
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}
	err := lottery.CheckPrize(data.Prizename)
	if err != nil {
		// 记录错误到Gin的上下文中
		//c.Error(err)
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.AddPrize(data)
	if err != nil {
		// 记录错误到Gin的上下文中
		//c.Error(err)
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list, msg.GetErrMsg(200)))
}



func EditPrize(c *gin.Context) {
	var data models_lottery.Prize
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(),msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.EditPrize(data)
	if err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list,msg.GetErrMsg(200)))
}
func GetPrize(c *gin.Context) {
	var data models_lottery.Prize
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(),msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.GetPrize(data.ID)
	if err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}
	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list,msg.GetErrMsg(200)))
}

func DelPrize(c *gin.Context) {
	var data models_lottery.Prize
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(msg.ERROR, err.Error(), msg.GetErrMsg(msg.ERROR)))
		return
	}
	code := lottery.DelPrize(data.Prizename)

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, code,msg.GetErrMsg(200)))

}