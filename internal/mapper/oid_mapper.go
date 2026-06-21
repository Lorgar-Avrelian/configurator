package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
	"encoding/json"
)

func OidToOidDto(d dao.Oid) dto.OidDto {
	var res dto.OidDto
	var mibDto dto.MibDto
	var vendorID int64
	var vendorName string
	var typeStr string
	var statusStr string
	var accessStr string
	var enumMap map[string]string
	res.ID = d.ID.String()
	if d.MibID.Valid {
		mibDto.ID = d.MibID.Int64
		if d.MibPath.Valid {
			mibDto.Path = d.MibPath.String
		}
		mibDto.Name = util.SqlNullStringToStringPtr(d.MibName)
		if d.MibVendor.Valid {
			vendorID = d.MibVendor.Int64
			vendorName = model.Vendor(vendorID).Data().Name
			mibDto.Vendor = &vendorName
		}
		res.Mib = &mibDto
	}
	typeStr = model.Asn1Type(d.Type).String()
	res.Type = typeStr
	res.Name = d.Name
	res.Number = util.SqlNullInt32ToInt32Ptr(d.Number)
	res.DotterNotation = d.DotterNotation
	res.ObjectDescriptor = d.ObjectDescriptor
	res.Syntax = util.SqlNullStringToStringPtr(d.Syntax)
	if len(d.Enum) > 0 {
		enumMap = make(map[string]string)
		_ = json.Unmarshal(d.Enum, &enumMap)
		res.Enum = enumMap
	}
	if d.Status.Valid {
		statusStr = model.OidStatus(d.Status.Int16).String()
		res.Status = &statusStr
	}
	if d.Access.Valid {
		accessStr = model.OidAccess(d.Access.Int16).String()
		res.Access = &accessStr
	}
	res.Units = util.SqlNullStringToStringPtr(d.Units)
	res.Description = util.SqlNullStringToStringPtr(d.Description)
	res.Category = util.SqlNullStringToStringPtr(d.Category)
	return res
}

func OidArrayToOidDtoArray(arr []dao.Oid) []dto.OidDto {
	var res []dto.OidDto
	var item dao.Oid
	res = []dto.OidDto{}
	for _, item = range arr {
		res = append(res, OidToOidDto(item))
	}
	return res
}
