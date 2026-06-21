package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
	"encoding/json"
)

func MappingCreateDtoToMappingDao(d dto.MappingCreateDto) dao.MappingDao {
	var res dao.MappingDao
	var parsedFreq model.PollingFrequency
	var parsedType model.OidType
	parsedFreq = model.ParsePollingFrequency(d.Frequency)
	res.Indicator = d.IndicatorID
	res.Param = d.ParamID
	res.Frequency = int16(parsedFreq)
	res.Value = util.StringPtrToSqlNullString(d.Value)
	res.Coefficient = util.Float64PtrToSqlNullFloat64(d.Coefficient)
	if len(d.Enum) > 0 {
		res.Enum, _ = json.Marshal(d.Enum)
	}
	res.Position = util.Int16PtrToSqlNullInt16(d.Position)
	if d.From != nil && *d.From > 0 {
		res.From = util.Int64PtrToSqlNullInt64(d.From)
	} else {
		res.From.Valid = false
		res.From.Int64 = 0
	}
	if d.PositionType != nil && *d.PositionType != "" {
		parsedType = model.ParseOidType(*d.PositionType)
		res.PositionType = util.Int16PtrToSqlNullInt16((*int16)(&parsedType))
	}
	return res
}

func MappingToMappingDto(d dao.Mapping) dto.MappingDto {
	var res dto.MappingDto
	var indDto dto.ParamIndicatorDto
	var oDto dto.OidDto
	var mDto dto.MibDto
	var pDto dto.ParamDto
	var vendorID int64
	var vendorName string
	var enumMap map[string]string
	var oidEnumMap map[string]string
	var freqStr string
	var posTypeStr string
	var typeStr string
	var statusStr string
	var accessStr string
	var pTypeStr string
	var pAccessStr string
	res.ID = d.ID
	freqStr = model.PollingFrequency(d.Frequency).String()
	res.Frequency = freqStr
	res.Value = util.SqlNullStringToStringPtr(d.Value)
	res.Coefficient = util.SqlNullFloat64ToFloat64Ptr(d.Coefficient)
	if len(d.Enum) > 0 {
		enumMap = make(map[string]string)
		_ = json.Unmarshal(d.Enum, &enumMap)
		res.Enum = enumMap
	}
	res.Position = util.SqlNullInt16ToInt16Ptr(d.Position)
	res.From = util.SqlNullInt64ToInt64Ptr(d.From)
	if d.PositionType.Valid {
		posTypeStr = model.OidType(d.PositionType.Int16).String()
		res.PositionType = &posTypeStr
	}
	indDto.ID = d.IndicatorID
	indDto.DotterNotation = util.SqlNullStringToStringPtr(d.IndDotterNotation)
	if d.IndOidID != nil {
		oDto.ID = d.IndOidID.String()
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
			oidEnumMap = make(map[string]string)
			_ = json.Unmarshal(d.OidEnum, &oidEnumMap)
			oDto.Enum = oidEnumMap
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
		indDto.Oid = &oDto
	}
	res.Indicator = &indDto
	pDto.ID = d.ParamID
	pDto.Title = d.ParamTitle
	pDto.NameEn = d.ParamNameEn
	pDto.NameRu = d.ParamNameRu
	pTypeStr = model.VarType(d.ParamType).String()
	pDto.Type = pTypeStr
	pDto.Value = util.SqlNullStringToStringPtr(d.ParamValue)
	pDto.DescriptionEn = util.SqlNullStringToStringPtr(d.ParamDescriptionEn)
	pDto.DescriptionRu = util.SqlNullStringToStringPtr(d.ParamDescriptionRu)
	pDto.UnitsEn = util.SqlNullStringToStringPtr(d.ParamUnitsEn)
	pDto.UnitsRu = util.SqlNullStringToStringPtr(d.ParamUnitsRu)
	pAccessStr = model.Access(d.ParamAccess).String()
	pDto.Access = pAccessStr
	pDto.Saved = d.ParamSaved
	pDto.Visible = d.ParamVisible
	res.Param = &pDto
	res.Children = []dto.MappingDto{}
	return res
}
