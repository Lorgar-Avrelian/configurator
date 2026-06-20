package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func BindParam(ctx context.Context, d dto.ComponentParamLinkDto) error {
	var compID int64
	var paramID int64
	var err error
	compID = d.ComponentID
	paramID = d.ParamID
	err = dao.BindParam(ctx, compID, paramID)
	return err
}

func UnbindParam(ctx context.Context, componentID int64, paramID int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.UnbindParam(ctx, componentID, paramID)
	return found, err
}

func GetComponentsByDirectParamID(ctx context.Context, paramID int64) ([]dto.ComponentDto, error) {
	var arr []dao.ComponentDao
	var err error
	arr, err = dao.GetComponentsByDirectParamID(ctx, paramID)
	if err != nil {
		return nil, err
	}
	var res []dto.ComponentDto
	res = mapper.ComponentDaoArrayToComponentDtoArray(arr)
	return res, nil
}
