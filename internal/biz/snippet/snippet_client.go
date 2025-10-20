package snippet

import (
	"context"
	"leo/api/snippet/v1"
	v1 "leo/api/snippet/v1"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Add(key, value string) {
	client().Put(context.Background(), &v1.Request{
		Key:   key,
		Value: value,
	})
}

func Rm(key, value string) {
	client().Delete(context.Background(), &v1.Request{
		Key:   key,
		Value: value,
	})
}

func Query(key, value string) []*v1.Response_Pair {
	resp, err := client().Query(context.Background(), &v1.Request{
		Key:   key,
		Value: value,
	})
	if err != nil {
		return nil
	}
	return resp.GetPairs()
}

func client() snippet.SnippetClient {
	dir, _ := os.UserHomeDir()
	socketPath := dir + "/.peon/leo.sock"
	conn, err := grpc.NewClient("unix://"+socketPath, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return snippet.NewSnippetClient(conn)
}
