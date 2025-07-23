package logging

import "text/template"

var GetOne = template.Must(template.New("get_one").Parse(`
{{ .Comment }} func (mw *loggingMiddleware) {{ .Method.GoName }}(ctx context.Context, in *{{ .Method.Input.GoIdent.GoName }}) (res *{{ .Method.Output.GoIdent.GoName }}, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"package", "{{ .Package }}",
			"method", "{{ .UnderscoreMethodname }}",
			"id", in.GetId(),
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.next.{{ .Method.GoName }}(ctx, in)
}
	`))
