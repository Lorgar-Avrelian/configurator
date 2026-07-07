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
	affectedParams = processAffectedParamsSelfIds(affectedParams, idMap)
	clear(idMap)
	affectedThresholds = processAffectedThresholdsSelfIds(affectedThresholds, idMap)
	clear(idMap)
	return components, params, deviceIndicators, paramIndicators, mappings, deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, thresholds, deviceSnmps, results, affectedThresholds, affectedParams, configInProcesses
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
		oldId = AntiquateResultComp(results[i], compMap)
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

func AntiquateResultComp(res dao.ResultDao, compMap map[int64]int64) int64 {
	var oldId int64
	var newId int64
	var found bool
	oldId = res.Component
	newId, found = compMap[oldId]
	if found {
		res.Component = newId
	}
	return oldId
}

func CompressParamDependentData(ctx context.Context) {
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
		logger.Error("Error while loading param dependent data:", err)
		return
	}
	tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc = overrideParamDependentIds(tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	err = DropParamDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while dropping param dependent tables constraints:", err)
		return
	}
	err = DropParamDependentTables(ctx)
	if err != nil {
		logger.Error("Error while dropping param dependent tables:", err)
		return
	}
	err = CreateParamDependentTables(ctx)
	if err != nil {
		logger.Error("Error while creating param dependent tables:", err)
		return
	}
	err = InsertParamDependentData(ctx, tParam, tCompParam, tDevInd, tParamInd, tMapping, tDevComp, tDevCompMap, tConfig, tDefConfig, tThreshold, tDevSnmp, tResult, tAffectedTh, tAffectedParam, tConfigProc)
	if err != nil {
		logger.Error("Error while inserting param dependent tables data:", err)
		return
	}
	err = CreateParamDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while creating param dependent tables constraints:", err)
		return
	}
	err = ResetParamDependentCounters(ctx)
	if err != nil {
		logger.Error("Error while resetting param dependent counters:", err)
	}
}

func overrideParamDependentIds(params []dao.ParamDao, componentParams []dao.ComponentParamDao, deviceIndicators []dao.DeviceIndicatorDao, paramIndicators []dao.ParamIndicatorDao, mappings []dao.MappingDao, deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, thresholds []dao.ThresholdDao, deviceSnmps []dao.DeviceSnmpDao, results []dao.ResultDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao, configInProcesses []dao.ConfigInProcessDao) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.DeviceIndicatorDao, []dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.ThresholdDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, []dao.ConfigInProcessDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	params, componentParams, mappings, results, affectedParams = processParams(params, componentParams, mappings, results, affectedParams, idMap)
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
	affectedParams = processAffectedParamsSelfIds(affectedParams, idMap)
	clear(idMap)
	affectedThresholds = processAffectedThresholdsSelfIds(affectedThresholds, idMap)
	clear(idMap)
	return params, componentParams, deviceIndicators, paramIndicators, mappings, deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, thresholds, deviceSnmps, results, affectedThresholds, affectedParams, configInProcesses
}

