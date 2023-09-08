package lottery

import (
	"gin-onter/models/models_lottery"
	"gin-onter/service/lottery"
	"gin-onter/utils/msg"
	"gin-onter/utils/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Addlottery(c *gin.Context) {
	var data models_lottery.Lottery
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}
	err := lottery.CheckLottery(data.Userid)
	if err != nil {
		// 记录错误到Gin的上下文中
		//c.Error(err)
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.AddLottery(data)
	if err != nil {
		// 记录错误到Gin的上下文中
		//c.Error(err)
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list, msg.GetErrMsg(200)))
}



func Editlottery(c *gin.Context) {
	var data models_lottery.Lottery
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(),msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.EditLottery(data)
	if err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list,msg.GetErrMsg(200)))
}
func Getlottery(c *gin.Context) {
	var data models_lottery.Lottery
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(),msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.GetLottery(data.ID)
	if err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}
	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list,msg.GetErrMsg(200)))
}
func Lotterymain(c *gin.Context) {
	var data struct{
		Numwinners int    `json:"numwinners"`
		Whatprize  string `json:"whatprize"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(),msg.GetErrMsg(500)))
		return
	}

	list, err := lottery.Lotterymain(data.Numwinners,data.Whatprize)
	if err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(500, err.Error(), msg.GetErrMsg(500)))
		return
	}
	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, list,msg.GetErrMsg(200)))
}

func Dellottery(c *gin.Context) {
	var data models_lottery.Lottery
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusOK, (&result.Result{}).Error(msg.ERROR, err.Error(), msg.GetErrMsg(msg.ERROR)))
		return
	}
	code := lottery.DelLottery(data.Userid)

	c.JSON(http.StatusOK, (&result.Result{}).Ok(200, code,msg.GetErrMsg(200)))

}