package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
)

type fileBlock struct {
	fileId int
	start  int
	len    int
}

func (f fileBlock) checksum() int {
	if f.fileId == -1 {
		return 0
	}
	arrSum := (f.len * (2*f.start + f.len - 1)) / 2
	return f.fileId * arrSum
}

func defrag(disk []fileBlock) []fileBlock {
	for i := 0; i < len(disk); i++ {
		if disk[i].fileId != -1 { //if its not empty we skip
			continue
		}
		for j := (len(disk) - 1); j > i; j-- { //search from back for files we can put here
			if disk[j].fileId == -1 { //if its empty we dont care
				continue
			}
			if disk[j].len == disk[i].len { //if its the same size we can just swap
				disk[i].fileId = disk[j].fileId
				disk[j].fileId = -1
				break
			} else if disk[j].len > disk[i].len { //if its bigger we can split it
				disk[i].fileId = disk[j].fileId
				disk[j].len -= disk[i].len
				break
			}
			if disk[j].len < disk[i].len { //if its smaller we can put it here
				remaining := fileBlock{-1, disk[i].start + disk[j].len, disk[i].len - disk[j].len}
				disk[i].fileId = disk[j].fileId
				disk[i].len = disk[j].len
				disk[j].fileId = -1
				disk = slices.Insert(disk, i+1, remaining)
				i++
				j++
			}
		}
	}
	return disk
}

func defrag2(disk []fileBlock) []fileBlock {
	for i := 0; i < len(disk); i++ {
		if disk[i].fileId != -1 { //if its not empty we skip
			continue
		}
		for j := (len(disk) - 1); j > i; j-- { //search from back for files we can put here
			if disk[j].fileId == -1 || disk[j].len > disk[i].len { //if its empty or bigger we don't care
				continue
			}
			disk[i].fileId = disk[j].fileId
			disk[j].fileId = -1
			if disk[i].len == disk[j].len { // fit there is no empty space left we can go
				break
			}
			// otherwise we need to split the block and search for another candidate for remaining space
			remaining := fileBlock{-1, disk[i].start + disk[j].len, disk[i].len - disk[j].len}
			disk[i].len = disk[j].len
			disk = slices.Insert(disk, i+1, remaining)
			i++
			j++
		}
	}
	return disk
}

func calcChecksum(disk []fileBlock) int {
	sum := 0
	for _, block := range disk {
		sum += block.checksum()
	}
	return sum
}

func fs1(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	disk := make([]fileBlock, 0)
	fileId := 0
	point := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i += 2 {
			lenFile := int(line[i] - 48)
			lenSpace := 0
			if i+1 < len(line) {
				lenSpace = int(line[i+1] - 48)
			}
			disk = append(disk, fileBlock{fileId, point, lenFile})
			point += lenFile
			if lenSpace > 0 {
				disk = append(disk, fileBlock{-1, point, lenSpace})
			}
			point += lenSpace
			fileId++
		}
	}
	return calcChecksum(defrag(disk))
}

func fs2(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	disk := make([]fileBlock, 0)
	fileId := 0
	point := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i += 2 {
			lenFile := int(line[i] - 48)
			lenSpace := 0
			if i+1 < len(line) {
				lenSpace = int(line[i+1] - 48)
			}
			disk = append(disk, fileBlock{fileId, point, lenFile})
			point += lenFile
			if lenSpace > 0 {
				disk = append(disk, fileBlock{-1, point, lenSpace})
			}
			point += lenSpace
			fileId++
		}
	}
	return calcChecksum(defrag2(disk))
}

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fs1(f)
}
