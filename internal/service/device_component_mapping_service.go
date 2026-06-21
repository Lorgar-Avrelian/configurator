package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"context"
)

func BindDeviceMapping(ctx context.Context, dcID int64, mID int64) (*dto.DeviceComponentDto, error) {
	var err error
	var res *dto.DeviceComponentDto
	err = dao.BindDeviceMapping(ctx, dcID, mID)
	if err != nil {
		return nil, err
	}
	res, err = GetDeviceComponentByIDOwn(ctx, dcID)
	return res, err
}

func UnbindDeviceMapping(ctx context.Context, dcID int64, mID int64) (*dto.DeviceComponentDto, error) {
	var found bool
	var err error
	var res *dto.DeviceComponentDto
	found, err = dao.UnbindDeviceMapping(ctx, dcID, mID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	res, err = GetDeviceComponentByIDOwn(ctx, dcID)
	return res, err
}
