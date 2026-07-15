package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
)

func ParamCreateDtoToParamDao(d dto.ParamCreateDto) dao.ParamDao {
	var res dao.ParamDao
	var parsedAccess model.Access
	parsedAccess = model.ParseAccess(d.Access)
	var parsedType model.VarType
	parsedType = model.ParseVarType(d.Type)
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.Type = int16(parsedType)
	res.Value = util.StringToSqlNullString(d.Value)
	res.DescriptionEn = util.StringToSqlNullString(d.DescriptionEn)
	res.DescriptionRu = util.StringToSqlNullString(d.DescriptionRu)
	res.UnitsEn = util.StringToSqlNullString(d.UnitsEn)
	res.UnitsRu = util.StringToSqlNullString(d.UnitsRu)
	res.Access = int16(parsedAccess)
	res.Saved = d.Saved
	res.Visible = d.Visible
	res.Diagram = d.Diagram
	return res
}

func ParamDaoToParamDto(d dao.ParamDao) dto.ParamDto {
	var res dto.ParamDto
	var modelAccess model.Access
	modelAccess = model.Access(d.Access)
	var modelType model.VarType
	modelType = model.VarType(d.Type)
	res.ID = d.ID
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.Type = modelType.String()
	res.Value = util.SqlNullStringToStringPtr(d.Value)
	res.DescriptionEn = util.SqlNullStringToStringPtr(d.DescriptionEn)
	res.DescriptionRu = util.SqlNullStringToStringPtr(d.DescriptionRu)
	res.UnitsEn = util.SqlNullStringToStringPtr(d.UnitsEn)
	res.UnitsRu = util.SqlNullStringToStringPtr(d.UnitsRu)
	res.Access = modelAccess.String()
	res.Saved = d.Saved
	res.Visible = d.Visible
	res.Diagram = d.Diagram
	return res
}

func ParamToParamDto(m dao.Param) dto.ParamDto {
	var res dto.ParamDto
	var modelAccess model.Access
	modelAccess = model.Access(m.Access)
	var modelType model.VarType
	modelType = model.VarType(m.Type)
	res.ID = m.ID
	res.Title = m.Title
	res.NameEn = m.NameEn
	res.NameRu = m.NameRu
	res.Type = modelType.String()
	res.Value = util.SqlNullStringToStringPtr(m.Value)
	res.DescriptionEn = util.SqlNullStringToStringPtr(m.DescriptionEn)
	res.DescriptionRu = util.SqlNullStringToStringPtr(m.DescriptionRu)
	res.UnitsEn = util.SqlNullStringToStringPtr(m.UnitsEn)
	res.UnitsRu = util.SqlNullStringToStringPtr(m.UnitsRu)
	res.Access = modelAccess.String()
	res.Saved = m.Saved
	res.Visible = m.Visible
	res.Diagram = m.Diagram
	return res
}

func ParamUpdateDtoToParamDao(d dto.ParamUpdateDto) dao.ParamDao {
	var res dao.ParamDao
	var parsedAccess model.Access
	parsedAccess = model.ParseAccess(d.Access)
	var parsedType model.VarType
	parsedType = model.ParseVarType(d.Type)
	res.ID = d.ID
	res.Title = d.Title
	res.NameEn = d.NameEn
	res.NameRu = d.NameRu
	res.Type = int16(parsedType)
	res.Value = util.StringPtrToSqlNullString(d.Value)
	res.DescriptionEn = util.StringPtrToSqlNullString(d.DescriptionEn)
	res.DescriptionRu = util.StringPtrToSqlNullString(d.DescriptionRu)
	res.UnitsEn = util.StringPtrToSqlNullString(d.UnitsEn)
	res.UnitsRu = util.StringPtrToSqlNullString(d.UnitsRu)
	res.Access = int16(parsedAccess)
	res.Saved = d.Saved
	res.Visible = d.Visible
	res.Diagram = d.Diagram
	return res
}

func ParamArrayToParamDtoArray(arr []dao.Param) []dto.ParamDto {
	var res []dto.ParamDto
	res = []dto.ParamDto{}
	var p dao.Param
	for _, p = range arr {
		res = append(res, ParamToParamDto(p))
	}
	return res
}
