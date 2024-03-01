package structs

import (
	"fmt"
	"os"
	"regexp"
	"unicode"
)

type SearchRegexes struct{
	HasCapitalLetterRegex 	   string
	HasLowercaselLetterRegex   string
	HasSpecialSymbolRegex  	   string
	HasNumbersRegex 		   string
	MinLengthRegex  	       string
	MaxLengthRegex   		   string
	CustomRegex 		       string
	ChooseGmailRegex           string
	ChooseTiktokRegex          string
	ChooseHotmailRegex         string
	ChooseWifiRegex            string
	ChooseSshRegex             string
	ChooseFacebookRegex        string
	ChooseInstagramRegex       string
	ChooseAlibabaRegex         string
	ChooseYahooRegex           string
	ChooseAWSRegex             string
	ChooseIBMRegex             string
	ChooseWindowsRegex         string
	ParseOptions               *ParseOptions
}

func (parser *SearchRegexes)Init(options *ParseOptions) *SearchRegexes {

    var regexes = SearchRegexes {
		HasCapitalLetterRegex 	: "[A-Z]",
		HasLowercaselLetterRegex: "[a-z]",
		HasSpecialSymbolRegex  	: "[!@#$%^&*(),.?\":{}|<>]",
		HasNumbersRegex 		: "\\d",
		MinLengthRegex  	    : fmt.Sprintf(".{%d}", options.MinLength),
		MaxLengthRegex		    : fmt.Sprintf("^.{%d}$", options.MaxLength),
		CustomRegex 			: options.CustomRegex,
		ChooseGmailRegex        : "-",
		ChooseTiktokRegex       : "-",
		ChooseHotmailRegex      : "-",
		ChooseYahooRegex        : "-",
		ChooseAWSRegex          : "-",
		ChooseIBMRegex          : "-",
		ChooseWifiRegex         : "-",
		ChooseSshRegex          : "-",
		ChooseFacebookRegex     : "-",
		ChooseInstagramRegex    : "-",
		ChooseAlibabaRegex      : "-", 
		ChooseWindowsRegex      : "-",
	}

	regexes.ParseOptions = options;

	return &regexes;
}
// start

func(entity *SearchRegexes) CheckSshPassword(password string) bool {
	if len(password) < 4  {
		return false
	}

	// Counters for different character types
	lowerCount := 0
	upperCount := 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			lowerCount++
		case unicode.IsUpper(char):
			upperCount++
		}
	}

	// Check if there are at least two types of characters
	typesCount := 0
	if lowerCount > 0 || upperCount >0  {
		typesCount++
	}

	return typesCount >= 1
}

func  (entity *SearchRegexes) CheckTikTokPassword(str string) bool {
	// Check the length of the string
	if len(str) < 8 || len(str) > 20 {
		return false
	}

	// Check if the string contains letters, numbers, and special characters
	hasLetter := false
	hasNumber := false
	hasSpecial := false

	for _, char := range str {
		switch {
		case (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z'):
			hasLetter = true
		case char >= '0' && char <= '9':
			hasNumber = true
		case regexp.MustCompile(`[[:punct:]]`).MatchString(string(char)):
			hasSpecial = true
		}
	}

	return hasLetter && hasNumber && hasSpecial
}

func(entity *SearchRegexes) CheckHotmailPassword(password string) bool {
	if (len(password) < 8) || (len(password) > 127) {
		return false
	}

	// Counters for different character types
	upperCount := 0
	lowerCount := 0
	digitCount := 0
	symbolCount := 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCount++
	
		case unicode.IsLower(char):
			lowerCount++
		
		case unicode.IsDigit(char):
			digitCount++
		
		case unicode.IsSymbol(char), unicode.IsPunct(char):
			symbolCount++
			
		}
	}

	// Check if there are at least two types of characters
	typesCount := 0
	if upperCount > 0 {
		typesCount++
	}
	if lowerCount > 0 {
		typesCount++
	}

	if digitCount > 0 {
		typesCount++
	}
	if symbolCount > 0 {
		typesCount++
	}

	return typesCount >= 2
}

func(entity *SearchRegexes) CheckYahooPassword(password string) bool {
	if (len(password) < 8) || (len(password) > 128) {
		return false
	}

	// Counters for different character types
	upperCount := 0
	lowerCount := 0
	digitCount := 0
	symbolCount := 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCount++
		case unicode.IsLower(char):
			lowerCount++
		case unicode.IsDigit(char):
			digitCount++
		case unicode.IsSymbol(char), unicode.IsPunct(char):
			symbolCount++
		}
	}

	// Check if there are at least two types of characters
	typesCount := 0
	if upperCount > 0 {
		typesCount++
	}
	if lowerCount > 0 {
		typesCount++
	}
	if digitCount > 0 {
		typesCount++
	}
	if symbolCount > 0 {
		typesCount++
	}

	if typesCount >= 2 {
		return true;
	} 
	if digitCount == 0 &&
	   symbolCount == 0 && 
	   upperCount == 0 || lowerCount == 0 &&
	   len(password) > 8{
		return true;
	}
  return false
}

func(entity *SearchRegexes) CheckAWSPassword(password string) bool {
	if (len(password) < 8) || (len(password) > 128) {
		return false
	}

	// Counters for different character types
	upperCount := 0
	lowerCount := 0
	digitCount := 0
	symbolCount := 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCount++
		case unicode.IsLower(char):
			lowerCount++
		case unicode.IsDigit(char):
			digitCount++
		case unicode.IsSymbol(char), unicode.IsPunct(char):
			symbolCount++
		}
	}

	// Check if there are at least 3 types of characters
	typesCount := 0
	if upperCount > 0 {
		typesCount++
	}
	if lowerCount > 0 {
		typesCount++
	}
	if digitCount > 0 {
		typesCount++
	}
	if symbolCount > 0 {
		typesCount++
	}

	return typesCount >= 3
}

