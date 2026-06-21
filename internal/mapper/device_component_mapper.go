package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/util"
)

func DeviceComponentCreateDtoToDeviceComponentDao(d dto.DeviceComponentCreateDto) dao.DeviceComponentDao {
	var res dao.DeviceComponentDao
	res.Model = d.ModelID
	res.InternalOrder = util.Int32PtrToSqlNullInt32(d.InternalOrder)
	res.Parent = util.Int64PtrToSqlNullInt64(d.Parent)
	return res
}

func DeviceComponentToDeviceComponentDto(d dao.DeviceComponent) dto.DeviceComponentDto {
	var res dto.DeviceComponentDto
	var info dto.ComponentInfoDto
	res.ID = d.ID
	info.ID = d.ModelID
	info.Title = d.CompTitle
	info.NameEn = d.CompNameEn
	info.NameRu = d.CompNameRu
	res.Component = info
	res.InternalOrder = util.SqlNullInt32ToInt32Ptr(d.InternalOrder)
	res.Parent = util.SqlNullInt64ToInt64Ptr(d.Parent)
	res.Mappings = []dto.MappingDto{}
	res.Children = []dto.DeviceComponentDto{}
	return res
}

func DeviceComponentArrayToDeviceComponentDtoArray(arr []dao.DeviceComponent) []dto.DeviceComponentDto {
	var res []dto.DeviceComponentDto
	var item dao.DeviceComponent
	res = []dto.DeviceComponentDto{}
	for _, item = range arr {
		res = append(res, DeviceComponentToDeviceComponentDto(item))
	}
	return res
}
