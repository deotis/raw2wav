package main

import (
	"fmt"
	"os"
	"os/exec"
	//"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func usage_print() {

	fmt.Println("Usage: raw2wav [input type] [dir path] [output_type]")
	fmt.Println("	- input type: alaw|mulaw")
	fmt.Println("	- output type: pcm_alaw|pcm_mulaw")
	fmt.Println("	ex)raw2wav alaw ./sound/*.vce pcm_alaw")
}

func main() {
	if (len(os.Args) < 4) {
		usage_print()
		os.Exit(0)
	}

	if (strings.Compare(os.Args[1], "alaw") != 0 && strings.Compare(os.Args[1], "mulaw") != 0) {
		usage_print()
		os.Exit(1)
	} 

	if (strings.Compare(os.Args[3], "pcm_alaw") != 0 && strings.Compare(os.Args[3], "pcm_mulaw") != 0) {
		usage_print()
		os.Exit(2)
	} 

	//fmt.Println(os.Args[1], os.Args[2])
	//files, err := ioutil.ReadDir("./")
	//if err != nil {
    //    log.Fatal(err)
    //}
	//fmt.Println(files)
	//for _, file := range files {
	//	fmt.Println(file.Name())
	//}

	files, err := filepath.Glob(os.Args[2])
 	if err != nil {
		log.Fatal(err)
		os.Exit(3)
	}

	if (len(files) == 0) {
		fmt.Println("there is no input files in", os.Args[2])
		os.Exit(4)
	}

	var out_files []string
	for _, value := range files {
		if (!strings.Contains(value, ".vce")) {
			continue
		}

		fmt.Println(value)
		str := strings.TrimSuffix(value, ".vce")
		str += ".wav"
		out_files = append(out_files, str)
		//fmt.Println(str)
	}
	fmt.Println(out_files)

	for i, out_file := range out_files {
		cmd := exec.Command("ffmpeg", "-f", os.Args[1], "-ar", "8000", "-ac", "1", "-i", files[i], "-acodec", os.Args[3], out_file)
		fmt.Println(cmd.Args)
		cmd.Run();
	}
}
