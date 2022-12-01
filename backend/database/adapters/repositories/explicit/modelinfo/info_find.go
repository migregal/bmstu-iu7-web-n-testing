package modelinfo

import (
	"fmt"
	"neural_storage/cube/core/entities/model"
	"neural_storage/cube/core/entities/structure"
	"neural_storage/cube/core/ports/repositories"
	dbmodel "neural_storage/database/core/entities/model"
)

func (r *Repository) Find(filter repositories.ModelInfoFilter) ([]*model.Info, int64, error) {
	query := r.db.DB
	if len(filter.IDs) > 0 {
		query = query.Where("id in ?", filter.IDs)
	}
	if len(filter.Owners) > 0 {
		query = query.Where("owner_id in ?", filter.Owners)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset * filter.Limit)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}

	var modelInfo []dbmodel.Model
	if err := query.Find(&modelInfo).Error; err != nil {
		return nil, 0, fmt.Errorf("model get error: %w", err)
	}

	var total int64
	r.db.Table(dbmodel.Model{}.TableName()).Count(&total)

	dbInfo := []*model.Info{}
	if len(filter.IDs) == 0 {
		for _, v := range modelInfo {
			st, err := r.getStructInfo(v.ID)
			if err != nil {
				return nil, 0, err
			}

			dbInfo = append(dbInfo, model.NewInfo(
				v.GetID(), v.GetOwnerID(),
				v.GetName(),
				structure.NewInfo(st.GetID(), st.GetName(), nil, nil, nil, nil)))
		}
		return dbInfo, total, nil
	}

	for _, v := range modelInfo {
		data, err := r.Get(v.GetID())
		if err != nil {
			return nil, 0, err
		}

		dbInfo = append(dbInfo, data)
	}

	return dbInfo, total, nil
}
