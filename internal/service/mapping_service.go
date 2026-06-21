package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/mapper"
	"context"
	"sort"
)

func CreateMapping(ctx context.Context, d dto.MappingCreateDto) (*dto.MappingDto, error) {
	var daoInput dao.MappingDao
	var insertedID int64
	var err error
	var res *dto.MappingDto
	daoInput = mapper.MappingCreateDtoToMappingDao(d)
	insertedID, err = dao.CreateMappingRaw(ctx, daoInput)
	if err != nil {
		return nil, err
	}
	res, err = GetMappingByID(ctx, insertedID)
	return res, err
}

func GetMappingByID(ctx context.Context, id int64) (*dto.MappingDto, error) {
	var arr []dao.Mapping
	var err error
	arr, err = dao.GetMappingByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(arr) == 0 {
		return nil, nil
	}
	var flatDto []dto.MappingDto
	var item dao.Mapping
	flatDto = []dto.MappingDto{}
	for _, item = range arr {
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
	var targetRoot dto.MappingDto
	var found bool
	found = false
	for _, d = range flatDto {
		if d.From == nil {
			targetRoot = d
			found = true
			break
		}
	}
	if !found {
		return nil, nil
	}
	buildMappingTree(&targetRoot, childrenMap)
	return &targetRoot, nil
}

func GetAllMappings(ctx context.Context) ([]dto.MappingDto, error) {
	var arr []dao.Mapping
	var err error
	arr, err = dao.GetAllMappings(ctx)
	if err != nil {
		return nil, err
	}
	var flatDto []dto.MappingDto
	var item dao.Mapping
	flatDto = []dto.MappingDto{}
	for _, item = range arr {
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
	return roots, nil
}

func UpdateMapping(ctx context.Context, id int64, d dto.MappingCreateDto) (*dto.MappingDto, error) {
	var daoInput dao.MappingDao
	var err error
	var res *dto.MappingDto
	daoInput = mapper.MappingCreateDtoToMappingDao(d)
	err = dao.UpdateMappingRaw(ctx, id, daoInput)
	if err != nil {
		return nil, err
	}
	res, err = GetMappingByID(ctx, id)
	return res, err
}

func DeleteMapping(ctx context.Context, id int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.DeleteMapping(ctx, id)
	return found, err
}

func buildMappingTree(parent *dto.MappingDto, childrenMap map[int64][]dto.MappingDto) {
	var directChildren []dto.MappingDto
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
	var child dto.MappingDto
	for i = 0; len(directChildren) > i; i++ {
		child = directChildren[i]
		buildMappingTree(&child, childrenMap)
		parent.Children = append(parent.Children, child)
	}
}

func GetMappingByIDOwn(ctx context.Context, id int64) (*dto.MappingDto, error) {
	var resPtr *dao.Mapping
	var err error
	var res dto.MappingDto
	resPtr, err = dao.GetMappingByIDOwn(ctx, id)
	if err != nil {
		return nil, err
	}
	if resPtr == nil {
		return nil, nil
	}
	res = mapper.MappingToMappingDto(*resPtr)
	return &res, nil
}
