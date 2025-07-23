package server

import "text/template"

var GetOne = template.Must(template.New("logging_get_one").Parse(`

var re{{ .Method.GoName }}Slug = regexp.MustCompile("{.*}")


{{ .Comment }} func (h *httpClient) {{ .Method.GoName }}(ctx context.Context, in *{{ .Method.Input.GoIdent.GoName }}) (*{{ .Method.Output.GoIdent.GoName }}, error) {
	var response  {{ .Method.Output.GoIdent.GoName }}

	var urlBuilder strings.Builder
	urlBuilder.WriteString(h.reqURL)
	urlBuilder.WriteString(re{{ .Method.GoName }}Slug.ReplaceAllString("{{ .Rule.URL }}", in.Id))

	req, err := http.NewRequestWithContext(ctx, "{{ .Rule.RequestType }}", urlBuilder.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := h.unmarshaler.Unmarshal(res.Body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
`))