func processParams(params []dao.ParamDao, componentParams []dao.ComponentParamDao, mappings []dao.MappingDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao, paramMap map[int64]int64) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.MappingDao, []dao.ResultDao, []dao.AffectedParamDao) {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range params {
		oldId = params[i].ID
		paramMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range params {
		oldId = params[i].ID
		newId, found = paramMap[oldId]
		if found {
			params[i].ID = newId
		}
	}
	componentParams, mappings, results, affectedParams = updateParamRefs(componentParams, mappings, results, affectedParams, paramMap)
	return params, componentParams, mappings, results, affectedParams
}

func updateParamRefs(componentParams []dao.ComponentParamDao, mappings []dao.MappingDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao, paramMap map[int64]int64) ([]dao.ComponentParamDao, []dao.MappingDao, []dao.ResultDao, []dao.AffectedParamDao) {
	var i int
	var oldId int64
	var newId int64
	var found bool
	for i = range componentParams {
		oldId = componentParams[i].ParamID
		newId, found = paramMap[oldId]
		if found {
			componentParams[i].ParamID = newId
		}
	}
	for i = range mappings {
		oldId = mappings[i].Param
		newId, found = paramMap[oldId]
		if found {
			mappings[i].Param = newId
		}
	}
	for i = range results {
		oldId = results[i].Param
		newId, found = paramMap[oldId]
		if found {
			results[i].Param = newId
		}
	}
	for i = range affectedParams {
		oldId = affectedParams[i].Param
		newId, found = paramMap[oldId]
		if found {
			affectedParams[i].Param = newId
		}
	}
	return componentParams, mappings, results, affectedParams
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

func processAffectedParamsSelfIds(affectedParams []dao.AffectedParamDao, affParMap map[int64]int64) []dao.AffectedParamDao {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range affectedParams {
		oldId = affectedParams[i].ID
		affParMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range affectedParams {
		oldId = affectedParams[i].ID
		newId, found = affParMap[oldId]
		if found {
			affectedParams[i].ID = newId
		}
	}
	return affectedParams
}

func processAffectedThresholdsSelfIds(affectedThresholds []dao.AffectedThresholdDao, affThreshMap map[int64]int64) []dao.AffectedThresholdDao {
	var i int
	var nextId int64
	var oldId int64
	var newId int64
	var found bool
	nextId = 1
	for i = range affectedThresholds {
		oldId = affectedThresholds[i].ID
		affThreshMap[oldId] = nextId
		nextId = nextId + 1
	}
	for i = range affectedThresholds {
		oldId = affectedThresholds[i].ID
		newId, found = affThreshMap[oldId]
		if found {
			affectedThresholds[i].ID = newId
		}
	}
	return affectedThresholds
}

func ResetComponentDependentCounters(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
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
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetParamDependentCounters(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(10)
	go func() {
		defer wg.Done()
		err = dao.DropParamDaoCounter(ctx)
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
	go func() {
		defer wg.Done()
		err = dao.DropAffectedParamDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func CompressDeviceIndicatorDependentData(ctx context.Context) {
	var err error
	var tDevInd []dao.DeviceIndicatorDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	tDevInd, tConfig, tDefConfig, tDevSnmp, err = GetDeviceIndicatorDependentData(ctx)
	if err != nil {
		logger.Error("Error while loading device indicator dependent data:", err)
		return
	}
	tDevInd, tConfig, tDefConfig, tDevSnmp = overrideDeviceIndicatorDependentIds(tDevInd, tConfig, tDefConfig, tDevSnmp)
	err = DropDeviceIndicatorDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while dropping device indicator constraints:", err)
		return
	}
	err = DropDeviceIndicatorDependentTables(ctx)
	if err != nil {
		logger.Error("Error while dropping device indicator tables:", err)
		return
	}
	err = CreateDeviceIndicatorDependentTables(ctx)
	if err != nil {
		logger.Error("Error while creating device indicator tables:", err)
		return
	}
	err = InsertDeviceIndicatorDependentData(ctx, tDevInd, tConfig, tDefConfig, tDevSnmp)
	if err != nil {
		logger.Error("Error while inserting device indicator data:", err)
		return
	}
	err = CreateDeviceIndicatorDependentConstraints(ctx)
	if err != nil {
		logger.Error("Error while creating device indicator constraints:", err)
		return
	}
	err = ResetDeviceIndicatorDependentCounters(ctx)
	if err != nil {
		logger.Error("Error while resetting device indicator counters:", err)
	}
}

func overrideDeviceIndicatorDependentIds(deviceIndicators []dao.DeviceIndicatorDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, deviceSnmps []dao.DeviceSnmpDao) ([]dao.DeviceIndicatorDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	deviceIndicators, configurations, defaultConfigurations = processDeviceIndicatorsOnly(deviceIndicators, configurations, defaultConfigurations, idMap)
	clear(idMap)
	configurations, deviceSnmps = processConfigurations(configurations, deviceSnmps, idMap)
	clear(idMap)
	defaultConfigurations = processDefaultConfigurations(defaultConfigurations, idMap)
	clear(idMap)
	return deviceIndicators, configurations, defaultConfigurations, deviceSnmps
}

func processDeviceIndicatorsOnly(deviceIndicators []dao.DeviceIndicatorDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, indMap map[int64]int64) ([]dao.DeviceIndicatorDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao) {
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
	return deviceIndicators, configurations, defaultConfigurations
}

func GetDeviceIndicatorDependentData(ctx context.Context) ([]dao.DeviceIndicatorDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tDevInd []dao.DeviceIndicatorDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	wg.Add(4)
	go func() {
		defer wg.Done()
		var res []dao.DeviceIndicatorDao
		res, err = dao.GetAllDeviceIndicatorDao(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevInd = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ConfigurationDao
		res, err = dao.GetAllConfigurationDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDefConfig = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceSnmpDao
		res, err = dao.GetAllDeviceSnmpDao(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevSnmp = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, errResult
	}
	return tDevInd, tConfig, tDefConfig, tDevSnmp, nil
}

func DropDeviceIndicatorDependentConstraints(ctx context.Context) error {
	var err error
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
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
	err = dao.DropDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_indicator: %v", err)
		return err
	}
	return nil
}

func DropDeviceIndicatorDependentTables(ctx context.Context) error {
	var err error
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
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
	err = dao.DropDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_indicator: %v", err)
		return err
	}
	return nil
}

func CreateDeviceIndicatorDependentTables(ctx context.Context) error {
	var err error
	err = dao.CreateDeviceIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_indicator: %v", err)
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
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	return nil
}

func CreateDeviceIndicatorDependentConstraints(ctx context.Context) error {
	var err error
	err = dao.CreateDeviceIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_indicator: %v", err)
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
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	return nil
}

func InsertDeviceIndicatorDependentData(ctx context.Context, tDevInd []dao.DeviceIndicatorDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tDevSnmp []dao.DeviceSnmpDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(4)
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceIndicatorDao(ctx, tDevInd)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportConfigurationDao(ctx, tConfig)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDefaultConfigurationDao(ctx, tDefConfig)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetDeviceIndicatorDependentCounters(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(5)
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
		err = dao.DropAffectedParamDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.DropAffectedThresholdDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}
