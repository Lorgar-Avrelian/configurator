package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func CreateComponent(ctx context.Context, d dto.ComponentCreateDto) (*dto.ComponentDto, error) {
	var daoInput dao.ComponentDao
	daoInput = mapper.ComponentCreateDtoToComponentDao(d)
	var compPtr *dao.Component
	var err error
	compPtr, err = dao.CreateComponent(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	var compFull dao.Component
	compFull = *compPtr
	var res dto.ComponentDto
	res = mapper.ComponentToComponentDto(compFull)
	return &res, nil
}

func GetComponentByID(ctx context.Context, id int64) (*dto.ComponentDto, error) {
	var compPtr *dao.Component
	var err error
	compPtr, err = dao.GetComponentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if compPtr == nil {
		return nil, nil
	}
	var compFull dao.Component
	compFull = *compPtr
	var res dto.ComponentDto
	res = mapper.ComponentToComponentDto(compFull)
	return &res, nil
}

func GetAllComponents(ctx context.Context) ([]dto.ComponentDto, error) {
	var arr []dao.Component
	var err error
	arr, err = dao.GetAllComponents(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ComponentDto
	res = mapper.ComponentArrayToComponentDtoArray(arr)
	return res, nil
}

func SearchComponents(ctx context.Context, query string) ([]dto.ComponentDto, error) {
	var arr []dao.Component
	var err error
	arr, err = dao.SearchComponents(ctx, query)
	if err != nil {
		return nil, err
	}
	var res []dto.ComponentDto
	res = mapper.ComponentArrayToComponentDtoArray(arr)
	return res, nil
}

func UpdateComponent(ctx context.Context, id int64, d dto.ComponentUpdateDto) (*dto.ComponentDto, error) {
	var daoInput dao.ComponentDao
	daoInput = mapper.ComponentUpdateDtoToComponentDao(d)
	var compPtr *dao.Component
	var err error
	compPtr, err = dao.UpdateComponent(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	if compPtr == nil {
		return nil, nil
	}
	var compFull dao.Component
	compFull = *compPtr
	var res dto.ComponentDto
	res = mapper.ComponentToComponentDto(compFull)
	return &res, nil
}

func DeleteComponent(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteComponent(ctx, id)
	return found, err
}
