package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
	"errors"

	"github.com/google/uuid"
)

func CreateParamIndicator(ctx context.Context, d dto.ParamIndicatorCreateDto) (*dto.ParamIndicatorDto, error) {
	var daoInput dao.ParamIndicatorDao
	var resPtr *dao.ParamIndicator
	var err error
	var res dto.ParamIndicatorDto
	var parsedUUID uuid.UUID
	var notation string
	if d.OidID != nil && *d.OidID != "" {
		parsedUUID, _ = uuid.Parse(*d.OidID)
		notation, err = dao.GetDotterNotationByOidID(ctx, parsedUUID)
		if err != nil {
			return nil, err
		}
		if notation == "" {
			return nil, errors.New("specified OID ID does not exist in the system")
		}
		d.DotterNotation = &notation
	}
	daoInput = mapper.ParamIndicatorCreateDtoToParamIndicatorDao(d)
	resPtr, err = dao.CreateParamIndicator(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	res = mapper.ParamIndicatorToParamIndicatorDto(*resPtr)
	return &res, nil
}

func GetParamIndicatorByID(ctx context.Context, id int64) (*dto.ParamIndicatorDto, error) {
	var resPtr *dao.ParamIndicator
	var err error
	var res dto.ParamIndicatorDto
	resPtr, err = dao.GetParamIndicatorByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if resPtr == nil {
		return nil, nil
	}
	res = mapper.ParamIndicatorToParamIndicatorDto(*resPtr)
	return &res, nil
}

func GetAllParamIndicators(ctx context.Context) ([]dto.ParamIndicatorDto, error) {
	var arr []dao.ParamIndicator
	var err error
	var res []dto.ParamIndicatorDto
	arr, err = dao.GetAllParamIndicators(ctx)
	if err != nil {
		return nil, err
	}
	res = mapper.ParamIndicatorArrayToParamIndicatorDtoArray(arr)
	return res, nil
}

func UpdateParamIndicator(ctx context.Context, id int64, d dto.ParamIndicatorCreateDto) (*dto.ParamIndicatorDto, error) {
	var daoInput dao.ParamIndicatorDao
	var resPtr *dao.ParamIndicator
	var err error
	var res dto.ParamIndicatorDto
	daoInput = mapper.ParamIndicatorCreateDtoToParamIndicatorDao(d)
	resPtr, err = dao.UpdateParamIndicator(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	if resPtr == nil {
		return nil, nil
	}
	res = mapper.ParamIndicatorToParamIndicatorDto(*resPtr)
	return &res, nil
}

func DeleteParamIndicator(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteParamIndicator(ctx, id)
	if err != nil {
		return false, err
	}
	CompressParamIndicatorDependentData(ctx)
	return found, err
}
