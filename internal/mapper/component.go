package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
)

func ComponentCreateDtoToComponentDao(d dto.ComponentCreateDto) dao.ComponentDao {
	var res dao.ComponentDao
	var parsedAccess model.Access
	parsedAccess = model.ParseAccess(d.Access)
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.PluralNameEn = d.PluralNameEn
	res.PluralNameRu = d.PluralNameRu
	res.BaseComponent = util.Int64PtrToSqlNullInt64(d.BaseComponent)
	res.DescriptionEn = util.StringToSqlNullString(d.DescriptionEn)
	res.DescriptionRu = util.StringToSqlNullString(d.DescriptionRu)
	res.Access = int16(parsedAccess)
	return res
}

func ComponentDaoToComponentDto(d dao.ComponentDao) dto.ComponentDto {
	var res dto.ComponentDto
	var modelAccess model.Access
	modelAccess = model.Access(d.Access)
	res.ID = d.ID
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.PluralNameEn = d.PluralNameEn
	res.PluralNameRu = d.PluralNameRu
	res.BaseComponent = util.SqlNullInt64ToInt64Ptr(d.BaseComponent)
	res.DescriptionEn = util.SqlNullStringToStringPtr(d.DescriptionEn)
	res.DescriptionRu = util.SqlNullStringToStringPtr(d.DescriptionRu)
	res.Access = modelAccess.String()
	res.Params = []dto.ParamDto{}
	return res
}

func ComponentDaoArrayToComponentDtoArray(arr []dao.ComponentDao) []dto.ComponentDto {
	var res []dto.ComponentDto
	res = []dto.ComponentDto{}
	var c dao.ComponentDao
	for _, c = range arr {
		res = append(res, ComponentDaoToComponentDto(c))
	}
	return res
}

func ComponentToComponentDto(m dao.Component) dto.ComponentDto {
	var res dto.ComponentDto
	var modelAccess model.Access
	modelAccess = model.Access(m.Access)
	res.ID = m.ID
	res.Title = m.Title
	res.NameEn = m.NameEn
	res.NameRu = m.NameRu
	res.PluralNameEn = m.PluralNameEn
	res.PluralNameRu = m.PluralNameRu
	res.BaseComponent = util.SqlNullInt64ToInt64Ptr(m.BaseComponent)
	res.DescriptionEn = util.SqlNullStringToStringPtr(m.DescriptionEn)
	res.DescriptionRu = util.SqlNullStringToStringPtr(m.DescriptionRu)
	res.Access = modelAccess.String()
	res.Params = ParamArrayToParamDtoArray(m.Params)
	return res
}

func ComponentUpdateDtoToComponentDao(d dto.ComponentUpdateDto) dao.ComponentDao {
	var res dao.ComponentDao
	var parsedAccess model.Access
	parsedAccess = model.ParseAccess(d.Access)
	res.ID = d.ID
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.PluralNameEn = d.PluralNameEn
	res.PluralNameRu = d.PluralNameRu
	res.BaseComponent = util.Int64PtrToSqlNullInt64(d.BaseComponent)
	res.DescriptionEn = util.StringPtrToSqlNullString(d.DescriptionEn)
	res.DescriptionRu = util.StringPtrToSqlNullString(d.DescriptionRu)
	res.Access = int16(parsedAccess)
	return res
}

func ComponentArrayToComponentDtoArray(arr []dao.Component) []dto.ComponentDto {
	var res []dto.ComponentDto
	res = []dto.ComponentDto{}
	var c dao.Component
	for _, c = range arr {
		res = append(res, ComponentToComponentDto(c))
	}
	return res
}
