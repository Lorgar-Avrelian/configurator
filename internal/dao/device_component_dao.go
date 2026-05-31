package dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"filler/internal/database"
	"filler/internal/dto"
	"filler/internal/model"

	"github.com/jackc/pgx/v5"
)

// Вспомогательный метод выгрузки плоского списка и сборки в дерево
func buildTreeFromRows(rows pgx.Rows) ([]model.DeviceComponent, error) {
	type rawNode struct {
		dc       model.DeviceComponent
		parentID sql.NullInt64
	}
	allNodes := make(map[int64]*rawNode)
	adjList := make(map[int64][]int64)
	for rows.Next() {
		rn := &rawNode{}
		var jsonBytes []byte
		err := rows.Scan(&rn.dc.ID, &rn.dc.ModelID, &rn.dc.InternalOrder, &rn.parentID, &jsonBytes)
		if err != nil {
			return nil, err
		}
		if rn.parentID.Valid {
			rn.dc.ParentID = &rn.parentID.Int64
			adjList[rn.parentID.Int64] = append(adjList[rn.parentID.Int64], rn.dc.ID)
		}
		rn.dc.Mappings = []model.Mapping{}
		rn.dc.Components = []model.DeviceComponent{}
		if len(jsonBytes) > 0 && string(jsonBytes) != "[null]" {
			_ = json.Unmarshal(jsonBytes, &rn.dc.Mappings)
		}
		allNodes[rn.dc.ID] = rn
	}
	var buildNode func(id int64) model.DeviceComponent
	buildNode = func(id int64) model.DeviceComponent {
		current := allNodes[id]
		childIDs := adjList[id]
		for _, childID := range childIDs {
			childNode := buildNode(childID)
			current.dc.Components = append(current.dc.Components, childNode)
		}

		return current.dc
	}
	var roots []model.DeviceComponent
	for id, rn := range allNodes {
		isRoot := !rn.parentID.Valid
		if !isRoot {
			_, parentExists := allNodes[rn.parentID.Int64]
			isRoot = !parentExists
		}
		if isRoot {
			roots = append(roots, buildNode(id))
		}
	}
	if len(allNodes) > 0 && len(roots) == 0 {
		for id := range allNodes {
			roots = append(roots, buildNode(id))
			break
		}
	}
	return roots, nil
}

func CreateDeviceComponent(ctx context.Context, d dto.DeviceComponentCreate) (*model.DeviceComponent, error) {
	conn := database.Get()
	query := `
		INSERT INTO public.device_component (model, internal_order, parent)
		VALUES ($1, $2, $3)
		RETURNING id, model, internal_order, parent`
	var parent sql.NullInt64
	if d.ParentID != nil {
		parent.Int64 = *d.ParentID
		parent.Valid = true
	}
	var dc model.DeviceComponent
	err := conn.QueryRow(ctx, query, d.ModelID, d.InternalOrder, parent).
		Scan(&dc.ID, &dc.ModelID, &dc.InternalOrder, &parent)
	if err != nil {
		return nil, err
	}

	if parent.Valid {
		dc.ParentID = &parent.Int64
	}
	dc.Mappings = []model.Mapping{}
	dc.Components = []model.DeviceComponent{}
	return &dc, nil
}

