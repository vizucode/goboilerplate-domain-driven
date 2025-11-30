package jsonplaceholder

import (
	"context"
	"goboilerplate-domain-driven/pkg/utils"
	"io"
	"net/http"
)

func (s *jsonPlaceHolder) FetchExternal(ctx context.Context) error {
	client := utils.NewClient(ctx, http.DefaultClient)

	// Build request
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts/1", nil)

	// Hit external
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Use response
	io.ReadAll(resp.Body)

	return nil
}
