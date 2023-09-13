package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"os"
)

var disk StorageMedia

type StorageMedia interface {
	Upload(path string, reader io.Reader) (err error)
	Download(path string, writer io.Writer) error
	Delete(path string) error
}

// STORAGE_ACCESS=wfIswdHUf1GJUXnWBS56;STORAGE_BUCKET=test;STORAGE_SECRET=b1Cf9FCugVkOY579KmodVIRZTEXp010IbSvzb8X7;STORAGE_TYPE=s3;STORAGE_ENDPOINT=http://localhost:9000
func InitStorageMedia() {
	storageType := viper.GetString("storage.type")
	switch storageType {
	case "s3":
		accessKey := viper.GetString("storage.access")
		secretKey := viper.GetString("storage.secret")
		endpoint := viper.GetString("storage.endpoint")
		bucket := viper.GetString("storage.bucket")
		newSession, err := session.NewSession(&aws.Config{
			Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
			Endpoint:         aws.String(endpoint),
			Region:           aws.String("cn-north-1"),
			DisableSSL:       aws.Bool(true),
			S3ForcePathStyle: aws.Bool(true),
		})
		if err != nil {
			panic(err)
		}
		if bucket == "" {
			panic("bucketName不能为空")
		}
		svc := s3.New(newSession)
		// 检查Bucket是否存在
		_, err = svc.HeadBucket(&s3.HeadBucketInput{
			Bucket: aws.String(bucket),
		})
		if err != nil {
			// 如果Bucket不存在，则创建新的Bucket
			_, err = svc.CreateBucket(&s3.CreateBucketInput{
				Bucket: aws.String(bucket),
			})
			if err != nil {

			}
			zap.S().Infof("Bucket '%s' 创建成功\n", bucket)
		}
		disk = &AWSS3{session: newSession, Bucket: bucket}
	default:
		disk = &DiskStorage{BasePath: "./"}
	}
}

type DiskStorage struct {
	BasePath string
}

func (d *DiskStorage) Upload(path string, reader io.Reader) (err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return
	}
	defer f.Close()
	_, err = io.Copy(f, reader)
	if err != nil {
		return
	}
	return
}
func (d *DiskStorage) Download(path string, writer io.Writer) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, f)
	if err != nil {
		return err
	}
	return nil
}
func (d *DiskStorage) Delete(path string) error {
	return os.Remove(path)
}

type AWSS3 struct {
	session *session.Session
	Bucket  string
}

func (a *AWSS3) Upload(path string, reader io.Reader) (err error) {
	uploader := s3manager.NewUploader(a.session)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(path),
		Body:   reader,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *AWSS3) Download(path string, writer io.Writer) error {
	buffer := aws.NewWriteAtBuffer([]byte{})
	_, err := s3manager.NewDownloader(a.session).Download(buffer,
		&s3.GetObjectInput{
			Bucket: aws.String(a.Bucket),
			Key:    aws.String(path),
		})
	if err != nil {
		return err
	}
	buffer.Bytes()
	_, err = writer.Write(buffer.Bytes())
	return err
}

func (a *AWSS3) Delete(path string) error {
	svc := s3.New(a.session)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(a.Bucket), Key: aws.String(path)})
	if err != nil {
		return err
	}

	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(a.Bucket),
		Key:    aws.String(path),
	})
	return err
}
func Upload(path string, reader io.Reader) (err error) {
	return disk.Upload(path, reader)
}
func Download(path string, writer io.Writer) error {
	return disk.Download(path, writer)
}
func Delete(path string) error {
	return disk.Delete(path)
}
