// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package handler

import (
	"log"
	"net/http"

	"github.com/Luftalian/Computer_software/model"
)

// serveWs handles websocket requests from the peer.
func ServeWs(hub *model.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := model.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := model.NewClient(hub, conn)
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
