package template

import (
	"bytes"
	"context"
	"html/template"

	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func RenderToStr(file string, payload any, ctx context.Context) (*string, error) {
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, payload); err != nil {
		return nil, err
	}
	return generic.GetPointer(buffer.String()), nil
}
