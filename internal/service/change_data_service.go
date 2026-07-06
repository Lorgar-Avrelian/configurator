package service

import (
	"configurator/internal/dao"
	"configurator/internal/logger"
	"configurator/internal/util"
	"context"
	"errors"
	"fmt"
	"sync"
)

func ChangeComponentData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	if prevId == newId {
		return true, nil
	}
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
	tComponent, tCompParam, tDevComp, tResult, tAffectedParam, err = changeComponentId(prevId, newId, tComponent, tCompParam, tDevComp, tResult, tAffectedParam)
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

func changeComponentId(prevId int64, newId int64, components []dao.ComponentDao, componentParams []dao.ComponentParamDao, deviceComponents []dao.DeviceComponentDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceComponentDao, []dao.ResultDao, []dao.AffectedParamDao, error) {
	var i int
	var err error
	var idMap map[int64]int64
	var item dao.ComponentDao
	var prevComponent dao.ComponentDao
	var foundPrev bool
	var baseIdPtr *int64
	var zero int64
	zero = int64(0)
	if newId <= zero || prevId <= zero {
		return nil, nil, nil, nil, nil, errors.New("ID cannot be less than or equal to zero")
	}
	if newId > int64(len(components)) {
		return nil, nil, nil, nil, nil, errors.New("ID cannot be greater than Components count")
	}
	for i = range components {
		item = components[i]
		if item.ID == prevId {
			prevComponent = item
			foundPrev = true
			break
		}
	}
	if !foundPrev {
		return nil, nil, nil, nil, nil, fmt.Errorf("Component with ID = %d is not exists", prevId)
	}
	baseIdPtr = util.SqlNullInt64ToInt64Ptr(prevComponent.BaseComponent)
	if baseIdPtr != nil && *baseIdPtr >= newId {
		return nil, nil, nil, nil, nil, errors.New("ID of the parent model cannot be less than the ID of the child model")
	}
	idMap, err = calculateIdMap(prevId, newId, components)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	updateComponentsIds(components, idMap)
	updateComponentParamsIds(componentParams, idMap)
	updateDeviceComponentsIds(deviceComponents, idMap)
	updateResultsIds(results, idMap)
	updateAffectedParamsIds(affectedParams, idMap)
	return components, componentParams, deviceComponents, results, affectedParams, nil
}

func calculateIdMap(prevId int64, newId int64, components []dao.ComponentDao) (map[int64]int64, error) {
	var idMap map[int64]int64
	var i int
	var length int
	var sortedIds []int64
	var currentId int64
	var firstSortedId int64
	idMap = make(map[int64]int64)
	length = len(components)
	sortedIds = make([]int64, length)
	for i = range components {
		sortedIds[i] = components[i].ID
	}
	sortIdsArray(sortedIds)
	currentId = int64(1)
	if length > 0 {
		firstSortedId = sortedIds[0]
		if currentId > firstSortedId {
			currentId = firstSortedId
		}
	}
	if currentId > newId {
		currentId = newId
	}
	buildDenseIdMap(idMap, sortedIds, prevId, newId, currentId)
	return idMap, nil
}

func sortIdsArray(ids []int64) {
	var j int
	var k int
	var tmp int64
	for j = range ids {
		for k = range ids {
			if ids[k] > ids[j] {
				tmp = ids[j]
				ids[j] = ids[k]
				ids[k] = tmp
			}
		}
	}
}

func buildDenseIdMap(idMap map[int64]int64, sortedIds []int64, prevId int64, newId int64, startId int64) {
	var i int
	var val int64
	var nextId int64
	nextId = startId
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			idMap[val] = newId
		}
	}
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			continue
		}
		if nextId == newId {
			nextId = nextId + 1
		}
		idMap[val] = nextId
		nextId = nextId + 1
	}
}

