package dao

/*// Вспомогательная функция для маппинга Null-типов в указатели Go структуры
func mapRowToIndicator(id int64, desc, obj, cont, name, loc sql.NullString, serv sql.NullInt16) model.DeviceIndicator {
	ind := model.DeviceIndicator{ID: id}
	if desc.Valid {
		ind.Description = &desc.String
	}
	if obj.Valid {
		ind.ObjectID = &obj.String
	}
	if cont.Valid {
		ind.Contact = &cont.String
	}
	if name.Valid {
		ind.Name = &name.String
	}
	if loc.Valid {
		ind.Location = &loc.String
	}
	if serv.Valid {
		ind.Services = &serv.Int16
	}
	return ind
}

// Вспомогательная функция перевода указателей Go в Null-типы для SQL
func toNullString(s *string) sql.NullString {
	if s == nil {
		return sql.NullString{}
	}
	return sql.NullString{String: *s, Valid: true}
}

func toNullInt16(i *int16) sql.NullInt16 {
	if i == nil {
		return sql.NullInt16{}
	}
	return sql.NullInt16{Int16: *i, Valid: true}
}

func CreateIndicator(ctx context.Context, d dto.DeviceIndicatorCreate) (*model.DeviceIndicator, error) {
	conn := database.Get()
	query := `
		INSERT INTO public.device_indicator (description, object_id, contact, name, location, services)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, description, object_id, contact, name, location, services`
	var id int64
	var desc, obj, cont, name, loc sql.NullString
	var serv sql.NullInt16
	err := conn.QueryRow(ctx, query, toNullString(d.Description), toNullString(d.ObjectID), toNullString(d.Contact), toNullString(d.Name), toNullString(d.Location), toNullInt16(d.Services)).
		Scan(&id, &desc, &obj, &cont, &name, &loc, &serv)
	if err != nil {
		return nil, err
	}
	res := mapRowToIndicator(id, desc, obj, cont, name, loc, serv)
	return &res, nil
}

func GetIndicatorByID(ctx context.Context, id int64) (*model.DeviceIndicator, error) {
	conn := database.Get()
	query := `SELECT id, description, object_id, contact, name, location, services FROM public.device_indicator WHERE id = $1`
	var desc, obj, cont, name, loc sql.NullString
	var serv sql.NullInt16
	err := conn.QueryRow(ctx, query, id).Scan(&id, &desc, &obj, &cont, &name, &loc, &serv)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	res := mapRowToIndicator(id, desc, obj, cont, name, loc, serv)
	return &res, nil
}

func GetAllIndicators(ctx context.Context) ([]model.DeviceIndicator, error) {
	conn := database.Get()
	query := `SELECT id, description, object_id, contact, name, location, services FROM public.device_indicator`
	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var indicators []model.DeviceIndicator
	for rows.Next() {
		var id int64
		var desc, obj, cont, name, loc sql.NullString
		var serv sql.NullInt16

		if err := rows.Scan(&id, &desc, &obj, &cont, &name, &loc, &serv); err != nil {
			return nil, err
		}
		indicators = append(indicators, mapRowToIndicator(id, desc, obj, cont, name, loc, serv))
	}
	return indicators, nil
}

func UpdateIndicator(ctx context.Context, id int64, d dto.DeviceIndicatorUpdate) (*model.DeviceIndicator, error) {
	conn := database.Get()
	query := `
		UPDATE public.device_indicator
		SET description = $1, object_id = $2, contact = $3, name = $4, location = $5, services = $6
		WHERE id = $7
		RETURNING id, description, object_id, contact, name, location, services`
	var desc, obj, cont, name, loc sql.NullString
	var serv sql.NullInt16
	err := conn.QueryRow(ctx, query, toNullString(d.Description), toNullString(d.ObjectID), toNullString(d.Contact), toNullString(d.Name), toNullString(d.Location), toNullInt16(d.Services), id).
		Scan(&id, &desc, &obj, &cont, &name, &loc, &serv)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	res := mapRowToIndicator(id, desc, obj, cont, name, loc, serv)
	return &res, nil
}

func DeleteIndicator(ctx context.Context, id int64) (bool, error) {
	conn := database.Get()
	query := `DELETE FROM public.device_indicator WHERE id = $1`
	commandTag, err := conn.Exec(ctx, query, id)
	if err != nil {
		return false, err
	}
	return commandTag.RowsAffected() > 0, nil
}*/
