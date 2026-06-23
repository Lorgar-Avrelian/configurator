package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func GetAllConfigInProcess(ctx context.Context) ([]dto.ConfigInProcessDto, error) {
	var list []dao.ConfigInProcessDao
	var err error
	list, err = dao.GetAllConfigInProcess(ctx)
	if err != nil {
		return nil, err
	}
	if 0 >= len(list) {
		return nil, nil
	}
	var res []dto.ConfigInProcessDto
	res = mapper.ConfigInProcessDaoArrayToConfigInProcessDtoArray(list)
	return res, nil
}

func SearchConfigInProcess(ctx context.Context, host *string, port *int32) ([]dto.ConfigInProcessDto, error) {
	var list []dao.ConfigInProcessDao
	var err error
	list, err = dao.SearchConfigInProcess(ctx, host, port)
	if err != nil {
		return nil, err
	}
	if 0 >= len(list) {
		return nil, nil
	}
	var res []dto.ConfigInProcessDto
	res = mapper.ConfigInProcessDaoArrayToConfigInProcessDtoArray(list)
	return res, nil
}
