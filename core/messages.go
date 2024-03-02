package core

var MenuLogo = `
▄︻デաօʀɖʟɨֆȶ օքȶʍɨʐɛʀ═══━一

Version 1.0.0

`
var MenuMessage = `This tool aims to optimize word lists to generate much more accurate attacks. Using regex and specific logic, it generates a much cleaner list of words`

var HelpMessage = `
Usage: wlopt [wordlist path] [save result path] [flags] 

-c   to include passwords with capital letters
-l   to include passwords with lowercase letters
-s   to include passwords with the following special symbols: [!@#$%^&*(),.?\":{}|<>] 
-r   followed by the custom regex to filter passwords
-n   to include passwords with numbers
-max followed by an integer to include passwords with maximum length
-min followed by an integer to include passwords with minimum length

Example:

wlopt /usr/share/wordlists/rockyou.txt /usr/share/wordlists -c -l -n -min 5 -max 8 -r "test"

The command above filters passwords with:

- capital letter (-c) 
- with lowercase letters (-l)
- with numbers (-n) 
- with min length of 5  (-min 5)
- with max length of 8 (-max 8)
- with "test" word in text

and saving in /usr/share/wordlists

The optimized file will have the following name optimized-xxx.txt where xxx is the name of the input wordlist file

The file with all passwords that did not meet the filter will be written to the removed-xxx.txt where xxx is the name of the input wordlist file


======> Filtering for popular patterns <======

If you want, you can filter directly for some popular password patterns:

wlopt [wordlist path] [save result path]  - service

where service can be: 

- gmail 
- hotmail
- tiktok
- wifi
- ssh
- facebook
- instagram
- alibaba
- yahoo
- aws
- ibm
- windows
- linkedIn

Example: 

wlopt /usr/share/wordlists/rockyou.txt /usr/share/wordlists  -wifi

OBS: You can select only one service 

`
