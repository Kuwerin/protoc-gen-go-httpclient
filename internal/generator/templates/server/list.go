package server

import "text/template"

var List = template.Must(template.New("logging_list").Parse(`
{{ .Comment }} func (h *httpClient) {{ .Method.GoName }}(ctx context.Context, in *{{ .Method.Input.GoIdent.GoName }}) (*{{ .Method.Output.GoIdent.GoName }}, error) {
	var response  {{ .Method.Output.GoIdent.GoName }}

	var urlBuilder strings.Builder
	urlBuilder.WriteString(h.reqURL)
	urlBuilder.WriteString("{{ .Rule.URL }}")
	urlBuilder.WriteRune('/')
	urlBuilder.WriteString("?limit=")
	urlBuilder.WriteString(strconv.Itoa(int(in.Limit)))
	urlBuilder.WriteString("&offset=")
	urlBuilder.WriteString(strconv.Itoa(int(in.Offset)))

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
