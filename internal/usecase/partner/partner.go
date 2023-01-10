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
		repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput]
	}
	PartnerGetFirstUseCase struct {
		pLUC usecase.IPartnerListUseCase
	}
)

func NewPartnerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerGetFirstUseCase {
	uc := &PartnerGetFirstUseCase{
		pLUC: repoList,
	}
	return uc
}

func NewPartnerListUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerListUseCase {
	uc := &PartnerListUseCase{
		repoList: repoList,
	}
	return uc
}
func (uc *PartnerGetFirstUseCase) GetFirst(ctx context.Context, o *model.PartnerOrderInput, w *model.PartnerWhereInput) (*model.Partner, error) {
	l, of := 1, 0
	entities, err := uc.pLUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (uc *PartnerListUseCase) List(ctx context.Context, limit, offset *int, o *model.PartnerOrderInput, w *model.PartnerWhereInput) ([]*model.Partner, error) {
	entities, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.partner.partner.PartnerListUseCase.List: %s", err))
	}
	return entities, nil
}
