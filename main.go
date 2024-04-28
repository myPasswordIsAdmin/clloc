package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Look for LoC in directory recursive
func lookUpDir(d string) (int, int, error) {

	loc := 0
	count := 0
	dir, err := os.ReadDir(d)
	if err != nil {
		log.Printf("Error while reading dir: %e", err)
		return 0, 0, err
	}

	for _, file := range dir {

		tempLoC := 0

		// Если директория - вызываем функцию рекурсивно
		if file.IsDir() {

			tempCount := 0
			tempLoC, tempCount, err = lookUpDir(d + file.Name() + "/")
			count += tempCount

			if err != nil {
				continue
			}

		} else {
			// Если файл, фильтруем и считаем LoC
			// ext := filepath.Ext(file.Name())

			//  Через slices.Contains в го 1.21
			if true {
				tempLoC = countLoC(d + file.Name())
				count += 1
			}

		}
		loc += tempLoC
	}

	return loc, count, nil
}

// Count loc in single file
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
	log.Printf("Scanned: %s (%d LoC)", f, loc)
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
		path += string(filepath.Separator)
	}

	result, count, err := lookUpDir(path)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Count of lines: %d, total files scanned: %d \n", result, count)
	fmt.Printf("Completed in  %f seconds", time.Since(startTime).Seconds())
}
