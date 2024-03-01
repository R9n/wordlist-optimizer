package structs

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type FileHandler struct{
	InputFilepath 	string;
	OutputFilePath  string;
	RemovedFilePath string;
	InputFile       *os.File;
	OutputFile      *os.File;
	RemovedFile     *os.File;
	Scanner       	*bufio.Scanner
}
var MAX_LINES_READ_BUFFER_SIZE = 100_000; 

func (fileHandler * FileHandler)OpenFile() {

    file, err := os.Open(fileHandler.InputFilepath)
    if err != nil {
		fmt.Println("Error opening wordlist file: "+ fileHandler.InputFilepath);
        log.Fatal(err);
		os.Exit(1);
    }
	fileHandler.InputFile = file;
	fileHandler.Scanner = bufio.NewScanner(file)
}

func (fileHandler * FileHandler) CreateOutputFile(){
	file, err := os.Create(fileHandler.OutputFilePath)
	if err != nil {
		fmt.Println("Error creating optimized wordlist file: "+ fileHandler.OutputFilePath);
        log.Fatal(err);
		os.Exit(1);
	}
	fileHandler.OutputFile = file;
}

func (fileHandler * FileHandler) CreateRemovedFile(){
	file, err := os.Create(fileHandler.RemovedFilePath)
	if err != nil {
		fmt.Println("Error creating removed file: " + fileHandler.OutputFilePath);
        log.Fatal(err);
		os.Exit(1);
	}
	fileHandler.RemovedFile = file;
}

func (fileHandler * FileHandler) CloseInputFile(){
	fileHandler.InputFile.Close();
}

func (fileHandler * FileHandler) CloseRemovedFile(){
	fileHandler.RemovedFile.Close();
}

func (fileHandler * FileHandler) CloseOutputFile(){
	fileHandler.OutputFile.Close();
}

func (fileHandler * FileHandler) WriteLineToOutputFile(line string){
	_, err := fileHandler.OutputFile.WriteString(line+"\n")
    if err != nil {
		fmt.Println("Error writing to output wordlist"+ fileHandler.OutputFilePath);
        log.Fatal(err);
		os.Exit(1);
    }
}

func (fileHandler * FileHandler) WriteLineToRemovedFile(line string){
	_, err := fileHandler.RemovedFile.WriteString(line+"\n")
    if err != nil {
		fmt.Println("Error writing to RemovedFile " + fileHandler.RemovedFilePath);
        log.Fatal(err);
		os.Exit(1);
    }
}

func (fileHandler * FileHandler)LoadNextChunk() []string {
	var readedLines = []string{};
	for i :=0; i < MAX_LINES_READ_BUFFER_SIZE; i++ {
		if fileHandler.Scanner.Scan(){
			nextLine := fileHandler.Scanner.Text();
			readedLines = append(readedLines,nextLine);
			continue;
		}
	    break;
	}
	return readedLines
}
