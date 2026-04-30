package main
/*import (
		"bufio"
		"strings"
		"os"
		"errors"
)
func GetInputMulti() (string, error){
	var input string
	scanner:= bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := scanner.Text()
		if line == strings.ToLower("end"){
			return string(input), nil
		}else{
			input+= "\n" + line
		}
	}

	return string(input), nil
}


func GetInputFile() (string, error){
	input,err := os.ReadFile("./library/"+ os.Args[1])
	if err != nil {
		return "", errors.New("ERROR: couldn't open the input file")
	}
	return string(input), nil
} */