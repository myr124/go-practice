
package main // Defines the main package for an executable Go program

import (
    "bufio"         // Provides buffered reading and writing
    "fmt"           // Used to print formatted output
    "io"            // Provides utilities and interfaces for input/output
    "os"            // Provides functions for interacting with the operating system
    "path/filepath" // Provides utilities for working with file paths
)

// Helper function for handling errors.
// Takes an error as input and panics if the error is not nil.
func check(e error) {
    if e != nil { // Check whether an error occurred
        panic(e) // Stop normal execution and report the error
    }
}

func main() {

    // ===================================
    // 1. READING AN ENTIRE FILE
    // ===================================

    // os.TempDir() returns the operating system's temporary directory.
    // filepath.Join() combines path components using the correct
    // separator for the operating system.
    //
    // Example on macOS: /var/folders/.../T/dat
    // Example on Linux: /tmp/dat
    path := filepath.Join(os.TempDir(), "dat")

    // ReadFile reads the entire file into memory.
    //
    // Returns:
    // dat = []byte containing the file's contents
    // err = error if the file cannot be read
    dat, err := os.ReadFile(path)

    check(err) // Panic if ReadFile failed

    // Convert the file's bytes into a string and print them.
    fmt.Print(string(dat))

    // ===================================
    // 2. OPENING A FILE FOR READING
    // ===================================

    // Open returns a *os.File representing an open file.
    //
    // Unlike ReadFile(), this doesn't load the entire file.
    // It lets us read the file incrementally.
    f, err := os.Open(path)

    check(err) // Check whether opening the file failed

    // ===================================
    // 3. READING THE FIRST 5 BYTES
    // ===================================

    // make([]byte, 5) creates a byte slice with length 5.
    //
    // Initially contains five zero-valued bytes.
    // This is our buffer where file data will be stored.
    b1 := make([]byte, 5)

    // Read attempts to read up to 5 bytes from the file.
    //
    // It writes the data directly into b1.
    //
    // Returns:
    // n1 = number of bytes actually read
    // err = error, possibly io.EOF if at the end of the file
    //
    // Reading also advances the file's current position.
    n1, err := f.Read(b1)

    check(err) // Handle any returned error

    // b1[:n1] selects only the bytes actually read.
    //
    // For example, if n1 = 3, we take b1[0:3].
    //
    // %d prints an integer; %s prints a string.
    fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

    // ===================================
    // 4. MOVING TO A SPECIFIC POSITION
    // ===================================

    // Seek changes the current position in the file.
    //
    // io.SeekStart means the offset is relative to
    // the beginning of the file.
    //
    // Seek(6, io.SeekStart) moves to byte offset 6.
    // Offsets begin at 0, so this is the 7th byte.
    //
    // o2 = resulting absolute file offset
    o2, err := f.Seek(6, io.SeekStart)

    check(err) // Check whether seeking failed

    // Create a buffer that can hold 2 bytes.
    b2 := make([]byte, 2)

    // Read up to 2 bytes starting from the current
    // file position (offset 6).
    n2, err := f.Read(b2)

    check(err) // Check for reading errors

    // Print the number of bytes and starting offset.
    //
    // No newline yet because we're using Printf
    // without \n at the end.
    fmt.Printf("%d bytes @ %d: ", n2, o2)

    // Print only the bytes actually read.
    // %v uses Go's default representation for the value.
    fmt.Printf("%v\n", string(b2[:n2]))

    // ===================================
    // 5. RELATIVE SEEKING
    // ===================================

    // io.SeekCurrent means the offset is relative to
    // the file's CURRENT position.
    //
    // Seek(2, io.SeekCurrent) moves forward 2 bytes.
    //
    // The underscore discards the resulting offset.
    _, err = f.Seek(2, io.SeekCurrent)

    check(err) // Check for seeking errors

    // io.SeekEnd means the offset is relative to
    // the END of the file.
    //
    // -4 moves the position to 4 bytes before EOF.
    //
    // Note: Moving the position does not read data.
    _, err = f.Seek(-4, io.SeekEnd)

    check(err) // Check for seeking errors

    // ===================================
    // 6. READING AN EXACT MINIMUM NUMBER
    // ===================================

    // Move back to byte offset 6 from the start.
    //
    // o3 stores the resulting position.
    o3, err := f.Seek(6, io.SeekStart)

    check(err) // Check for seeking errors

    // Allocate a buffer that holds 2 bytes.
    b3 := make([]byte, 2)

    // io.ReadAtLeast attempts to read until AT LEAST
    // the specified minimum number of bytes is obtained.
    //
    // Arguments:
    // f  = source to read from
    // b3 = destination buffer
    // 2  = minimum required number of bytes
    //
    // Returns:
    // n3  = number of bytes read
    // err = error if the minimum cannot be satisfied
    //
    // Because b3 has length 2 and minimum is 2,
    // this requires exactly 2 bytes to succeed.
    n3, err := io.ReadAtLeast(f, b3, 2)

    check(err) // Panic if fewer than 2 bytes are available

    // Print the number of bytes, starting offset,
    // and contents of the buffer.
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // ===================================
    // 7. RESETTING THE FILE POSITION
    // ===================================

    // Move the file's position back to the beginning.
    //
    // This is useful when we want to reread data.
    _, err = f.Seek(0, io.SeekStart)

    check(err) // Check for seeking errors

    // ===================================
    // 8. BUFFERED READING AND PEEKING
    // ===================================

    // Create a buffered reader around the file.
    //
    // bufio.NewReader(f) lets us read through an
    // internal memory buffer instead of making
    // a separate underlying read for every small request.
    r4 := bufio.NewReader(f)

    // Peek returns the next 5 bytes without advancing
    // the buffered reader's logical position.
    //
    // Important:
    // Peek may read ahead from the underlying file,
    // so f's actual file offset can still advance.
    //
    // b4 references the reader's internal buffer,
    // so it should not be kept after later reads.
    b4, err := r4.Peek(5)

    check(err) // Check whether peeking failed

    // Convert the peeked bytes into a string and print.
    fmt.Printf("5 bytes: %s\n", string(b4))

    // ===================================
    // 9. CLOSING THE FILE
    // ===================================

    // Release the open file's operating system resources.
    f.Close()
}
