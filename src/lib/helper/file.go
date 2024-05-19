package helper

import (
	"context"
	"fmt"
	"github.com/kodekoding/phastos/v2/go/helper"
	"github.com/kodekoding/phastos/v2/go/storage"
	"github.com/pkg/errors"
	"io"
	"io/fs"

	"os"
)

func GetGCSFile(ctx context.Context, filePath string, gcs ...storage.Buckets) (fs.File, error) {
	var gcsObj storage.Buckets
	var err error
	if len(gcs) > 0 {
		gcsObj = gcs[0]
	} else if gcs == nil {
		gcsObj, err = storage.NewGCS(ctx, os.Getenv("HEYKUDO_BUCKET_NAME"))
		if err != nil {
			return nil, errors.Wrap(err, "heykudo.lib.helper.file.GetGCSFile.InitGCS")
		}
		defer gcsObj.Close()
	}

	content, err := gcsObj.GetFileFS(ctx, filePath)
	if err != nil {
		return nil, errors.Wrap(err, "error get file GCS")
	}

	return content, nil
}

func CopyToTempFolder(sourceFile fs.File, destFileName *string) error {
	baseFolder := "/tmp/copy"
	if IsLocal() {
		baseFolder = fmt.Sprintf("files%s", baseFolder)
	}

	helper.CheckFolder(baseFolder)

	*destFileName = fmt.Sprintf("%s/%s", baseFolder, *destFileName)
	if _, err := os.Stat(*destFileName); err != nil {
		// if the file isn't exists, then create new
		tmpFile, err := os.Create(*destFileName)
		if err != nil {
			return errors.Wrap(err, "heykudo.lib.helper.file.CopyToTempFolder.CreateNewFile")
		}

		buf := make([]byte, 1000)
		for {
			n, err := sourceFile.Read(buf)
			if err != nil && err != io.EOF {
				return err
			}
			if n == 0 {
				break
			}

			if _, err = tmpFile.Write(buf[:n]); err != nil {
				return errors.Wrap(err, "heykudo.lib.helper.file.CopyToTempFolder.WriteTempFile")
			}
		}
	}

	return nil
}
