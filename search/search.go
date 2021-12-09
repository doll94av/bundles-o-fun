package search

import (
	 //"io/ioutil"
	 //"log"
	 "path/filepath"
	 "os"
)

func SearchDirectory() {
	
// 	files, err := ioutil.ReadDir("./unzippedbundles")
// 	if err != nil {
//         log.Fatal(err)
// 	}

// 	for _, test := range files {
// 		println(test.Name)
// 	}
// }
	var files []string
	err := filepath.Walk("./unzippedbundles", func(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		files = append(files, path)
	}
	return nil
	})
	if err != nil {

	}
	for _, test := range files {
		println(test)
	}
}
