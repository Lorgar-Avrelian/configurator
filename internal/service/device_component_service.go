package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
	"sort"
)

func CreateDeviceComponent(ctx context.Context, d dto.DeviceComponentCreateDto) (*dto.DeviceComponentDto, error) {
	var daoInput dao.DeviceComponentDao
	var insertedID int64
	var err error
	var res *dto.DeviceComponentDto
	daoInput = mapper.DeviceComponentCreateDtoToDeviceComponentDao(d)
	insertedID, err = dao.CreateDeviceComponentRaw(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	res, err = GetDeviceComponentByID(ctx, insertedID)
	return res, err
}

func GetDeviceComponentByID(ctx context.Context, id int64) (*dto.DeviceComponentDto, error) {
	var arr []dao.DeviceComponent
	var err error
	arr, err = dao.GetDeviceComponentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	var flatDto []dto.DeviceComponentDto
	var item dao.DeviceComponent
	flatDto = []dto.DeviceComponentDto{}
	for _, item = range arr {
		flatDto = append(flatDto, mapper.DeviceComponentToDeviceComponentDto(item))
	}
	var childrenMap map[int64][]dto.DeviceComponentDto
	var d dto.DeviceComponentDto
	childrenMap = make(map[int64][]dto.DeviceComponentDto)
	for _, d = range flatDto {
		if d.Parent != nil {
			childrenMap[*d.Parent] = append(childrenMap[*d.Parent], d)
		}
	}
	var targetRoot dto.DeviceComponentDto
	var found bool
	found = false
	for _, d = range flatDto {
		if d.Parent == nil {
			targetRoot = d
			found = true
			break
		}
	}
	if !found {
		return nil, nil
	}
	buildDeviceComponentTree(ctx, &targetRoot, childrenMap)
	return &targetRoot, nil
}

func GetAllDeviceComponents(ctx context.Context) ([]dto.DeviceComponentDto, error) {
	var arr []dao.DeviceComponent
	var err error
	arr, err = dao.GetAllDeviceComponents(ctx)
	if err != nil {
		return nil, err
	}
	var flatDto []dto.DeviceComponentDto
	var item dao.DeviceComponent
	flatDto = []dto.DeviceComponentDto{}
	for _, item = range arr {
		flatDto = append(flatDto, mapper.DeviceComponentToDeviceComponentDto(item))
	}
	var childrenMap map[int64][]dto.DeviceComponentDto
	var d dto.DeviceComponentDto
	childrenMap = make(map[int64][]dto.DeviceComponentDto)
	for _, d = range flatDto {
		if d.Parent != nil {
			childrenMap[*d.Parent] = append(childrenMap[*d.Parent], d)
		}
	}
	var roots []dto.DeviceComponentDto
	roots = []dto.DeviceComponentDto{}
	for _, d = range flatDto {
		if d.Parent == nil {
			buildDeviceComponentTree(ctx, &d, childrenMap)
			roots = append(roots, d)
		}
	}
	sort.Slice(roots, func(i, j int) bool {
		return roots[j].ID > roots[i].ID
	})
	return roots, nil
}

func UpdateDeviceComponent(ctx context.Context, id int64, d dto.DeviceComponentCreateDto) (*dto.DeviceComponentDto, error) {
	var daoInput dao.DeviceComponentDao
	var err error
	var res *dto.DeviceComponentDto
	daoInput = mapper.DeviceComponentCreateDtoToDeviceComponentDao(d)
	err = dao.UpdateDeviceComponentRaw(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	res, err = GetDeviceComponentByID(ctx, id)
	return res, err
}

func DeleteDeviceComponent(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteDeviceComponent(ctx, id)
	return found, err
}

func buildDeviceComponentTree(ctx context.Context, parent *dto.DeviceComponentDto, childrenMap map[int64][]dto.DeviceComponentDto) {
	enrichComponentWithMappings(ctx, parent)
	var directChildren []dto.DeviceComponentDto
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
	var child dto.DeviceComponentDto
	for i = 0; len(directChildren) > i; i++ {
		child = directChildren[i]
		buildDeviceComponentTree(ctx, &child, childrenMap)
		parent.Children = append(parent.Children, child)
	}
}

func enrichComponentWithMappings(ctx context.Context, node *dto.DeviceComponentDto) {
	var dbMappings []dao.Mapping
	var err error
	dbMappings, err = dao.GetMappingsByDeviceComponentID(ctx, node.ID)
	if err != nil || len(dbMappings) == 0 {
		node.Mappings = []dto.MappingDto{}
		return
	}
	var flatDto []dto.MappingDto
	var item dao.Mapping
	flatDto = []dto.MappingDto{}
	for _, item = range dbMappings {
		flatDto = append(flatDto, mapper.MappingToMappingDto(item))
	}
	var childrenMap map[int64][]dto.MappingDto
	var d dto.MappingDto
	childrenMap = make(map[int64][]dto.MappingDto)
	for _, d = range flatDto {
		if d.From != nil {
			childrenMap[*d.From] = append(childrenMap[*d.From], d)
		}
	}
	var roots []dto.MappingDto
	roots = []dto.MappingDto{}
	for _, d = range flatDto {
		if d.From == nil {
			buildMappingTree(&d, childrenMap)
			roots = append(roots, d)
		}
	}
	sort.Slice(roots, func(i, j int) bool {
		return roots[j].ID > roots[i].ID
	})
	node.Mappings = roots
}
