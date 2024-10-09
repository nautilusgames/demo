package handler

import (
	"encoding/binary"
	"fmt"
	"io"
	"net/http"
)

func (h *Handler) HandleUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("read body error %v \n", err)
			w.Write([]byte("read body error"))
		}

		w.Write([]byte(fmt.Sprintf("upload done, size %d", binary.Size(bytes))))
	}
}
