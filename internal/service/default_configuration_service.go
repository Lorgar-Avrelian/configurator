package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
	"errors"
)

func CreateDefaultConfiguration(ctx context.Context, indicatorID int64, devCompID *int64) (*dto.DefaultConfigurationDto, error) {
	var exists bool
	var err error
	exists, err = dao.CheckDefaultConfigurationExistsByIndicator(ctx, indicatorID, nil)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("default configuration already exists for this indicator")
	}
	var rows []dao.DefaultConfiguration
	rows, err = dao.CreateDefaultConfiguration(ctx, indicatorID, devCompID)
	if err != nil {
		return nil, err
	}
	var res *dto.DefaultConfigurationDto
	res = mapper.DefaultConfigurationToDefaultConfigurationDto(rows)
	return res, nil
}

func GetDefaultConfigurationByID(ctx context.Context, id int64) (*dto.DefaultConfigurationDto, error) {
	var rows []dao.DefaultConfiguration
	var err error
	rows, err = dao.GetDefaultConfigurationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var res *dto.DefaultConfigurationDto
	res = mapper.DefaultConfigurationToDefaultConfigurationDto(rows)
	return res, nil
}

func GetAllDefaultConfigurations(ctx context.Context) ([]dto.DefaultConfigurationDto, error) {
	var list []dao.DefaultConfiguration
	var err error
	list, err = dao.GetAllDefaultConfigurations(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.DefaultConfigurationDto
	res = mapper.DefaultConfigurationArrayToDefaultConfigurationDtoArray(list)
	return res, nil
}

func UpdateDefaultConfiguration(ctx context.Context, id int64, indicatorID int64, devCompID *int64) (*dto.DefaultConfigurationDto, error) {
	var exists bool
	var err error
	exists, err = dao.CheckDefaultConfigurationExistsByIndicator(ctx, indicatorID, &id)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("default configuration already exists for this indicator")
	}
	var rows []dao.DefaultConfiguration
	rows, err = dao.UpdateDefaultConfiguration(ctx, id, indicatorID, devCompID)
	if err != nil {
		return nil, err
	}
	var res *dto.DefaultConfigurationDto
	res = mapper.DefaultConfigurationToDefaultConfigurationDto(rows)
	return res, nil
}

func DeleteDefaultConfiguration(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteDefaultConfiguration(ctx, id)
	if err != nil {
		return false, err
	}
	CompressDefaultConfigurationDependentData(ctx)
	return found, err
}
