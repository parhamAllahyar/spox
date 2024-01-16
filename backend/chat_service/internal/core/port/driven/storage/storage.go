package storage

type Storage interface {
	ObjectRemove() error
	ObjectUpload(path string) error
	ObjectDownloadPerSign() (string, error)
}
