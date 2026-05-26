package repository

import (
	"context"

	internal "github.com/gui-marc/concurra/repository/internal/generated/pgx"
)

type ComponentType string

const (
	ComponentTypeTenant     ComponentType = "tenant"
	ComponentTypeThirdParty ComponentType = "third-party"
)

type Component struct {
	ID   string
	Name string
	Team string
	Area string
	Type ComponentType
}

type CreateComponentParams struct {
	Name string
	Team string
	Area string
	Type ComponentType
}

func componentFromInternal(component internal.Component) Component {
	return Component{
		ID:   component.ID.String(),
		Name: component.Name,
		Team: component.Team,
		Area: component.Area,
		Type: ComponentType(component.Type),
	}
}

type ComponentRepository interface {
	GetAllComponents(context context.Context, limit, offset int) ([]Component, error)
	GetComponentByID(ctx context.Context, id string) (Component, error)
	CreateComponent(ctx context.Context, params CreateComponentParams) (Component, error)
	GetComponentByNameTeamArea(ctx context.Context, team, area, name string) (Component, error)
}

type pgxComponentRepository struct {
	queries *internal.Queries
}

func NewPGXComponentRepository(queries *internal.Queries) ComponentRepository {
	return &pgxComponentRepository{
		queries: queries,
	}
}

func (r *pgxComponentRepository) GetAllComponents(ctx context.Context, limit, offset int) ([]Component, error) {
	components, err := r.queries.GetComponents(ctx, internal.GetComponentsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	result := make([]Component, len(components))
	for i, component := range components {
		result[i] = componentFromInternal(component)
	}
	return result, nil
}

func (r *pgxComponentRepository) GetComponentByID(ctx context.Context, id string) (Component, error) {
	component, err := r.queries.GetComponentByID(ctx, toPgUUID(id))
	if err != nil {
		return Component{}, err
	}
	return componentFromInternal(component), nil
}

func (r *pgxComponentRepository) CreateComponent(ctx context.Context, params CreateComponentParams) (Component, error) {
	component, err := r.queries.CreateComponent(ctx, internal.CreateComponentParams{
		Name: params.Name,
		Team: params.Team,
		Area: params.Area,
		Type: internal.ComponentType(params.Type),
	})
	if err != nil {
		return Component{}, err
	}
	return componentFromInternal(component), nil
}

func (r *pgxComponentRepository) GetComponentByNameTeamArea(ctx context.Context, team, area, name string) (Component, error) {
	component, err := r.queries.GetComponentByNameTeamArea(ctx, internal.GetComponentByNameTeamAreaParams{
		Team: team,
		Area: area,
		Name: name,
	})
	if err != nil {
		return Component{}, err
	}
	return componentFromInternal(component), nil
}
