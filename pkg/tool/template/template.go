package template

import (
	"bytes"
	"context"
	"html/template"
	"strings"

	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func RenderFileToStr(file string, payload any, ctx context.Context) (*string, error) {
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

func RenderToStr(temp string, payload any, ctx context.Context) (*string, error) {
	t := template.Must(template.New("temp").Parse(temp))
	builder := &strings.Builder{}
	if err := t.Execute(builder, payload); err != nil {
		return nil, err
	}
	s := builder.String()
	return &s, nil
}
