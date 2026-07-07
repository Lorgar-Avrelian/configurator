package service

import (
	"configurator/internal/dao"
	"configurator/internal/logger"
	"context"
	"sync"
)

func CompressComponentDependentData(ctx context.Context) {
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
		logger.Error("Error while loading component dependent data:", err)
		return
	}
	tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc = overrideComponentDependentIds(tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	err = DropComponentDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while dropping component dependent tables constraints:", err)
		return
	}
	err = DropComponentDependentTables(ctx)
	if err != nil {
		logger.Error("Error while dropping component dependent tables:", err)
		return
	}
	err = CreateComponentDependentTables(ctx)
	if err != nil {
		logger.Error("Error while creating component dependent tables:", err)
		return
	}
	err = InsertComponentDependentData(ctx, tComponent, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		logger.Error("Error while inserting component dependent tables data:", err)
		return
	}
	err = CreateComponentDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while creating component dependent tables constraints:", err)
		return
	}
	err = ResetComponentDependentCounters(ctx)
	if err != nil {
		logger.Error("Error while resetting component dependent counters:", err)
	}
}

func overrideComponentDependentIds(components []dao.ComponentDao, params []dao.ComponentParamDao, deviceIndicators []dao.DeviceIndicatorDao, paramIndicators []dao.ParamIndicatorDao, mappings []dao.MappingDao, deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, thresholds []dao.ThresholdDao, deviceSnmps []dao.DeviceSnmpDao, results []dao.ResultDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao, configInProcesses []dao.ConfigInProcessDao) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceIndicatorDao, []dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	components, params, deviceComponents, results, affectedParams = processComponents(components, params, deviceComponents, results, affectedParams, idMap)
	clear(idMap)
	deviceIndicators, mappings, configurations, defaultConfigurations = processDeviceIndicators(deviceIndicators, mappings, configurations, defaultConfigurations, idMap)
	clear(idMap)
	paramIndicators, mappings = processParamIndicators(paramIndicators, mappings, idMap)
	clear(idMap)
	thresholds, affectedThresholds = processThresholds(thresholds, affectedThresholds, idMap)
	clear(idMap)
	mappings, deviceComponentMappings = processMappings(mappings, deviceComponentMappings, idMap)
	clear(idMap)
	deviceComponents, deviceComponentMappings, configurations, defaultConfigurations = processDeviceComponents(deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, idMap)
	clear(idMap)
	configurations, deviceSnmps = processConfigurations(configurations, deviceSnmps, idMap)
	clear(idMap)
	defaultConfigurations = processDefaultConfigurations(defaultConfigurations, idMap)
	clear(idMap)
	return components, params, deviceIndicators, paramIndicators, mappings, deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, thresholds, deviceSnmps, results, affectedThresholds, affectedParams, configInProcesses
}

