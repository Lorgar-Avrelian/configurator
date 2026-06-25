package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
)

func DeviceSnmpToDeviceSnmpDto(m dao.DeviceSnmp, configDto *dto.ConfigurationDto) dto.DeviceSnmpDto {
	var res dto.DeviceSnmpDto
	var vStr string
	var aStr string
	var pStr string
	var thresholdList []dto.ThresholdDto
	var componentList []dto.ComponentDto
	var t dao.ThresholdDao
	var c dao.Component
	vStr = model.VersionSnmp(m.Device.Version).String()
	aStr = model.AuthProtocolSnmp(m.Device.Authentication).String()
	pStr = model.PrivacyProtocolSnmp(m.Device.Privacy).String()
	res.Host = m.Device.Host
	res.Port = m.Device.Port
	res.Community = m.Device.Community
	res.Version = vStr
	res.Login = util.SqlNullStringToStringPtr(m.Device.Login)
	res.Password = util.SqlNullStringToStringPtr(m.Device.Password)
	res.Authentication = aStr
	res.Privacy = pStr
	res.ID = m.Device.ID.String()
	res.Configuration = configDto
	thresholdList = []dto.ThresholdDto{}
	for _, t = range m.Thresholds {
		thresholdList = append(thresholdList, ThresholdDaoToThresholdDto(t))
	}
	res.Thresholds = thresholdList
	componentList = []dto.ComponentDto{}
	for _, c = range m.Components {
		componentList = append(componentList, ComponentToComponentDto(c))
	}
	res.Components = componentList
	return res
}
