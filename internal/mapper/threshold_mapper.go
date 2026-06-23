package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/util"
	"time"
)

func ThresholdCreateDtoToThresholdDao(d dto.ThresholdCreateDto) dao.ThresholdDao {
	var res dao.ThresholdDao
	res.Name = d.Name
	res.Description = util.StringPtrToSqlNullString(d.Description)
	res.Author = util.StringPtrToSqlNullString(d.Author)
	res.Query = d.Query
	return res
}

func ThresholdDaoToThresholdDto(d dao.ThresholdDao) dto.ThresholdDto {
	var res dto.ThresholdDto
	res.ID = d.ID
	res.Name = d.Name
	res.Description = util.SqlNullStringToStringPtr(d.Description)
	res.Author = util.SqlNullStringToStringPtr(d.Author)
	res.Query = d.Query
	if d.Created.Valid {
		var timeStr string
		timeStr = d.Created.Time.Format(time.RFC3339)
		res.Created = &timeStr
	}
	return res
}

func ThresholdDaoArrayToThresholdDtoArray(arr []dao.ThresholdDao) []dto.ThresholdDto {
	var res []dto.ThresholdDto
	res = []dto.ThresholdDto{}
	var i int
	var item dao.ThresholdDao
	for i = 0; len(arr) > i; i++ {
		item = arr[i]
		res = append(res, ThresholdDaoToThresholdDto(item))
	}
	return res
}
