package repository

import (
	"LineOA/internal/models"
	"LineOA/pkg/database"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/line/line-bot-sdk-go/linebot"
)

var ErrNoRows = errors.New("no rows found")
var db *sql.DB
var once sync.Once

func GetDB() *sql.DB {
	once.Do(func() {
		var err error
		db, err = database.ConnectToDB()
		if err != nil {
			log.Fatalf("Error connecting to database: %v", err)
		}
	})
	return db
}

func GetPatientInfoByName(db *sql.DB, name_ string) (*models.PatientInfo, error) {
	query := `SELECT name_, patiet_id, age, sex, blood, phone_numbers FROM patient_info WHERE name_ = ?`
	row := db.QueryRow(query, name_)

	var patientInfo models.PatientInfo
	err := row.Scan(&patientInfo.Name, &patientInfo.PatientID, &patientInfo.Age, &patientInfo.Sex, &patientInfo.Blood, &patientInfo.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ไม่พบข้อมูลผู้ป่วยที่มีชื่อ %s", name_)
		}
		return nil, err
	}
	log.Println("ข้อมูลผู้ป่วยที่ดึงมา:", &patientInfo)
	return &patientInfo, nil
}

func GetServiceInfoByName(db *sql.DB, name string) ([]models.ServiceInfo, error) {
	query := `SELECT patient_info.id, patient_info.patiet_id, patient_info.name_, service_info.activity
	FROM patient_info
	INNER JOIN service_info 
	ON patient_info.id = service_info.service_id WHERE patient_info.name_ = ?`
	rows, err := db.Query(query, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var serviceInfos []models.ServiceInfo
	for rows.Next() {
		var serviceInfo models.ServiceInfo
		err := rows.Scan(&serviceInfo.PatientInfo.ID, &serviceInfo.PatientInfo.PatientID, &serviceInfo.PatientInfo.Name, &serviceInfo.Activity)
		if err != nil {
			return nil, err
		}
		serviceInfos = append(serviceInfos, serviceInfo)
	}

	if len(serviceInfos) == 0 {
		return nil, fmt.Errorf("ไม่พบข้อมูลกิจกรรมสำหรับผู้ป่วย: %s", name)
	}

	return serviceInfos, nil

}

func GetAllActivity(db *sql.DB, patientID string) ([]models.ServiceInfo, error) {
	query := `SELECT activity FROM service_info WHERE patient_id = ?`
	rows, err := db.Query(query, patientID)

	if err != nil {
		return nil, fmt.Errorf("ไม่สามารถดึงข้อมูลกิจกรรมได้: %v", err)
	}
	defer rows.Close()

	var activities []models.ServiceInfo

	for rows.Next() {
		var activity models.ServiceInfo
		err := rows.Scan(&activity.Activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil

}

func FormatPatientInfo(patient *models.PatientInfo) string {
	return fmt.Sprintf("ข้อมูลผู้ป่วย:\nชื่อ: %s\nรหัสผู้ป่วย: %s\nอายุ: %d\nเพศ: %s\nหมู่เลือด: %s\nหมายเลขโทรศัพท์: %s",
		patient.Name, patient.PatientID, patient.Age, patient.Sex, patient.Blood, patient.PhoneNumber)
}

func FormatServiceInfo(serviceInfo []models.ServiceInfo) string {
	if len(serviceInfo) == 0 {
		return "ไม่พบกิจกรรมสำหรับผู้ป่วยนี้ กรุณาลองใหม่อีกครั้ง"
	}

	message := fmt.Sprintf("ชื่อผู้ป่วย: %s\nกิจกรรมที่สำเร็จแล้ว:\n", serviceInfo[0].PatientInfo.Name)
	for _, info := range serviceInfo {
		message += fmt.Sprintf("- %s\n", info.Activity)
	}

	activities := []string{
		"แช่เท้า", "นวด/ประคบ", "ฝังเข็ม", "คาราโอเกะ", "ครอบแก้ว",
		"ทำอาหาร", "นั่งสมาธิ", "เล่าสู่กัน", "ซุโดกุ", "จับคู่ภาพ",
	}
	message += "\nเลือกกิจกรรมที่คุณต้องการเพิ่ม:\n"
	for _, activity := range activities {
		message += fmt.Sprintf("- %s\n", activity)
	}
	return message
}

func ReplyErrorFormat(bot *linebot.Client, replyToken string) {
	if _, err := bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage("กรุณากรอกรูปแบบข้อความให้ถูกต้อง เช่น 'นางสมหวัง สดใส'"),
	).Do(); err != nil {
		log.Println("เกิดข้อผิดพลาดในการส่งข้อความ:", err)
	}
}

func ReplyDataNotFound(bot *linebot.Client, replyToken string) {
	notFoundMessage := "ไม่พบข้อมูลผู้สูงอายุตามชื่อ กรุณาลองใหม่อีกครั้ง"
	if _, err := bot.ReplyMessage(replyToken, linebot.NewTextMessage(notFoundMessage)).Do(); err != nil {
		log.Println("Error sending not found message:", err)
	}
}

func GetServiceHistory(db *sql.DB, text string) {

}
