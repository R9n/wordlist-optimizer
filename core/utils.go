package core

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"wl-optimizer/core/structs"

	"github.com/fatih/color"
)

func PrintWhite(message string){
	println(message);
}

func PrintRed(message string){
	color.Red(message);
}

func PrintBlue(message string){
	color.Blue(message);
}

func PrintGreen(message string){
	color.Green(message);
}

func PrintCyan(message string){
	color.Cyan(message);
}

func PrintYellow(message string){
	color.Yellow(message);
}

func FileExists(path string) bool {
    _, err := os.Stat(path)
    if err == nil {
        return true
    }
    if os.IsNotExist(err) {
        return false 
    }
    fmt.Println("Erro ao verificar a existência do arquivo:", err)
    return false
}

func convertToNumber(arg string)(int,error){
	num, err := strconv.Atoi(arg)
  	return num, err
}

func Contains(arr []string, target string) bool {
    for _, element := range arr {
        if element == target {
            return true
        }
    }
    return false
}

func ParseArgs(args []string) (*structs.ParseOptions, string, string){

	if len(args) == 3{
		PrintRed("Malformed Command Error!  At last one flag must be provided");
		os.Exit(1);
	  }
	  
	var maxArgsIndex = len((args)) - 1
	var executablePathIndex = 0;
	var resultTargetPathIndex = 2;
	var wordlistPathindex = 1;
	var canJumpNextInfo = false;
	var defaultServicePatternsIndex = 3;
	var defaultServices = []string{"-linkedin","-gmail", "-hotmail", "-tiktok", "-wifi", "-ssh", "-facebook", "-instagram", "-alibaba", "-yahoo", "-aws", "-ibm", "-windows"}

	var searchOptions = structs.ParseOptions{
		HasCapitalLetter: false,
		HasLowercaselLetter: false,
		HasNumbers: false,
		HasSpecialSymbol: false,
		MinLength: 0,
		MaxLength: 0,
		SplitResultLength: 0,
		CustomRegex: "",
	}

	for index,flag := range args{
		if canJumpNextInfo == true {
			canJumpNextInfo = false;
			continue;
		}

		if index == executablePathIndex || index == resultTargetPathIndex {
			continue;
		}

		if index == 1 {
			wordlistPath := args[1]
			if(FileExists(wordlistPath)) == true {
				continue;
			}
			PrintRed("The file "+wordlistPath+ " does not exists" )
			os.Exit(1)
		}

		if Contains(defaultServices,args[defaultServicePatternsIndex]){
			ParseServiceFilterFlag(index, args, maxArgsIndex, &searchOptions);
			break;
		}

		switch flag {
		case "-c":
			ParseCFlag(index, args, maxArgsIndex, &searchOptions)
		case "-l":
			ParseLFlag(index, args, maxArgsIndex, &searchOptions)
		case "-n":
			ParseNFlag(index, args, maxArgsIndex, &searchOptions)
		case "-r":
			ParseRegexFlag(index, args, maxArgsIndex, &searchOptions)
			canJumpNextInfo = true;
		case "-s":
			ParseSFlag(index, args, maxArgsIndex, &searchOptions)
		case "-min":
			ParseMinFlag(index, args, maxArgsIndex, &searchOptions)
			canJumpNextInfo = true;
		case "-max":
			ParseMaxFlag(index, args, maxArgsIndex, &searchOptions)
			canJumpNextInfo = true;
		default:
			PrintRed("Malformed Command Error! Invalid flag: " + flag)
			os.Exit(1)
		}
	}

	return &searchOptions, args[wordlistPathindex], args[resultTargetPathIndex];
}

func PrintHelp(){
	PrintRed(MenuLogo);
	PrintGreen(MenuMessage);
	PrintWhite(HelpMessage);
}
func PrintProgramHeader(){
	PrintRed(MenuLogo);
	PrintGreen(MenuMessage);
}

func PrintLoading() {
	chars := []rune{'|', '/', '-', '\\'}
	go func (){
	for {
		for _, char := range chars {
			fmt.Printf("\r%c", char)
			time.Sleep(100 * time.Millisecond)
		}
    }
	}()
}


