package mapper

import (
	"configurator/internal/dao"
	"configurator/internal/dto"
	"configurator/internal/logger"
	"configurator/internal/model"
	"configurator/internal/util"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	res.PollingProtocol = t.PollingProtocol
	res.Host = t.Host
	res.Port = t.Port
	res.Target = parseTargetNode(t.Target)
	return &res
}

func pStr(s string) *string {
	if s == "*" || s == "ANY" || s == "" {
		return nil
	}
	if s == "#" {
		var empty string
		empty = ""
		return &empty
	}
	return &s
}

func pInt(s string) *int32 {
	var clean string
	clean = s
	if strings.HasPrefix(clean, "[") && strings.HasSuffix(clean, "]") {
		clean = clean[1 : len(clean)-1]
	}
	if clean == "*" || clean == "ANY" {
		return nil
	}
	if clean == "#" {
		var zero int32
		zero = 0
		return &zero
	}
	var val int64
	var err error
	val, err = strconv.ParseInt(clean, 10, 32)
	if err != nil {
		return nil
	}
	var v32 int32
	v32 = int32(val)
	return &v32
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

func parseTargetDto(s string) (*dto.ThresholdTargetDto, error) {
	var parts []string
	parts = strings.SplitN(s, ".", 2)
	if len(parts) < 2 {
		return nil, errors.New("invalid target string format missing elements path")
	}
	var connection string
	connection = parts[0]
	var path string
	path = parts[1]
	var connParts []string
	connParts = strings.Split(connection, ":")
	if len(connParts) < 3 {
		return nil, errors.New("invalid definition of connection prefix protocol host port")
	}
	var res dto.ThresholdTargetDto
	if connParts[0] != "*" && connParts[0] != "ANY" {
		if connParts[0] == "#" {
			var empty string
			empty = ""
			res.PollingProtocol = &empty
		} else {
			res.PollingProtocol = &connParts[0]
		}
	}
	if connParts[1] != "*" && connParts[1] != "ANY" {
		if connParts[1] == "#" {
			var empty string
			empty = ""
			res.Host = &empty
		} else {
			res.Host = &connParts[1]
		}
	}
	if connParts[2] != "*" && connParts[2] != "ANY" {
		if connParts[2] == "#" {
			var zero int32
			zero = 0
			res.Port = &zero
		} else {
			var pVal int64
			var err error
			pVal, err = strconv.ParseInt(connParts[2], 10, 32)
			if err == nil {
				var p32 int32
				p32 = int32(pVal)
				res.Port = &p32
			}
		}
	}
	var nodeParts []string
	nodeParts = strings.Split(path, ".")
	res.Target = buildTargetNodes(nodeParts, 0)
	return &res, nil
}

func buildTargetNodes(parts []string, idx int) *dto.ThresholdTargetNodeDto {
	if idx >= len(parts) {
		return nil
	}
	var curr string
	curr = parts[idx]
	var res dto.ThresholdTargetNodeDto
	var nextIdx int
	nextIdx = idx + 1
	if curr == "*" || curr == "ANY" || curr == "#" || (strings.HasPrefix(curr, "[") && strings.HasSuffix(curr, "]")) {
		return nil
	}
	if len(parts) > nextIdx {
		var nextPart string
		nextPart = parts[nextIdx]
		if nextPart == "*" || nextPart == "ANY" || nextPart == "#" || (strings.HasPrefix(nextPart, "[") && strings.HasSuffix(nextPart, "]")) {
			var order *int32
			order = pInt(nextPart)
			res.InternalOrder = order
			nextIdx++
		}
	}
	if strings.HasSuffix(curr, "_level") || curr == "percentage" || curr == "value" {
		res.Field = pStr(curr)
	} else if curr == "voltage" || curr == "amperage" || curr == "charge_level" || curr == "load_percent" {
		res.Param = pStr(curr)
	} else {
		res.Component = pStr(curr)
	}
	res.Next = buildTargetNodes(parts, nextIdx)
	return &res
}

func tokenizeExpression(s string) []string {
	var cleanStr string
	cleanStr = strings.TrimSpace(s)
	cleanStr = strings.TrimPrefix(cleanStr, "\ufeff")
	var prepared strings.Builder
	var inQuotes bool
	var i int
	var rawRunes []rune
	rawRunes = []rune(cleanStr)
	for i = 0; len(rawRunes) > i; i++ {
		var r rune
		r = rawRunes[i]
		if r == '"' || r == '\'' {
			inQuotes = !inQuotes
			prepared.WriteRune(r)
			continue
		}
		if inQuotes {
			prepared.WriteRune(r)
			continue
		}
		if r == '(' || r == ')' {
			prepared.WriteString(" ")
			prepared.WriteRune(r)
			prepared.WriteString(" ")
			continue
		}
		if r == '=' || r == '>' || r == '<' || r == '!' {
			if len(rawRunes) > i+1 && rawRunes[i+1] == '=' {
				prepared.WriteString(" ")
				prepared.WriteRune(r)
				prepared.WriteRune(rawRunes[i+1])
				prepared.WriteString(" ")
				i++
			} else if len(rawRunes) > i+1 && r == '<' && rawRunes[i+1] == '>' {
				prepared.WriteString(" ")
				prepared.WriteRune(r)
				prepared.WriteRune(rawRunes[i+1])
				prepared.WriteString(" ")
				i++
			} else {
				prepared.WriteString(" ")
				prepared.WriteRune(r)
				prepared.WriteString(" ")
			}
			continue
		}
		prepared.WriteRune(r)
	}
	var words []string
	var fields []string
	fields = strings.Fields(prepared.String())
	var currentWord strings.Builder
	inQuotes = false
	for i = 0; len(fields) > i; i++ {
		var f string
		f = fields[i]
		if (strings.HasPrefix(f, "\"") || strings.HasPrefix(f, "'")) && !inQuotes {
			if (strings.HasSuffix(f, "\"") || strings.HasSuffix(f, "'")) && len(f) > 1 {
				words = append(words, f)
				continue
			}
			inQuotes = true
			currentWord.WriteString(f)
			continue
		}
		if inQuotes {
			currentWord.WriteString(" ")
			currentWord.WriteString(f)
			if strings.HasSuffix(f, "\"") || strings.HasSuffix(f, "'") {
				inQuotes = false
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
			continue
		}
		words = append(words, f)
	}
	var mergedWords []string
	for i = 0; len(words) > i; i++ {
		if len(words) > i+1 {
			var combined string
			combined = words[i] + " " + words[i+1]
			var op model.LogicOperator
			op = model.ParseLogicOperator(combined)
			if op.Data() != nil && strings.ToUpper(op.String()) == strings.ToUpper(combined) && op.Data().Type == "COMPARISON" {
				mergedWords = append(mergedWords, combined)
				i++
				continue
			}
		}
		mergedWords = append(mergedWords, words[i])
	}
	return mergedWords
}

func parseExpressionWords(words []string) (*dto.ThresholdNodeDto, error) {
	return parseExpressionSubRange(words, 0, len(words))
}

func parseExpressionSubRange(words []string, start int, end int) (*dto.ThresholdNodeDto, error) {
	var outputQueue []string
	var operatorStack []string
	var subExprs map[int]*dto.ThresholdNodeDto
	subExprs = make(map[int]*dto.ThresholdNodeDto)
	var i int
	for i = start; end > i; i++ {
		var w string
		w = words[i]
		if w == "(" {
			var depth int
			depth = 1
			var j int
			for j = i + 1; end > j; j++ {
				if words[j] == "(" {
					depth++
				} else if words[j] == ")" {
					depth--
				}
				if 0 == depth {
					var subNode *dto.ThresholdNodeDto
					var err error
					subNode, err = parseExpressionSubRange(words, i+1, j)
					if err != nil {
						return nil, err
					}
					var placeholder string
					placeholder = fmt.Sprintf("__SUB_EXPR_%d__", len(subExprs))
					subExprs[len(subExprs)] = subNode
					outputQueue = append(outputQueue, placeholder)
					i = j
					break
				}
			}
			if 0 < depth {
				return nil, errors.New("mismatched bracket parentheses found in query string")
			}
			continue
		}
		var op model.LogicOperator
		op = model.ParseLogicOperator(w)
		if op.Data() != nil && strings.ToUpper(op.String()) == strings.ToUpper(w) {
			var tData *model.LogicOperatorData
			tData = op.Data()
			for len(operatorStack) > 0 {
				var top string
				top = operatorStack[len(operatorStack)-1]
				if top == "(" {
					break
				}
				var topOp model.LogicOperator
				topOp = model.ParseLogicOperator(top)
				if topOp.Data() == nil {
					break
				}
				var topData *model.LogicOperatorData
				topData = topOp.Data()
				if topData.Precedence >= tData.Precedence {
					outputQueue = append(outputQueue, top)
					operatorStack = operatorStack[:len(operatorStack)-1]
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, w)
			continue
		}
		outputQueue = append(outputQueue, w)
	}
	for len(operatorStack) > 0 {
		var top string
		top = operatorStack[len(operatorStack)-1]
		outputQueue = append(outputQueue, top)
		operatorStack = operatorStack[:len(operatorStack)-1]
	}
	return buildASTFromQueueWithSub(outputQueue, subExprs)
}

func buildASTFromQueueWithSub(queue []string, subExprs map[int]*dto.ThresholdNodeDto) (*dto.ThresholdNodeDto, error) {
	var stack []*dto.ThresholdNodeDto
	var i int
	for i = 0; len(queue) > i; i++ {
		var val string
		val = queue[i]
		if strings.HasPrefix(val, "__SUB_EXPR_") && strings.HasSuffix(val, "__") {
			var idxStr string
			idxStr = val[11 : len(val)-2]
			var idx64 int64
			var err error
			idx64, err = strconv.ParseInt(idxStr, 10, 32)
			if err != nil {
				return nil, err
			}
			var subNode *dto.ThresholdNodeDto
			subNode = subExprs[int(idx64)]
			var n dto.ThresholdNodeDto
			var elem dto.ThresholdElementDto
			var expr dto.ThresholdExpressionDto
			expr.Root = subNode
			elem.Expression = &expr
			n.Element = &elem
			stack = append(stack, &n)
			continue
		}
		var op model.LogicOperator
		op = model.ParseLogicOperator(val)
		if op.Data() == nil || strings.ToUpper(op.String()) != strings.ToUpper(val) {
			var n dto.ThresholdNodeDto
			var elem dto.ThresholdElementDto
			var comp dto.ThresholdComparisonDto
			comp.Value = val
			elem.Comparison = &comp
			n.Element = &elem
			stack = append(stack, &n)
			continue
		}
		var tData *model.LogicOperatorData
		tData = op.Data()
		if tData.Type == "COMPARISON" {
			if len(stack) < 2 {
				return nil, errors.New("insufficient operand structures on stack layout for target comparison")
			}
			var right *dto.ThresholdNodeDto
			var left *dto.ThresholdNodeDto
			right = stack[len(stack)-1]
			left = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var targetObj *dto.ThresholdTargetDto
			var err error
			targetObj, err = parseTargetDto(left.Element.Comparison.Value)
			if err != nil {
				return nil, err
			}
			var resNode dto.ThresholdNodeDto
			var elem dto.ThresholdElementDto
			var comp dto.ThresholdComparisonDto
			comp.Target = targetObj
			comp.Operator = op.String()
			var cleanVal string
			cleanVal = strings.Trim(right.Element.Comparison.Value, "\"'")
			comp.Value = cleanVal
			elem.Comparison = &comp
			resNode.Element = &elem
			stack = append(stack, &resNode)
		} else if tData.Type == "LOGICAL" {
			if len(stack) < 2 {
				return nil, errors.New("insufficient condition logical subtrees on stack for logical operation")
			}
			var right *dto.ThresholdNodeDto
			var left *dto.ThresholdNodeDto
			right = stack[len(stack)-1]
			left = stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var rootLeft *dto.ThresholdNodeDto
			rootLeft = left
			for rootLeft.Next != nil {
				rootLeft = rootLeft.Next
			}
			var opStr string
			opStr = op.String()
			rootLeft.Operator = &opStr
			rootLeft.Next = right
			stack = append(stack, left)
		}
	}
	if len(stack) != 1 {
		return nil, errors.New("invalid syntax tree resolution layout stack parameters match error")
	}
	return stack[0], nil
}

func StringToThresholdCreateDto(s string, name string, desc *string, author string) (dto.ThresholdCreateDto, error) {
	var res dto.ThresholdCreateDto
	var words []string
	words = tokenizeExpression(s)
	if len(words) == 0 || strings.ToUpper(words[0]) != "IF" {
		return res, errors.New("input logical expression must start with valid IF condition")
	}
	var thenIdx int
	thenIdx = -1
	var i int
	for i = 0; len(words) > i; i++ {
		if strings.ToUpper(words[i]) == "THEN" {
			thenIdx = i
			break
		}
	}
	if thenIdx == -1 {
		return res, errors.New("missing mandatory THEN action separator statement")
	}
	var queryWords []string
	queryWords = words[1:thenIdx]
	var actionWords []string
	actionWords = words[thenIdx+1:]
	if len(actionWords) != 3 {
		return res, errors.New("invalid action statement layout structure definition length after THEN")
	}
	var op model.LogicOperator
	op = model.ParseLogicOperator(actionWords[1])
	if op.Data() == nil || op.Data().Type != "COMPARISON" {
		return res, errors.New("invalid operation operator inside action THEN block segment definition")
	}
	var rootNode *dto.ThresholdNodeDto
	var err error
	rootNode, err = parseExpressionWords(queryWords)
	if err != nil {
		return res, err
	}
	var actionTarget *dto.ThresholdTargetDto
	actionTarget, err = parseTargetDto(actionWords[0])
	if err != nil {
		return res, err
	}
	res.Name = name
	res.Description = desc
	res.Author = author
	res.Query.Root = rootNode
	res.Target = *actionTarget
	var cleanValue string
	cleanValue = strings.Trim(actionWords[2], "\"'")
	res.Value = cleanValue
	return res, nil
}

func renderTarget(t *dto.ThresholdTargetDto) string {
	if t == nil {
		return "*:*:*"
	}
	var proto string
	proto = "*"
	if t.PollingProtocol != nil {
		if *t.PollingProtocol == "" {
			proto = "#"
		} else {
			proto = *t.PollingProtocol
		}
	}
	var host string
	host = "*"
	if t.Host != nil {
		if *t.Host == "" {
			host = "#"
		} else {
			host = *t.Host
		}
	}
	var port string
	port = "*"
	if t.Port != nil {
		if *t.Port == 0 {
			port = "#"
		} else {
			port = fmt.Sprintf("%d", *t.Port)
		}
	}
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s:%s:%s", proto, host, port))
	var curr *dto.ThresholdTargetNodeDto
	curr = t.Target
	for curr != nil {
		sb.WriteString(".")
		if curr.Component != nil && *curr.Component != "" {
			sb.WriteString(*curr.Component)
			sb.WriteString(".")
			if curr.InternalOrder == nil {
				sb.WriteString("*")
			} else if *curr.InternalOrder == 0 {
				sb.WriteString("#")
			} else {
				sb.WriteString(fmt.Sprintf("[%d]", *curr.InternalOrder))
			}
		} else if curr.Param != nil {
			sb.WriteString(*curr.Param)
		} else if curr.Field != nil {
			sb.WriteString(*curr.Field)
		} else {
			sb.WriteString("*")
		}
		curr = curr.Next
	}
	return sb.String()
}

func renderExpressionNode(n *dto.ThresholdNodeDto) string {
	if n == nil {
		return ""
	}
	var sb strings.Builder
	if n.Element != nil {
		if n.Element.Comparison != nil {
			var comp *dto.ThresholdComparisonDto
			comp = n.Element.Comparison
			var cleanVal string
			cleanVal = comp.Value
			var _, numErr = strconv.ParseFloat(cleanVal, 64)
			if numErr != nil && cleanVal != "true" && cleanVal != "false" && cleanVal != "null" {
				if !strings.HasPrefix(cleanVal, "\"") && !strings.HasSuffix(cleanVal, "\"") {
					cleanVal = fmt.Sprintf("\"%s\"", cleanVal)
				}
			}
			sb.WriteString(fmt.Sprintf("%s %s %s", renderTarget(comp.Target), comp.Operator, cleanVal))
		} else if n.Element.Expression != nil && n.Element.Expression.Root != nil {
			sb.WriteString("( ")
			sb.WriteString(renderExpressionNode(n.Element.Expression.Root))
			sb.WriteString(" )")
		}
	}
	if n.Operator != nil && n.Next != nil {
		sb.WriteString(fmt.Sprintf(" %s ", *n.Operator))
		sb.WriteString(renderExpressionNode(n.Next))
	}
	return sb.String()
}

func ThresholdDtoToString(d dto.ThresholdDto) string {
	var sb strings.Builder
	sb.WriteString("IF ")
	sb.WriteString(renderExpressionNode(d.Query.Root))
	sb.WriteString(" THEN ")
	sb.WriteString(renderTarget(&d.Target))
	sb.WriteString(" = ")
	var finalVal string
	finalVal = d.Value
	var _, numErr = strconv.ParseFloat(finalVal, 64)
	if numErr != nil && finalVal != "true" && finalVal != "false" && finalVal != "null" {
		if !strings.HasPrefix(finalVal, "\"") && !strings.HasSuffix(finalVal, "\"") {
			finalVal = fmt.Sprintf("\"%s\"", finalVal)
		}
	}
	sb.WriteString(finalVal)
	return sb.String()
}

func ThresholdDtoToThresholdStringDto(d dto.ThresholdDto) dto.ThresholdStringDto {
	var res dto.ThresholdStringDto
	res.ID = d.ID
	res.Name = d.Name
	res.Description = d.Description
	res.Author = d.Author
	res.Expression = ThresholdDtoToString(d)
	res.Created = d.Created
	return res
}
