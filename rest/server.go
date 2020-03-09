package rest

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// HandleShutdownReq for handling the metric request
func HandleShutdownReq(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid http method used. Use GET for health metrics"))
		return
	}

	w.WriteHeader(http.StatusOK)
	start := "Received shutdown request at" + time.Now().String()
	fmt.Println(start)
	w.Write([]byte(start))
	// Sleep for 120 seconds, and then send OK
	time.Sleep(120 * time.Second)

	end := "Sent ok response at" + time.Now().String()
	fmt.Println(end)
	w.Write([]byte("\n"))
	w.Write([]byte(end))

	return
}

// readPostReq  checks if this is a post requests and reads body of the request
func readPostReq(req *http.Request) (data []byte, err error) {
	if req.Method != http.MethodPost {
		err = errors.New("Invalid http method used. Use POST")
	}

	data, err = ioutil.ReadAll(req.Body)
	return
}

// writeError writes http error code with error message to the response
func writeError(httpStatus int, err error, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	w.Write([]byte(err.Error()))
	log.Println("Error while handling" + err.Error())
}

func handlePrintReq(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	data, err := readPostReq(req)

	if err != nil {
		writeError(http.StatusBadRequest, err, w)
	}
	fmt.Println(string(data))
}

//StartServer starts HTTP server with given port
func StartServer(ctx context.Context, port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/readyforshutdown", HandleShutdownReq)
	mux.HandleFunc("/printdata", handlePrintReq)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	//channel to exit goroutine
	doneCh := make(chan struct{})

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Server is shutting down ")
			shutdownCtx, cancel := context.WithTimeout(
				context.Background(),
				time.Second*5,
			)
			defer cancel()
			server.Shutdown(shutdownCtx)
		case <-doneCh:
		}

	}()

	log.Fatal(server.ListenAndServe())
	close(doneCh)
}
