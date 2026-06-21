package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func CreateDeviceIndicator(ctx context.Context, d dto.DeviceIndicatorCreateDto) (*dto.DeviceIndicatorDto, error) {
	var daoInput dao.DeviceIndicatorDao
	var resPtr *dao.DeviceIndicatorDao
	var err error
	var res dto.DeviceIndicatorDto
	daoInput = mapper.DeviceIndicatorCreateDtoToDeviceIndicatorDao(d)
	resPtr, err = dao.CreateIndicator(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	res = mapper.DeviceIndicatorDaoToDeviceIndicatorDto(*resPtr)
	return &res, nil
}

func GetDeviceIndicatorByID(ctx context.Context, id int64) (*dto.DeviceIndicatorDto, error) {
	var resPtr *dao.DeviceIndicatorDao
	var err error
	var res dto.DeviceIndicatorDto
	resPtr, err = dao.GetIndicatorByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if resPtr == nil {
		return nil, nil
	}
	res = mapper.DeviceIndicatorDaoToDeviceIndicatorDto(*resPtr)
	return &res, nil
}

func GetAllDeviceIndicators(ctx context.Context) ([]dto.DeviceIndicatorDto, error) {
	var arr []dao.DeviceIndicatorDao
	var err error
	var res []dto.DeviceIndicatorDto
	arr, err = dao.GetAllIndicators(ctx)
	if err != nil {
		return nil, err
	}
	res = mapper.DeviceIndicatorDaoArrayToDeviceIndicatorDtoArray(arr)
	return res, nil
}

func UpdateDeviceIndicator(ctx context.Context, id int64, d dto.DeviceIndicatorCreateDto) (*dto.DeviceIndicatorDto, error) {
	var daoInput dao.DeviceIndicatorDao
	var resPtr *dao.DeviceIndicatorDao
	var err error
	var res dto.DeviceIndicatorDto
	daoInput = mapper.DeviceIndicatorCreateDtoToDeviceIndicatorDao(d)
	resPtr, err = dao.UpdateIndicator(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	if resPtr == nil {
		return nil, nil
	}
	res = mapper.DeviceIndicatorDaoToDeviceIndicatorDto(*resPtr)
	return &res, nil
}

func DeleteDeviceIndicator(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteIndicator(ctx, id)
	return found, err
}
