package cmd

import (
	"context"
	"log"
	"net"

	"git.kicoe.com/blog/write/models/protodata"
	"git.kicoe.com/blog/write/modules/setting"
	"git.kicoe.com/blog/write/services/code"
	"git.kicoe.com/blog/write/modules/se"
	"github.com/gookit/color"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcS struct {
	protodata.CodeServer
}

func RunRpc(ctx *cli.Context) error {
	listener, err := net.Listen("tcp", ":"+setting.RPC.Port)
	if err != nil {
		log.Fatalf("error start rpc server: %v", err)
	}

	se.InitRiot()
	// reindex

	s := grpc.NewServer()
	reflection.Register(s)
	protodata.RegisterCodeServer(s, &grpcS{})
	color.Printf("⇨ grpc server started on <cyan>[::]:%s</>\n", setting.RPC.Port)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return nil
}

func (s *grpcS) Search(ctx context.Context, request *protodata.SearchRequest) (*protodata.SearchResponse, error) {
	codes, _, err := code.SearchDoc(request.Text, 1, 30)
	if err != nil {
		return nil, err
	}
	result := &protodata.SearchResponse{}
	for _, code := range codes {
		result.Data = append(result.Data, &protodata.CodeDetail{
			Id:      code.ID,
			Title:   code.Description,
			Lang:    code.Lang,
			Tags:    code.Tags,
			Content: code.Content,
		})
	}
	return result, nil
}
