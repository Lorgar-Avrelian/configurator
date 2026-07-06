package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/mapper"
	"context"
)

func CreateParam(ctx context.Context, d dto.ParamCreateDto) (*dto.ParamDto, error) {
	var daoInput dao.ParamDao
	daoInput = mapper.ParamCreateDtoToParamDao(d)
	var paramPtr *dao.ParamDao
	var err error
	paramPtr, err = dao.CreateParam(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	var paramFull dao.ParamDao
	paramFull = *paramPtr
	var res dto.ParamDto
	res = mapper.ParamDaoToParamDto(paramFull)
	return &res, nil
}

func GetParamByID(ctx context.Context, id int64) (*dto.ParamDto, error) {
	var paramPtr *dao.ParamDao
	var err error
	paramPtr, err = dao.GetParamByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if paramPtr == nil {
		return nil, nil
	}
	var paramFull dao.ParamDao
	paramFull = *paramPtr
	var res dto.ParamDto
	res = mapper.ParamDaoToParamDto(paramFull)
	return &res, nil
}

func GetAllParams(ctx context.Context) ([]dto.ParamDto, error) {
	var arr []dao.ParamDao
	var err error
	arr, err = dao.GetAllParams(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ParamDto
	res = []dto.ParamDto{}
	var p dao.ParamDao
	for _, p = range arr {
		res = append(res, mapper.ParamDaoToParamDto(p))
	}
	return res, nil
}

func UpdateParam(ctx context.Context, id int64, d dto.ParamUpdateDto) (*dto.ParamDto, error) {
	var daoInput dao.ParamDao
	daoInput = mapper.ParamUpdateDtoToParamDao(d)
	var paramPtr *dao.ParamDao
	var err error
	paramPtr, err = dao.UpdateParam(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	if paramPtr == nil {
		return nil, nil
	}
	var paramFull dao.ParamDao
	paramFull = *paramPtr
	var res dto.ParamDto
	res = mapper.ParamDaoToParamDto(paramFull)
	return &res, nil
}

func DeleteParam(ctx context.Context, id int64) (bool, error) {
	var err error
	var params []dao.ParamDao
	var found bool
	params, err = dao.GetAllParamDao(ctx)
	if err != nil {
		logger.Errorf("Error getting params from table public.param: %v", err)
		return false, err
	}
	found, err = dao.DeleteParam(ctx, id)
	if err != nil {
		logger.Errorf("Error deleting param ID %d: %v", id, err)
		return false, err
	}
	if !found {
		return false, nil
	}
	if id < int64(len(params)) {
		_, err = ChangeParamData(ctx, id+1, id)
		if err != nil {
			logger.Errorf("Error shifting ID from %d to %d: %v", id+1, id, err)
			return true, err
		}
	}
	return true, nil
}

func GetUnattachedParams(ctx context.Context) ([]dto.ParamDto, error) {
	var arr []dao.ParamDao
	var err error
	arr, err = dao.GetUnattachedParams(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ParamDto
	res = []dto.ParamDto{}
	var p dao.ParamDao
	for _, p = range arr {
		res = append(res, mapper.ParamDaoToParamDto(p))
	}
	return res, nil
}

func SearchParams(ctx context.Context, query string) ([]dto.ParamDto, error) {
	var arr []dao.ParamDao
	var err error
	arr, err = dao.SearchParams(ctx, query)
	if err != nil {
		return nil, err
	}
	var res []dto.ParamDto
	res = []dto.ParamDto{}
	var p dao.ParamDao
	for _, p = range arr {
		res = append(res, mapper.ParamDaoToParamDto(p))
	}
	return res, nil
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
