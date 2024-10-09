package store

import (
	"gvadmin_v3/core/config"
	"sync"
)

var (
	storeClient StoreClient
	once        sync.Once
)

type StoreClient interface {
	UploadFile(dstFileName string, localFilePath string) (string, error)
	DeleteFile(dstFileName string) error
}

func Instance() StoreClient {
	if storeClient == nil {
		once.Do(func() {
			switch config.Instance().Store.StoreType {
			case "minio":
				storeClient = newMinioClient()
			case "oss":
				storeClient = newOssClient()
			default:
			}
		})
	}
	return storeClient
}
