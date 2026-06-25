package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
)

func GetWorkingConfiguration(ctx context.Context, host string, port int32) (*dto.DeviceSnmpDto, error) {
	var snmpDao *dao.DeviceSnmpDao
	var err error
	snmpDao, err = dao.GetDeviceSnmpByHostPort(ctx, host, port)
	if err != nil {
		return nil, err
	}
	if snmpDao == nil {
		return nil, nil
	}
	var thresholds []dao.ThresholdDao
	thresholds, err = dao.GetThresholdsByHostPort(ctx, host, port)
	if err != nil {
		return nil, err
	}
	var components []dao.Component
	components, err = dao.GetComponentsWithTargetParamByHostPort(ctx, host, port)
	if err != nil {
		return nil, err
	}
	var configDto *dto.ConfigurationDto
	configDto = nil
	if snmpDao.Config.Valid {
		var cfgID int64
		cfgID = snmpDao.Config.Int64
		configDto, err = GetConfigurationByID(ctx, cfgID)
		if err != nil {
			return nil, err
		}
	}
	var aggregated dao.DeviceSnmp
	aggregated.Device = *snmpDao
	aggregated.Thresholds = thresholds
	aggregated.Components = components
	var finalDto dto.DeviceSnmpDto
	finalDto = mapper.DeviceSnmpToDeviceSnmpDto(aggregated, configDto)
	return &finalDto, nil
}
