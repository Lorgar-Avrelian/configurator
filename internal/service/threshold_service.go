package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/mapper"
	"context"
)

func CreateThreshold(ctx context.Context, input dto.ThresholdCreateDto) (*dto.ThresholdDto, error) {
	var thresholdDao dao.ThresholdDao
	thresholdDao = mapper.ThresholdCreateDtoToThresholdDao(input)
	var insertedID int64
	var err error
	insertedID, err = dao.CreateThreshold(ctx, thresholdDao)
	if err != nil {
		return nil, err
	}
	var res *dto.ThresholdDto
	res, err = GetThresholdByID(ctx, insertedID)
	return res, err
}

func GetThresholdByID(ctx context.Context, id int64) (*dto.ThresholdDto, error) {
	var th *dao.ThresholdDao
	var err error
	th, err = dao.GetThresholdByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if th == nil {
		return nil, nil
	}
	var res dto.ThresholdDto
	res = mapper.ThresholdDaoToThresholdDto(*th)
	return &res, nil
}

func GetAllThresholds(ctx context.Context) ([]dto.ThresholdDto, error) {
	var list []dao.ThresholdDao
	var err error
	list, err = dao.GetAllThresholds(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ThresholdDto
	res = mapper.ThresholdDaoArrayToThresholdDtoArray(list)
	return res, nil
}

func UpdateThreshold(ctx context.Context, id int64, input dto.ThresholdCreateDto) (*dto.ThresholdDto, error) {
	var thresholdDao dao.ThresholdDao
	thresholdDao = mapper.ThresholdCreateDtoToThresholdDao(input)
	var updated bool
	var err error
	updated, err = dao.UpdateThreshold(ctx, id, thresholdDao)
	if err != nil {
		return nil, err
	}
	if !updated {
		return nil, nil
	}
	var res *dto.ThresholdDto
	res, err = GetThresholdByID(ctx, id)
	return res, err
}

func DeleteThreshold(ctx context.Context, id int64) (bool, error) {
	var err error
	var found bool
	found, err = dao.DeleteThreshold(ctx, id)
	if err != nil {
		logger.Errorf("Error deleting threshold ID %d: %v", id, err)
		return false, err
	}
	if !found {
		return false, nil
	}
	CompressThresholdDependentData(ctx)
	return true, nil
}
