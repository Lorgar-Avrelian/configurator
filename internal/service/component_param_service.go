package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func BindParam(ctx context.Context, componentID int64, paramID int64) (*dto.ComponentDto, error) {
	var err error
	var compPtr *dao.Component
	var res dto.ComponentDto
	err = dao.BindParam(ctx, componentID, paramID)
	if err != nil {
		return nil, err
	}
	compPtr, err = dao.GetComponentByID(ctx, componentID)
	if err != nil {
		return nil, err
	}
	if compPtr == nil {
		return nil, nil
	}
	res = mapper.ComponentToComponentDto(*compPtr)
	return &res, nil
}

func UnbindParam(ctx context.Context, componentID int64, paramID int64) (*dto.ComponentDto, error) {
	var found bool
	var err error
	var compPtr *dao.Component
	var res dto.ComponentDto
	found, err = dao.UnbindParam(ctx, componentID, paramID)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, nil
	}
	compPtr, err = dao.GetComponentByID(ctx, componentID)
	if err != nil {
		return nil, err
	}
	if compPtr == nil {
		return nil, nil
	}
	res = mapper.ComponentToComponentDto(*compPtr)
	return &res, nil
}
