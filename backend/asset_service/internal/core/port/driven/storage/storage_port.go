package storage

import "io"

type Storage interface {
	ObjectRemove() error
	ObjectUpload(io.Reader) error
	ObjectDownloadPerSign(string) (string, error)
}
