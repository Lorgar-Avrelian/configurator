package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/util"
	"encoding/json"
	"time"
)

func ThresholdCreateDtoToThresholdDao(d dto.ThresholdCreateDto) dao.ThresholdDao {
	var res dao.ThresholdDao
	var queryBytes []byte
	var targetBytes []byte
	var err error
	queryBytes, err = json.Marshal(d.Query)
	if err != nil {
		logger.Errorf("Failed to marshal query expression to json: %v", err)
		queryBytes = []byte("{}")
	}
	targetBytes, err = json.Marshal(d.Target)
	if err != nil {
		logger.Errorf("Failed to marshal target to json: %v", err)
		targetBytes = []byte("{}")
	}
	res.Name = d.Name
	res.Description = util.StringPtrToSqlNullString(d.Description)
	res.Author = d.Author
	res.Query = queryBytes
	res.Target = targetBytes
	res.Value = d.Value
	return res
}

func ThresholdDaoToThresholdDto(d dao.ThresholdDao) dto.ThresholdDto {
	var res dto.ThresholdDto
	var qExpr dto.ThresholdExpressionDto
	var tTarg dto.ThresholdTargetDto
	var err error
	err = json.Unmarshal(d.Query, &qExpr)
	if err != nil {
		logger.Errorf("Failed to unmarshal query expression from json: %v", err)
	}
	err = json.Unmarshal(d.Target, &tTarg)
	if err != nil {
		logger.Errorf("Failed to unmarshal target from json: %v", err)
	}
	res.ID = d.ID
	res.Name = d.Name
	res.Description = util.SqlNullStringToStringPtr(d.Description)
	res.Author = d.Author
	res.Query = qExpr
	res.Target = tTarg
	res.Value = d.Value
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
