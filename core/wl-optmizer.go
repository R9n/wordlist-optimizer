package core

import (
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"wl-optimizer/core/structs"
)

type WlOptimizer struct{
 fileHandler    *structs.FileHandler;
 filter  		*structs.SearchRegexes;
}

func (wlOptimizer *WlOptimizer) Init(inputFilepath string, outputFilePath string, parsedOptions *structs.ParseOptions){
	
    filenName := filepath.Base(inputFilepath);

    wlOptimizer.fileHandler = &structs.FileHandler{
		InputFilepath:   inputFilepath,
		OutputFilePath:  outputFilePath + "/optimized-" + filenName,
        RemovedFilePath: outputFilePath + "/removed-" + filenName,
	}
    searchFilter := structs.SearchRegexes{}

	wlOptimizer.filter = searchFilter.Init(parsedOptions);
}

func (wlOptimizer *WlOptimizer) ProcessChunk(chunck []string, channel chan int){
    var removedPasswords = 0;

    if wlOptimizer.filter.ParseOptions.ServiceSetSearch != ""{
         for _, password := range chunck {
            if wlOptimizer.filter.MatchDefaultServices(password) {
                wlOptimizer.fileHandler.WriteLineToOutputFile(password);
            }else{
                removedPasswords++;
                wlOptimizer.fileHandler.WriteLineToRemovedFile(password);
            }
        }
	}else{
        for _, password := range chunck {
            if wlOptimizer.filter.MatchFilters(password){
                wlOptimizer.fileHandler.WriteLineToOutputFile(password);
            }else{
                removedPasswords++;
                wlOptimizer.fileHandler.WriteLineToRemovedFile(password);
            }
        }
    }

	channel <- removedPasswords;
}

func calcAlignSpaces(strings ...string) int {
    maxLen := 0
    for _, s := range strings {
        if len(s) > maxLen {
            maxLen = len(s)
        }
    }
    return maxLen;
}

func fillTabAlign(stringToFill string, length int ) string{
    if length < 0 {
        return "";
    }
    return strings.Repeat(stringToFill, length)
}

func (wlOptimizer *WlOptimizer) PrintStatistics(totalPasswords int, removedPasswords int, startTime time.Time, endTime time.Time) {
    var duration = endTime.Sub(startTime)
    reductionPercentage := 0.0
    if totalPasswords > 0 {
        reductionPercentage = (float64(removedPasswords) / float64(totalPasswords)) * 100
    }

    var durationString = duration.String();
    var totalPasswordsString = strconv.Itoa(totalPasswords);
    var totalremovedPasswordsString = strconv.Itoa(removedPasswords);
    var totalMatchedPasswordsString = strconv.Itoa(totalPasswords-removedPasswords)
    var reductionPercentageString = strconv.FormatFloat(reductionPercentage, 'f', 2, 64) + "%";

   var maxAlignLength = calcAlignSpaces(durationString,
    totalPasswordsString,totalremovedPasswordsString,
    totalMatchedPasswordsString,reductionPercentageString);

    PrintGreen("\n======> ✅ Finished ✅ <======\n\n")
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("|             Statistics:    "+ fillTabAlign(" ",maxAlignLength +1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("| Total Time:                | " + durationString  + fillTabAlign(" ", maxAlignLength - len(durationString)-1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("| Total Processed Password:  | " + totalPasswordsString + fillTabAlign(" ", maxAlignLength - len(totalPasswordsString)-1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("| Total Removed Passwords:   | " + totalremovedPasswordsString + fillTabAlign(" ", maxAlignLength - len(totalremovedPasswordsString)-1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("| Matched Passwords:         | " + totalMatchedPasswordsString + fillTabAlign(" ", maxAlignLength - len(totalMatchedPasswordsString)-1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
    PrintCyan("| Reduction of:              | " + reductionPercentageString + fillTabAlign(" ", maxAlignLength - len(reductionPercentageString)-1) + "|");
    PrintCyan("+----------------------------"+ fillTabAlign("-",maxAlignLength +1) + "+");
}



func (wlOptimizer *WlOptimizer) OptimizeWordlist() {
    PrintProgramHeader();
    PrintGreen("\n 1 - Opening file: " + wlOptimizer.fileHandler.InputFilepath );
    wlOptimizer.fileHandler.OpenFile();

    PrintGreen("\n 2 - Creating output file: " + wlOptimizer.fileHandler.OutputFilePath );
    wlOptimizer.fileHandler.CreateOutputFile();

    PrintGreen("\n 3 - Creating removed passwords file: " + wlOptimizer.fileHandler.RemovedFilePath );
    wlOptimizer.fileHandler.CreateRemovedFile();

    PrintCyan("\n Processing 🚀... Please, await ⏰ \n");

    PrintLoading();

    
    var wg sync.WaitGroup;
    channel := make(chan int);


    var startTime = time.Now();
    var totalProcessedPasswords = 0;
    var totalRemovedPasswords = 0;

    for {
        nextChunk := wlOptimizer.fileHandler.LoadNextChunk();
        var nextChunkLength = len(nextChunk);
        if  nextChunkLength == 0 {
            break;
        }
        wg.Add(1);
        totalProcessedPasswords += nextChunkLength;
        go func(chunk []string) {
            defer wg.Done();
            wlOptimizer.ProcessChunk(chunk, channel);
        }(nextChunk);
    }

    go func() {
        wg.Wait();
        close(channel);
    }()

    for removedPass := range channel {
        totalRemovedPasswords +=removedPass;
    }

    var endTime = time.Now();

    wlOptimizer.PrintStatistics(totalProcessedPasswords,totalRemovedPasswords,startTime,endTime);

    wlOptimizer.fileHandler.CloseInputFile();
    wlOptimizer.fileHandler.CloseOutputFile();
    wlOptimizer.fileHandler.CloseRemovedFile();
}





