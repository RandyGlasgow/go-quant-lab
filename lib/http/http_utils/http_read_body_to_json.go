package http_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ExtractJsonFromBody[T any](response *http.Response, v *T) error {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return nil
}
