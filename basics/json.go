package main // Defines the main package (entry point for executable Go programs)

import (
	"bytes"            // Provides buffers for storing bytes in memory
	"encoding/json/v2" // Converts Go values to and from JSON
	"fmt"              // Used to print output
	"strings"          // Provides utilities for working with strings
)

// Defines a struct representing a response.
// Without JSON tags, exported field names are used as JSON keys.
type response1 struct {
	Page   int      // Integer field
	Fruits []string // Slice (dynamic array) of strings
}

// Same struct, but with JSON tags to customize the JSON keys.
type response2 struct {
	Page   int      `json:"page"`   // "Page" becomes "page" in JSON
	Fruits []string `json:"fruits"` // "Fruits" becomes "fruits" in JSON
}

func main() {

	// ==============================
	// 1. MARSHALING BASIC DATA TYPES
	// ==============================

	// Marshal converts a Go value into JSON bytes.
	// It returns ([]byte, error).
	// The underscore (_) discards the error.

	bolB, _ := json.Marshal(true) // Convert boolean true to JSON
	fmt.Println(string(bolB))     // Convert bytes to string and print: true

	intB, _ := json.Marshal(1) // Convert integer 1 to JSON
	fmt.Println(string(intB))  // Output: 1

	fltB, _ := json.Marshal(2.34) // Convert floating-point number to JSON
	fmt.Println(string(fltB))     // Output: 2.34

	strB, _ := json.Marshal("gopher") // Convert Go string to JSON string
	fmt.Println(string(strB))         // Output: "gopher"

	// ===========================
	// 2. MARSHALING SLICES
	// ===========================

	// []string creates a slice containing strings.
	slcD := []string{"apple", "peach", "pear"}

	// Convert the Go slice into a JSON array.
	slcB, _ := json.Marshal(slcD)

	// Output: ["apple","peach","pear"]
	fmt.Println(string(slcB))

	// ===========================
	// 3. MARSHALING MAPS
	// ===========================

	// map[string]int means:
	// Keys are strings, values are integers.
	mapD := map[string]int{
		"apple":   5,
		"lettuce": 7,
	}

	// Convert the Go map into a JSON object.
	mapB, _ := json.Marshal(mapD)

	// Output: {"apple":5,"lettuce":7}
	fmt.Println(string(mapB))

	// ===========================
	// 4. MARSHALING STRUCTS
	// ===========================

	// &response1{} creates a new response1 struct
	// and returns a pointer to it.
	res1D := &response1{
		Page:   1,                                  // Set Page field to 1
		Fruits: []string{"apple", "peach", "pear"}, // Set Fruits slice
	}

	// Convert the struct into JSON.
	res1B, _ := json.Marshal(res1D)

	// Output: {"Page":1,"Fruits":["apple","peach","pear"]}
	// Notice the capitalized JSON keys.
	fmt.Println(string(res1B))

	// Create a pointer to a response2 struct.
	// This struct uses JSON tags.
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}

	// Marshal the struct using its JSON tag names.
	res2B, _ := json.Marshal(res2D)

	// Output: {"page":1,"fruits":["apple","peach","pear"]}
	// Notice the lowercase keys from the struct tags.
	fmt.Println(string(res2B))

	// ===========================
	// 5. UNMARSHALING JSON
	// ===========================

	// Backticks (`) create a raw string literal.
	// This lets us include double quotes without escaping them.
	//
	// []byte(...) converts the JSON string into a byte slice.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Declare a map with string keys and values of any type.
	//
	// any is an alias for interface{}.
	// Values could be strings, numbers, slices, etc.
	var dat map[string]any

	// Unmarshal converts JSON bytes into Go data.
	//
	// &dat passes a pointer so Unmarshal can populate dat.
	//
	// if err := ... declares err and checks it immediately.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err) // Stop execution if JSON decoding fails
	}

	// Print the decoded Go map.
	// Example: map[num:6.13 strs:[a b]]
	fmt.Println(dat)

	// ===========================
	// 6. TYPE ASSERTIONS
	// ===========================

	// dat["num"] retrieves the value stored under "num".
	//
	// Since dat is map[string]any, the value's static type is any.
	//
	// .(float64) asserts that the actual value is a float64.
	// JSON numbers decoded into any default to float64.
	num := dat["num"].(float64)

	fmt.Println(num) // Output: 6.13

	// dat["strs"] accesses the JSON array.
	//
	// .([]any) asserts that its concrete type is a slice of any.
	// Arrays decoded into any become []any.
	strs := dat["strs"].([]any)

	// strs[0] accesses the first element: "a".
	//
	// .(string) asserts that the element contains a string.
	str1 := strs[0].(string)

	fmt.Println(str1) // Output: a

	// ===========================
	// 7. UNMARSHALING INTO STRUCTS
	// ===========================

	// JSON represented as a raw string literal.
	str := `{"page": 1, "fruits": ["apple", "peach"]}`

	// Create an empty response2 struct.
	// Initially Page = 0 and Fruits = nil.
	res := response2{}

	// Convert str to []byte and decode into res.
	//
	// &res passes its address, allowing Unmarshal
	// to modify the original struct.
	//
	// The underscore discards the returned error.
	_ = json.Unmarshal([]byte(str), &res)

	// Print the populated struct.
	// Output: {1 [apple peach]}
	fmt.Println(res)

	// Access the first element of the Fruits slice.
	// Output: apple
	fmt.Println(res.Fruits[0])

	// ===========================
	// 8. WRITING JSON TO A BUFFER
	// ===========================

	// Create a map with string keys and integer values.
	d := map[string]int{"apple": 5, "lettuce": 7}

	// Create an empty bytes.Buffer.
	//
	// A buffer is temporary storage in memory
	// that can grow as data is written into it.
	var buf bytes.Buffer

	// MarshalWrite converts d to JSON and writes it to buf.
	//
	// &buf passes a pointer to the buffer.
	//
	// The buffer implements the io.Writer interface.
	_ = json.MarshalWrite(&buf, d)

	// Convert the buffer's bytes into a string and print.
	// Output: {"apple":5,"lettuce":7}
	fmt.Println(buf.String())

	// ===========================
	// 9. READING JSON FROM A READER
	// ===========================

	// Create another empty response2 struct.
	res1 := response2{}

	// strings.NewReader(str) creates an io.Reader
	// that reads data from the JSON string.
	//
	// json.UnmarshalRead reads JSON from that reader
	// and decodes it into the struct pointed to by &res1.
	_ = json.UnmarshalRead(strings.NewReader(str), &res1)

	// Print the populated struct.
	// Output: {1 [apple peach]}
	fmt.Println(res1)
}
