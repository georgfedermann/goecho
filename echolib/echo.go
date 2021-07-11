package echolib

func ReadArgsIntoString(args []string) string {
	var output, separator string = "", ""
	for _, token := range args {
		output += separator + token
		separator = " "
	}
	return output
}
