package server

import "text/template"

var BatchGet = template.Must(template.New("logging_batch_get").Parse(`
{{ .Comment }} func (h *httpClient) {{ .Method.GoName }}(ctx context.Context, in *{{ .Method.Input.GoIdent.GoName }}) (*{{ .Method.Output.GoIdent.GoName }}, error) {
	var response  {{ .Method.Output.GoIdent.GoName }}

	var responseBody = map[string][]string{"ids": in.Ids}

	var urlBuilder strings.Builder
	urlBuilder.WriteString(h.reqURL)
	urlBuilder.WriteString("{{ .Rule.URL }}")
	urlBuilder.WriteRune('/')

	jsonBody, err := json.Marshal(responseBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "{{ .Rule.RequestType }}", urlBuilder.String(), bytes.NewBuffer(jsonBody))
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
