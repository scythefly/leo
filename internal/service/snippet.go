package service

import (
	"context"
	"leo/api/snippet/v1"
	"leo/internal/data"
)

func NewSnippetService(data *data.Data) *SnippetService {
	return &SnippetService{data: data}
}

// SnippetService is a snippet service.
type SnippetService struct {
	snippet.UnimplementedSnippetServer

	data *data.Data
}

func (s *SnippetService) Query(_ context.Context, req *snippet.Request) (*snippet.Response, error) {
	resp := s.data.SnippetQuery(req.GetKey(), req.GetValue())
	return resp, nil
}
func (s *SnippetService) Put(_ context.Context, req *snippet.Request) (*snippet.Response, error) {
	return s.data.SnippetPut(req.GetKey(), req.GetValue()), nil
}
func (s *SnippetService) Delete(_ context.Context, req *snippet.Request) (*snippet.Response, error) {
	return s.data.SnippetDelete(req.GetKey(), req.GetValue()), nil
}
