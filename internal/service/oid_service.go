package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"configurator/internal/model"
	"context"
	"strings"
)

func GetOidsByExactNotation(ctx context.Context, notation string) ([]dto.OidDto, error) {
	var target string
	target = notation
	if !strings.HasPrefix(target, ".") {
		target = "." + target
	}
	var arr []dao.Oid
	var err error
	arr, err = dao.GetOidsByExactDotter(ctx, target)
	if err != nil {
		return nil, err
	}
	var res []dto.OidDto
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}

func GetOidsByPrefixNotation(ctx context.Context, prefix string, page int) ([]dto.OidDto, error) {
	var target string
	target = prefix
	if !strings.HasPrefix(target, ".") {
		target = "." + target
	}
	var arr []dao.Oid
	var err error
	arr, err = dao.GetOidsByDotterPrefixPaged(ctx, target, page)
	if err != nil {
		return nil, err
	}
	var res []dto.OidDto
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}

func GetOidsByMib(ctx context.Context, name string) ([]dto.OidDto, error) {
	var arr []dao.Oid
	var err error
	arr, err = dao.GetOidsByMibName(ctx, name)
	if err != nil {
		return nil, err
	}
	var res []dto.OidDto
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}

func GetOidsByVendor(ctx context.Context, vendor string, page int) ([]dto.OidDto, error) {
	var vKey model.Vendor
	var vendorID int64
	var filterByVendor bool
	var arr []dao.Oid
	var err error
	var res []dto.OidDto
	filterByVendor = false
	vendorID = 0
	if vendor != "" {
		vKey = model.ParseVendor(vendor)
		vendorID = int64(vKey)
		filterByVendor = true
	}
	arr, err = dao.GetOidsByVendorIDPaged(ctx, vendorID, filterByVendor, page)
	if err != nil {
		return nil, err
	}
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}

func GetOidsByDotterMibAndVendor(ctx context.Context, notation string, mibName string, vendor *string) ([]dto.OidDto, error) {
	var target string
	target = notation
	if !strings.HasPrefix(target, ".") {
		target = "." + target
	}
	var arr []dao.Oid
	var err error
	arr, err = dao.GetOidsByDotterMibAndVendor(ctx, target, mibName, vendor)
	if err != nil {
		return nil, err
	}
	var res []dto.OidDto
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}
