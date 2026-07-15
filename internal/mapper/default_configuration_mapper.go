package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/model"
	"configurator/internal/util"
	"encoding/json"
	"sort"
)

func DefaultConfigurationToDefaultConfigurationDto(rows []dao.DefaultConfiguration) *dto.DefaultConfigurationDto {
	if 0 >= len(rows) {
		return nil
	}
	var first dao.DefaultConfiguration
	first = rows[0]
	var res dto.DefaultConfigurationDto
	res.ID = first.ID
	var indDto dto.DeviceIndicatorDto
	indDto.ID = first.IndicatorID
	indDto.Description = util.SqlNullStringToStringPtr(first.IndDescription)
	indDto.ObjectID = first.IndObjectID
	indDto.Contact = util.SqlNullStringToStringPtr(first.IndContact)
	indDto.Name = util.SqlNullStringToStringPtr(first.IndName)
	indDto.Location = util.SqlNullStringToStringPtr(first.IndLocation)
	indDto.Services = util.SqlNullInt16ToInt16Ptr(first.IndServices)
	res.Indicator = &indDto
	if !first.DeviceComponentID.Valid {
		return &res
	}
	var dcMap map[int64]*dto.DeviceComponentDto
	var compMaps map[int64]map[int64]*dto.MappingDto
	dcMap = make(map[int64]*dto.DeviceComponentDto)
	compMaps = make(map[int64]map[int64]*dto.MappingDto)
	var i int
	var r dao.DefaultConfiguration
	for i = 0; len(rows) > i; i++ {
		r = rows[i]
		processComponentRow(r, dcMap, compMaps)
	}
	res.DeviceComponent = assembleComponentTree(first.DeviceComponentID.Int64, dcMap, compMaps)
	return &res
}

func processComponentRow(r dao.DefaultConfiguration, dcMap map[int64]*dto.DeviceComponentDto, compMaps map[int64]map[int64]*dto.MappingDto) {
	if !r.DeviceComponentID.Valid {
		return
	}
	var dcID int64
	dcID = r.DeviceComponentID.Int64
	var ok bool
	_, ok = dcMap[dcID]
	if !ok {
		var dcDto dto.DeviceComponentDto
		var info dto.ComponentInfoDto
		info.ID = r.CompModel.Int64
		if r.CompTitle.Valid {
			info.Title = r.CompTitle.String
		}
		if r.CompNameEn.Valid {
			info.NameEn = r.CompNameEn.String
		}
		if r.CompNameRu.Valid {
			info.NameRu = r.CompNameRu.String
		}
		dcDto.ID = dcID
		dcDto.Component = info
		dcDto.InternalOrder = util.SqlNullInt32ToInt32Ptr(r.CompInternalOrder)
		dcDto.Parent = util.SqlNullInt64ToInt64Ptr(r.CompParent)
		dcDto.Mappings = []dto.MappingDto{}
		dcDto.Children = []dto.DeviceComponentDto{}
		dcMap[dcID] = &dcDto
		compMaps[dcID] = make(map[int64]*dto.MappingDto)
	}
	processMappingRow(r, compMaps[dcID])
}

