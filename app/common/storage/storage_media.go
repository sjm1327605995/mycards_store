package storage

import (
	"io"
	"os"
)

var disk StorageMedia

type StorageMedia interface {
	Upload(path string, reader io.Reader) (err error)
	Download(path string) (io.ReadCloser, error)
	Delete(path string) error
}

func InitStorageMedia(s StorageMedia) {
	disk = s
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
func (d *DiskStorage) Download(path string) (io.ReadCloser, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (d *DiskStorage) Delete(path string) error {
	return os.Remove(path)
}
