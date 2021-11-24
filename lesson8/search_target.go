package main

import (
	"errors"
	"os"
	"path/filepath"
	"sync"
)

var (
	ErrorWrongFileName = errors.New("wrong file name")
)

type SearchTarget struct {
	Name string
	Path string
	Size int64
	wg   *sync.WaitGroup
}

type Checker interface {
	Check(info os.FileInfo, filePath string) bool
	WgAdd()
	WgDone()
	WgWait()
}

func (s *SearchTarget) Check(fi os.FileInfo, filePath string) bool {
	if s.Path == filePath {
		return false
	}
	return s.isCopy(fi)
}

func (s *SearchTarget) isCopy(fi os.FileInfo) bool {
	return s.Name == fi.Name() && s.Size == fi.Size()
}

func (s *SearchTarget) WgAdd() {
	s.wg.Add(1)
}

func (s *SearchTarget) WgDone() {
	s.wg.Done()
}

func (s *SearchTarget) WgWait() {
	s.wg.Wait()
}

func NewSearchTarget(fileName string) (*SearchTarget, error) {
	if fileName == "" {
		return nil, ErrorWrongFileName
	}

	fi, err := os.Stat(fileName)
	if err != nil {
		return nil, err
	}

	path, err := filepath.Abs(fileName)
	if err != nil {
		return nil, err
	}

	return &SearchTarget{
		Name: fi.Name(),
		Path: path,
		Size: fi.Size(),
		wg:   new(sync.WaitGroup),
	}, nil
}