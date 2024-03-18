package util

import (
	"Tmage/controller/status"
	"Tmage/pkg/grpc"
	"fmt"
	"github.com/listenGrey/TmagegRpcPKG/images"

	"context"
)

func GetFiles(tags []string) (status.Code, []string) {
	client := grpc.ClientServer(grpc.GetImages)
	if client == status.StatusConnGrpcServerErr {
		return status.StatusConnGrpcServerErr, []string{}
	}
	sendTag := &images.TagRequest{Tag: tags}
	res, err := client.(images.FileServiceClient).GetFileListByTag(context.Background(), sendTag)
	if err != nil {
		fmt.Printf("Failed to receive info from gRpc server; %v\n", err)
		return status.StatusRecvGrpcSerInfoErr, []string{}
	}
	var filenames []string
	for _, item := range res.GetFiles() {
		item.GetFilename()
		item.GetFormat()
		filenames = append(filenames, item.GetFilename()+"."+item.GetFormat())
	}
	return status.StatusSuccess, filenames
}
