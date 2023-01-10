package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetPartnerListRepository(
	client *ent.Client,
) ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput] {
	return ent.NewPartnerReadRepository(client)
}
