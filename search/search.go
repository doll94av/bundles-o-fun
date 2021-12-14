package search

import (
	 "log"
	 "path/filepath"
	 "os"
	 "bufio"
	 "strings"
	 "time"
)

//loads in keys from the 'phrases' file and returns a map of each line as a value in the map
func loadPhrases()map[string]int {
	file, err := os.Open("./phrases")
	var loadingMap  = make(map[string]int)


	if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
    // optionally, resize scanner's capacity for lines over 64K, see next example
    for scanner.Scan() {
    	println(scanner.Text())
		loadingMap[scanner.Text()] = 0
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return loadingMap
}
func SearchDirectory() {

	//setup map for counting occurances as well as loading keys
	var phraseMap map[string]int
	phraseMap = loadPhrases()
	for key, value := range phraseMap {
		println("Key:", key, "Value:", value)
	}

	var files []string
	err := filepath.Walk("./unzippedbundles", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
		files = append(files, path)
	}
		return nil
	})
	
	if err != nil {
		log.Fatal(err)
	}

	//loads in each file in the dir
	startTIme := time.Now()

	for _, currentFile := range files {
		//println(test)
		openedFile, _ := os.Open(currentFile)
		if err != nil {
			log.Fatal(err)
		}
		defer openedFile.Close()

		scanner := bufio.NewScanner(openedFile)
		for scanner.Scan() {
			//println(scanner.Text())
			for key, _ := range phraseMap {
				// if strings.Contains(scanner.Text(), key) {
				// 	println("Key ", key, " found in ", currentFile)
				// 	println(scanner.Text())
				// 	phraseMap[key] = value + 1
				// }
				go checkKey(key, scanner.Text(), currentFile)
			}
		}
		
	}
	endTime := time.Now()

	//print out final counts of occurances
	println("start time: ", startTIme.String(), " end time: ", endTime.String())

}

func checkKey(key string, logLine string, fileName string){
	if strings.Contains(logLine, key) {
		println("Key ", key, " found in ", fileName)
		println(logLine)
	}
}
