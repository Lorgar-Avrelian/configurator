package service

import (
	"configurator/internal/dao"
	"configurator/internal/logger"
	"context"
	"database/sql"
	"sync"
)

func CompressComponentDependentData(ctx context.Context) {
	var err error
	var tComponent []dao.ComponentDao
	var tCompParam []dao.ComponentParamDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	tComponent, tCompParam, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tResult, tAffectedTh, tAffectedParam, err = GetComponentDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading component dependent data:", err)
		return
	}
	tComponent, tCompParam, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tResult, tAffectedTh, tAffectedParam = overrideComponentDependentIds(tComponent, tCompParam, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tResult, tAffectedTh, tAffectedParam)
	err = DropComponentDependentConstraintsForComponent(ctx)
	if err != nil {
		logger.Error("Error while dropping component dependent constraints:", err)
		return
	}
	err = DropComponentDependentTablesForComponent(ctx)
	if err != nil {
		logger.Error("Error while dropping component dependent tables:", err)
		return
	}
	err = CreateComponentDependentTablesForComponent(ctx)
	if err != nil {
		logger.Error("Error while creating component dependent tables:", err)
		return
	}
	err = InsertComponentDependentDataForComponent(ctx, tComponent, tCompParam, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	if err != nil {
		logger.Error("Error while inserting component dependent data:", err)
		return
	}
	err = CreateComponentDependentConstraintsForComponent(ctx)
	if err != nil {
		logger.Error("Error while creating component dependent constraints:", err)
		return
	}
	err = ResetComponentDependentCountersForComponent(ctx)
	if err != nil {
		logger.Error("Error while resetting component dependent counters:", err)
	}
}

func overrideComponentDependentIds(components []dao.ComponentDao, params []dao.ComponentParamDao, deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, deviceSnmps []dao.DeviceSnmpDao, results []dao.ResultDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	components, params, deviceComponents, results, affectedParams = processComponents(components, params, deviceComponents, results, affectedParams, idMap)
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
	return components, params, deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, deviceSnmps, results, affectedThresholds, affectedParams
}

