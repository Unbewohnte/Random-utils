package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type FileServer struct {
	server          *http.Server
	serveMux        *http.ServeMux
	Port            uint
	ServingLocation string
}

func NewFileServer(port uint) *FileServer {
	serveMux := http.NewServeMux()

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        serveMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &FileServer{
		server:   server,
		serveMux: serveMux,
	}
}

// makes given directory and its subdirectories available
func (s *FileServer) ServeDirectory(pattern, dir string) {
	s.serveMux.Handle(pattern,
		http.StripPrefix(pattern, http.FileServer(http.Dir(dir))))
}

// makes available a certain file
func (s *FileServer) ServeFile(pattern, file string) {
	s.serveMux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})
}

func (s *FileServer) Start() error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

// thanks to https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func GetOutboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}
