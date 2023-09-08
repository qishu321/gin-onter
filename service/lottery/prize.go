package lottery

import (
	"errors"
	"gin-onter/models"
	"gin-onter/models/models_lottery"
	"gin-onter/utils/msg"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func AddPrize(data models_lottery.Prize)(list interface{},err error)  {
	dao := models_lottery.Prize{
		Prizename: data.Prizename,
		Whatprize: data.Whatprize,
		Number: data.Number,
		Label: data.Label,
	}
	if err := models.Db.Create(&dao).Error; err != nil {
		return nil, err // 返回错误
	}
	return dao, nil

}

func EditPrize(data models_lottery.Prize)(list interface{},err error)  {
	err = models.Db.Select("id").Where("prizename = ?", data.Prizename).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "Prize，修改失败，请重新提交", err
		}
		return msg.InvalidParams, err
	}
	dao := models_lottery.Prize{
		Whatprize: data.Whatprize,
		Number: data.Number,
		Label: data.Label,
	}
	if err := models.Db.Model(&dao).Where("prizename = ?",data.Prizename).Updates(dao).Error; err != nil {
		return nil, err // 返回错误
	}
	return dao, nil

}

func GetPrize(id int)([]models_lottery.Prize,error)  {
	var list []models_lottery.Prize

	if id != 0 {
		res := models.Db.Debug().Preload("Whatprizes").Where("id = ?", id).First(&list)
		return list, res.Error
	} else {
		res := models.Db.Debug().Preload("Whatprizes").Find(&list)
		return list, res.Error
	}
}

func CheckPrize(prizename string) (err error) {
	var data models_lottery.Prize
	models.Db.Select("id").Where("prizename = ?", prizename).First(&data)

	if data.ID > 0 {
		err = errors.New("prizename重复，请重新生成")
		log.Printf("Response: %d - %s - Data: %+v", http.StatusInternalServerError, err.Error(), data)
	}

	return err
}


func DelPrize(prizename string)(code int)  {
	var data models_lottery.Prize
	models.Db.Select("id").Where("prizename = ?", prizename).First(&data)

	if data.ID > 0 {
		res := models.Db.Where("Lotteries = ?",prizename).Delete(&data).Error
		if res != nil {
			return msg.ERROR
		}
		return msg.SUCCSE
	} else {
		return msg.InvalidParams
	}


}