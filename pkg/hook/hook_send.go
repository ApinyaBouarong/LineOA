package hook

import (
	"LineOA/internal/repository"
	linebotConfig "LineOA/pkg/linebot"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
)

// ใช้ map หรือ session เพื่อเก็บสถานะการสนทนา
var userState = make(map[string]string)

func HandleLineWebhook(c *gin.Context) {
	bot := linebotConfig.GetLineBot()
	events, err := bot.ParseRequest(c.Request)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			c.Writer.WriteHeader(http.StatusBadRequest)
		} else {
			c.Writer.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	db := repository.GetDB()

	for _, lineevent := range events {
		if lineevent.Type == linebot.EventTypeMessage {
			message, ok := lineevent.Message.(*linebot.TextMessage)
			if ok {
				text := strings.TrimSpace(message.Text)
				log.Printf("Received message: %s", text)

				userID := lineevent.Source.UserID

				delete(userState, userID)

				switch text {
				case "ข้อมูลผู้สูงอายุ":
					userState[userID] = "awaiting_patient_name"
					bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("กรุณากรอกชื่อผู้สูงอายุ")).Do()

				case "ลงเวลาการทำงาน":
					userState[userID] = "awaiting_work_time"
					bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("กรุณากรอกเวลาการทำงานของคุณ")).Do()

				case "ประวัติการเข้ารับบริการ":
					userState[userID] = "awaiting_patient_name_for_history"
					bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("กรุณากรอกชื่อผู้สูงอายุเพื่อดูประวัติการเข้ารับบริการ")).Do()

				case "บันทึกการเข้ารับบริการ":
					userState[userID] = "awaiting_service_data"
					bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("กรุณากรอกชื่อผู้สูงอายุและข้อมูลการเข้ารับบริการ")).Do()

				default:
					switch userState[userID] {
					case "awaiting_patient_name":
						patientInfo, err := repository.GetPatientInfoByName(db, text)
						if err != nil {
							if err == repository.ErrNoRows {
								repository.ReplyDataNotFound(bot, lineevent.ReplyToken)
							} else {
								log.Println("Error fetching patient info:", err)
								bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด กรุณาลองใหม่")).Do()
							}
						} else {
							infoMessage := repository.FormatPatientInfo(patientInfo)
							bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage(infoMessage)).Do()
						}

					case "awaiting_work_time":
						workTime := text
						log.Printf("บันทึกเวลาการทำงาน: %s", workTime)
						bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("บันทึกเวลาการทำงานสำเร็จ")).Do()

					case "awaiting_patient_name_for_history":
						// serviceHistory, err := repository.GetServiceHistory(db, text)
						// if err != nil {
						// 	if err == repository.ErrNoRows {
						// 		repository.ReplyDataNotFound(bot, lineevent.ReplyToken)
						// 	} else {
						// 		log.Println("Error fetching service history:", err)
						// 		bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("เกิดข้อผิดพลาด กรุณาลองใหม่")).Do()
						// 	}
						// } else {
						// 	historyMessage := repository.FormatServiceInfo(serviceHistory)
						// 	bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage(historyMessage)).Do()
						// }

					case "awaiting_service_data":
						serviceData := text
						log.Printf("บันทึกข้อมูลการเข้ารับบริการ: %s", serviceData)
						bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("บันทึกข้อมูลการเข้ารับบริการสำเร็จ")).Do()

					default:
						bot.ReplyMessage(lineevent.ReplyToken, linebot.NewTextMessage("กรุณาเลือกคำสั่งที่ถูกต้อง")).Do()
					}
				}
			}
		}
	}
	c.Writer.WriteHeader(http.StatusOK)
	log.Println("Webhook response sent with status 200")
}
