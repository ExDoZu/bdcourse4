

package bdcourse

import (
	"bdcourse/internal/server"
)

func StartServer(dbconfig, serverAddr string) error {

	server, err := server.New(dbconfig)
	if err != nil {
		return err
	}

	return server.Run(serverAddr)
	
}
