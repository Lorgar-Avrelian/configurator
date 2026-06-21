package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
	"encoding/json"

	"github.com/google/uuid"
)

func ParamIndicatorCreateDtoToParamIndicatorDao(d dto.ParamIndicatorCreateDto) dao.ParamIndicatorDao {
	var res dao.ParamIndicatorDao
	var parsedUUID uuid.UUID
	if d.OidID != nil && *d.OidID != "" {
		parsedUUID, _ = uuid.Parse(*d.OidID)
		res.OidID = parsedUUID
	} else {
		res.OidID = uuid.Nil
	}
	res.DotterNotation = util.StringPtrToSqlNullString(d.DotterNotation)
	return res
}

func ParamIndicatorToParamIndicatorDto(d dao.ParamIndicator) dto.ParamIndicatorDto {
	var res dto.ParamIndicatorDto
	var oDto dto.OidDto
	var mDto dto.MibDto
	var vendorID int64
	var vendorName string
	var enumMap map[string]string
	var typeStr string
	var statusStr string
	var accessStr string
	res.ID = d.ID
	res.DotterNotation = util.SqlNullStringToStringPtr(d.DotterNotation)
	if d.OidID != nil {
		oDto.ID = d.OidID.String()
		if d.OidMibID.Valid {
			mDto.ID = d.OidMibID.Int64
			if d.OidMibPath.Valid {
				mDto.Path = d.OidMibPath.String
			}
			mDto.Name = util.SqlNullStringToStringPtr(d.OidMibName)
			if d.OidMibVendor.Valid {
				vendorID = d.OidMibVendor.Int64
				vendorName = model.Vendor(vendorID).Data().Name
				mDto.Vendor = &vendorName
			}
			oDto.Mib = &mDto
		}
		if d.OidType.Valid {
			typeStr = model.Asn1Type(d.OidType.Int16).String()
			oDto.Type = typeStr
		}
		if d.OidName.Valid {
			oDto.Name = d.OidName.String
		}
		oDto.Number = util.SqlNullInt32ToInt32Ptr(d.OidNumber)
		if d.OidDotterNotation.Valid {
			oDto.DotterNotation = d.OidDotterNotation.String
		}
		if d.OidObjectDescriptor.Valid {
			oDto.ObjectDescriptor = d.OidObjectDescriptor.String
		}
		oDto.Syntax = util.SqlNullStringToStringPtr(d.OidSyntax)
		if len(d.OidEnum) > 0 {
			enumMap = make(map[string]string)
			_ = json.Unmarshal(d.OidEnum, &enumMap)
			oDto.Enum = enumMap
		}
		if d.OidStatus.Valid {
			statusStr = model.OidStatus(d.OidStatus.Int16).String()
			oDto.Status = &statusStr
		}
		if d.OidAccess.Valid {
			accessStr = model.OidAccess(d.OidAccess.Int16).String()
			oDto.Access = &accessStr
		}
		oDto.Units = util.SqlNullStringToStringPtr(d.OidUnits)
		oDto.Description = util.SqlNullStringToStringPtr(d.OidDescription)
		oDto.Category = util.SqlNullStringToStringPtr(d.OidCategory)
		res.Oid = &oDto
	}
	return res
}

func ParamIndicatorArrayToParamIndicatorDtoArray(arr []dao.ParamIndicator) []dto.ParamIndicatorDto {
	var res []dto.ParamIndicatorDto
	var item dao.ParamIndicator
	res = []dto.ParamIndicatorDto{}
	for _, item = range arr {
		res = append(res, ParamIndicatorToParamIndicatorDto(item))
	}
	return res
}