// GetDeviceComponentByID рекурсивно собирает всю ветку компонента вниз по иерархии со всеми маппингами для каждого уровня
func GetDeviceComponentByID(ctx context.Context, id int64) (*model.DeviceComponent, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE device_tree AS (
			SELECT id, model, internal_order, parent FROM public.device_component WHERE id = $1
			UNION ALL
			SELECT c.id, c.model, c.internal_order, c.parent
			FROM public.device_component c
			JOIN device_tree dt ON c.parent = dt.id
		)
		SELECT dt.id, dt.model, dt.internal_order, dt.parent,
		       COALESCE(json_strip_nulls(json_agg(json_build_object(
				    'id', m.id,
				    'frequency', pf.value,
				    'coefficient', m.coefficient::text,
				    'enum', m.enum,
				    'param', json_build_object(
					    'id', p.id,
					    'title', p.title,
					    'name_en', p.name_en,
					    'name_ru', p.name_ru,
					    'type', p_vt.value,
					    'value', p.value,
					    'description_en', p.description_en,
					    'description_ru', p.description_ru,
					    'units_en', p.units_en,
					    'units_ru', p.units_ru,
					    'access', p_ac.value,
					    'saved', p.saved,
					    'visible', p.visible
				    ),
				    'indicator', json_build_object(
					    'id', pi.id,
					    'oid_id', pi.oid_id,
					    'dotter_notation', pi.dotter_notation,
					    'oid', json_build_object(
						    'id', o.id,
						    'mib_id', o.mib,
						    'type', o_at.value,
						    'name', o.name,
						    'number', o.number,
						    'dotter_notation', o.dotter_notation,
						    'object_descriptor', o.object_descriptor,
						    'syntax', o.syntax,
						    'enum', o.enum,
						    'status', o_st.value,
						    'access', o_oac.value,
						    'units', o.units,
						    'description', o.description,
						    'category', o.category
					    )
				    )
				)) FILTER (WHERE m.id IS NOT NULL)), '[]'::json) AS mappings_json
		FROM device_tree dt
		LEFT JOIN public.device_component_mapping dcm ON dt.id = dcm.device_component_id
		LEFT JOIN public.mapping m ON dcm.mapping_id = m.id
		LEFT JOIN public.polling_frequency pf ON m.frequency = pf.id
		LEFT JOIN public.param p ON m.param = p.id
		LEFT JOIN public.var_type p_vt ON p.type = p_vt.id
		LEFT JOIN public.access p_ac ON p.access = p_ac.id
		LEFT JOIN public.param_indicator pi ON m.indicator = pi.id
		LEFT JOIN public.oid o ON pi.oid_id = o.id
		LEFT JOIN public.asn1_type o_at ON o.type = o_at.id
		LEFT JOIN public.oid_status o_st ON o.status = o_st.id
		LEFT JOIN public.oid_access o_oac ON o.access = o_oac.id
		GROUP BY dt.id, dt.model, dt.internal_order, dt.parent`
	rows, err := conn.Query(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	tree, err := buildTreeFromRows(rows)
	if err != nil {
		return nil, err
	}
	if len(tree) == 0 {
		return nil, nil
	}
	for _, node := range tree {
		if node.ID == id {
			return &node, nil
		}
	}
	return &tree[0], nil
}

// GetAllDeviceComponents возвращает полный массив иерархических деревьев устройств со всеми маппингами
func GetAllDeviceComponents(ctx context.Context) ([]model.DeviceComponent, error) {
	conn := database.Get()
	query := `
		WITH RECURSIVE device_tree AS (
			SELECT id, model, internal_order, parent FROM public.device_component WHERE parent IS NULL
			UNION ALL
			SELECT c.id, c.model, c.internal_order, c.parent
			FROM public.device_component c
			JOIN device_tree dt ON c.parent = dt.id
		)
		SELECT dt.id, dt.model, dt.internal_order, dt.parent,
		       COALESCE(json_strip_nulls(json_agg(json_build_object(
				    'id', m.id,
				    'frequency', pf.value,
				    'coefficient', m.coefficient::text,
				    'enum', m.enum,
				    'param', json_build_object(
					    'id', p.id,
					    'title', p.title,
					    'name_en', p.name_en,
					    'name_ru', p.name_ru,
					    'type', p_vt.value,
					    'value', p.value,
					    'description_en', p.description_en,
					    'description_ru', p.description_ru,
					    'units_en', p.units_en,
					    'units_ru', p.units_ru,
					    'access', p_ac.value,
					    'saved', p.saved,
					    'visible', p.visible
				    ),
				    'indicator', json_build_object(
					    'id', pi.id,
					    'oid_id', pi.oid_id,
					    'dotter_notation', pi.dotter_notation,
					    'oid', json_build_object(
						    'id', o.id,
						    'mib_id', o.mib,
						    'type', o_at.value,
						    'name', o.name,
						    'number', o.number,
						    'dotter_notation', o.dotter_notation,
						    'object_descriptor', o.object_descriptor,
						    'syntax', o.syntax,
						    'enum', o.enum,
						    'status', o_st.value,
						    'access', o_oac.value,
						    'units', o.units,
						    'description', o.description,
						    'category', o.category
					    )
				    )
				)) FILTER (WHERE m.id IS NOT NULL)), '[]'::json) AS mappings_json
		FROM device_tree dt
		LEFT JOIN public.device_component_mapping dcm ON dt.id = dcm.device_component_id
		LEFT JOIN public.mapping m ON dcm.mapping_id = m.id
		LEFT JOIN public.polling_frequency pf ON m.frequency = pf.id
		LEFT JOIN public.param p ON m.param = p.id
		LEFT JOIN public.var_type p_vt ON p.type = p_vt.id
		LEFT JOIN public.access p_ac ON p.access = p_ac.id
		LEFT JOIN public.param_indicator pi ON m.indicator = pi.id
		LEFT JOIN public.oid o ON pi.oid_id = o.id
		LEFT JOIN public.asn1_type o_at ON o.type = o_at.id
		LEFT JOIN public.oid_status o_st ON o.status = o_st.id
		LEFT JOIN public.oid_access o_oac ON o.access = o_oac.id
		GROUP BY dt.id, dt.model, dt.internal_order, dt.parent`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return buildTreeFromRows(rows)
}

func UpdateDeviceComponent(ctx context.Context, id int64, d dto.DeviceComponentUpdate) (*model.DeviceComponent, error) {
	conn := database.Get()
	query := `
		UPDATE public.device_component 
		SET model = $1, internal_order = $2, parent = $3
		WHERE id = $4
		RETURNING id, model, internal_order, parent`
	var parent sql.NullInt64
	if d.ParentID != nil {
		parent.Int64 = *d.ParentID
		parent.Valid = true
	}
	var dc model.DeviceComponent
	err := conn.QueryRow(ctx, query, d.ModelID, d.InternalOrder, parent, id).
		Scan(&dc.ID, &dc.ModelID, &dc.InternalOrder, &parent)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if parent.Valid {
		dc.ParentID = &parent.Int64
	}
	return &dc, nil
}

func DeleteDeviceComponent(ctx context.Context, id int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.device_component WHERE id = $1`
	commandTag, err := conn.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}

func BindDeviceMapping(ctx context.Context, d dto.BindDeviceMappingRequest) error {
	conn := database.Get()
	query := `INSERT INTO public.device_component_mapping (device_component_id, mapping_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := conn.Exec(ctx, query, d.DeviceComponentID, d.MappingID)
	return err
}

func UnbindDeviceMapping(ctx context.Context, dcID, mID int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.device_component_mapping WHERE device_component_id = $1 AND mapping_id = $2`
	commandTag, err := conn.Exec(ctx, query, dcID, mID)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}
