package utils

import (
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetRandInt(anchor int) int {
	rand.Seed(int64(anchor))
	return rand.Intn(100000)
}

func GetFile(filePath string) (f *os.File, err error) {
	if PathExists(filePath) {
		f, err = os.OpenFile(filePath, os.O_APPEND, 0666)
	} else {
		f, err = os.Create(filePath)
	}
	if err != nil {
		log.Println(err)
		log.Println("open file err")
	}
	return f, err
}

func RandIntGen(filePath string, numVol int) {
	//var numVol = 100000
	var f *os.File
	var err error
	//var filePath = "IntegerArray.txt"
	if f,err = GetFile(filePath);err != nil{
		return
	}
	for i := 0 ; i < numVol ; i++ {
		//if _, err = io.WriteString(f, strconv.Itoa(GetRandInt(i))+"\n"); err != nil {
		//	panic(err)
		//}
		io.WriteString(f, strconv.Itoa(GetRandInt(i))+"\n")
	}
	log.Println("done")
}