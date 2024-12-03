package flex

import (
	"encoding/json"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

func IntPtr(i int) *int {
	return &i
}

func LogFlexMessage(flexContainer *linebot.BubbleContainer) {
	jsonData, err := json.Marshal(flexContainer)
	if err != nil {
		log.Println("Error marshaling Flex Message JSON:", err)
		return
	}
	log.Println("Generated Flex Message JSON:", string(jsonData))
}

func createRichMenu(name_, patientID, age, sex, blood, phone_numbers string) *linebot.FlexMessage {
	if name_ == "" || patientID == "" || age == "" || sex == "" || blood == "" || phone_numbers == "" {
		log.Println("Invalid input data for Flex Message")
		return nil
	}

	flexContainer := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "ข้อมูลผู้ป่วย",
					Size:   linebot.FlexTextSizeTypeXl,
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#000000",
				},
				&linebot.SeparatorComponent{
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  "ชื่อ - สกุล: " + name_,
					Size:  linebot.FlexTextSizeTypeMd,
					Color: "#333333",
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  "รหัสผู้ป่วย: " + patientID,
					Size:  linebot.FlexTextSizeTypeMd,
					Color: "#333333",
				},
				&linebot.BoxComponent{
					Type:   linebot.FlexComponentTypeBox,
					Layout: linebot.FlexBoxLayoutTypeHorizontal,
					Contents: []linebot.FlexComponent{
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  "อายุ: " + age,
							Size:  linebot.FlexTextSizeTypeSm,
							Color: "#666666",
							Flex:  IntPtr(1),
						},
						&linebot.TextComponent{
							Type:  linebot.FlexComponentTypeText,
							Text:  "เพศ: " + sex,
							Size:  linebot.FlexTextSizeTypeSm,
							Color: "#666666",
							Flex:  IntPtr(1),
						},
					},
				},
				&linebot.TextComponent{
					Type:  linebot.FlexComponentTypeText,
					Text:  "หมู่เลือด: " + blood,
					Size:  linebot.FlexTextSizeTypeSm,
					Color: "#666666",
				},
				&linebot.SeparatorComponent{
					Margin: linebot.FlexComponentMarginTypeMd,
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "เบอร์โทร: " + phone_numbers,
					Size:   linebot.FlexTextSizeTypeMd,
					Color:  "#0000FF",
					Weight: linebot.FlexTextWeightTypeBold,
				},
			},
		},
	}

	LogFlexMessage(flexContainer)
	return linebot.NewFlexMessage("ข้อมูลผู้ป่วย", flexContainer)
}
