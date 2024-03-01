package core

import (
	"os"
	"wl-optimizer/core/structs"
)

func ParseCFlag(index int, args []string, maxArgsIndex int, parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		parseOption.HasCapitalLetter = true;
		return
	}
	nextIndex := index + 1;
	nextFlag  := args[nextIndex];
    cFollowSet := []string{"-max","-min","-sp","-l","-s","-r","-n"};
	if Contains(cFollowSet,nextFlag) == false {
		PrintRed("Malformed Command Error! the -c flag can only be followed by -max, -min, -sp, -l, -s, -r or -n")
		os.Exit(1)
	}
	parseOption.HasCapitalLetter = true;
}

func ParseLFlag(index int, args []string, maxArgsIndex int, parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		parseOption.HasLowercaselLetter = true;
		return
	}
	nextIndex := index + 1;
	nextFlag  := args[nextIndex];
    cFollowSet := []string{"-max","-min","-sp","-c","-s","-r","-n"};
	if Contains(cFollowSet,nextFlag) == false {
		PrintRed("Malformed Command Error! the -l flag can only be followed by -max, -min, -sp, -l, -s, -r or -n")
		os.Exit(1)
	}
	parseOption.HasLowercaselLetter = true;
}

func ParseSFlag(index int, args []string, maxArgsIndex int, parseOption * structs.ParseOptions){
	if index == maxArgsIndex {
		parseOption.HasSpecialSymbol = true
		return
	}
	nextIndex := index + 1;
	nextFlag  := args[nextIndex];
    cFollowSet := []string{"-max","-min","-sp","-c","-l","-r","-n"};
	if Contains(cFollowSet,nextFlag) == false {
		PrintRed("Malformed Command Error! the -s flag can only be followed by -max, -min, -sp, -l, -c, -r or -n")
		os.Exit(1)
	}
	parseOption.HasSpecialSymbol = true
}

func ParseRegexFlag(index int, args []string, maxArgsIndex int, parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		PrintRed("Malformed Command Error! The -r flag must be folowed by the custom regex to filter passwords")
		os.Exit(1)
	}
	nextIndex := index + 1;
	regex := args[nextIndex];
	parseOption.CustomRegex = regex;
}

func ParseNFlag(index int, args []string, maxArgsIndex int, parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		parseOption.HasNumbers = true;
		return
	}
	nextIndex := index + 1;
	nextFlag  := args[nextIndex];
    cFollowSet := []string{"-max","-min","-sp","-c","-s","-r","-l"};
	if Contains(cFollowSet,nextFlag) == false {
		PrintRed("Malformed Command Error! the -n flag can only be followed by -max, -min, -sp, -l, -s, -r or -c")
		os.Exit(1)
	}
	parseOption.HasNumbers = true;
}

func ParseMinFlag(index int, args []string, maxArgsIndex int,parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		PrintRed("Malformed Command Error! The -min flag must be folowed by a valid integer")
		os.Exit(1)
	}
	nextIndex := index + 1;
	number,err := convertToNumber(args[nextIndex]);

	if err != nil{
		PrintRed("Malformed Command Error! The -min flag must be folowed by a valid integer")
		os.Exit(1)
	}

   if number < 1{
		PrintRed("Malformed Command Error! the -min flag must be followed by an interger greater or equal than 1")
		os.Exit(1)
	}

	parseOption.MinLength = number;
}

func ParseMaxFlag(index int, args []string, maxArgsIndex int,parseOption *structs.ParseOptions){
	if index == maxArgsIndex {
		PrintRed("Malformed Command Error! The -max flag must be folowed by the min length to filter")
		os.Exit(1)
	}
	nextIndex := index + 1;
	number,err := convertToNumber(args[nextIndex]);
	
	if err != nil {
		PrintRed("Malformed Command Error! the -max flag must be followed by an interger")
		os.Exit(1)
	}

   if number < 1 {
		PrintRed("Malformed Command Error! the -max flag must be followed by an interger greater or equal than 1")
		os.Exit(1)
	}
	parseOption.MaxLength = number;
}

func ParseServiceFilterFlag(index int, args []string, maxArgsIndex int, parseOption *structs.ParseOptions){

   if len(args) > 4 {
		PrintRed("Malformed Command Error! You can only select one service when filter for default services")
		os.Exit(1)
	}

	parseOption.ServiceSetSearch = args[index];
}