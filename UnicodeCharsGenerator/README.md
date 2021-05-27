# UnicodeCharsGenerator
## Ever seen some of these weird looking Unicode characters on the Internet ? Time to find more )

# Usage
## There are 2 flags to specify - `-inline` and `-limit`
`-inline` flag takes an integer. If not negative - each line in the output file will contain specified number of characters
`-limit` flag takes an integer. If not negative - the output file will contain all Unicode characters up to specified one   

## Examples

### `-inline`
1. `./UnicodeCharsGenerator -inline=20` - the generated output file will **try** to contain 20 characters in one line
2. `./UnicodeCharsGenerator -inline=-1` - all characters will be put in one line
3. `./UnicodeCharsGenerator` - as in 2., all characters will be put in one line  

### `-limit`
1. `./UnicodeCharsGenerator -limit=20000` - the generated output file will contain all Unicode characters up to 20000 one
2. `./UnicodeCharsGenerator -limit=-727` - the limit is ignored and you will get **all 10FFFF** Unicode characters

### general
1. `./UnicodeCharsGenerator` - all Unicode characters, all in one line
2. `./UnicodeCharsGenerator -limit=900 -inline=5` - 900 first characters, only 5 characters on one line