func processMappingRow(r dao.DefaultConfiguration, mapMap map[int64]*dto.MappingDto) {
	if !r.MapID.Valid {
		return
	}
	var mID int64
	mID = r.MapID.Int64
	var ok bool
	_, ok = mapMap[mID]
	if ok {
		return
	}
	var mDto dto.MappingDto
	mDto.ID = mID
	mDto.Frequency = model.PollingFrequency(r.MapFrequency.Int16).String()
	mDto.Value = util.SqlNullStringToStringPtr(r.MapValue)
	mDto.Coefficient = util.SqlNullFloat64ToFloat64Ptr(r.MapCoefficient)
	if len(r.MapEnum) > 0 {
		var enumMap map[string]string
		enumMap = make(map[string]string)
		_ = json.Unmarshal(r.MapEnum, &enumMap)
		mDto.Enum = enumMap
	}
	mDto.Position = util.SqlNullInt16ToInt16Ptr(r.MapPosition)
	mDto.From = util.SqlNullInt64ToInt64Ptr(r.MapFrom)
	if r.MapPositionType.Valid {
		var posTypeStr string
		posTypeStr = model.OidType(r.MapPositionType.Int16).String()
		mDto.PositionType = &posTypeStr
	}
	var indDto dto.ParamIndicatorDto
	indDto.ID = r.MapIndicatorID.Int64
	indDto.DotterNotation = util.SqlNullStringToStringPtr(r.MapIndDotter)
	if r.MapIndOidID != nil {
		var oDto dto.OidDto
		oDto.ID = r.MapIndOidID.String()
		if r.MapOidMibID.Valid {
			var mibDto dto.MibDto
			mibDto.ID = r.MapOidMibID.Int64
			if r.MapOidMibPath.Valid {
				mibDto.Path = r.MapOidMibPath.String
			}
			mibDto.Name = util.SqlNullStringToStringPtr(r.MapOidMibName)
			if r.MapOidMibVendor.Valid {
				var vName string
				vName = model.Vendor(r.MapOidMibVendor.Int64).Data().Name
				mibDto.Vendor = &vName
			}
			oDto.Mib = &mibDto
		}
		if r.MapOidType.Valid {
			oDto.Type = model.Asn1Type(r.MapOidType.Int16).String()
		}
		if r.MapOidName.Valid {
			oDto.Name = r.MapOidName.String
		}
		oDto.Number = util.SqlNullInt32ToInt32Ptr(r.MapOidNumber)
		if r.MapOidDotter.Valid {
			oDto.DotterNotation = r.MapOidDotter.String
		}
		if r.MapOidDescriptor.Valid {
			oDto.ObjectDescriptor = r.MapOidDescriptor.String
		}
		oDto.Syntax = util.SqlNullStringToStringPtr(r.MapOidSyntax)
		if len(r.MapOidEnum) > 0 {
			var oEnumMap map[string]string
			oEnumMap = make(map[string]string)
			_ = json.Unmarshal(r.MapOidEnum, &oEnumMap)
			oDto.Enum = oEnumMap
		}
		if r.MapOidStatus.Valid {
			var sStr string
			sStr = model.OidStatus(r.MapOidStatus.Int16).String()
			oDto.Status = &sStr
		}
		if r.MapOidAccess.Valid {
			var aStr string
			aStr = model.OidAccess(r.MapOidAccess.Int16).String()
			oDto.Access = &aStr
		}
		oDto.Units = util.SqlNullStringToStringPtr(r.MapOidUnits)
		oDto.Description = util.SqlNullStringToStringPtr(r.MapOidDescription)
		oDto.Category = util.SqlNullStringToStringPtr(r.MapOidCategory)
		indDto.Oid = &oDto
	}
	mDto.Indicator = &indDto
	var pDto dto.ParamDto
	pDto.ID = r.MapParamID.Int64
	if r.MapParamTitle.Valid {
		pDto.Title = r.MapParamTitle.String
	}
	if r.MapParamNameEn.Valid {
		pDto.NameEn = r.MapParamNameEn.String
	}
	if r.MapParamNameRu.Valid {
		pDto.NameRu = r.MapParamNameRu.String
	}
	if r.MapParamType.Valid {
		pDto.Type = model.VarType(r.MapParamType.Int16).String()
	}
	pDto.Value = util.SqlNullStringToStringPtr(r.MapParamValue)
	pDto.DescriptionEn = util.SqlNullStringToStringPtr(r.MapParamDescEn)
	pDto.DescriptionRu = util.SqlNullStringToStringPtr(r.MapParamDescRu)
	pDto.UnitsEn = util.SqlNullStringToStringPtr(r.MapParamUnitsEn)
	pDto.UnitsRu = util.SqlNullStringToStringPtr(r.MapParamUnitsRu)
	if r.MapParamAccess.Valid {
		pDto.Access = model.Access(r.MapParamAccess.Int16).String()
	}
	pDto.Saved = r.MapParamSaved.Bool
	pDto.Visible = r.MapParamVisible.Bool
	pDto.Diagram = r.MapParamDiagram.Bool
	mDto.Param = &pDto
	mDto.Children = []dto.MappingDto{}
	mapMap[mID] = &mDto
}

