package lottery

import (
	"errors"
	"gin-onter/models"
	"gin-onter/models/models_lottery"
	"gin-onter/utils/msg"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func AddLottery(data models_lottery.Lottery)(list interface{},err error)  {
	if data.Userid == 0 || data.Username == "" {
		return nil, errors.New("userid和username字段都是必填的")
	}
	dao := models_lottery.Lottery{
		Userid: data.Userid,
		Username: data.Username,
		Label: data.Label,
	}
	if err := models.Db.Create(&dao).Error; err != nil {
		return nil, err // 返回错误
	}
	return dao, nil

}

func EditLottery(data models_lottery.Lottery)(list interface{},err error)  {
	err = models.Db.Select("id").Where("userid = ?", data.Userid).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "没有这个userid，修改失败，请重新提交", err
		}
		return msg.InvalidParams, err
	}
	dao := models_lottery.Lottery{
		Username: data.Username,
		Label: data.Label,
	}
	if err := models.Db.Model(&dao).Where("userid = ?",data.Userid).Updates(dao).Error; err != nil {
		return nil, err // 返回错误
	}
	return dao, nil

}
// Lotterymain 执行抽奖主要逻辑
func Lotterymain(numWinners int,whatprize string) ([]models_lottery.Lottery, error) {
	// 获取所有 status 为 0 的人
	var candidates []models_lottery.Lottery
	if err := models.Db.Debug().Where("Status = ?", 0).Find(&candidates).Error; err != nil {
		return nil, err
	}

	if len(candidates) == 0 {
		return nil, nil // 没有符合条件的人，可以根据需要返回相应的信息
	}

	rand.Seed(time.Now().UnixNano()) // 设置随机种子

	var winners []models_lottery.Lottery
	for i := 0; i < numWinners && len(candidates) > 0; i++ {
		winnerIndex := rand.Intn(len(candidates))
		winner := candidates[winnerIndex]

		// 更新中奖人的信息
		winner.Status = 1
		winner.Whatprize = whatprize

		if err := models.Db.Model(&winner).Updates(winner).Error; err != nil {
			return nil, err
		}

		winners = append(winners, winner)

		// 从候选人中移除已经中奖的人
		candidates = append(candidates[:winnerIndex], candidates[winnerIndex+1:]...)
	}

	return winners, nil
}



func GetLottery(id int)(list interface{},err error)  {
	var data []models_lottery.Lottery

	if id != 0 {
		res := models.Db.Debug().Where("id = ?",id).Find(&data)
		return data, res.Error
	} else {
		res := models.Db.Debug().Find(&data)
		return data, res.Error
	}
}

func CheckLottery(userid int64) (err error) {
	var data models_lottery.Lottery
	models.Db.Select("id").Where("userid = ?", userid).First(&data)

	if data.ID > 0 {
		err = errors.New("userid重复，请重新生成")
		log.Printf("Response: %d - %s - Data: %+v", http.StatusInternalServerError, err.Error(), data)
	}

	return err
}


func DelLottery(userid int64)(code int)  {
	var data models_lottery.Lottery
	models.Db.Select("id").Where("userid = ?", userid).First(&data)

	if data.ID > 0 {
		res := models.Db.Where("userid = ?",userid).Delete(&data).Error
		if res != nil {
			return msg.ERROR
		}
		return msg.SUCCSE
	} else {
		return msg.InvalidParams
	}


}