func updateComponentsIds(components []dao.ComponentDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	var baseIdPtr *int64
	var newBaseId int64
	for i = range components {
		oldId = components[i].ID
		newId, exists = idMap[oldId]
		if exists {
			components[i].ID = newId
		}
		baseIdPtr = util.SqlNullInt64ToInt64Ptr(components[i].BaseComponent)
		if baseIdPtr != nil {
			oldId = *baseIdPtr
			newId, exists = idMap[oldId]
			if exists {
				newBaseId = newId
				components[i].BaseComponent = util.Int64PtrToSqlNullInt64(&newBaseId)
			}
		}
	}
}

func updateComponentParamsIds(componentParams []dao.ComponentParamDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range componentParams {
		oldId = componentParams[i].ComponentID
		newId, exists = idMap[oldId]
		if exists {
			componentParams[i].ComponentID = newId
		}
	}
}

func updateDeviceComponentsIds(deviceComponents []dao.DeviceComponentDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range deviceComponents {
		oldId = deviceComponents[i].Model
		newId, exists = idMap[oldId]
		if exists {
			deviceComponents[i].Model = newId
		}
	}
}

func updateResultsIds(results []dao.ResultDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range results {
		oldId = results[i].Component
		newId, exists = idMap[oldId]
		if exists {
			results[i].Component = newId
		}
	}
}

func updateAffectedParamsIds(affectedParams []dao.AffectedParamDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range affectedParams {
		oldId = affectedParams[i].Component
		newId, exists = idMap[oldId]
		if exists {
			affectedParams[i].Component = newId
		}
	}
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
	var err error
	err = dao.DropConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
		return err
	}
	err = dao.DropMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
		return err
	}
	err = dao.DropParamIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.param_indicator: %v", err)
		return err
	}
	err = dao.DropDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_indicator: %v", err)
		return err
	}
	err = dao.DropComponentParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.component_param: %v", err)
		return err
	}
	err = dao.DropComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.component: %v", err)
		return err
	}
	return nil
}

func DropComponentDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component: %v", err)
		return err
	}
	err = dao.DropMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.mapping: %v", err)
		return err
	}
	err = dao.DropParamIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.param_indicator: %v", err)
		return err
	}
	err = dao.DropDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_indicator: %v", err)
		return err
	}
	err = dao.DropComponentParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.component_param: %v", err)
		return err
	}
	err = dao.DropComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.component: %v", err)
		return err
	}
	return nil
}

func CreateComponentDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.component: %v", err)
		return err
	}
	err = dao.CreateComponentParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.component_param: %v", err)
		return err
	}
	err = dao.CreateDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_indicator: %v", err)
		return err
	}
	err = dao.CreateParamIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.param_indicator: %v", err)
		return err
	}
	err = dao.CreateMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.config_in_process: %v", err)
		return err
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
	var err error
	err = dao.CreateComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.component: %v", err)
		return err
	}
	err = dao.CreateComponentParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.component_param: %v", err)
		return err
	}
	err = dao.CreateDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_indicator: %v", err)
		return err
	}
	err = dao.CreateParamIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.param_indicator: %v", err)
		return err
	}
	err = dao.CreateMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
		return err
	}
	return nil
}

func ChangeParamData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	if prevId == newId {
		return true, nil
	}
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
	tParam, tCompParam, tMapping, tResult, tAffectedParam, err = changeParamId(prevId, newId, tParam, tCompParam, tMapping, tResult, tAffectedParam)
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

