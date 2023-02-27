//lint:file-ignore ST1006 because I like my code rustic, and I name my parameters however I want

package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	// "io/ioutil" //deprecated as of Go 1.16
)

func main(){
	payload := []byte("Hello, if you have made it here. Email me");
	hashAndBroadcast(NewHashReader(payload));
}

type HashReader interface {
	io.Reader
	hash() string
}


//struct embedding => composability

type hashReader struct {
	*bytes.Reader
	buffer *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader : bytes.NewReader(b),
		buffer : bytes.NewBuffer(b),

	};
}

func (self *hashReader) hash() string {
	return hex.EncodeToString(self.buffer.Bytes());
}

func hashAndBroadcast(reader HashReader) error {

	hash := reader.hash()
	fmt.Println(hash);
	// broadcast will fail, because 
	return broadcast(reader);
}

func broadcast(reader io.Reader) error {
	bytes, err := io.ReadAll(reader);

	if err != nil {
		return err;
	}

	fmt.Println("string of the bytes", string(bytes))
	return nil;
}


//! bad 
// func hashAndBroadcast(reader io.Reader) error {
// 	//reader is being read, it gets empty
// 	bytes, err := ioutil.ReadAll(reader);

// 	if err != nil {
// 		return err;
// 	}

// 	hash := sha1.Sum(bytes);
// 	fmt.Println(hex.EncodeToString(hash[:]));
// 	// broadcast will fail, because 
// 	return broadcast(reader);
// }

// func broadcast(reader io.Reader) error {
// 	bytes, err := ioutil.ReadAll(reader);

// 	if err != nil {
// 		return err;
// 	}

// 	fmt.Println("string of the bytes", string(bytes))
// 	return nil;
// }