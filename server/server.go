package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"unsafe"

	library "github.com/aikon001/color-api/colorapi"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var (
	enableTls       = flag.Bool("enable_tls", true, "Use TLS - required for HTTP2.")
	tlsCertFilePath = flag.String("tls_cert_file", "cert/localhost.crt", "Path to the CRT/PEM file.")
	tlsKeyFilePath  = flag.String("tls_key_file", "cert/localhost.key", "Path to the private key file.")
)

type color struct {
	Name string `json:"Name"`
	Hex  string `json:"Hex"`
	R    uint8  `json:"R"`
	G    uint8  `json:"G"`
	B    uint8  `json:"B"`
}

type server struct {
	library.UnimplementedColorsServer
}

func reverseByte(Array []byte) []byte {
	newArray := make([]byte, 0, len(Array))
	for i := len(Array) - 1; i >= 0; i-- {
		newArray = append(newArray, Array[i])
	}
	return newArray
}

func (s *server) AddColorFromHexOrRgb(ctx context.Context, c *library.Color) (reply *library.Reply, e error) {
	jsonFile, err := os.OpenFile("data/colorlist.json", os.O_WRONLY, 0644)
	if err != nil {
		return &library.Reply{Message: "Error reading json :"}, err
	}
	stat, _ := jsonFile.Stat()

	var colorRef color

	if len(c.Hex) != 0 {
		byt, _ := hex.DecodeString(c.Hex)
		colorRef = color{
			c.Name,
			c.Hex,
			byt[0],
			byt[1],
			byt[2],
		}
	} else if unsafe.Sizeof(c.RGB) != 0 {
		rgb := (*[3]byte)(unsafe.Pointer(&c.RGB))[:]
		rgb = reverseByte(rgb)

		hexadecimal := hex.EncodeToString(rgb)

		colorRef = color{
			c.Name,
			hexadecimal,
			rgb[0],
			rgb[1],
			rgb[2],
		}
	} else {
		return &library.Reply{Message: "No hexadecimal provided [Neither RGB provided!]"}, nil
	}

	byt, _ := json.Marshal(colorRef)

	jsonFile.Seek(stat.Size()-1, 0)

	jsonFile.WriteString(",")
	jsonFile.Write(byt)
	jsonFile.WriteString("]")

	defer jsonFile.Close()
	return &library.Reply{Message: "Ok"}, nil
}

func (s *server) PickAllColors(index *library.StartIndex, stream library.Colors_PickAllColorsServer) error {

	jsonFile, err := os.Open("data/colorlist.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var Allcolors []color
	json.Unmarshal(byteValue, &Allcolors)

	for _, element := range Allcolors {

		res := &library.Color{
			Name: element.Name,
			Hex:  element.Hex,
			RGB:  uint32(element.R | element.G | element.B),
		}

		err := stream.Send(&library.ColorResponse{ColorResp: res})

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	flag.Parse()

	port := 9090
	if *enableTls {
		port = 9091
	}

	grpcServer := grpc.NewServer()
	library.RegisterColorsServer(grpcServer, &server{})
	grpclog.SetLogger(log.New(os.Stdout, "exampleserver: ", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)
	handler := func(resp http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(resp, req)
	}

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Printf("Starting server. http port: %d, with TLS: %v", port, *enableTls)

	if *enableTls {
		if err := httpServer.ListenAndServeTLS(*tlsCertFilePath, *tlsKeyFilePath); err != nil {
			grpclog.Fatalf("failed starting http2 server: %v", err)
		}
	} else {
		if err := httpServer.ListenAndServe(); err != nil {
			grpclog.Fatalf("failed starting http server: %v", err)
		}
	}
}