func changeParamId(prevId int64, newId int64, params []dao.ParamDao, componentParams []dao.ComponentParamDao, mappings []dao.MappingDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.MappingDao, []dao.ResultDao, []dao.AffectedParamDao, error) {
	var err error
	var idMap map[int64]int64
	var i int
	var foundPrev bool
	var zero int64
	zero = int64(0)
	if newId <= zero || prevId <= zero {
		return nil, nil, nil, nil, nil, errors.New("ID cannot be less than or equal to zero")
	}
	if newId > int64(len(params)) {
		return nil, nil, nil, nil, nil, errors.New("ID cannot be greater than Params count")
	}
	for i = range params {
		if params[i].ID == prevId {
			foundPrev = true
			break
		}
	}
	if !foundPrev {
		return nil, nil, nil, nil, nil, fmt.Errorf("Param with ID = %d is not exists", prevId)
	}
	idMap, err = calculateParamIdMap(prevId, newId, params)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	updateParamsIds(params, idMap)
	updateCompParamsParamIds(componentParams, idMap)
	updateMappingsIds(mappings, idMap)
	updateResultsParamIds(results, idMap)
	updateAffectedParamsParamIds(affectedParams, idMap)
	return params, componentParams, mappings, results, affectedParams, nil
}

func calculateParamIdMap(prevId int64, newId int64, params []dao.ParamDao) (map[int64]int64, error) {
	var idMap map[int64]int64
	var i int
	var length int
	var sortedIds []int64
	var currentId int64
	var firstSortedId int64
	idMap = make(map[int64]int64)
	length = len(params)
	sortedIds = make([]int64, length)
	for i = range params {
		sortedIds[i] = params[i].ID
	}
	sortIdsArray(sortedIds)
	currentId = int64(1)
	if length > 0 {
		firstSortedId = sortedIds[0]
		if currentId > firstSortedId {
			currentId = firstSortedId
		}
	}
	if currentId > newId {
		currentId = newId
	}
	buildDenseParamIdMap(idMap, sortedIds, prevId, newId, currentId)
	return idMap, nil
}

func buildDenseParamIdMap(idMap map[int64]int64, sortedIds []int64, prevId int64, newId int64, startId int64) {
	var i int
	var val int64
	var nextId int64
	nextId = startId
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			idMap[val] = newId
		}
	}
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			continue
		}
		if nextId == newId {
			nextId = nextId + 1
		}
		idMap[val] = nextId
		nextId = nextId + 1
	}
}

func updateParamsIds(params []dao.ParamDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range params {
		oldId = params[i].ID
		newId, exists = idMap[oldId]
		if exists {
			params[i].ID = newId
		}
	}
}

func updateCompParamsParamIds(componentParams []dao.ComponentParamDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range componentParams {
		oldId = componentParams[i].ParamID
		newId, exists = idMap[oldId]
		if exists {
			componentParams[i].ParamID = newId
		}
	}
}

func updateMappingsIds(mappings []dao.MappingDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range mappings {
		oldId = mappings[i].Param
		newId, exists = idMap[oldId]
		if exists {
			mappings[i].Param = newId
		}
	}
}

func updateResultsParamIds(results []dao.ResultDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range results {
		oldId = results[i].Param
		newId, exists = idMap[oldId]
		if exists {
			results[i].Param = newId
		}
	}
}

func updateAffectedParamsParamIds(affectedParams []dao.AffectedParamDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range affectedParams {
		oldId = affectedParams[i].Param
		newId, exists = idMap[oldId]
		if exists {
			affectedParams[i].Param = newId
		}
	}
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
	var err error
	err = dao.DropConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
		return err
	}
	err = dao.DropMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
		return err
	}
	err = dao.DropParamIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.param_indicator: %v", err)
		return err
	}
	err = dao.DropDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_indicator: %v", err)
		return err
	}
	err = dao.DropComponentParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.component_param: %v", err)
		return err
	}
	err = dao.DropParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.param: %v", err)
		return err
	}
	return nil
}

func DropParamDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component: %v", err)
		return err
	}
	err = dao.DropMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.mapping: %v", err)
		return err
	}
	err = dao.DropParamIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.param_indicator: %v", err)
		return err
	}
	err = dao.DropDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_indicator: %v", err)
		return err
	}
	err = dao.DropComponentParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.component_param: %v", err)
		return err
	}
	err = dao.DropParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.param: %v", err)
		return err
	}
	return nil
}

func CreateParamDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.param: %v", err)
		return err
	}
	err = dao.CreateComponentParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.component_param: %v", err)
		return err
	}
	err = dao.CreateDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_indicator: %v", err)
		return err
	}
	err = dao.CreateParamIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.param_indicator: %v", err)
		return err
	}
	err = dao.CreateMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.config_in_process: %v", err)
		return err
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
	var err error
	err = dao.CreateParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.param: %v", err)
		return err
	}
	err = dao.CreateComponentParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.component_param: %v", err)
		return err
	}
	err = dao.CreateDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_indicator: %v", err)
		return err
	}
	err = dao.CreateParamIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.param_indicator: %v", err)
		return err
	}
	err = dao.CreateMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
		return err
	}
	return nil
}

func ChangeDeviceComponentData(ctx context.Context, prevId int64, newId int64) (bool, error) {
	if prevId == newId {
		return true, nil
	}
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
	tDevComp, tDevCompMap, tConfig, tDefConfig, err = changeDeviceComponentId(prevId, newId, tDevComp, tDevCompMap, tConfig, tDefConfig)
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

func changeDeviceComponentId(prevId int64, newId int64, deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao) ([]dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, error) {
	var i int
	var err error
	var idMap map[int64]int64
	var item dao.DeviceComponentDao
	var prevComponent dao.DeviceComponentDao
	var foundPrev bool
	var parentIdPtr *int64
	var zero int64
	zero = int64(0)
	if newId <= zero || prevId <= zero {
		return nil, nil, nil, nil, errors.New("ID cannot be less than or equal to zero")
	}
	if newId > int64(len(deviceComponents)) {
		return nil, nil, nil, nil, errors.New("ID cannot be greater than Device Components count")
	}
	for i = range deviceComponents {
		item = deviceComponents[i]
		if item.ID == prevId {
			prevComponent = item
			foundPrev = true
			break
		}
	}
	if !foundPrev {
		return nil, nil, nil, nil, fmt.Errorf("Device Component with ID = %d is not exists", prevId)
	}
	parentIdPtr = util.SqlNullInt64ToInt64Ptr(prevComponent.Parent)
	if parentIdPtr != nil && *parentIdPtr >= newId {
		return nil, nil, nil, nil, errors.New("ID of the parent device component cannot be less than the ID of the child device component")
	}
	idMap, err = calculateDeviceCompIdMap(prevId, newId, deviceComponents)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	updateDeviceCompIds(deviceComponents, idMap)
	updateDeviceCompMappingsIds(deviceComponentMappings, idMap)
	updateConfigurationsIds(configurations, idMap)
	updateDefaultConfigurationsIds(defaultConfigurations, idMap)
	return deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, nil
}

func calculateDeviceCompIdMap(prevId int64, newId int64, deviceComponents []dao.DeviceComponentDao) (map[int64]int64, error) {
	var idMap map[int64]int64
	var i int
	var length int
	var sortedIds []int64
	var currentId int64
	var firstSortedId int64
	idMap = make(map[int64]int64)
	length = len(deviceComponents)
	sortedIds = make([]int64, length)
	for i = range deviceComponents {
		sortedIds[i] = deviceComponents[i].ID
	}
	sortIdsArray(sortedIds)
	currentId = int64(1)
	if length > 0 {
		firstSortedId = sortedIds[0]
		if currentId > firstSortedId {
			currentId = firstSortedId
		}
	}
	if currentId > newId {
		currentId = newId
	}
	buildDenseDeviceCompIdMap(idMap, sortedIds, prevId, newId, currentId)
	return idMap, nil
}

func buildDenseDeviceCompIdMap(idMap map[int64]int64, sortedIds []int64, prevId int64, newId int64, startId int64) {
	var i int
	var val int64
	var nextId int64
	nextId = startId
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			idMap[val] = newId
		}
	}
	for i = range sortedIds {
		val = sortedIds[i]
		if val == prevId {
			continue
		}
		if nextId == newId {
			nextId = nextId + 1
		}
		idMap[val] = nextId
		nextId = nextId + 1
	}
}

