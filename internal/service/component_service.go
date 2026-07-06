package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
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
	var err error
	var i int
	var closestIdx int64
	var components []dao.ComponentDao
	var found bool
	components, err = dao.GetAllComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error getting components from table public.component: %v", err)
		return false, err
	}
	found, err = dao.DeleteComponent(ctx, id)
	if err != nil {
		logger.Errorf("Error deleting component ID %d: %v", id, err)
		return false, err
	}
	if !found {
		return false, nil
	}
	if id < int64(len(components)) {
		closestIdx = 0
		components, err = dao.GetAllComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error getting components from table public.component: %v", err)
			return false, err
		}
		for i = range components {
			if components[i].ID <= id {
				continue
			}
			closestIdx = components[i].ID
			break
		}
		if closestIdx != 0 {
			_, err = ChangeComponentData(ctx, id+1, id)
			if err != nil {
				logger.Errorf("Error shifting ID from %d to %d: %v", closestIdx, id, err)
				return true, err
			}
		}
	}
	return true, nil
}
