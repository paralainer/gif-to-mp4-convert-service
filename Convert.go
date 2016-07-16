package main

import (
	"os/exec"
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		template, _ := ioutil.ReadFile("index.html");
		w.Write(template);
	} else {
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("file")
		if (err != nil){
			writeError(w)
			return
		}
		defer file.Close()

		bytes, err := ioutil.ReadAll(file)
		if (err != nil) {
			writeError(w)
			return
		}

		res, err := convert(bytes);
		if (err != nil) {
			writeError(w)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=result.mp4")
		w.Header().Set("Content-Type", "video/mp4")
		w.Write(res)
	}
}

func writeError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("error"))
}

func main() {
	http.HandleFunc("/convert", handler)
	http.ListenAndServe(":8080", nil)
}

func convert(input []byte) ([]byte, error) {
	os.MkdirAll("convert/input", os.ModePerm);
	os.MkdirAll("convert/output", os.ModePerm);

	fileName := uuid.NewV4().String()
	inputFileName := "convert/input/" + fileName + ".gif"
	outputFileName := "convert/output/" + fileName + ".mp4"

	ioutil.WriteFile(inputFileName, input, os.ModePerm);

	cmd := "-f gif -i " + inputFileName + " -pix_fmt yuv420p -c:v libx264 -movflags +faststart -filter:v crop='floor(in_w/2)*2:floor(in_h/2)*2' " + outputFileName;
	output, err := exec.Command("ffmpeg", strings.Split(cmd, " ")...).Output()
	if (err != nil) {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(string(output[:]))

	result, err := ioutil.ReadFile(outputFileName)
	if (err != nil){
		return nil, err
	}

	os.Remove(inputFileName)
	os.Remove(outputFileName)

	return result, nil
}
