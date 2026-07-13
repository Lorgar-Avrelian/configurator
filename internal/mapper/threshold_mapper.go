package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/util"
	"encoding/json"
	"errors"
	"time"
)

func validateTargetNode(n *dto.ThresholdTargetNodeDto) error {
	if n == nil {
		return nil
	}
	var count int
	count = 0
	if n.Component != nil && *n.Component != "" {
		count++
	}
	if n.Param != nil && *n.Param != "" {
		count++
	}
	if n.Field != nil && *n.Field != "" {
		count++
	}
	if count > 1 {
		return errors.New("only one of component, param, or field can be set per target node")
	}
	if 1 > count {
		return errors.New("at least one of component, param, or field must be set per target node")
	}
	return validateTargetNode(n.Next)
}

func parseTargetNode(n *dto.ThresholdTargetNodeDto) *dto.ThresholdTargetNodeDto {
	if n == nil {
		return nil
	}
	var res dto.ThresholdTargetNodeDto
	res.Component = n.Component
	res.InternalOrder = n.InternalOrder
	res.Param = n.Param
	res.Field = n.Field
	res.Next = parseTargetNode(n.Next)
	return &res
}

func parseTarget(t *dto.ThresholdTargetDto) *dto.ThresholdTargetDto {
	if t == nil {
		return nil
	}
	var res dto.ThresholdTargetDto
	if t.PollingProtocol != nil && *t.PollingProtocol != "" {
		res.PollingProtocol = t.PollingProtocol
	}
	if t.Host != nil && *t.Host != "" {
		res.Host = t.Host
	}
	if t.Port != nil && *t.Port != 0 {
		res.Port = t.Port
	}
	res.Target = parseTargetNode(t.Target)
	return &res
}

func parseElement(e *dto.ThresholdElementDto) *dto.ThresholdElementDto {
	if e == nil {
		return nil
	}
	var res dto.ThresholdElementDto
	if e.Comparison != nil {
		var comp dto.ThresholdComparisonDto
		comp.Operator = e.Comparison.Operator
		comp.Value = e.Comparison.Value
		comp.Target = parseTarget(e.Comparison.Target)
		res.Comparison = &comp
	}
	if e.Expression != nil {
		var expr dto.ThresholdExpressionDto
		expr.Root = parseExpressionNode(e.Expression.Root)
		res.Expression = &expr
	}
	return &res
}

func parseExpressionNode(n *dto.ThresholdNodeDto) *dto.ThresholdNodeDto {
	if n == nil {
		return nil
	}
	var res dto.ThresholdNodeDto
	res.Element = parseElement(n.Element)
	res.Next = parseExpressionNode(n.Next)
	if res.Next != nil {
		res.Operator = n.Operator
	} else {
		res.Operator = nil
	}
	return &res
}

func validateExpressionNode(n *dto.ThresholdNodeDto) error {
	if n == nil {
		return nil
	}
	if n.Element != nil && n.Element.Comparison != nil && n.Element.Comparison.Target != nil {
		var err error
		err = validateTargetNode(n.Element.Comparison.Target.Target)
		if err != nil {
			return err
		}
	}
	if n.Element != nil && n.Element.Expression != nil {
		var err error
		err = validateExpressionNode(n.Element.Expression.Root)
		if err != nil {
			return err
		}
	}
	return validateExpressionNode(n.Next)
}

func ThresholdCreateDtoToThresholdDao(d dto.ThresholdCreateDto) (dao.ThresholdDao, error) {
	var res dao.ThresholdDao
	var err error
	err = validateTargetNode(d.Target.Target)
	if err != nil {
		return res, err
	}
	err = validateExpressionNode(d.Query.Root)
	if err != nil {
		return res, err
	}
	var queryBytes []byte
	var targetBytes []byte
	var cleanQuery dto.ThresholdExpressionDto
	var cleanTargetObj *dto.ThresholdTargetDto
	cleanQuery.Root = parseExpressionNode(d.Query.Root)
	cleanTargetObj = parseTarget(&d.Target)
	queryBytes, err = json.Marshal(cleanQuery)
	if err != nil {
		logger.Errorf("Failed to marshal query expression to json: %v", err)
		queryBytes = []byte("{}")
	}
	targetBytes, err = json.Marshal(cleanTargetObj)
	if err != nil {
		logger.Errorf("Failed to marshal target to json: %v", err)
		targetBytes = []byte("{}")
	}
	res.Name = d.Name
	res.Description = util.StringPtrToSqlNullString(d.Description)
	res.Author = d.Author
	res.Query = queryBytes
	res.Target = targetBytes
	res.Value = d.Value
	return res, nil
}

func ThresholdDaoToThresholdDto(d dao.ThresholdDao) dto.ThresholdDto {
	var res dto.ThresholdDto
	var qExpr dto.ThresholdExpressionDto
	var tTarg dto.ThresholdTargetDto
	var err error
	err = json.Unmarshal(d.Query, &qExpr)
	if err != nil {
		logger.Errorf("Failed to unmarshal query expression from json: %v", err)
	}
	err = json.Unmarshal(d.Target, &tTarg)
	if err != nil {
		logger.Errorf("Failed to unmarshal target from json: %v", err)
	}
	res.ID = d.ID
	res.Name = d.Name
	res.Description = util.SqlNullStringToStringPtr(d.Description)
	res.Author = d.Author
	res.Query.Root = parseExpressionNode(qExpr.Root)
	var parsedTarg *dto.ThresholdTargetDto
	parsedTarg = parseTarget(&tTarg)
	if parsedTarg != nil {
		res.Target = *parsedTarg
	}
	res.Value = d.Value
	if d.Created.Valid {
		var timeStr string
		timeStr = d.Created.Time.Format(time.RFC3339)
		res.Created = &timeStr
	}
	return res
}

func ThresholdDaoArrayToThresholdDtoArray(arr []dao.ThresholdDao) []dto.ThresholdDto {
	var res []dto.ThresholdDto
	res = []dto.ThresholdDto{}
	var i int
	var item dao.ThresholdDao
	for i = 0; len(arr) > i; i++ {
		item = arr[i]
		res = append(res, ThresholdDaoToThresholdDto(item))
	}
	return res
}
