package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"shanker.com/grpc/invoicer"
)

type myInvoicerServer struct{
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error){
	return &invoicer.CreateResponse{
		Pdf: []byte("test"),
		Docx: []byte("test"),
	},nil
}




func main(){
	lis,err := net.Listen("tcp",":8089")

	if(err!=nil){
		log.Fatalf("Cannot create listner: %s",err)
	}

	serverRegistar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistar,service)

	err = serverRegistar.Serve(lis)

	if(err!=nil){
		log.Fatalf("Not possible to serve: %s",err)
	}


}