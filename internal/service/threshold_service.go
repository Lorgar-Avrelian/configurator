package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/mapper"
	"context"
	"errors"
)

func CreateThreshold(ctx context.Context, input dto.ThresholdCreateDto) (*dto.ThresholdDto, error) {
	var thresholdDao dao.ThresholdDao
	var err error
	thresholdDao, err = mapper.ThresholdCreateDtoToThresholdDao(input)
	if err != nil {
		logger.Errorf("Validation failed during threshold creation mapping: %v", err)
		return nil, err
	}
	var insertedID int64
	insertedID, err = dao.CreateThreshold(ctx, thresholdDao)
	if err != nil {
		logger.Errorf("Service layer failed to create threshold: %v", err)
		return nil, err
	}
	var res *dto.ThresholdDto
	res, err = GetThresholdByID(ctx, insertedID)
	if err != nil {
		logger.Errorf("Service layer failed to retrieve created threshold: %v", err)
		return nil, err
	}
	return res, nil
}

func GetThresholdByID(ctx context.Context, id int64) (*dto.ThresholdDto, error) {
	var th *dao.ThresholdDao
	var err error
	th, err = dao.GetThresholdByID(ctx, id)
	if err != nil {
		logger.Errorf("Service layer failed to get threshold by ID %d: %v", id, err)
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
		logger.Errorf("Service layer failed to get all thresholds: %v", err)
		return nil, err
	}
	var res []dto.ThresholdDto
	res = mapper.ThresholdDaoArrayToThresholdDtoArray(list)
	return res, nil
}

func UpdateThreshold(ctx context.Context, id int64, input dto.ThresholdCreateDto) (*dto.ThresholdDto, error) {
	var thresholdDao dao.ThresholdDao
	var err error
	thresholdDao, err = mapper.ThresholdCreateDtoToThresholdDao(input)
	if err != nil {
		logger.Errorf("Validation failed during threshold update mapping: %v", err)
		return nil, err
	}
	var updated bool
	updated, err = dao.UpdateThreshold(ctx, id, thresholdDao)
	if err != nil {
		logger.Errorf("Service layer failed to update threshold ID %d: %v", id, err)
		return nil, err
	}
	if !updated {
		return nil, nil
	}
	var res *dto.ThresholdDto
	res, err = GetThresholdByID(ctx, id)
	if err != nil {
		logger.Errorf("Service layer failed to retrieve updated threshold ID %d: %v", id, err)
		return nil, err
	}
	return res, nil
}

func DeleteThreshold(ctx context.Context, id int64) (bool, error) {
	var err error
	var found bool
	found, err = dao.DeleteThreshold(ctx, id)
	if err != nil {
		logger.Errorf("Service error deleting threshold ID %d: %v", id, err)
		return false, err
	}
	if !found {
		return false, nil
	}
	CompressThresholdDependentData(ctx)
	return true, nil
}

func CreateThresholdFromString(ctx context.Context, expr string) (dto.ThresholdDto, error) {
	var createDto dto.ThresholdCreateDto
	var err error
	createDto, err = mapper.StringToThresholdCreateDto(expr)
	if err != nil {
		logger.Errorf("Failed to parse string expression for creation: %v", err)
		return dto.ThresholdDto{}, err
	}
	var res *dto.ThresholdDto
	res, err = CreateThreshold(ctx, createDto)
	if err != nil {
		return dto.ThresholdDto{}, err
	}
	if res == nil {
		return dto.ThresholdDto{}, errors.New("received empty entity threshold dto pointer data")
	}
	return *res, nil
}

func UpdateThresholdFromString(ctx context.Context, id int64, expr string) (dto.ThresholdDto, error) {
	var createDto dto.ThresholdCreateDto
	var err error
	createDto, err = mapper.StringToThresholdCreateDto(expr)
	if err != nil {
		logger.Errorf("Failed to parse string expression for update: %v", err)
		return dto.ThresholdDto{}, err
	}
	var daoObj dao.ThresholdDao
	daoObj, err = mapper.ThresholdCreateDtoToThresholdDao(createDto)
	if err != nil {
		return dto.ThresholdDto{}, err
	}
	var success bool
	success, err = dao.UpdateThreshold(ctx, id, daoObj)
	if err != nil {
		return dto.ThresholdDto{}, err
	}
	if !success {
		return dto.ThresholdDto{}, errors.New("threshold configuration update target entity not found")
	}
	var updated *dao.ThresholdDao
	updated, err = dao.GetThresholdByID(ctx, id)
	if err != nil {
		return dto.ThresholdDto{}, err
	}
	if updated == nil {
		return dto.ThresholdDto{}, errors.New("failed to fetch updated threshold entity by specified id")
	}
	return mapper.ThresholdDaoToThresholdDto(*updated), nil
}

func GetThresholdStringByID(ctx context.Context, id int64) (string, error) {
	var daoObj *dao.ThresholdDao
	var err error
	daoObj, err = dao.GetThresholdByID(ctx, id)
	if err != nil {
		logger.Errorf("Failed to fetch threshold ID %d for string rendering: %v", id, err)
		return "", err
	}
	if daoObj == nil {
		return "", errors.New("requested threshold configuration entry data not found by id")
	}
	var d dto.ThresholdDto
	d = mapper.ThresholdDaoToThresholdDto(*daoObj)
	var expr string
	expr = mapper.ThresholdDtoToString(d)
	return expr, nil
}
