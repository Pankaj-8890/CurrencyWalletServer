package Currencyconverter

import (
	"context"
	cv "currencyServer/conversion"
	pb "currencyServer/proto"
	"errors"
	"fmt"
	"log"
	"math"
	"net"

	"google.golang.org/grpc"
)


var add string = "0.0.0.0:9090"


type Server struct{
	pb.CurrencyServer
}


func Init(){
	lis, err := net.Listen("tcp",add)

	if err != nil{
		log.Fatalf("Failed to Listen :%v\n",err)
	}

	log.Printf("listening %s\n",add)

	s := grpc.NewServer()
	pb.RegisterCurrencyServer(s,&Server{})
	if err = s.Serve(lis); err != nil{
		log.Fatalf("failed to server %v\n",err)
	}
	
}

func(s *Server)ConvertToINR(ctx context.Context, req *pb.Request) (*pb.Response, error){

	fmt.Println("inside ConvertToINR")

	currency := req.Currency
	value := req.Value

	// fmt.Println(req)

	currencyType,isPresent := cv.GetCurrencyType(currency)

	// fmt.Println(currencyType)

	if(!isPresent){
		return nil,errors.New("CurrencyType Not found")
	}


	val := float32(math.Round((currencyType.ConversionFactor) * 100)/100)
	return &pb.Response{Value: value * val},nil;

}

func(s *Server)ConvertFromINR(ctx context.Context, req *pb.Request) (*pb.Response, error){


	fmt.Printf("inside ConvertFromINR")
	targetCurrency := req.TargetCurrency
	value := req.Value

	// fmt.Println(req)

	currencyType,isPresent := cv.GetCurrencyType(targetCurrency)

	// fmt.Println(currencyType)

	if(!isPresent){
		return nil,errors.New("CurrencyType Not found")
	}

	val := float32(currencyType.ConversionFactor)
	finalVal := float32(math.Round(float64(value/val) * 100)/100)
	return &pb.Response{Value: finalVal},nil;


}