package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func CreateConfiguration(ctx context.Context, indicatorID int64, devCompID *int64) (*dto.ConfigurationDto, error) {
	var rows []dao.Configuration
	var err error
	rows, err = dao.CreateConfiguration(ctx, indicatorID, devCompID)
	if err != nil {
		return nil, err
	}
	var res *dto.ConfigurationDto
	res = mapper.ConfigurationToConfigurationDto(rows)
	return res, nil
}

func GetConfigurationByID(ctx context.Context, id int64) (*dto.ConfigurationDto, error) {
	var rows []dao.Configuration
	var err error
	rows, err = dao.GetConfigurationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	var res *dto.ConfigurationDto
	res = mapper.ConfigurationToConfigurationDto(rows)
	return res, nil
}

func GetAllConfigurations(ctx context.Context) ([]dto.ConfigurationDto, error) {
	var list []dao.Configuration
	var err error
	list, err = dao.GetAllConfigurations(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ConfigurationDto
	res = mapper.ConfigurationArrayToConfigurationDtoArray(list)
	return res, nil
}

func UpdateConfiguration(ctx context.Context, id int64, indicatorID int64, devCompID *int64) (*dto.ConfigurationDto, error) {
	var rows []dao.Configuration
	var err error
	rows, err = dao.UpdateConfiguration(ctx, id, indicatorID, devCompID)
	if err != nil {
		return nil, err
	}
	var res *dto.ConfigurationDto
	res = mapper.ConfigurationToConfigurationDto(rows)
	return res, nil
}

func DeleteConfiguration(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteConfiguration(ctx, id)
	if err != nil {
		return false, err
	}
	CompressConfigurationDependentData(ctx)
	return found, err
}
