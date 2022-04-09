package utils

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

type Config map[string]string

func ReadProperties(fileName string) (Config, error) {
	waitgroup := sync.WaitGroup{}
	lineChannel := make(chan string)
	configChannel := make(chan Config)
	waitgroup.Add(2)
	go ReadFileLines(lineChannel, &waitgroup, fileName)
	go parsePropertiesLine(lineChannel, configChannel, &waitgroup)
	var actualConf Config
	for conf := range configChannel {
		actualConf = conf
	}
	waitgroup.Wait()
	if actualConf != nil {
		return actualConf, nil
	}
	return nil, errors.New("could not parse the configuration")
}

func ReadFileLines(lineChannel chan string, wg *sync.WaitGroup, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println(fmt.Sprintf("could not open file %s", filePath))
		wg.Done()
	}
	//close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("could not close the file")
		}
	}(file)

	//close channel in the end
	defer close(lineChannel)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineChannel <- scanner.Text()
	}

	wg.Done()
}

func parsePropertiesLine(lineCh chan string, configChannel chan Config, wg *sync.WaitGroup) {
	config := Config{}
	for line := range lineCh {
		parts := strings.Split(line, "=")
		config[parts[0]] = parts[1]
	}
	configChannel <- config
	close(configChannel)
	wg.Done()
}
