package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/partner"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewPartnerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerGetFirstUseCase {
	uc := &partner.PartnerGetFirstUseCase{
		UC1: NewPartnerListUseCase(repoList),
	}
	return uc
}

func NewPartnerListUseCase(
	repoList repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
) usecase.IPartnerListUseCase {
	uc := &partner.PartnerListUseCase{
		RepoList: repoList,
	}
	return uc
}
