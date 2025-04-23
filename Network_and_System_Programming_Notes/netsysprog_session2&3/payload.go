package main

// the real background here, is that we want to explore the dynamic approach of
//allocating the size of buffer like we did in the method 1

// what will happen then ? it's gonna ease the pain of updating, delete, initializing the buffer size

import (
	"bytes"
	"encoding/binary"
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
  String() string
}

/* 
this type Binary []byte equals to this in java, since go doesn't have any classes to begin with
the type, actually implements things like this;

public final class Binary {
    private final byte[] data;

    public Binary(byte[] data) {
        this.data = data;
    }
}
*/
type Binary []byte


// this method returns the Binary itself (the underlying []byte).
// `m` is just the receiver — like `this` in Java or `self` in Python.
func (m Binary) Bytes() []byte{ 
  return m
}

func (m Binary) String() string{
  return string(m)
}


// `WriteTo` takes the receiver `m` (type Binary), and writes it to an io.Writer `w`.
// This is Go's way of attaching methods to types without using classes.

// this one i guessed that this one takes from m Binary, need that parameter which w io.writer
// for receiver 
// and got mapped for int64 / error
func (m Binary) WriteTo(w io.Writer)( int64, error){
  // in here you input the connection, and then you encode it with binary.BigEndian
  // the binary.BigEndian -> to place the very biggest byte to the very first address ascending

  err := binary.Write(w, binary.BigEndian, BinaryType) // Write the payload type as 1 byte (BinaryType)

  if err != nil {
    return 0, err // it can be continue
  }

  err = binary.Write(w, binary.BigEndian, uint32(len(m))) //4 byte for length 

  if err != nil {
    return 0, err // it can be continue
  }

  n, err := w.Write(m)
  
  //1 byte (type) + 4 bytes (length) + data.
  return int64(n+5),err
}

func (m *Binary) ReadFrom(r io.Reader)(int64, error){
    var typ uint8 
    err := binary.Read(r, binary.BigEndian, &typ) // 1 byte

    if err != nil{
      return 0, err
  } 

        
    if typ != BinaryType{
        return 1, nil
    }

     var n uint32 = 1 // because the write data use uint32, then logically, u need to read as uint32


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

type String string

func (m String) Bytes() []byte{
  return []byte(m)
}

func (m String) String() string{
  return string(m)
}



// this one below is to handle if string dude, siad that the binary handling having an err
// `WriteTo` takes the receiver `m` (type String), and writes it to an io.Writer `w`.
// This is Go's way of attaching methods to types without using classes.

// this one i guessed that this one takes from m Binary, need that parameter which w io.writer
// for receiver 
// and got mapped for int64 / error
func (m String) WriteTo(w io.Writer)( int64, error){
  // in here you input the connection, and then you encode it with binary.BigEndian
  // the binary.BigEndian -> to place the very biggest byte to the very first address ascending

  err := binary.Write(w, binary.BigEndian, StringType) // Write the payload type as 1 byte (BinaryType)

  if err != nil {
    return 0, err // it can be continue
  }

  err = binary.Write(w, binary.BigEndian, uint32(len(m))) //4 byte for length 

  if err != nil {
    return 0, err // it can be continue
  }

  n, err := w.Write([]byte(m))
  
  //1 byte (type) + 4 bytes (length) + data.
  return int64(n+5),err
}


// u read not assign, so it needs that address to know which struct you reading on, getter shit
func (m *String) ReadFrom(r io.Reader)(int64, error){
    var typ uint8 
    err := binary.Read(r, binary.BigEndian, &typ) // 1 byte

    if err != nil{
      return 0, err
  } 

        
    if typ != StringType{
        return 1, nil
    }

    var n uint32 = 1 // because the write data use uint32, then logically, u need to read as uint32


  err = binary.Read(r, binary.BigEndian, &n) // 4 byte
  
  if err != nil{
    return 0, err
  }

  // so if we want to ReadFrom
  // we can just point to m, and use that m for alocating the byte size
  buf := make([]byte, n)
  o, err := r.Read(buf)     

  if err != nil{
    return 0, err
  }

  *m = String(buf)

  return int64(o+5),err
}


// we're gonna cook decode function below here
func decode(r io.Reader)(Payload, error){
  // finding the type
  var typ uint8
  err := binary.Read(r, binary.BigEndian, &typ) // only catching error


  if err != nil {
    return nil, err // because u know, see if the err is exist, it means no data retrieved right ?
  }

  var payload Payload

  switch(typ){
  case BinaryType:
    payload = new(Binary)
  case StringType:
    payload = new(String)
  }
  
  // to read multiple things once, you can use the MultiReader
  // this one below is to determine which type to use, the stringtype or the binarytype
  // notes : if you don't use the value, and maybe you just want to catch the err, just use the _
  _, err = payload.ReadFrom(io.MultiReader(bytes.NewReader([]byte{typ}),r))
  
  if err != nil{
    return nil, err
  }

  return payload, nil 
}
