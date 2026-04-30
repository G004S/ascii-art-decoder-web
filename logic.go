package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Decode(input string) (string, error) {
	var res string
	bracketsClosed := true
	var reps int
	var chars string
	for i := 0; i < len(input); i++ {
		if input[i] == ']' {
			return "", errors.New("ERROR: square brackets are unbalanced")
		}
		if input[i] == '[' {
			bracketsClosed = false
			//get the encoded piece
			var piece string
			for j := i; j < len(input); j++ {
				if (j == len(input) && !bracketsClosed) || (j != i && input[j] == '[') {
					return "", errors.New("ERROR: square brackets are unbalanced")
				}
				if input[j] == ']' {
					piece += string(input[j])
					bracketsClosed = true
					break
				}
				piece += string(input[j])
			}
			//get number of reps and the string that needs to be repeated
			var splitterSpace int
			for i, v := range piece {
				if v == ' ' {
					splitterSpace = i
					break
				} else if i+1 == len(piece) && v != ' ' {
					return "", errors.New("ERROR: space separator missing")
				}
			}
			var err error
			reps, err = strconv.Atoi(piece[1:splitterSpace])
			if err != nil {
				return "", errors.New("ERROR: repetition number faulty")
			}
			chars = piece[splitterSpace+1 : len(piece)-1]
			if chars == "" {
				return "", errors.New("ERROR: repeatable charachers missing")
			}

			//repeat the piece N times
			for k := 0; k < reps; k++ {
				res += chars
			}

			//shift the index forward to the length of the encoded piece
			i += len(piece) - 1
			continue

		}

		res += string(input[i])
	}

	return res, nil
}
 
func Encode(input string) (string, error) {
	var res string
	for i := 0; i < len(input); i++ {
		// check for pattern of 3 chars
		if i+6 <= len(input) {
			//if all 3 chars are the same:
			if input[i:i+3] == input[i+3:i+6] {
				reps := 1
				scope := input[i:i+3]
				//run the single char counter
				if input[i] == input[i+1] && input[i+1] == input[i+2] {
					reps, scope = SingleCharCounter(i, input)
					res += fmt.Sprintf("[%d %s]", reps, scope)
					i += reps-1
					continue
				}
				//run the triple char pattern counter
				matched:= false
				for j:= i+3; j+3 <= len(input); j += 3 {
					if input[j:j+3] == scope {
						reps++
						matched = true
						continue
					}else{
						break
					}
				}
				if matched{
					res += fmt.Sprintf("[%d %s]", reps, scope)
					i += reps*3-1
					continue
				}
			}
		}
		//check for pattern of 2 chars
		if i+5 <= len(input) {
			if input[i:i+2] == input[i+2:i+4] {
				//check if both chars are the same and run the single char counter
				if input[i] == input[i+1] {
					reps, scope := SingleCharCounter(i, input)
					res += fmt.Sprintf("[%d %s]", reps, scope)
					i += reps-1
					continue
				}
				reps := 1
				scope := input[i:i+2]
				matched := false
				for j := i+2; j+2<= len(input); j += 2 {
					if input[j:j+2] == scope {
						reps++
						matched = true
					}else{
						break
					}
				}
				if matched{
					res += fmt.Sprintf("[%d %s]", reps, scope)
					i += reps * 2-1
					continue
				}
			}
		}
		// check for repeating single chars
		if i+2 <= len(input) {
			if input[i] == input[i+1] {
				reps, scope := SingleCharCounter(i, input)
				res += fmt.Sprintf("[%d %s]", reps, scope)
				i += reps-1
				continue
			}
		}
		res+= string(input[i])
	}
	return res, nil
}

func SingleCharCounter(i int, input string) (reps int, char string) {
	rep := 1
	scop := input[i]
	for j:= i; j+1 < len(input); j++ {
		if input[j+1] == scop {
			rep++
			continue
		}
		break
	}
	
	return rep, string(scop)
}