func assembleComponentTree(rootID int64, dcMap map[int64]*dto.DeviceComponentDto, compMaps map[int64]map[int64]*dto.MappingDto) *dto.DeviceComponentDto {
	var childrenMap map[int64][]*dto.DeviceComponentDto
	childrenMap = make(map[int64][]*dto.DeviceComponentDto)
	var dc *dto.DeviceComponentDto
	for _, dc = range dcMap {
		if dc.Parent != nil {
			childrenMap[*dc.Parent] = append(childrenMap[*dc.Parent], dc)
		}
		var flatMaps map[int64]*dto.MappingDto
		flatMaps = compMaps[dc.ID]
		var mapChildren map[int64][]*dto.MappingDto
		mapChildren = make(map[int64][]*dto.MappingDto)
		var m *dto.MappingDto
		for _, m = range flatMaps {
			if m.From != nil {
				mapChildren[*m.From] = append(mapChildren[*m.From], m)
			}
		}
		var mapRoots []*dto.MappingDto
		mapRoots = []*dto.MappingDto{}
		for _, m = range flatMaps {
			if m.From == nil {
				buildFlatMappingTree(m, mapChildren)
				mapRoots = append(mapRoots, m)
			}
		}
		sort.Slice(mapRoots, func(i, j int) bool {
			return mapRoots[j].ID > mapRoots[i].ID
		})
		var valM dto.MappingDto
		for _, m = range mapRoots {
			valM = *m
			dc.Mappings = append(dc.Mappings, valM)
		}
	}
	var root *dto.DeviceComponentDto
	root = dcMap[rootID]
	if root == nil {
		return nil
	}
	buildFlatComponentTree(root, childrenMap)
	return root
}

func buildFlatComponentTree(parent *dto.DeviceComponentDto, childrenMap map[int64][]*dto.DeviceComponentDto) {
	var directChildren []*dto.DeviceComponentDto
	var ok bool
	directChildren, ok = childrenMap[parent.ID]
	if !ok {
		return
	}
	sort.Slice(directChildren, func(i, j int) bool {
		if directChildren[i].Component.ID != directChildren[j].Component.ID {
			return directChildren[j].Component.ID > directChildren[i].Component.ID
		}
		var ordI int32
		var ordJ int32
		ordI = 0
		ordJ = 0
		if directChildren[i].InternalOrder != nil {
			ordI = *directChildren[i].InternalOrder
		}
		if directChildren[j].InternalOrder != nil {
			ordJ = *directChildren[j].InternalOrder
		}
		return ordJ > ordI
	})
	var i int
	var child *dto.DeviceComponentDto
	var valDc dto.DeviceComponentDto
	for i = 0; len(directChildren) > i; i++ {
		child = directChildren[i]
		buildFlatComponentTree(child, childrenMap)
		valDc = *child
		parent.Children = append(parent.Children, valDc)
	}
}

func buildFlatMappingTree(parent *dto.MappingDto, childrenMap map[int64][]*dto.MappingDto) {
	var directChildren []*dto.MappingDto
	var ok bool
	directChildren, ok = childrenMap[parent.ID]
	if !ok {
		return
	}
	sort.Slice(directChildren, func(i, j int) bool {
		var posI int16
		var posJ int16
		posI = 0
		posJ = 0
		if directChildren[i].Position != nil {
			posI = *directChildren[i].Position
		}
		if directChildren[j].Position != nil {
			posJ = *directChildren[j].Position
		}
		return posJ > posI
	})
	var i int
	var child *dto.MappingDto
	var valM dto.MappingDto
	for i = 0; len(directChildren) > i; i++ {
		child = directChildren[i]
		buildFlatMappingTree(child, childrenMap)
		valM = *child
		parent.Children = append(parent.Children, valM)
	}
}

func DefaultConfigurationArrayToDefaultConfigurationDtoArray(arr []dao.DefaultConfiguration) []dto.DefaultConfigurationDto {
	var groups map[int64][]dao.DefaultConfiguration
	groups = make(map[int64][]dao.DefaultConfiguration)
	var keys []int64
	keys = []int64{}
	var i int
	var r dao.DefaultConfiguration
	for i = 0; len(arr) > i; i++ {
		r = arr[i]
		var ok bool
		_, ok = groups[r.ID]
		if !ok {
			keys = append(keys, r.ID)
		}
		groups[r.ID] = append(groups[r.ID], r)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[j] > keys[i]
	})
	var res []dto.DefaultConfigurationDto
	res = []dto.DefaultConfigurationDto{}
	var k int64
	for _, k = range keys {
		var itemDto *dto.DefaultConfigurationDto
		itemDto = DefaultConfigurationToDefaultConfigurationDto(groups[k])
		if itemDto != nil {
			res = append(res, *itemDto)
		}
	}
	return res
}
