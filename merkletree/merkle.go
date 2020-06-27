package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

type merkleTree struct {
	hash     string
	file     string
	fileMap  map[string]*merkleTree
}

func getHashForFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (m *merkleTree) generateHashFromChildren() string {
	h := md5.New()
	for _, child := range m.fileMap {
		if _, err := io.WriteString(h, child.hash); err != nil {
			panic(err)
		}
	}
	m.hash = fmt.Sprintf("%x", h.Sum(nil))
	return m.hash
}


func generateMerkleTree(root *merkleTree, dirPath string) {
	files, err := ioutil.ReadDir(dirPath)
	nonDirFiles := make([]string, 0)
	if err != nil {
		// log here
		panic(err)
	}
	for _, file := range files {
		fullPath := path.Join(dirPath, file.Name())
		if file.IsDir() {
			dirMerkleTree := &merkleTree{fileMap:make(map[string]*merkleTree), file:fullPath}
			generateMerkleTree(dirMerkleTree, fullPath)
			dirMerkleTree.generateHashFromChildren()
			root.fileMap[file.Name()] = dirMerkleTree
		} else {
			nonDirFiles = append(nonDirFiles, fullPath)
		}
	}
	for _, file := range nonDirFiles {
		md5Hash := getHashForFile(file)
		root.fileMap[path.Base(file)] = &merkleTree{hash:md5Hash, file:file}
	}
}

func diff(merkleNodeOne, merkleNodeTwo *merkleTree) []string {
	differedFiles := make([]string, 0)
	if merkleNodeOne.hash == merkleNodeTwo.hash {
		return []string{}
	}

	for k, mOne := range merkleNodeOne.fileMap {
		if mTwo, ok := merkleNodeTwo.fileMap[k]; ok {
			differedFiles = append(differedFiles, diff(mOne, mTwo)...)
		} else {
			differedFiles = append(differedFiles, k)
		}
	}
	if len(merkleNodeOne.fileMap) == 0 {
		differedFiles = append(differedFiles, merkleNodeOne.file)
	}
	return differedFiles
}

func main() {
	mOne := &merkleTree{fileMap:make(map[string]*merkleTree)}
	generateMerkleTree(mOne, "/home/ndakota/temp_dev/server_extension")
	fmt.Println(mOne.generateHashFromChildren())
	mTwo := &merkleTree{fileMap:make(map[string]*merkleTree)}
	generateMerkleTree(mTwo, "/home/ndakota/temp_dev/server_extension_copy")
	fmt.Println(mTwo.generateHashFromChildren())
	fmt.Println(diff(mOne, mTwo))
}