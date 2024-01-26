package loader

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/kodekoding/phastos/v2/go/mail"
	"github.com/kodekoding/phastos/v2/go/storage"

	"godem/infrastructure/common"
)

var (
	smtp mail.SMTPs
	gcs  storage.Buckets
	err  error
)

func loadDataSources() {
	gcs, err = storage.NewGCS(context.Background(), common.BucketName)
	if err != nil {
		log.Fatalln("Failed to init Google Cloud Storage:", err.Error())
	}

	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtp = mail.NewSMTP(
		mail.WithHost(os.Getenv("SMTP_HOST")),
		mail.WithEmail(os.Getenv("SMTP_EMAIL")),
		mail.WithEmailPassword(os.Getenv("SMTP_PASSWORD")),
		mail.WithPort(smtpPort),
		mail.WithSender(os.Getenv("SMTP_SENDER_NAME")),
	)

}
