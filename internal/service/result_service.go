package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func GetAllParamResults(ctx context.Context) ([]dto.ParamResultDto, error) {
	var list []dao.Result
	var err error
	list, err = dao.GetAllParamResults(ctx)
	if err != nil {
		return nil, err
	}
	var res []dto.ParamResultDto
	res = mapper.ResultArrayToParamResultDtoArray(list)
	return res, nil
}

func GetParamResultsByFilter(ctx context.Context, d dto.ParamResultGetDto) ([]dto.ParamResultDto, error) {
	var mask string
	mask = ""
	if d.Host != nil {
		mask = mask + "1"
	} else {
		mask = mask + "0"
	}
	if d.Port != nil {
		mask = mask + "1"
	} else {
		mask = mask + "0"
	}
	if d.ComponentTitle != nil {
		mask = mask + "1"
	} else {
		mask = mask + "0"
	}
	if d.InternalOrder != nil {
		mask = mask + "1"
	} else {
		mask = mask + "0"
	}
	if d.ParamTitle != nil {
		mask = mask + "1"
	} else {
		mask = mask + "0"
	}
	var list []dao.Result
	var err error
	switch mask {
	case "00000":
		list, err = dao.GetAllParamResults(ctx)
	case "00001":
		list, err = dao.GetParamResultsByParamTitle(ctx, *d.ParamTitle)
	case "00010":
		list, err = dao.GetParamResultsByInternalOrder(ctx, *d.InternalOrder)
	case "00011":
		list, err = dao.GetParamResultsByInternalOrderAndParamTitle(ctx, *d.InternalOrder, *d.ParamTitle)
	case "00100":
		list, err = dao.GetParamResultsByComponentTitle(ctx, *d.ComponentTitle)
	case "00101":
		list, err = dao.GetParamResultsByComponentTitleAndParamTitle(ctx, *d.ComponentTitle, *d.ParamTitle)
	case "00110":
		list, err = dao.GetParamResultsByComponentTitleAndInternalOrder(ctx, *d.ComponentTitle, *d.InternalOrder)
	case "00111":
		list, err = dao.GetParamResultsByComponentTitleAndInternalOrderAndParamTitle(ctx, *d.ComponentTitle, *d.InternalOrder, *d.ParamTitle)
	case "01000":
		list, err = dao.GetParamResultsByPort(ctx, *d.Port)
	case "01001":
		list, err = dao.GetParamResultsByPortAndParamTitle(ctx, *d.Port, *d.ParamTitle)
	case "01010":
		list, err = dao.GetParamResultsByPortAndInternalOrder(ctx, *d.Port, *d.InternalOrder)
	case "01011":
		list, err = dao.GetParamResultsByPortAndInternalOrderAndParamTitle(ctx, *d.Port, *d.InternalOrder, *d.ParamTitle)
	case "01100":
		list, err = dao.GetParamResultsByPortAndComponentTitle(ctx, *d.Port, *d.ComponentTitle)
	case "01101":
		list, err = dao.GetParamResultsByPortAndComponentTitleAndParamTitle(ctx, *d.Port, *d.ComponentTitle, *d.ParamTitle)
	case "01110":
		list, err = dao.GetParamResultsByPortAndComponentTitleAndInternalOrder(ctx, *d.Port, *d.ComponentTitle, *d.InternalOrder)
	case "01111":
		list, err = dao.GetParamResultsByPortAndComponentTitleAndInternalOrderAndParamTitle(ctx, *d.Port, *d.ComponentTitle, *d.InternalOrder, *d.ParamTitle)
	case "10000":
		list, err = dao.GetParamResultsByHost(ctx, *d.Host)
	case "10001":
		list, err = dao.GetParamResultsByHostAndParamTitle(ctx, *d.Host, *d.ParamTitle)
	case "10010":
		list, err = dao.GetParamResultsByHostAndInternalOrder(ctx, *d.Host, *d.InternalOrder)
	case "10011":
		list, err = dao.GetParamResultsByHostAndInternalOrderAndParamTitle(ctx, *d.Host, *d.InternalOrder, *d.ParamTitle)
	case "10100":
		list, err = dao.GetParamResultsByHostAndComponentTitle(ctx, *d.Host, *d.ComponentTitle)
	case "10101":
		list, err = dao.GetParamResultsByHostAndComponentTitleAndParamTitle(ctx, *d.Host, *d.ComponentTitle, *d.ParamTitle)
	case "10110":
		list, err = dao.GetParamResultsByHostAndComponentTitleAndInternalOrder(ctx, *d.Host, *d.ComponentTitle, *d.InternalOrder)
	case "10111":
		list, err = dao.GetParamResultsByHostAndComponentTitleAndInternalOrderAndParamTitle(ctx, *d.Host, *d.ComponentTitle, *d.InternalOrder, *d.ParamTitle)
	case "11000":
		list, err = dao.GetParamResultsByHostAndPort(ctx, *d.Host, *d.Port)
	case "11001":
		list, err = dao.GetParamResultsByHostAndPortAndParamTitle(ctx, *d.Host, *d.Port, *d.ParamTitle)
	case "11010":
		list, err = dao.GetParamResultsByHostAndPortAndInternalOrder(ctx, *d.Host, *d.Port, *d.InternalOrder)
	case "11011":
		list, err = dao.GetParamResultsByHostAndPortAndInternalOrderAndParamTitle(ctx, *d.Host, *d.Port, *d.InternalOrder, *d.ParamTitle)
	case "11100":
		list, err = dao.GetParamResultsByHostAndPortAndComponentTitle(ctx, *d.Host, *d.Port, *d.ComponentTitle)
	case "11101":
		list, err = dao.GetParamResultsByHostAndPortAndComponentTitleAndParamTitle(ctx, *d.Host, *d.Port, *d.ComponentTitle, *d.ParamTitle)
	case "11110":
		list, err = dao.GetParamResultsByHostAndPortAndComponentTitleAndInternalOrder(ctx, *d.Host, *d.Port, *d.ComponentTitle, *d.InternalOrder)
	case "11111":
		list, err = dao.GetParamResultsByHostAndPortAndComponentTitleAndInternalOrderAndParamTitle(ctx, *d.Host, *d.Port, *d.ComponentTitle, *d.InternalOrder, *d.ParamTitle)
	}
	if err != nil {
		return nil, err
	}
	var res []dto.ParamResultDto
	res = mapper.ResultArrayToParamResultDtoArray(list)
	return res, nil
}
