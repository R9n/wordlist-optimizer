package structs

type ParseOptions struct{
	HasCapitalLetter 	  bool
	HasLowercaselLetter   bool
	HasSpecialSymbol 	  bool
	HasNumbers			  bool
	MinLength 			  int
	MaxLength  			  int
	SplitResultLength 	  int
	CustomRegex 		  string
	ServiceSetSearch      string
} 