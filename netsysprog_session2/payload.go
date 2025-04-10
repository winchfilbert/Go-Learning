package main

// the real background here, is that we want to explore the dynamic approach of
//allocating the size of buffer like we did in the method 1

// what will happen then ? it's gonna ease the pain of updating, delete, initializing the buffer size

import (
	"encoding/binary"
	"fmt"
	"io"
)

// making constants -> plural const variable
const (

  // this below code takes long to type, and however, you can fasten this up with iota
  // BinaryType uint8 = 1
  // StringType uint8 = 2
  // ...
  //...



  // said that iota is to auto increment stuff, iota starts with 0, so you need to add up with 1
  BinaryType uint8 = iota + 1 
  StringType // you don't have to define it again, the iota do the stuff 
)

type Payload interface{
  io.ReaderFrom
  io.WriterTo
  Bytes() []byte
  String() []string
}

type Binary []byte

// to make me not confused, .... idk the m binary ..... this function named Bytes, with 0 parameters, that returns Binary
func (m Binary) Bytes() Binary{ 
  return m
}

func (m Binary) String() string{
  return string(m)
}

// somehow this code line it's too english, by getting the m which is the type Binary
// we can write and define the purpose and w which is from io.writer be that pointer for purpose


// this one i guessed that this one takes from m Binary, need that parameter which w io.writer
// for receiver 
// and got mapped for int64 / error


func (m Binary) WriteTo(w io.Writer)( int64, error){
  // in here you input the connection, and then you encode it with binary.BigEndian
  // the binary.BigEndian -> to place the very biggest byte to the very first address ascending
  err := binary.Write(w, binary.BigEndian, 1) // 1 byte for the size

  if err != nil {
    return 0, err // it can be continue
  }

  error := binary.Write(w, binary.BigEndian, uint32(len(m))) //4 byte / size

  if error != nil {
    return 0, err // it can be continue
  }

  n, err := w.Write(m)

  return int64(n+5),err
}

func (m *Binary) ReadFrom(r io.Reader)(int64, error){
    var typ uint8 
    err := binary.Read(r, binary.BigEndian, &typ) // 1 byte

    if err != nil{
      return 0, err
  } 

    var n int64 = 1
    
    if typ != BinaryType{
        return n, nil
    }

  err = binary.Read(r, binary.BigEndian, &n) // 4 byte
  
  if err != nil{
    return 0, err
  }

  // so if we want to ReadFrom
  // we can just point to m, and use that m for alocating the byte size
  *m = make([]byte, n)

  o, err := r.Read(*m)

  return int64(o+5),err
}