func(entity *SearchRegexes) CheckWifiPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true;
}

func(entity *SearchRegexes) CheckGmailPassword(password string) bool {
	if len(password) < 8 || len(password) > 100 {
		return false
	}
	return true;
}

func(entity *SearchRegexes) CheckInstagramPassword(password string) bool {
	if len(password) < 6 || len(password) > 255 {
		return false
	}
	return true;
}

func(entity *SearchRegexes) CheckFacebookPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true;
}

func(entity *SearchRegexes) CheckIBMPassword(password string) bool {
	if (len(password) < 12) || (len(password) > 63) {
		return false
	}

	// Counters for different character types
	upperCount := 0
	lowerCount := 0
	digitCount := 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCount++
		case unicode.IsLower(char):
			lowerCount++
		case unicode.IsDigit(char):
			digitCount++
		}
	}

	// Check if there are at least two types of characters
	typesCount := 0
	if upperCount > 0 {
		typesCount++
	}
	if lowerCount > 0 {
		typesCount++
	}
	if digitCount > 0 {
		typesCount++
	}

	return typesCount >= 3
}

func(entity *SearchRegexes) CheckWindowsPassword(password string) bool {
	if (len(password) < 4) {
		return false
	}
	return true;
}

func(entity *SearchRegexes) CheckAlibabaPassword(password string) bool {
	if (len(password) < 6) || (len(password) > 20) {
		return false
	}

	// Counters for different character types
	upperCount := 0
	lowerCount := 0
	digitCount := 0
	symbolCount:= 0

	// Loop through each character in the password
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			upperCount++
		case unicode.IsLower(char):
			lowerCount++
		case unicode.IsDigit(char):
			digitCount++
		case unicode.IsSymbol(char):
			symbolCount++
			break
		case unicode.IsSpace(char):
			return false;
		}
	}

	// Check if there are at least two types of characters
	typesCount := 0

	if upperCount > 0 || lowerCount > 0 {
		typesCount++
	}

	if digitCount > 0 {
		typesCount++
	}
	if symbolCount > 0 {
		typesCount++
	}

	return typesCount >= 3
}

func (entity *SearchRegexes) MatchDefaultServices(password string) bool {
	
	switch entity.ParseOptions.ServiceSetSearch {
    case "-gmail":
        return entity.CheckGmailPassword(password);
    case "-hotmail":
        return entity.CheckHotmailPassword(password);
    case "-tiktok":
		return entity.CheckTikTokPassword(password);
    case "-wifi":
        return entity.CheckWifiPassword(password);
    case "-ssh":
        return entity.CheckSshPassword(password);
    case "-facebook":
        return entity.CheckFacebookPassword(password);
    case "-instagram":
        return entity.CheckInstagramPassword(password);
    case "-alibaba":
        return entity.CheckAlibabaPassword(password);
    case "-yahoo":
        return entity.CheckYahooPassword(password);
    case "-aws":
        return entity.CheckAWSPassword(password);
    case "-ibm":
        return entity.CheckIBMPassword(password);
    case "-windows":
        return entity.CheckWindowsPassword(password);
    case "-linkedin":
        return entity.CheckLinkenidPassword(password);
    default:
        fmt.Println("Unknown service");
		os.Exit(1)
		return false
    }	
}

func(entity *SearchRegexes) CheckLinkenidPassword(password string) bool {
	if len(password) < 6 ||len(password) > 200  {
		return false
	}
	return true;
}


// end

func (entity *SearchRegexes) MatchCustomRegex(password string) bool {
	found, _ := regexp.MatchString(entity.CustomRegex, password);
	return found;
}

func (entity *SearchRegexes) MatchMaxlength(password string) bool {
	found, _ := regexp.MatchString(entity.MaxLengthRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchMinlength(password string) bool {
	
	found, _ := regexp.MatchString(entity.MinLengthRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchNumbers(password string) bool {

	found, _ := regexp.MatchString(entity.HasNumbersRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchSpecialSymbol(password string) bool {

	found, _ := regexp.MatchString(entity.HasSpecialSymbolRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchCapitalLetter(password string) bool {

	found, _ := regexp.MatchString(entity.HasCapitalLetterRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchLowerCaseLetter(password string) bool {

	found, _ := regexp.MatchString(entity.HasLowercaselLetterRegex,password);
	return found;
}

func (entity *SearchRegexes) MatchFilters(password string) bool{

    if entity.ParseOptions.HasCapitalLetter && !entity.MatchCapitalLetter(password){
       return false;
    }

    if entity.ParseOptions.HasLowercaselLetter && !entity.MatchLowerCaseLetter(password) {
		return false;
    }
	
    if entity.ParseOptions.HasSpecialSymbol  && !entity.MatchSpecialSymbol(password){
		return false;
    }
	
	if entity.ParseOptions.HasNumbers && !entity.MatchNumbers(password) {
        return false;
    }
	
	if entity.ParseOptions.MinLength != 0 && !entity.MatchMinlength(password)  {
		return false;
	}
	
	if entity.ParseOptions.MaxLength != 0  && !entity.MatchMaxlength(password){
		return false;
	}

	if entity.ParseOptions.CustomRegex != "" && !entity.MatchCustomRegex(password)  {
		return false;
    }
	
	return true;
}
