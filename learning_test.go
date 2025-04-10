package main;

import (
  "testing"
)
//session 3: we're learning some unit test here
// the *testing referring from library, so i think it's kinda that solid syntax mode
func TestAddNumber(t *testing.T){
  result := addNumber(3,4)
  answer := 7

  if(result != answer){
    // can be format as if using the Print and Printf function
    // (we're talking go programming language, so the method should be start's with capital)
    // the equal is to use the error and the errorf

    t.Error("the result isn't as expected")
    t.Errorf("the result output is %d, but the expected value should be %d", result, answer)
  } 
}
