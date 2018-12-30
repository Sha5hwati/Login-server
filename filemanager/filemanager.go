/*****************
Authors - Shashwati Shradha, Manasi Paste, Garfield Tong
The file handling the loading and saving data into json files
 */

package filemanager

import (
	//allows manipulation of bytes
	"bytes"

	//Allows synchronization
	"sync"

	//handles json files
	"encoding/json"

	//handles input and output
	"io"

	//a platform-independent interface to operating system functionality
	"os"
)

/*
type Manager struct
A structure storing the locks to make
sure that there is no conflict which loading and
saving data. The data is package private */
type Manager struct {
   lock sync.Mutex
}

/*
Marshal( v interface{}) (io.Reader, error)
Marshal is a function that marshals the object into an
io.Reader using JSON Marshaller.
Parameter - v - an empty interface that handles values of
            unknown types
Return - io.Reader -  data from the reader
       error - returns anytype of errors
*/
var Marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}


/*
Unmarshal (r io.Reader, v interface{}) error
Unmarshal is a function that unmarshals the data from the
reader into the specified value using the JSON unmarshaller.
Parameter - r - data from the reader
         v - an empty interface that handles values of
            unknown types
Return - error - returns anytype of errors
*/
var Unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}


/*
func (manager Manager) Save(path string, v interface{}) error
Save saves a representation of v to the file at path.
Parameter - path - path of the file to be saved rel to
            project root directory
         v - an empty interface that handles values of
            unknown types
Return - error - returns anytype of errors
*/
func (manager Manager) Save(path string, v interface{}) error {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r, err := Marshal(v)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}


/*
func (manager Manager) Load(path string, v interface{}) error
Load loads the file at path into v.
Use os.IsNotExist() to see if the returned error is due
to the file being missing.
Parameter - path - path of the file to be loaded rel to
            project root directory
         v - an empty interface that handles values of
            unknown types
Return - error - returns anytype of errors
*/
func (manager Manager) Load(path string, v interface{}) error {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return Unmarshal(f, v)
}

