package service

import (
	"configurator/internal/dao"
	"configurator/internal/logger"
	"context"
	"sync"
)

func ChangeComponentData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var tComponent []dao.ComponentDao
	var tCompParam []dao.ComponentParamDao
	var tDevInd []dao.DeviceIndicatorDao
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, err = GetComponentDependentData(ctx)
	if err != nil {
		return false, err
	}
	err = DropComponentDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	err = DropComponentDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = CreateComponentDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = InsertComponentDependentData(ctx, tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		return false, err
	}
	err = CreateComponentDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetComponentDependentData(ctx context.Context) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceIndicatorDao, []dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tComponent []dao.ComponentDao
	var tCompParam []dao.ComponentParamDao
	var tDevInd []dao.DeviceIndicatorDao
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	wg.Add(15)
	go func() {
		defer wg.Done()
		var res []dao.ComponentDao
		res, err = dao.GetAllComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tComponent = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ComponentParamDao
		res, err = dao.GetAllComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tCompParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceIndicatorDao
		res, err = dao.GetAllDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevInd = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ParamIndicatorDao
		res, err = dao.GetAllParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tParamInd = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.MappingDao
		res, err = dao.GetAllMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tMapping = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentDao
		res, err = dao.GetAllDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevComp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentMappingDao
		res, err = dao.GetAllDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigurationDao
		res, err = dao.GetAllConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DefaultConfigurationDao
		res, err = dao.GetAllDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDefConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ThresholdDao
		res, err = dao.GetAllThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tThreshold = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceSnmpDao
		res, err = dao.GetAllDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevSnmp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tResult = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedTh = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedParamDao
		res, err = dao.GetAllAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigInProcessDao
		res, err = dao.GetAllConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfigProc = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errResult
	}
	return tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, nil
}

func DropComponentDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.DropComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropComponentParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropParamIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func DropComponentDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.DropComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateComponentDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.CreateComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func InsertComponentDependentData(ctx context.Context, tComponent []dao.ComponentDao, tCompParam []dao.ComponentParamDao, tDevInd []dao.DeviceIndicatorDao, tParamInd []dao.ParamIndicatorDao, tMapping []dao.MappingDao, tDevComp []dao.DeviceComponentDao, tDevCompMap []dao.DeviceComponentMappingDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tThreshold []dao.ThresholdDao, tDevSnmp []dao.DeviceSnmpDao, tResult []dao.ResultDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao, tConfigProc []dao.ConfigInProcessDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.ImportComponentDao(ctx, tComponent)
		if err != nil {
			logger.Errorf("Error importing data into table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportComponentParamDao(ctx, tCompParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceIndicatorDao(ctx, tDevInd)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportParamIndicatorDao(ctx, tParamInd)
		if err != nil {
			logger.Errorf("Error importing data into table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			logger.Errorf("Error importing data into table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentDao(ctx, tDevComp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigurationDao(ctx, tConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDefaultConfigurationDao(ctx, tDefConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportThresholdDao(ctx, tThreshold)
		if err != nil {
			logger.Errorf("Error importing data into table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportResultDao(ctx, tResult)
		if err != nil {
			logger.Errorf("Error importing data into table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigInProcessDao(ctx, tConfigProc)
		if err != nil {
			logger.Errorf("Error importing data into table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateComponentDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.CreateComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateComponentParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateParamIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func ChangeParamData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var tParam []dao.ParamDao
	var tCompParam []dao.ComponentParamDao
	var tDevInd []dao.DeviceIndicatorDao
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, err = GetParamDependentData(ctx)
	if err != nil {
		return false, err
	}
	err = DropParamDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	err = DropParamDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = CreateParamDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = InsertParamDependentData(ctx, tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		return false, err
	}
	err = CreateParamDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetParamDependentData(ctx context.Context) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.DeviceIndicatorDao, []dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tParam []dao.ParamDao
	var tCompParam []dao.ComponentParamDao
	var tDevInd []dao.DeviceIndicatorDao
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	wg.Add(15)
	go func() {
		defer wg.Done()
		var res []dao.ParamDao
		res, err = dao.GetAllParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ComponentParamDao
		res, err = dao.GetAllComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tCompParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceIndicatorDao
		res, err = dao.GetAllDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevInd = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ParamIndicatorDao
		res, err = dao.GetAllParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tParamInd = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.MappingDao
		res, err = dao.GetAllMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tMapping = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentDao
		res, err = dao.GetAllDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevComp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentMappingDao
		res, err = dao.GetAllDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigurationDao
		res, err = dao.GetAllConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DefaultConfigurationDao
		res, err = dao.GetAllDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDefConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ThresholdDao
		res, err = dao.GetAllThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tThreshold = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceSnmpDao
		res, err = dao.GetAllDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevSnmp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tResult = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedTh = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedParamDao
		res, err = dao.GetAllAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigInProcessDao
		res, err = dao.GetAllConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfigProc = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errResult
	}
	return tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, nil
}

func DropParamDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.DropParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropComponentParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropParamIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func DropParamDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.DropParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateParamDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.CreateParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateComponentParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateParamIndicatorDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func InsertParamDependentData(ctx context.Context, tParam []dao.ParamDao, tCompParam []dao.ComponentParamDao, tDevInd []dao.DeviceIndicatorDao, tParamInd []dao.ParamIndicatorDao, tMapping []dao.MappingDao, tDevComp []dao.DeviceComponentDao, tDevCompMap []dao.DeviceComponentMappingDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tThreshold []dao.ThresholdDao, tDevSnmp []dao.DeviceSnmpDao, tResult []dao.ResultDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao, tConfigProc []dao.ConfigInProcessDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.ImportParamDao(ctx, tParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportComponentParamDao(ctx, tCompParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceIndicatorDao(ctx, tDevInd)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportParamIndicatorDao(ctx, tParamInd)
		if err != nil {
			logger.Errorf("Error importing data into table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			logger.Errorf("Error importing data into table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentDao(ctx, tDevComp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigurationDao(ctx, tConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDefaultConfigurationDao(ctx, tDefConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportThresholdDao(ctx, tThreshold)
		if err != nil {
			logger.Errorf("Error importing data into table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportResultDao(ctx, tResult)
		if err != nil {
			logger.Errorf("Error importing data into table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigInProcessDao(ctx, tConfigProc)
		if err != nil {
			logger.Errorf("Error importing data into table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateParamDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(15)
	go func() {
		defer wg.Done()
		err = dao.CreateParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateComponentParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.component_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateParamIndicatorDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.param_indicator: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func ChangeDeviceComponentData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, err = GetDeviceComponentDependentData(ctx)
	if err != nil {
		return false, err
	}
	err = DropDeviceComponentDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	err = DropDeviceComponentDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = CreateDeviceComponentDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = InsertDeviceComponentDependentData(ctx, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		return false, err
	}
	err = CreateDeviceComponentDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetDeviceComponentDependentData(ctx context.Context) ([]dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	wg.Add(10)
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentDao
		res, err = dao.GetAllDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevComp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentMappingDao
		res, err = dao.GetAllDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigurationDao
		res, err = dao.GetAllConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DefaultConfigurationDao
		res, err = dao.GetAllDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDefConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ThresholdDao
		res, err = dao.GetAllThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tThreshold = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceSnmpDao
		res, err = dao.GetAllDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevSnmp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tResult = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedTh = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedParamDao
		res, err = dao.GetAllAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigInProcessDao
		res, err = dao.GetAllConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfigProc = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errResult
	}
	return tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, nil
}

func DropDeviceComponentDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func DropDeviceComponentDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateDeviceComponentDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func InsertDeviceComponentDependentData(ctx context.Context, tDevComp []dao.DeviceComponentDao, tDevCompMap []dao.DeviceComponentMappingDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tThreshold []dao.ThresholdDao, tDevSnmp []dao.DeviceSnmpDao, tResult []dao.ResultDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao, tConfigProc []dao.ConfigInProcessDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentDao(ctx, tDevComp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigurationDao(ctx, tConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDefaultConfigurationDao(ctx, tDefConfig)
		if err != nil {
			logger.Errorf("Error importing data into table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportThresholdDao(ctx, tThreshold)
		if err != nil {
			logger.Errorf("Error importing data into table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportResultDao(ctx, tResult)
		if err != nil {
			logger.Errorf("Error importing data into table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigInProcessDao(ctx, tConfigProc)
		if err != nil {
			logger.Errorf("Error importing data into table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateDeviceComponentDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func ChangeMappingData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	tMapping, tDevCompMap, err = GetMappingDependentData(ctx)
	if err != nil {
		return false, err
	}
	err = DropMappingDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	err = DropMappingDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = CreateMappingDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = InsertMappingDependentData(ctx, tMapping, tDevCompMap)
	if err != nil {
		return false, err
	}
	err = CreateMappingDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetMappingDependentData(ctx context.Context) ([]dao.MappingDao, []dao.DeviceComponentMappingDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	wg.Add(2)
	go func() {
		defer wg.Done()
		var res []dao.MappingDao
		res, err = dao.GetAllMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tMapping = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentMappingDao
		res, err = dao.GetAllDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, errResult
	}
	return tMapping, tDevCompMap, nil
}

func DropMappingDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.DropMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func DropMappingDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.DropMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateMappingDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func InsertMappingDependentData(ctx context.Context, tMapping []dao.MappingDao, tDevCompMap []dao.DeviceComponentMappingDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			logger.Errorf("Error importing data into table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateMappingDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.CreateMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func ChangeThresholdData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	var err error
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, err = GetThresholdDependentData(ctx)
	if err != nil {
		return false, err
	}
	err = DropThresholdDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	err = DropThresholdDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = CreateThresholdDependentTables(ctx)
	if err != nil {
		return false, err
	}
	err = InsertThresholdDependentData(ctx, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		return false, err
	}
	err = CreateThresholdDependentConstraints(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetThresholdDependentData(ctx context.Context) ([]dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tThreshold []dao.ThresholdDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	var tConfigProc []dao.ConfigInProcessDao
	wg.Add(6)
	go func() {
		defer wg.Done()
		var res []dao.ThresholdDao
		res, err = dao.GetAllThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tThreshold = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceSnmpDao
		res, err = dao.GetAllDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevSnmp = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tResult = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedTh = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedParamDao
		res, err = dao.GetAllAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigInProcessDao
		res, err = dao.GetAllConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error getting data from table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tConfigProc = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, errResult
	}
	return tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc, nil
}

func DropThresholdDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func DropThresholdDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropResultDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error dropping table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateThresholdDependentTables(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDao(ctx)
		if err != nil {
			logger.Errorf("Error creating table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func InsertThresholdDependentData(ctx context.Context, tThreshold []dao.ThresholdDao, tDevSnmp []dao.DeviceSnmpDao, tResult []dao.ResultDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao, tConfigProc []dao.ConfigInProcessDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
	go func() {
		defer wg.Done()
		err = dao.ImportThresholdDao(ctx, tThreshold)
		if err != nil {
			logger.Errorf("Error importing data into table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			logger.Errorf("Error importing data into table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportResultDao(ctx, tResult)
		if err != nil {
			logger.Errorf("Error importing data into table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			logger.Errorf("Error importing data into table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigInProcessDao(ctx, tConfigProc)
		if err != nil {
			logger.Errorf("Error importing data into table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}

func CreateThresholdDependentConstraints(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
	go func() {
		defer wg.Done()
		err = dao.CreateThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateDeviceSnmpDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateResultDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.result: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedThresholdDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateAffectedParamDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.CreateConfigInProcessDaoConstraints(ctx)
		if err != nil {
			logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	if errResult != nil {
		return errResult
	}
	return nil
}
