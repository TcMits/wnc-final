package template_test

import (
	"context"
	"testing"

	"github.com/TcMits/wnc-final/pkg/tool/template"
	"github.com/stretchr/testify/require"
)

func TestRenderToStr(t *testing.T) {
	file := "template_test.html"
	res, err := template.RenderToStr(file, map[string]string{
		"text": "foo",
	}, context.Background())
	require.Nil(t, err)
	require.NotNil(t, res)
}
