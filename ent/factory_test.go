package ent_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/datastore"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/stretchr/testify/require"
)

func TestPartnerFactory(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client)
	}{
		{
			name:  "no opt",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client) {
				ent.EmbedClient(&ctx, c)
				eAny, err := ent.PartnerFactory().CreateWithContext(ctx)
				e, ok := eAny.(*model.Partner)
				require.Nil(t, err)
				require.True(t, ok)
				require.NotEmpty(t, e.APIKey)
			},
		},
		{
			name:  "opt available",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client) {
				ent.EmbedClient(&ctx, c)
				eAny, err := ent.PartnerFactory(ent.Opt{"Name", generic.GetPointer("foo")}).CreateWithContext(ctx)
				e, ok := eAny.(*model.Partner)
				require.Nil(t, err)
				require.True(t, ok)
				require.Equal(t, e.Name, "foo")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			tt.setUp(t, ctx, c)
			tt.expect(t, ctx, c)
		})
	}
}

func TestMustTransactionFactory(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		setUp  func(*testing.T, context.Context, *ent.Client)
		expect func(*testing.T, context.Context, *ent.Client)
	}{
		{
			name:  "no opt",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client) {
				e, err := ent.MustTransactionFactory().CreateWithClient(ctx, c)
				require.Nil(t, err)
				require.NotEmpty(t, e.SenderID)
			},
		},
		{
			name:  "opt available",
			setUp: func(t *testing.T, ctx context.Context, c *ent.Client) {},
			expect: func(t *testing.T, ctx context.Context, c *ent.Client) {
				sender, _ := ent.MustBankAccountFactory(ent.Opt{"IsForPayment", generic.GetPointer(true)}).CreateWithClient(ctx, c)
				e, err := ent.MustTransactionFactory(ent.Opt{"SenderID", generic.GetPointer(sender.ID)}).CreateWithClient(ctx, c)
				require.Nil(t, err)
				require.Equal(t, *e.SenderID, sender.ID)
				require.Equal(t, e.SenderBankAccountNumber, sender.AccountNumber)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := datastore.NewClientTestConnection(t)
			defer c.Close()
			ctx := context.Background()
			tt.setUp(t, ctx, c)
			tt.expect(t, ctx, c)
		})
	}
}
