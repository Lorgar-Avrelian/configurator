package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/util"
)

func ResultToParamResultDto(m dao.Result) dto.ParamResultDto {
	var res dto.ParamResultDto
	res.Host = m.Host
	res.Port = m.Port
	res.ComponentTitle = m.ComponentTitle
	res.InternalOrder = util.SqlNullInt32ToInt32Ptr(m.InternalOrder)
	res.ParamTitle = m.ParamTitle
	res.Value = util.SqlNullStringToStringPtr(m.Value)
	return res
}

func ResultArrayToParamResultDtoArray(arr []dao.Result) []dto.ParamResultDto {
	var res []dto.ParamResultDto
	res = []dto.ParamResultDto{}
	var item dao.Result
	for _, item = range arr {
		res = append(res, ResultToParamResultDto(item))
	}
	return res
}
