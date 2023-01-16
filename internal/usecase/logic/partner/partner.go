package partner

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	PartnerListUseCase struct {
		RepoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput]
	}
	PartnerGetFirstUseCase struct {
		UC1 usecase.IPartnerListUseCase
	}
)

func (s *PartnerGetFirstUseCase) GetFirst(ctx context.Context, o *model.PartnerOrderInput, w *model.PartnerWhereInput) (*model.Partner, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *PartnerListUseCase) List(ctx context.Context, limit, offset *int, o *model.PartnerOrderInput, w *model.PartnerWhereInput) ([]*model.Partner, error) {
	entities, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.partner.partner.PartnerListUseCase.List: %s", err))
	}
	return entities, nil
}