func updateDeviceCompIds(deviceComponents []dao.DeviceComponentDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	var parentIdPtr *int64
	var newParentId int64
	for i = range deviceComponents {
		oldId = deviceComponents[i].ID
		newId, exists = idMap[oldId]
		if exists {
			deviceComponents[i].ID = newId
		}
		parentIdPtr = util.SqlNullInt64ToInt64Ptr(deviceComponents[i].Parent)
		if parentIdPtr != nil {
			oldId = *parentIdPtr
			newId, exists = idMap[oldId]
			if exists {
				newParentId = newId
				deviceComponents[i].Parent = util.Int64PtrToSqlNullInt64(&newParentId)
			}
		}
	}
}

func updateDeviceCompMappingsIds(mappings []dao.DeviceComponentMappingDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	for i = range mappings {
		oldId = mappings[i].DeviceComponentID
		newId, exists = idMap[oldId]
		if exists {
			mappings[i].DeviceComponentID = newId
		}
	}
}

func updateConfigurationsIds(configurations []dao.ConfigurationDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	var idPtr *int64
	var newNullId int64
	for i = range configurations {
		idPtr = util.SqlNullInt64ToInt64Ptr(configurations[i].DeviceComponentID)
		if idPtr != nil {
			oldId = *idPtr
			newId, exists = idMap[oldId]
			if exists {
				newNullId = newId
				configurations[i].DeviceComponentID = util.Int64PtrToSqlNullInt64(&newNullId)
			}
		}
	}
}

func updateDefaultConfigurationsIds(defaultConfigurations []dao.DefaultConfigurationDao, idMap map[int64]int64) {
	var i int
	var oldId int64
	var newId int64
	var exists bool
	var idPtr *int64
	var newNullId int64
	for i = range defaultConfigurations {
		idPtr = util.SqlNullInt64ToInt64Ptr(defaultConfigurations[i].DeviceComponentID)
		if idPtr != nil {
			oldId = *idPtr
			newId, exists = idMap[oldId]
			if exists {
				newNullId = newId
				defaultConfigurations[i].DeviceComponentID = util.Int64PtrToSqlNullInt64(&newNullId)
			}
		}
	}
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
	var err error
	err = dao.DropConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component: %v", err)
		return err
	}
	return nil
}

func DropDeviceComponentDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.threshold: %v", err)
		return err
	}
	err = dao.DropDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.default_configuration: %v", err)
		return err
	}
	err = dao.DropConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.configuration: %v", err)
		return err
	}
	err = dao.DropDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component: %v", err)
		return err
	}
	return nil
}

func CreateDeviceComponentDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateDeviceComponentDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.config_in_process: %v", err)
		return err
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
	var err error
	err = dao.CreateDeviceComponentDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.CreateConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
		return err
	}
	err = dao.CreateThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
		return err
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
	var err error
	err = dao.DropDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.mapping: %v", err)
		return err
	}
	return nil
}

func DropMappingDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_component_mapping: %v", err)
		return err
	}
	err = dao.DropMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.mapping: %v", err)
		return err
	}
	return nil
}

func CreateMappingDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component_mapping: %v", err)
		return err
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
	var err error
	err = dao.CreateMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.mapping: %v", err)
		return err
	}
	err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
		return err
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
	var err error
	err = dao.DropConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.threshold: %v", err)
		return err
	}
	return nil
}

func DropThresholdDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.config_in_process: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropResultDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.result: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.threshold: %v", err)
		return err
	}
	return nil
}

func CreateThresholdDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.config_in_process: %v", err)
		return err
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
	var err error
	err = dao.CreateThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.threshold: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateResultDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.result: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateConfigInProcessDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.config_in_process: %v", err)
		return err
	}
	return nil
}
