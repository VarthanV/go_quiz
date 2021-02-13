package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type Problem struct{
	question string
	answer string
}

func main(){
	timeLimit:= ParseFlags()
	timer := time.NewTimer(time.Duration(timeLimit)* time.Second)
	lines,err := ReadCsv()
	if err != nil{
		fmt.Println("Error in reading the file....")
	}
	problems := ParseQuestions(lines)
	score := 0
	answerChannel := make(chan string)
	for i,prob := range problems{
		fmt.Printf("Problem %d  %s =", i+1, prob.question)
		go func(){
			var answer string
			fmt.Scanf("%s\n",&answer)
			answerChannel <- answer
		}()
		select{
		case <- timer.C:
			fmt.Printf("\nThe test is ended !! , Your score is %d\n",score)
			return
		case answer := <- answerChannel:	
			if answer == prob.answer{
				score+=1
			}
		}
	}
	fmt.Printf("Your score is %d \n",score)
}

func ReadCsv() (lines[][]string,err error){
	file, err := os.Open("problems.csv")
	if err != nil{
		return nil ,err
	}
	r := csv.NewReader(file)
	lines,err = r.ReadAll()
	return lines,nil
}

func ParseQuestions(lines[][]string) ([]Problem){
	resp := make([]Problem,len(lines))
	for i,line :=range(lines){
		resp[i]= Problem{
			question: line[0],
			answer: strings.TrimSpace(line[1]),
		}
	}
	return resp
}

func ParseFlags() (timeLimit int){
	timeVal:= flag.Int("time",30,"Time limit for the quiz")
	flag.Parse()
	return *timeVal
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}