package service

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"context"
)

func BindParam(ctx context.Context, d dto.ComponentParamLinkDto) error {
	var compID int64
	var paramID int64
	var err error
	compID = d.ComponentID
	paramID = d.ParamID
	err = dao.BindParam(ctx, compID, paramID)
	return err
}

func UnbindParam(ctx context.Context, componentID int64, paramID int64) (bool, error) {
	var found bool
	var err error
	found, err = dao.UnbindParam(ctx, componentID, paramID)
	return found, err
}
