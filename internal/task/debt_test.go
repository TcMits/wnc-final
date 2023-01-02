package task_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	gomock "github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestDebtCreateTaskExecutor(t *testing.T) {
	t.Parallel()
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	b := sse.NewMockINotify(mockCtl)
	b.EXPECT().Notify(gomock.Any()).Return(nil).AnyTimes()
	pl := &task.DebtNotifyPayload{
		UserID: uuid.New(),
	}
	l := logger.New(logger.DebugLevel)
	handler := task.GetDebtTaskExecutor(b, l)
	err := handler.ExecuteTask(context.Background(), pl)
	require.Nil(t, err)
}
