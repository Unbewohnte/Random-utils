package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unbewohnte/Tiny-Utils/fileserver/logger"
	"github.com/unbewohnte/Tiny-Utils/fileserver/server"
)

var (
	PORT *int    = flag.Int("port", 8000, "Specifies the port")
	DIR  *string = flag.String("dir", ".", "Specifies the directory that will be served")
)

func init() {
	flag.Parse()
}

func main() {
	fs := server.NewFileServer(uint(*PORT))

	// check if directory does exist
	_, err := os.Stat(*DIR)
	if err != nil {
		logger.LogError(true, "Given directory does not exist")
	}

	fs.ServeDirectory("/", *DIR)

	logger.LogInfo("Created a new file server")

	localAddr, err := server.GetOutboundIP()
	if err != nil {
		logger.LogError(true, fmt.Sprintf("Could not retrieve your local IP address: %s", err))
	}

	logger.LogInfo(fmt.Sprintf("Serving \"%s\" at %s:%d", *DIR, localAddr.String(), *PORT))
	fmt.Println((fmt.Sprintf("Serving \"%s\" at %s:%d", *DIR, localAddr.String(), *PORT)))
	err = fs.Start()
	if err != nil {
		logger.LogError(true, "Fatal server error: ", err)
	}
}
