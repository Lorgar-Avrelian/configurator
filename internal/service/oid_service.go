package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
	"strings"
)

type filterFlags struct {
	hasNotation bool
	hasMib      bool
	hasVendor   bool
	isPrefix    bool
}

func GetOids(ctx context.Context, req dto.OidRequestDto) ([]dto.OidDto, error) {
	var flags filterFlags
	var target string
	var pageNum int
	var sizeNum int
	var arr []dao.Oid
	var err error
	var res []dto.OidDto
	target = req.DotterNotation
	if target != "" {
		if !strings.HasPrefix(target, ".") {
			target = "." + target
		}
	}
	flags.hasNotation = target != ""
	flags.hasMib = req.Mib != ""
	flags.hasVendor = req.Vendor != ""
	flags.isPrefix = true
	if req.Prefix != nil {
		flags.isPrefix = *req.Prefix
	}
	pageNum = 1
	if req.Page != nil {
		if *req.Page > 0 {
			pageNum = *req.Page
		}
	}
	sizeNum = 100
	if req.PageSize != nil {
		if *req.PageSize > 0 {
			sizeNum = *req.PageSize
		}
	}
	arr, err = executeDaoByFlags(ctx, flags, target, req.Mib, req.Vendor, pageNum, sizeNum)
	if err != nil {
		return nil, err
	}
	res = mapper.OidArrayToOidDtoArray(arr)
	return res, nil
}

func executeDaoByFlags(ctx context.Context, f filterFlags, notation string, mib string, vendor string, page int, size int) ([]dao.Oid, error) {
	var arr []dao.Oid
	var err error
	switch {
	case f.hasNotation && f.isPrefix && f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByPrefixMibVendorPaged(ctx, notation, mib, vendor, page, size)
	case f.hasNotation && !f.isPrefix && f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByExactMibVendorPaged(ctx, notation, mib, vendor, page, size)
	case f.hasNotation && f.isPrefix && f.hasMib && !f.hasVendor:
		arr, err = dao.GetOidsByPrefixMibPaged(ctx, notation, mib, page, size)
	case f.hasNotation && !f.isPrefix && f.hasMib && !f.hasVendor:
		arr, err = dao.GetOidsByExactMibPaged(ctx, notation, mib, page, size)
	case f.hasNotation && f.isPrefix && !f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByPrefixVendorPaged(ctx, notation, vendor, page, size)
	case f.hasNotation && !f.isPrefix && !f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByExactVendorPaged(ctx, notation, vendor, page, size)
	case f.hasNotation && f.isPrefix && !f.hasMib && !f.hasVendor:
		arr, err = dao.GetOidsByPrefixOnlyPaged(ctx, notation, page, size)
	case f.hasNotation && !f.isPrefix && !f.hasMib && !f.hasVendor:
		arr, err = dao.GetOidsByExactOnlyPaged(ctx, notation, page, size)
	case !f.hasNotation && f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByMibVendorPaged(ctx, mib, vendor, page, size)
	case !f.hasNotation && f.hasMib && !f.hasVendor:
		arr, err = dao.GetOidsByMibOnlyPaged(ctx, mib, page, size)
	case !f.hasNotation && !f.hasMib && f.hasVendor:
		arr, err = dao.GetOidsByVendorOnlyPaged(ctx, vendor, page, size)
	default:
		arr, err = dao.GetAllOidsPaged(ctx, page, size)
	}
	return arr, err
}
