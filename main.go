package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

// LoC in directory
func lookUpDir(d string) (int, error) {

	loc := 0
	dir, err := os.ReadDir(d)
	if err != nil {
		log.Printf("Error while reading dir: %e", err)
		return 0, err
	}

	for _, file := range dir {

		tempLoC := 0
		fmt.Println(d + file.Name())

		if file.IsDir() {

			tempLoC, err = lookUpDir(d + file.Name() + "/")

			if err != nil {
				continue
			}

		} else {

			tempLoC = countLoC(d + file.Name())
		}
		loc += tempLoC
	}

	return loc, nil
}

func countLoC(f string) int {

	loc := 0
	file, err := os.Open(f)

	if err != nil {
		log.Printf("Error while reading file %s: %e", f, err)
		return loc
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		loc += 1
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error while scanning file %s: %e", f, err)
	}

	return loc
}

func main() {
	startTime := time.Now()
	// Парсим флаг пути
	var path string
	flag.StringVar(&path, "path", "./", "Path to inspected directory. Current directory for default")
	flag.Parse()

	// Добавляем слеш в конец пути
	if path[len(path)-1] != '/' {
		path += "/"
	}

	result, err := lookUpDir(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count of lines in %s: %d \n", path, result)
	fmt.Printf("Completed in  %f seconds", time.Since(startTime).Seconds())
}
