package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/util"
)

func DeviceIndicatorCreateDtoToDeviceIndicatorDao(d dto.DeviceIndicatorCreateDto) dao.DeviceIndicatorDao {
	var res dao.DeviceIndicatorDao
	res.Description = util.StringPtrToSqlNullString(d.Description)
	res.ObjectID = d.ObjectID
	res.Contact = util.StringPtrToSqlNullString(d.Contact)
	res.Name = util.StringPtrToSqlNullString(d.Name)
	res.Location = util.StringPtrToSqlNullString(d.Location)
	res.Services = util.Int16PtrToSqlNullInt16(d.Services)
	return res
}

func DeviceIndicatorDaoToDeviceIndicatorDto(d dao.DeviceIndicatorDao) dto.DeviceIndicatorDto {
	var res dto.DeviceIndicatorDto
	res.ID = d.ID
	res.Description = util.SqlNullStringToStringPtr(d.Description)
	res.ObjectID = d.ObjectID
	res.Contact = util.SqlNullStringToStringPtr(d.Contact)
	res.Name = util.SqlNullStringToStringPtr(d.Name)
	res.Location = util.SqlNullStringToStringPtr(d.Location)
	res.Services = util.SqlNullInt16ToInt16Ptr(d.Services)
	return res
}

func DeviceIndicatorDaoArrayToDeviceIndicatorDtoArray(arr []dao.DeviceIndicatorDao) []dto.DeviceIndicatorDto {
	var res []dto.DeviceIndicatorDto
	var item dao.DeviceIndicatorDao
	res = []dto.DeviceIndicatorDto{}
	for _, item = range arr {
		res = append(res, DeviceIndicatorDaoToDeviceIndicatorDto(item))
	}
	return res
}
