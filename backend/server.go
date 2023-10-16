package main

import (
	"context"
	"log"
	"net"

	"github.com/Lichttugend/Go-blog/backend/protos/blog"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) ListPosts(ctx context.Context, in *blog.Empty) (*blog.PostList, error) {
	// ここでデータベースから記事のリストを取得するロジックを実装します。
	// 今回はダミーデータを返す例とします。
	return &blog.PostList{
		Posts: []*blog.Post{
			{Id: "1", Title: "First Post", Content: "This is the first post.", Author: "Author1"},
			{Id: "2", Title: "Second Post", Content: "This is the second post.", Author: "Author2"},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	blog.RegisterBlogServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
