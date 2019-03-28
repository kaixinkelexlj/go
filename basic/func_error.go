package basic

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type PathError struct {
	path       string
	op         string
	message    string
	createTime string
}

func (p *PathError) Error() string {
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}

func OpenFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return &PathError{fileName, "read", err.Error(), fmt.Sprintf("%v", time.Now())}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		println(err)
	}

	return nil
}
