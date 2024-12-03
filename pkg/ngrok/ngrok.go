package ngrok

import (
	"LineOA/pkg/database"
	"fmt"
	"log"
)

func GetNgrokAuthToken() string {
	config, err := database.LoadConfig()
	if err != nil {
		log.Fatalf("ไม่สามารถโหลดไฟล์ config ได้: %v", err)
	}

	authtoken := config.Agent.Authtoken
	fmt.Println("Ngrok Auth Token:", authtoken)
	return authtoken
}
