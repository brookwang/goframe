package service

import (
    "github.com/goframe/model"
    "github.com/goframe/dao"
    "github.com/goframe/cache"
)

type DataService struct {
    Mysql      *dao.Mysql
    Redis      *dao.Redis
    LocalCache *cache.LocalCache
}

func (dataService *DataService) GetBookInfoByName(name string, bookInfoMap model.BookInfoMap) error {
    return nil
}

func (dataService *DataService) GetLocalCache() map[string]int {
    return nil
}