func ResetComponentDependentCounters(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(8)
	go func() {
		defer wg.Done()
		err = dao.DropComponentDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceIndicatorDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropParamIndicatorDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropMappingDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDeviceComponentDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropConfigurationDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropDefaultConfigurationDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropThresholdDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func processComponents(components []dao.ComponentDao, params []dao.ComponentParamDao, deviceComponents []dao.DeviceComponentDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao, compMap map[int64]int64) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceComponentDao, []dao.ResultDao, []dao.AffectedParamDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	var oldBaseId int64
	nextId = 1
	for i = range components {
		oldId = components[i].ID
		compMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range components {
		oldId = components[i].ID
		newId, found = compMap[oldId]
		if found {
			components[i].ID = newId
		}
		if components[i].BaseComponent.Valid {
			oldBaseId = components[i].BaseComponent.Int64
			newId, found = compMap[oldBaseId]
			if found {
				components[i].BaseComponent.Int64 = newId
			} else {
				components[i].BaseComponent.Valid = false
				components[i].BaseComponent.Int64 = 0
			}
		}
	}
	params, deviceComponents, results, affectedParams = updateComponentRefs(params, deviceComponents, results, affectedParams, compMap)
	return components, params, deviceComponents, results, affectedParams
}

func updateComponentRefs(params []dao.ComponentParamDao, deviceComponents []dao.DeviceComponentDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao, compMap map[int64]int64) ([]dao.ComponentParamDao, []dao.DeviceComponentDao, []dao.ResultDao, []dao.AffectedParamDao) {
	var i int
	var oldId int64
	var newId int64
	var found bool
	for i = range params {
		oldId = params[i].ComponentID
		newId, found = compMap[oldId]
		if found {
			params[i].ComponentID = newId
		}
	}
	for i = range deviceComponents {
		oldId = deviceComponents[i].Model
		newId, found = compMap[oldId]
		if found {
			deviceComponents[i].Model = newId
		}
	}
	for i = range results {
		oldId = results[i].Component
		newId, found = compMap[oldId]
		if found {
			results[i].Component = newId
		}
	}
	for i = range affectedParams {
		oldId = affectedParams[i].Component
		newId, found = compMap[oldId]
		if found {
			affectedParams[i].Component = newId
		}
	}
	return params, deviceComponents, results, affectedParams
}

func processDeviceIndicators(deviceIndicators []dao.DeviceIndicatorDao, mappings []dao.MappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, indMap map[int64]int64) ([]dao.DeviceIndicatorDao, []dao.MappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range deviceIndicators {
		oldId = deviceIndicators[i].ID
		indMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range deviceIndicators {
		oldId = deviceIndicators[i].ID
		newId, found = indMap[oldId]
		if found {
			deviceIndicators[i].ID = newId
		}
	}
	for i = range mappings {
		oldId = mappings[i].Indicator
		newId, found = indMap[oldId]
		if found {
			mappings[i].Indicator = newId
		}
	}
	for i = range configurations {
		oldId = configurations[i].Indicator
		newId, found = indMap[oldId]
		if found {
			configurations[i].Indicator = newId
		}
	}
	for i = range defaultConfigurations {
		oldId = defaultConfigurations[i].Indicator
		newId, found = indMap[oldId]
		if found {
			defaultConfigurations[i].Indicator = newId
		}
	}
	return deviceIndicators, mappings, configurations, defaultConfigurations
}

func processParamIndicators(paramIndicators []dao.ParamIndicatorDao, mappings []dao.MappingDao, pIndMap map[int64]int64) ([]dao.ParamIndicatorDao, []dao.MappingDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range paramIndicators {
		oldId = paramIndicators[i].ID
		pIndMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range paramIndicators {
		oldId = paramIndicators[i].ID
		newId, found = pIndMap[oldId]
		if found {
			paramIndicators[i].ID = newId
		}
	}
	return paramIndicators, mappings
}

func processThresholds(thresholds []dao.ThresholdDao, affectedThresholds []dao.AffectedThresholdDao, threshMap map[int64]int64) ([]dao.ThresholdDao, []dao.AffectedThresholdDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range thresholds {
		oldId = thresholds[i].ID
		threshMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range thresholds {
		oldId = thresholds[i].ID
		newId, found = threshMap[oldId]
		if found {
			thresholds[i].ID = newId
		}
	}
	for i = range affectedThresholds {
		oldId = affectedThresholds[i].Threshold
		newId, found = threshMap[oldId]
		if found {
			affectedThresholds[i].Threshold = newId
		}
	}
	return thresholds, affectedThresholds
}

func processMappings(mappings []dao.MappingDao, deviceComponentMappings []dao.DeviceComponentMappingDao, mapIdMap map[int64]int64) ([]dao.MappingDao, []dao.DeviceComponentMappingDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	var oldFromId int64
	nextId = 1
	for i = range mappings {
		oldId = mappings[i].ID
		mapIdMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range mappings {
		oldId = mappings[i].ID
		newId, found = mapIdMap[oldId]
		if found {
			mappings[i].ID = newId
		}
		if mappings[i].From.Valid {
			oldFromId = mappings[i].From.Int64
			newId, found = mapIdMap[oldFromId]
			if found {
				mappings[i].From.Int64 = newId
			} else {
				mappings[i].From.Valid = false
				mappings[i].From.Int64 = 0
			}
		}
	}
	for i = range deviceComponentMappings {
		oldId = deviceComponentMappings[i].MappingID
		newId, found = mapIdMap[oldId]
		if found {
			deviceComponentMappings[i].MappingID = newId
		}
	}
	return mappings, deviceComponentMappings
}

func processDeviceComponents(deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, devCompMap map[int64]int64) ([]dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	var oldParentId int64
	nextId = 1
	for i = range deviceComponents {
		oldId = deviceComponents[i].ID
		devCompMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range deviceComponents {
		oldId = deviceComponents[i].ID
		newId, found = devCompMap[oldId]
		if found {
			deviceComponents[i].ID = newId
		}
		if deviceComponents[i].Parent.Valid {
			oldParentId = deviceComponents[i].Parent.Int64
			newId, found = devCompMap[oldParentId]
			if found {
				deviceComponents[i].Parent.Int64 = newId
			} else {
				deviceComponents[i].Parent.Valid = false
				deviceComponents[i].Parent.Int64 = 0
			}
		}
	}
	deviceComponentMappings, configurations, defaultConfigurations = updateDeviceComponentRefs(deviceComponentMappings, configurations, defaultConfigurations, devCompMap)
	return deviceComponents, deviceComponentMappings, configurations, defaultConfigurations
}

func updateDeviceComponentRefs(deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, devCompMap map[int64]int64) ([]dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao) {
	var i int
	var oldId int64
	var newId int64
	var found bool
	for i = range deviceComponentMappings {
		oldId = deviceComponentMappings[i].DeviceComponentID
		newId, found = devCompMap[oldId]
		if found {
			deviceComponentMappings[i].DeviceComponentID = newId
		}
	}
	for i = range configurations {
		if configurations[i].DeviceComponentID.Valid {
			oldId = configurations[i].DeviceComponentID.Int64
			newId, found = devCompMap[oldId]
			if found {
				configurations[i].DeviceComponentID.Int64 = newId
			} else {
				configurations[i].DeviceComponentID.Valid = false
				configurations[i].DeviceComponentID.Int64 = 0
			}
		}
	}
	for i = range defaultConfigurations {
		if defaultConfigurations[i].DeviceComponentID.Valid {
			oldId = defaultConfigurations[i].DeviceComponentID.Int64
			newId, found = devCompMap[oldId]
			if found {
				defaultConfigurations[i].DeviceComponentID.Int64 = newId
			} else {
				defaultConfigurations[i].DeviceComponentID.Valid = false
				defaultConfigurations[i].DeviceComponentID.Int64 = 0
			}
		}
	}
	return deviceComponentMappings, configurations, defaultConfigurations
}

func processConfigurations(configurations []dao.ConfigurationDao, deviceSnmps []dao.DeviceSnmpDao, configMap map[int64]int64) ([]dao.ConfigurationDao, []dao.DeviceSnmpDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range configurations {
		oldId = configurations[i].ID
		configMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range configurations {
		oldId = configurations[i].ID
		newId, found = configMap[oldId]
		if found {
			configurations[i].ID = newId
		}
	}
	for i = range deviceSnmps {
		if deviceSnmps[i].Config.Valid {
			oldId = deviceSnmps[i].Config.Int64
			newId, found = configMap[oldId]
			if found {
				deviceSnmps[i].Config.Int64 = newId
			} else {
				deviceSnmps[i].Config.Valid = false
				deviceSnmps[i].Config.Int64 = 0
			}
		}
	}
	return configurations, deviceSnmps
}

func processDefaultConfigurations(defaultConfigurations []dao.DefaultConfigurationDao, defaultConfigMap map[int64]int64) []dao.DefaultConfigurationDao {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range defaultConfigurations {
		oldId = defaultConfigurations[i].ID
		defaultConfigMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range defaultConfigurations {
		oldId = defaultConfigurations[i].ID
		newId, found = defaultConfigMap[oldId]
		if found {
			defaultConfigurations[i].ID = newId
		}
	}
	return defaultConfigurations
}
