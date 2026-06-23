package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
)

func ConfigInProcessDaoToConfigInProcessDto(d dao.ConfigInProcessDao) dto.ConfigInProcessDto {
	var res dto.ConfigInProcessDto
	res.Host = d.Host
	res.Port = d.Port
	res.Protocol = model.PollingProtocol(d.Protocol).String()
	return res
}

func ConfigInProcessDaoArrayToConfigInProcessDtoArray(arr []dao.ConfigInProcessDao) []dto.ConfigInProcessDto {
	var res []dto.ConfigInProcessDto
	res = []dto.ConfigInProcessDto{}
	var i int
	var item dao.ConfigInProcessDao
	for i = 0; len(arr) > i; i++ {
		item = arr[i]
		res = append(res, ConfigInProcessDaoToConfigInProcessDto(item))
	}
	return res
}
