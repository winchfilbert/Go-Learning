package handler;

import "fmt"

func ErrorHandler(err error){
	if err != nil{
		fmt.Println("no value found");
		return 
	}

}