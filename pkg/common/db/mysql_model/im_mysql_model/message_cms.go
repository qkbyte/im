package im_mysql_model

import (
	"Open_IM/pkg/common/db"
	"fmt"
)

func GetChatLog(chatLog db.ChatLog, pageNumber, showNumber int32) ([]db.ChatLog, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var chatLogs []db.ChatLog
	if err != nil {
		return chatLogs, err
	}
	dbConn.LogMode(true)
	err = dbConn.Table("chat_logs").Where(fmt.Sprintf(" content like '%%%s%%'", chatLog.Content)).
		Limit(showNumber).Offset(showNumber * (pageNumber - 1)).Find(&chatLogs).Error
	return chatLogs, err
}

func GetChatLogCount(chatLog db.ChatLog) (int64, error) {
	dbConn, err := db.DB.MysqlDB.DefaultGormDB()
	var count int64
	if err != nil {
		return count, err
	}
	dbConn.LogMode(true)
	err = dbConn.Table("chat_logs").Where(fmt.Sprintf(" content like '%%%s%%' and ", chatLog.Content)).Count(&count).Error
	return count, err
}