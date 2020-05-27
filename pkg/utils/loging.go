package utils

import (
	"io"
	"log"
	"os"
	"time"
)

// LoggingSettings ログの設定をします。
func LoggingSettings() {

	// 日付単位でログファイルを作成する。
	day := time.Now()
	const layout = "2006-01-02"
	filePath := "./log/" + day.Format(layout) + ".log"

	// ファイルに権限を付与する。
	logfile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=logFile err=%s", err.Error())
	}

	// 標準出力とファイルの両方を出力先に設定する。
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}