func GetComponentDependentDataOnly(ctx context.Context) ([]dao.ComponentDao, []dao.ComponentParamDao, []dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.ResultDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tComponent []dao.ComponentDao
	var tCompParam []dao.ComponentParamDao
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tResult []dao.ResultDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	wg.Add(10)
	go func() {
		defer wg.Done()
		var res []dao.ComponentDao
		res, err = dao.GetAllComponentDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tCompParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentDao
		res, err = dao.GetAllDeviceComponentDao(ctx)
		if err != nil {
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
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, errResult
	}
	return tComponent, tCompParam, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tResult, tAffectedTh, tAffectedParam, nil
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
				components[i].BaseComponent = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				components[i].BaseComponent = sql.NullInt64{Valid: false}
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
				deviceComponents[i].Parent = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				deviceComponents[i].Parent = sql.NullInt64{Valid: false}
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
				configurations[i].DeviceComponentID = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				configurations[i].DeviceComponentID = sql.NullInt64{Valid: false}
			}
		}
	}
	for i = range defaultConfigurations {
		if defaultConfigurations[i].DeviceComponentID.Valid {
			oldId = defaultConfigurations[i].DeviceComponentID.Int64
			newId, found = devCompMap[oldId]
			if found {
				defaultConfigurations[i].DeviceComponentID = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				defaultConfigurations[i].DeviceComponentID = sql.NullInt64{Valid: false}
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
				deviceSnmps[i].Config = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				deviceSnmps[i].Config = sql.NullInt64{Valid: false}
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

func DropComponentDependentConstraintsForComponent(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
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

func DropComponentDependentTablesForComponent(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
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

func CreateComponentDependentTablesForComponent(ctx context.Context) error {
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
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func CreateComponentDependentConstraintsForComponent(ctx context.Context) error {
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
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func InsertComponentDependentDataForComponent(ctx context.Context, tComponent []dao.ComponentDao, tCompParam []dao.ComponentParamDao, tDevComp []dao.DeviceComponentDao, tDevCompMap []dao.DeviceComponentMappingDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tDevSnmp []dao.DeviceSnmpDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(9)
	go func() {
		defer wg.Done()
		err = dao.ImportComponentDao(ctx, tComponent)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportComponentParamDao(ctx, tCompParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentDao(ctx, tDevComp)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
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
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetComponentDependentCountersForComponent(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
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

func CompressParamDependentData(ctx context.Context) {
	var err error
	var tParam []dao.ParamDao
	var tCompParam []dao.ComponentParamDao
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tResult []dao.ResultDao
	var tAffectedParam []dao.AffectedParamDao
	tParam, tCompParam, tMapping, tDevCompMap, tResult, tAffectedParam, err = GetParamDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading param dependent data:", err)
		return
	}
	tParam, tCompParam, tMapping, tDevCompMap, tResult, tAffectedParam = overrideParamDependentIds(tParam, tCompParam, tMapping, tDevCompMap, tResult, tAffectedParam)
	err = DropParamDependentConstraintsForParam(ctx)
	if err != nil {
		logger.Error("Error while dropping param dependent constraints:", err)
		return
	}
	err = DropParamDependentTablesForParam(ctx)
	if err != nil {
		logger.Error("Error while dropping param dependent tables:", err)
		return
	}
	err = CreateParamDependentTablesForParam(ctx)
	if err != nil {
		logger.Error("Error while creating param dependent tables:", err)
		return
	}
	err = InsertParamDependentDataForParam(ctx, tParam, tCompParam, tMapping, tDevCompMap, tAffectedParam)
	if err != nil {
		logger.Error("Error while inserting param dependent data:", err)
		return
	}
	err = CreateParamDependentConstraintsForParam(ctx)
	if err != nil {
		logger.Error("Error while creating param dependent constraints:", err)
		return
	}
	err = ResetParamDependentCountersForParam(ctx)
	if err != nil {
		logger.Error("Error while resetting param dependent counters:", err)
	}
}

func overrideParamDependentIds(params []dao.ParamDao, componentParams []dao.ComponentParamDao, mappings []dao.MappingDao, deviceComponentMappings []dao.DeviceComponentMappingDao, results []dao.ResultDao, affectedParams []dao.AffectedParamDao) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.MappingDao, []dao.DeviceComponentMappingDao, []dao.ResultDao, []dao.AffectedParamDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	params, componentParams, mappings, results, affectedParams = processParams(params, componentParams, mappings, results, affectedParams, idMap)
	clear(idMap)
	mappings, deviceComponentMappings = processMappings(mappings, deviceComponentMappings, idMap)
	clear(idMap)
	affectedParams = processAffectedParamsSelfIds(affectedParams, idMap)
	clear(idMap)
	return params, componentParams, mappings, deviceComponentMappings, results, affectedParams
}

func GetParamDependentDataOnly(ctx context.Context) ([]dao.ParamDao, []dao.ComponentParamDao, []dao.MappingDao, []dao.DeviceComponentMappingDao, []dao.ResultDao, []dao.AffectedParamDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tParam []dao.ParamDao
	var tCompParam []dao.ComponentParamDao
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tResult []dao.ResultDao
	var tAffectedParam []dao.AffectedParamDao
	wg.Add(6)
	go func() {
		defer wg.Done()
		var res []dao.ParamDao
		res, err = dao.GetAllParamDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tCompParam = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.MappingDao
		res, err = dao.GetAllMappingDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.ResultDao
		res, err = dao.GetAllResultDao(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tResult = res
	}()
	go func() {
		defer wg.Done()
		var res []dao.AffectedParamDao
		res, err = dao.GetAllAffectedParamDao(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, errResult
	}
	return tParam, tCompParam, tMapping, tDevCompMap, tResult, tAffectedParam, nil
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
				mappings[i].From = sql.NullInt64{Int64: newId, Valid: true}
			} else {
				mappings[i].From = sql.NullInt64{Valid: false}
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

func DropParamDependentConstraintsForParam(ctx context.Context) error {
	var err error
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
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

func DropParamDependentTablesForParam(ctx context.Context) error {
	var err error
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
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

func CreateParamDependentTablesForParam(ctx context.Context) error {
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
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	return nil
}

func CreateParamDependentConstraintsForParam(ctx context.Context) error {
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
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	return nil
}

func InsertParamDependentDataForParam(ctx context.Context, tParam []dao.ParamDao, tCompParam []dao.ComponentParamDao, tMapping []dao.MappingDao, tDevCompMap []dao.DeviceComponentMappingDao, tAffectedParam []dao.AffectedParamDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(5)
	go func() {
		defer wg.Done()
		err = dao.ImportParamDao(ctx, tParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportComponentParamDao(ctx, tCompParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetParamDependentCountersForParam(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(3)
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
		err = dao.DropMappingDaoCounter(ctx)
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
	wg.Wait()
	return errResult
}

func CompressDeviceIndicatorDependentData(ctx context.Context) {
	var err error
	var tDevInd []dao.DeviceIndicatorDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	tDevInd, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam, err = GetDeviceIndicatorDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading device indicator dependent data:", err)
		return
	}
	tDevInd, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam = overrideDeviceIndicatorDependentIds(tDevInd, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	err = DropDeviceIndicatorDependentConstraintsForDeviceIndicator(ctx)
	if err != nil {
		logger.Error("Error while dropping device indicator dependent constraints:", err)
		return
	}
	err = DropDeviceIndicatorDependentTablesForDeviceIndicator(ctx)
	if err != nil {
		logger.Error("Error while dropping device indicator dependent tables:", err)
		return
	}
	err = CreateDeviceIndicatorDependentTablesForDeviceIndicator(ctx)
	if err != nil {
		logger.Error("Error while creating device indicator dependent tables:", err)
		return
	}
	err = InsertDeviceIndicatorDependentDataForDeviceIndicator(ctx, tDevInd, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	if err != nil {
		logger.Error("Error while inserting device indicator dependent data:", err)
		return
	}
	err = CreateDeviceIndicatorDependentConstraintsForDeviceIndicator(ctx)
	if err != nil {
		logger.Error("Error while creating device indicator dependent constraints:", err)
		return
	}
	err = ResetDeviceIndicatorDependentCountersForDeviceIndicator(ctx)
	if err != nil {
		logger.Error("Error while resetting device indicator dependent counters:", err)
	}
}

func overrideDeviceIndicatorDependentIds(deviceIndicators []dao.DeviceIndicatorDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, deviceSnmps []dao.DeviceSnmpDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao) ([]dao.DeviceIndicatorDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	deviceIndicators, configurations, defaultConfigurations = processDeviceIndicatorsOnly(deviceIndicators, configurations, defaultConfigurations, idMap)
	clear(idMap)
	configurations, deviceSnmps = processConfigurations(configurations, deviceSnmps, idMap)
	clear(idMap)
	defaultConfigurations = processDefaultConfigurations(defaultConfigurations, idMap)
	clear(idMap)
	affectedParams = processAffectedParamsSelfIds(affectedParams, idMap)
	clear(idMap)
	affectedThresholds = processAffectedThresholdsSelfIds(affectedThresholds, idMap)
	clear(idMap)
	return deviceIndicators, configurations, defaultConfigurations, deviceSnmps, affectedThresholds, affectedParams
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

func GetDeviceIndicatorDependentDataOnly(ctx context.Context) ([]dao.DeviceIndicatorDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tDevInd []dao.DeviceIndicatorDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	wg.Add(6)
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
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, errResult
	}
	return tDevInd, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam, nil
}

func DropDeviceIndicatorDependentConstraintsForDeviceIndicator(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
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

func DropDeviceIndicatorDependentTablesForDeviceIndicator(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
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

func CreateDeviceIndicatorDependentTablesForDeviceIndicator(ctx context.Context) error {
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
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func CreateDeviceIndicatorDependentConstraintsForDeviceIndicator(ctx context.Context) error {
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
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func InsertDeviceIndicatorDependentDataForDeviceIndicator(ctx context.Context, tDevInd []dao.DeviceIndicatorDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tDevSnmp []dao.DeviceSnmpDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(6)
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
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetDeviceIndicatorDependentCountersForDeviceIndicator(ctx context.Context) error {
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

func CompressParamIndicatorDependentData(ctx context.Context) {
	var err error
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	tParamInd, tMapping, tDevCompMap, err = GetParamIndicatorDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading param indicator dependent data:", err)
		return
	}
	tParamInd, tMapping, tDevCompMap = overrideParamIndicatorDependentIds(tParamInd, tMapping, tDevCompMap)
	err = DropParamIndicatorDependentConstraintsForParamIndicator(ctx)
	if err != nil {
		logger.Error("Error while dropping param indicator dependent constraints:", err)
		return
	}
	err = DropParamIndicatorDependentTablesForParamIndicator(ctx)
	if err != nil {
		logger.Error("Error while dropping param indicator dependent tables:", err)
		return
	}
	err = CreateParamIndicatorDependentTablesForParamIndicator(ctx)
	if err != nil {
		logger.Error("Error while creating param indicator dependent tables:", err)
		return
	}
	err = InsertParamIndicatorDependentDataForParamIndicator(ctx, tParamInd, tMapping, tDevCompMap)
	if err != nil {
		logger.Error("Error while inserting param indicator dependent data:", err)
		return
	}
	err = CreateParamIndicatorDependentConstraintsForParamIndicator(ctx)
	if err != nil {
		logger.Error("Error while creating param indicator dependent constraints:", err)
		return
	}
	err = ResetParamIndicatorDependentCountersForParamIndicator(ctx)
	if err != nil {
		logger.Error("Error while resetting param indicator dependent counters:", err)
	}
}

func overrideParamIndicatorDependentIds(paramIndicators []dao.ParamIndicatorDao, mappings []dao.MappingDao, deviceComponentMappings []dao.DeviceComponentMappingDao) ([]dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentMappingDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	paramIndicators, mappings = processParamIndicators(paramIndicators, mappings, idMap)
	clear(idMap)
	mappings, deviceComponentMappings = processMappings(mappings, deviceComponentMappings, idMap)
	clear(idMap)
	return paramIndicators, mappings, deviceComponentMappings
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

func GetParamIndicatorDependentDataOnly(ctx context.Context) ([]dao.ParamIndicatorDao, []dao.MappingDao, []dao.DeviceComponentMappingDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tParamInd []dao.ParamIndicatorDao
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	wg.Add(3)
	go func() {
		defer wg.Done()
		var res []dao.ParamIndicatorDao
		res, err = dao.GetAllParamIndicatorDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tDevCompMap = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, errResult
	}
	return tParamInd, tMapping, tDevCompMap, nil
}

func DropParamIndicatorDependentConstraintsForParamIndicator(ctx context.Context) error {
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
	err = dao.DropParamIndicatorDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.param_indicator: %v", err)
		return err
	}
	return nil
}

func DropParamIndicatorDependentTablesForParamIndicator(ctx context.Context) error {
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
	err = dao.DropParamIndicatorDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.param_indicator: %v", err)
		return err
	}
	return nil
}

func CreateParamIndicatorDependentTablesForParamIndicator(ctx context.Context) error {
	var err error
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
	err = dao.CreateDeviceComponentMappingDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_component_mapping: %v", err)
		return err
	}
	return nil
}

func CreateParamIndicatorDependentConstraintsForParamIndicator(ctx context.Context) error {
	var err error
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
	err = dao.CreateDeviceComponentMappingDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_component_mapping: %v", err)
		return err
	}
	return nil
}

func InsertParamIndicatorDependentDataForParamIndicator(ctx context.Context, tParamInd []dao.ParamIndicatorDao, tMapping []dao.MappingDao, tDevCompMap []dao.DeviceComponentMappingDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(3)
	go func() {
		defer wg.Done()
		err = dao.ImportParamIndicatorDao(ctx, tParamInd)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetParamIndicatorDependentCountersForParamIndicator(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
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
	wg.Wait()
	return errResult
}

func CompressDeviceComponentDependentData(ctx context.Context) {
	var err error
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam, err = GetDeviceComponentDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading device component dependent data:", err)
		return
	}
	tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam = overrideDeviceComponentDependentIds(tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	err = DropDeviceComponentDependentConstraintsForDeviceComponent(ctx)
	if err != nil {
		logger.Error("Error while dropping device component dependent constraints:", err)
		return
	}
	err = DropDeviceComponentDependentTablesForDeviceComponent(ctx)
	if err != nil {
		logger.Error("Error while dropping device component dependent tables:", err)
		return
	}
	err = CreateDeviceComponentDependentTablesForDeviceComponent(ctx)
	if err != nil {
		logger.Error("Error while creating device component dependent tables:", err)
		return
	}
	err = InsertDeviceComponentDependentDataForDeviceComponent(ctx, tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	if err != nil {
		logger.Error("Error while inserting device component dependent data:", err)
		return
	}
	err = CreateDeviceComponentDependentConstraintsForDeviceComponent(ctx)
	if err != nil {
		logger.Error("Error while creating device component dependent constraints:", err)
		return
	}
	err = ResetDeviceComponentDependentCountersForDeviceComponent(ctx)
	if err != nil {
		logger.Error("Error while resetting device component dependent counters:", err)
	}
}

func overrideDeviceComponentDependentIds(deviceComponents []dao.DeviceComponentDao, deviceComponentMappings []dao.DeviceComponentMappingDao, configurations []dao.ConfigurationDao, defaultConfigurations []dao.DefaultConfigurationDao, deviceSnmps []dao.DeviceSnmpDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao) ([]dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
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
	return deviceComponents, deviceComponentMappings, configurations, defaultConfigurations, deviceSnmps, affectedThresholds, affectedParams
}

func GetDeviceComponentDependentDataOnly(ctx context.Context) ([]dao.DeviceComponentDao, []dao.DeviceComponentMappingDao, []dao.ConfigurationDao, []dao.DefaultConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tDevComp []dao.DeviceComponentDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	var tConfig []dao.ConfigurationDao
	var tDefConfig []dao.DefaultConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	wg.Add(7)
	go func() {
		defer wg.Done()
		var res []dao.DeviceComponentDao
		res, err = dao.GetAllDeviceComponentDao(ctx)
		if err != nil {
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
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, nil, nil, nil, errResult
	}
	return tDevComp, tDevCompMap, tConfig, tDefConfig, tDevSnmp, tAffectedTh, tAffectedParam, nil
}

func DropDeviceComponentDependentConstraintsForDeviceComponent(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
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

func DropDeviceComponentDependentTablesForDeviceComponent(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
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

func CreateDeviceComponentDependentTablesForDeviceComponent(ctx context.Context) error {
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
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func CreateDeviceComponentDependentConstraintsForDeviceComponent(ctx context.Context) error {
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
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func InsertDeviceComponentDependentDataForDeviceComponent(ctx context.Context, tDevComp []dao.DeviceComponentDao, tDevCompMap []dao.DeviceComponentMappingDao, tConfig []dao.ConfigurationDao, tDefConfig []dao.DefaultConfigurationDao, tDevSnmp []dao.DeviceSnmpDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(7)
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentDao(ctx, tDevComp)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
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
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetDeviceComponentDependentCountersForDeviceComponent(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(5)
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

func CompressMappingDependentData(ctx context.Context) {
	var err error
	var tMapping []dao.MappingDao
	var tDevCompMap []dao.DeviceComponentMappingDao
	tMapping, tDevCompMap, err = GetMappingDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading mapping dependent data:", err)
		return
	}
	tMapping, tDevCompMap = overrideMappingDependentIds(tMapping, tDevCompMap)
	err = DropMappingDependentConstraintsForMapping(ctx)
	if err != nil {
		logger.Error("Error while dropping mapping dependent constraints:", err)
		return
	}
	err = DropMappingDependentTablesForMapping(ctx)
	if err != nil {
		logger.Error("Error while dropping mapping dependent tables:", err)
		return
	}
	err = CreateMappingDependentTablesForMapping(ctx)
	if err != nil {
		logger.Error("Error while creating mapping dependent tables:", err)
		return
	}
	err = InsertMappingDependentDataForMapping(ctx, tMapping, tDevCompMap)
	if err != nil {
		logger.Error("Error while inserting mapping dependent data:", err)
		return
	}
	err = CreateMappingDependentConstraintsForMapping(ctx)
	if err != nil {
		logger.Error("Error while creating mapping dependent constraints:", err)
		return
	}
	err = ResetMappingDependentCountersForMapping(ctx)
	if err != nil {
		logger.Error("Error while resetting mapping dependent counters:", err)
	}
}

func overrideMappingDependentIds(mappings []dao.MappingDao, deviceComponentMappings []dao.DeviceComponentMappingDao) ([]dao.MappingDao, []dao.DeviceComponentMappingDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	mappings, deviceComponentMappings = processMappings(mappings, deviceComponentMappings, idMap)
	clear(idMap)
	return mappings, deviceComponentMappings
}

func GetMappingDependentDataOnly(ctx context.Context) ([]dao.MappingDao, []dao.DeviceComponentMappingDao, error) {
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

func DropMappingDependentConstraintsForMapping(ctx context.Context) error {
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

func DropMappingDependentTablesForMapping(ctx context.Context) error {
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

func CreateMappingDependentTablesForMapping(ctx context.Context) error {
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

func CreateMappingDependentConstraintsForMapping(ctx context.Context) error {
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

func InsertMappingDependentDataForMapping(ctx context.Context, tMapping []dao.MappingDao, tDevCompMap []dao.DeviceComponentMappingDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(2)
	go func() {
		defer wg.Done()
		err = dao.ImportMappingDao(ctx, tMapping)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportDeviceComponentMappingDao(ctx, tDevCompMap)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetMappingDependentCountersForMapping(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = dao.DropMappingDaoCounter(ctx)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func CompressConfigurationDependentData(ctx context.Context) {
	var err error
	var tConfig []dao.ConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	tConfig, tDevSnmp, tAffectedTh, tAffectedParam, err = GetConfigurationDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading configuration dependent data:", err)
		return
	}
	tConfig, tDevSnmp, tAffectedTh, tAffectedParam = overrideConfigurationDependentIds(tConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	err = DropConfigurationDependentConstraintsForConfiguration(ctx)
	if err != nil {
		logger.Error("Error while dropping configuration dependent constraints:", err)
		return
	}
	err = DropConfigurationDependentTablesForConfiguration(ctx)
	if err != nil {
		logger.Error("Error while dropping configuration dependent tables:", err)
		return
	}
	err = CreateConfigurationDependentTablesForConfiguration(ctx)
	if err != nil {
		logger.Error("Error while creating configuration dependent tables:", err)
		return
	}
	err = InsertConfigurationDependentDataForConfiguration(ctx, tConfig, tDevSnmp, tAffectedTh, tAffectedParam)
	if err != nil {
		logger.Error("Error while inserting configuration dependent data:", err)
		return
	}
	err = CreateConfigurationDependentConstraintsForConfiguration(ctx)
	if err != nil {
		logger.Error("Error while creating configuration dependent constraints:", err)
		return
	}
	err = ResetConfigurationDependentCountersForConfiguration(ctx)
	if err != nil {
		logger.Error("Error while resetting configuration dependent counters:", err)
	}
}

func overrideConfigurationDependentIds(configurations []dao.ConfigurationDao, deviceSnmps []dao.DeviceSnmpDao, affectedThresholds []dao.AffectedThresholdDao, affectedParams []dao.AffectedParamDao) ([]dao.ConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao) {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	configurations, deviceSnmps = processConfigurations(configurations, deviceSnmps, idMap)
	clear(idMap)
	affectedParams = processAffectedParamsSelfIds(affectedParams, idMap)
	clear(idMap)
	affectedThresholds = processAffectedThresholdsSelfIds(affectedThresholds, idMap)
	clear(idMap)
	return configurations, deviceSnmps, affectedThresholds, affectedParams
}

func GetConfigurationDependentDataOnly(ctx context.Context) ([]dao.ConfigurationDao, []dao.DeviceSnmpDao, []dao.AffectedThresholdDao, []dao.AffectedParamDao, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	var tConfig []dao.ConfigurationDao
	var tDevSnmp []dao.DeviceSnmpDao
	var tAffectedTh []dao.AffectedThresholdDao
	var tAffectedParam []dao.AffectedParamDao
	wg.Add(4)
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
	go func() {
		defer wg.Done()
		var res []dao.AffectedThresholdDao
		res, err = dao.GetAllAffectedThresholdDao(ctx)
		if err != nil {
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
			mu.Lock()
			errResult = err
			mu.Unlock()
			return
		}
		tAffectedParam = res
	}()
	wg.Wait()
	if errResult != nil {
		return nil, nil, nil, nil, errResult
	}
	return tConfig, tDevSnmp, tAffectedTh, tAffectedParam, nil
}

func DropConfigurationDependentConstraintsForConfiguration(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.configuration: %v", err)
		return err
	}
	return nil
}

func DropConfigurationDependentTablesForConfiguration(ctx context.Context) error {
	var err error
	err = dao.DropAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_threshold: %v", err)
		return err
	}
	err = dao.DropAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.affected_param: %v", err)
		return err
	}
	err = dao.DropDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.device_snmp: %v", err)
		return err
	}
	err = dao.DropConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.configuration: %v", err)
		return err
	}
	return nil
}

func CreateConfigurationDependentTablesForConfiguration(ctx context.Context) error {
	var err error
	err = dao.CreateConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func CreateConfigurationDependentConstraintsForConfiguration(ctx context.Context) error {
	var err error
	err = dao.CreateConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.configuration: %v", err)
		return err
	}
	err = dao.CreateDeviceSnmpDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.device_snmp: %v", err)
		return err
	}
	err = dao.CreateAffectedParamDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_param: %v", err)
		return err
	}
	err = dao.CreateAffectedThresholdDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.affected_threshold: %v", err)
		return err
	}
	return nil
}

func InsertConfigurationDependentDataForConfiguration(ctx context.Context, tConfig []dao.ConfigurationDao, tDevSnmp []dao.DeviceSnmpDao, tAffectedTh []dao.AffectedThresholdDao, tAffectedParam []dao.AffectedParamDao) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(4)
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
		err = dao.ImportDeviceSnmpDao(ctx, tDevSnmp)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedParamDao(ctx, tAffectedParam)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		err = dao.ImportAffectedThresholdDao(ctx, tAffectedTh)
		if err != nil {
			mu.Lock()
			errResult = err
			mu.Unlock()
		}
	}()
	wg.Wait()
	return errResult
}

func ResetConfigurationDependentCountersForConfiguration(ctx context.Context) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errResult error
	var err error
	wg.Add(3)
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

func CompressDefaultConfigurationDependentData(ctx context.Context) {
	var err error
	var tDefConfig []dao.DefaultConfigurationDao
	tDefConfig, err = GetDefaultConfigurationDependentDataOnly(ctx)
	if err != nil {
		logger.Error("Error while loading default configuration dependent data:", err)
		return
	}
	tDefConfig = overrideDefaultConfigurationDependentIds(tDefConfig)
	err = DropDefaultConfigurationDependentConstraintsForDefaultConfiguration(ctx)
	if err != nil {
		logger.Error("Error while dropping default configuration dependent constraints:", err)
		return
	}
	err = DropDefaultConfigurationDependentTablesForDefaultConfiguration(ctx)
	if err != nil {
		logger.Error("Error while dropping default configuration dependent tables:", err)
		return
	}
	err = CreateDefaultConfigurationDependentTablesForDefaultConfiguration(ctx)
	if err != nil {
		logger.Error("Error while creating default configuration dependent tables:", err)
		return
	}
	err = InsertDefaultConfigurationDependentDataForDefaultConfiguration(ctx, tDefConfig)
	if err != nil {
		logger.Error("Error while inserting default configuration dependent data:", err)
		return
	}
	err = CreateDefaultConfigurationDependentConstraintsForDefaultConfiguration(ctx)
	if err != nil {
		logger.Error("Error while creating default configuration dependent constraints:", err)
		return
	}
	err = ResetDefaultConfigurationDependentCountersForDefaultConfiguration(ctx)
	if err != nil {
		logger.Error("Error while resetting default configuration dependent counters:", err)
	}
}

func overrideDefaultConfigurationDependentIds(defaultConfigurations []dao.DefaultConfigurationDao) []dao.DefaultConfigurationDao {
	var idMap map[int64]int64
	idMap = make(map[int64]int64)
	defaultConfigurations = processDefaultConfigurations(defaultConfigurations, idMap)
	clear(idMap)
	return defaultConfigurations
}

func GetDefaultConfigurationDependentDataOnly(ctx context.Context) ([]dao.DefaultConfigurationDao, error) {
	var err error
	var tDefConfig []dao.DefaultConfigurationDao
	var res []dao.DefaultConfigurationDao
	res, err = dao.GetAllDefaultConfigurationDao(ctx)
	if err != nil {
		return nil, err
	}
	tDefConfig = res
	return tDefConfig, nil
}

func DropDefaultConfigurationDependentConstraintsForDefaultConfiguration(ctx context.Context) error {
	var err error
	err = dao.DropDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error dropping constraints for table public.default_configuration: %v", err)
		return err
	}
	return nil
}

func DropDefaultConfigurationDependentTablesForDefaultConfiguration(ctx context.Context) error {
	var err error
	err = dao.DropDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error dropping table public.default_configuration: %v", err)
		return err
	}
	return nil
}

func CreateDefaultConfigurationDependentTablesForDefaultConfiguration(ctx context.Context) error {
	var err error
	err = dao.CreateDefaultConfigurationDao(ctx)
	if err != nil {
		logger.Errorf("Error creating table public.default_configuration: %v", err)
		return err
	}
	return nil
}

func CreateDefaultConfigurationDependentConstraintsForDefaultConfiguration(ctx context.Context) error {
	var err error
	err = dao.CreateDefaultConfigurationDaoConstraints(ctx)
	if err != nil {
		logger.Errorf("Error creating constraints for table public.default_configuration: %v", err)
		return err
	}
	return nil
}

func InsertDefaultConfigurationDependentDataForDefaultConfiguration(ctx context.Context, tDefConfig []dao.DefaultConfigurationDao) error {
	var err error
	err = dao.ImportDefaultConfigurationDao(ctx, tDefConfig)
	if err != nil {
		return err
	}
	return nil
}

func ResetDefaultConfigurationDependentCountersForDefaultConfiguration(ctx context.Context) error {
	var err error
	err = dao.DropDefaultConfigurationDaoCounter(ctx)
	if err != nil {
		return err
	}
	return nil
}
