package main

// Similar parenthesis based grouping is also allowed for "var" and "const" statements
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

// - Semicolons are optional at end of statements
// - Any identified begining with a capital letter gets automatically exported

func swap(a, b string) (string, string) {
	return b, a
}

func variableDeclarations() {
	// Ways of defining integers
	i := 0 // can only be used within functions, not at file scope, cannot be used for consts
	var j = 1
	var k int
	var l int = 1

	fmt.Println("Variables:", i, j, k, l)

	// Basic variable types are as follows:
	// bool, string, float32, float64, complex64, complex128
	// byte (alias for uint8), rune (alias for int32)
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr

	var (
		isInit       bool   = true
		errorMessage string = "Do not exit"
		tcpPort      int16  = 8080
	)

	fmt.Printf("Variables:Group: %T %v, %T %v, %T %v\n", isInit, isInit, errorMessage, errorMessage, tcpPort, tcpPort)

	// Type conversions are explicit
	num := 10
	numf := float32(num)
	numu := uint(numf)
	const pi = 3.14

	fmt.Printf("Variables: %T %v, %T %v, %T %v, %T %v\n", num, num, numf, numf, numu, numu, pi, pi)

	return
}

func functionsInGo() {
	// Functions cannot be nested but they can be assigned to variables.
	// Functions can be used before they are defined, within the same file.

	// Function with multiple return values (unnamed)
	var swap = func(x, y string) (string, string) {
		return y, x
	}
	var first, second = swap("hello", "world")
	fmt.Println("Functions:Swap:", first, second)

	// Function with multiple return values (named)
	var square = func(a int, b int) (x int, y int) {
		x = a * a
		y = b * b
		return
	}
	var numA, numB = square(4, 5)
	fmt.Println("Functions:Square:", numA, numB)

	// Callback functions
	var callCallback = func(f func(string, string) (string, string)) {
		a, b := f("one", "two")
		fmt.Println("Functions:callback", a, b)
	}
	callCallback(swap)

	// Variadic parameters. Here "nums" is available as a []int slice.
	var sumOfNums = func(nums ...int) int {
		total := 0
		for _, num := range nums {
			total += num
		}
		return total
	}
	fmt.Println("Functions:sumofnums", sumOfNums(10, 20, 30, 40, 50))
	sampleInputs := []int{100, 200, 300, 400, 500}
	fmt.Println("Functions:sumofnums", sumOfNums(sampleInputs...))
	sampleInputsArray := [...]int{100, 200, 300, 400, 500}
	sampleInputs2 := sampleInputsArray[:]
	fmt.Println("Functions:sumofnums", sumOfNums(sampleInputs2[1:]...))
}

func controlFlow() {
	fmt.Printf("Ctrlflow:")
	// Variable has local scope, braces are mandatory
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	// The for loop is also the while loop as semicolons are optional
	// Omiting the loop condition creates an infinite loop: for { ... }
	j := 10
	for j < 20 {
		fmt.Printf("%d", j)
		j++
	}
	fmt.Println()

	// The break and continue keywords also take an optional label to specify
	// which loop to break.
	fmt.Print("CtrlFlow:labels ")
J:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i > 5 {
				break J
			}
		}
		fmt.Print(i)
	}
	fmt.Println()

	// "if" also supports a variable initialition similar to "for"
	if v := math.Pow(3, 2); v < 10 {
		fmt.Println("Ctrlflow:If less than 10:", v)
	} else {
		// "v" is also within scope of the "else" block
		fmt.Println("Ctrlflow:If greater than 10:", v)
	}

	// Switch also has an optional initialization, no break needed, cases must be const
	// Case values can be function calls.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Ctrlflow:OS X")
	case "linux":
		fmt.Println("Ctrlflow:OS Linux")
	case "some other OS", "yes another OS": // Multiple cases can be separated by a comma
		fallthrough // Use fallthrough explicitly to fall through to next case
	default:
		fmt.Println("Ctrlflow:", os)
	}

	// The switch condition can be skipped to write long if/else chains
	t := time.Now()
	switch {
	case t.Hour() >= 0 && t.Hour() < 12:
		fmt.Println("Ctrlflow:Good morning")
	case t.Hour() >= 12 && t.Hour() < 17:
		fmt.Println("Ctrlflow:Good afternoon")
	default:
		fmt.Println("Ctrlflow:Good evening")
	}
}

func deferredFunctions() {
	// A defer keyword before a function defers the function call until the surrounding
	// function returns. And if there are multiple functions then they are called
	// in a reverse (stack based) order.
	fmt.Println("Defer:Start")
	for i := 0; i < 5; i++ {
		defer fmt.Println("Defer:Count", i)
	}
	fmt.Println("Defer:End")

	// A deferred function can also modify named return values
	var doSomething = func() (ret int) {
		defer func() { ret = 20 }()
		// Following return value will be overwritten by the deferred call above
		return 1
	}
	x := doSomething()
	fmt.Println("Defer:return value modified", x)
}

func panicAndRecover() {
	// panic(): is a built-in function that stops the ordinary flow of control and begins panicking.
	// When the function F calls panic, execution of F stops, any deferred functions in F are
	// executed normally, and then F returns to its caller. To the caller, F then behaves like a
	// call to panic. The process continues up the stack until all functions in the current goroutine
	// have returned, at which point the program crashes. Panics can be initiated by invoking panic
	// directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.
	//
	// recover(): is a built-in function that regains control of a panicking goroutine. Recover is
	// only useful inside deferred functions. During normal execution, a call to recover will return
	// nil and have no other effect. If the current goroutine is panicking, a call to recover will
	// capture the value given to panic and resume normal execution.

	// Checks whether a given function will panic
	willPanic := func(f func()) (b bool) {
		defer func() {
			if value := recover(); value != nil {
				b = true
			}
		}()
		f()
		return
	}

	panicyFunc := func() {
		var nums []int
		nums[4] = 10
	}

	fmt.Println("Panic:willPanic", willPanic(panicyFunc))
}

func pointersInGo() {
	i, j := 10, 2189
	ptrToI := &i
	*ptrToI += 10

	var ptrToJ *int = &j
	*ptrToJ = *ptrToJ / 17

	fmt.Println("Pointers:", i, j)

	// Pointer aritematic is not allowed in Go
}

func structDataType() {
	type Vertex struct {
		// Field names starting with uppercase are accessible from other packages
		// and those starting with lowercase are only accessible in the current package.
		X, Y int
	}

	var v1 = Vertex{1, 2}
	v2 := Vertex{X: 3} // Y is zero implicitly
	v3 := &Vertex{3, 4}

	// Pointers also use dot notation for struct members (v3.X)
	fmt.Printf("Struct:Vertex: v3.X=%d, %v, %v\n", v3.X, v1, v2)

	type Student struct {
		// These are anonymous field names. Their name is the same as their type.
		int
		string
		// Having another "string" field in this struct would be an error due to a name conflict.
	}

	student1 := Student{int: 10, string: "Fred"}
	fmt.Println("Struct:Student", student1)
}

func arrayDataType() {
	var myStrings [2]string
	myStrings[0] = "hello"
	myStrings[1] = "world"
	fmt.Println("Arrays:", myStrings)

	// You can also skip the array size by using ... instead
	primes := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Arrays:", primes)

	// Slices are a window into an array which can grow (have separate len and cap)
	var slicedPrimes []int = primes[1:4]
	fmt.Println("Arrays:Sliced:", slicedPrimes)
	slicedPrimes[0] = 100
	fmt.Println("Arrays:", primes)

	// Note: arrays are always passed by value to functions and slices
	// are always passed by reference. Assigning one array to another copies
	// its values and assigning one slice to another makes both slices
	// point to the same underlying array.

	cars := [4]string{
		"Toyota",
		"Honda",
		"Ford",
		"Suzuki", // A comma is always needed at the end of each item, even the last one
	}
	fmt.Println("Arrays:cars:", cars)

	// A slice literal first creates an array and then its slice
	carsSlice := []string{"Toyota", "Vitz", "Nissan", "Prez"}
	fmt.Println("Array:carslice:", carsSlice)

	carFactory := []struct {
		model int
		make  string
	}{
		{2019, "Toyota"},
		{2000, "Honda"},
		{1995, "Suzuki"},
	}

	fmt.Println("Arrays:carfactory:", carFactory)
}

func moreOnSlices() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("Slices:s1:", s1, "Cap:", cap(s1), "Len:", len(s1))
	s2 := s1[2:6]
	fmt.Println("Slices:s2:", s2, "Cap:", cap(s2), "Len:", len(s2))
	// s2[4] = 20 -- this will cause a panic

	// This will not extend the length of the slice
	s2 = s2[0:]
	fmt.Println("Slices:s2:", s2, "Cap:", cap(s2), "Len:", len(s2))

	// Extend the length of the slice explicitly to maximum capacity
	// which is equivalent to s2 = s2[:6]
	s2 = s2[:cap(s2)]
	fmt.Println("Slices:s2:", s2, "Cap:", cap(s2), "Len:", len(s2))

	// The zero value of a slice is nil (zero len zero cap). This cannot
	// be defined using the short syntax: s3 := []int
	var s3 []int
	if s3 == nil {
		fmt.Println("Slices:s3 is nil")
	}

	// Dynamically sized arrays can be creating by using the make(type, len, cap)
	// function to create slices.
	dynSlice1 := make([]int, 5, 10)
	fmt.Println("Slices:dynSlice:", dynSlice1, "Cap:", cap(dynSlice1), "Len:", len(dynSlice1))

	// Slices of slices (multidimensional)
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[0][2] = "O"
	board[1][1] = "X"
	board[2][0] = "O"
	board[1][2] = "X"
	board[2][2] = "O"
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], " "))
	}

	// The append(slice, values...) function allows appending to a slice and
	// grows the slice if required.
	var growingSlice []int
	growingSlice = append(growingSlice, 10, 11, 12)
	fmt.Println("Slices:growingSlice", growingSlice, cap(growingSlice))
	growingSlice = append(growingSlice, 20, 30, 40)
	fmt.Println("Slices:growingSlice", growingSlice, cap(growingSlice))

	// The copy function allows copying one slice to another.
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s4 = make([]int, 6)
	copylen := copy(s4, a[0:])
	fmt.Println("Slices:copy1:", copylen, s4) // prints 6 [0 1 2 3 4 5]
	copylen = copy(s4, s4[2:])
	fmt.Println("Slices:copy2:", copylen, s4) // prints 4 [2 3 4 5 4 5]
}

func mapDataType() {
	monthnames := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31, "Apr": 30,
		"May": 31, "Jun": 30, "Jul": 31, "Aug": 31,
		"Sep": 30, "Oct": 31, "Nov": 30, "Dec": 31, // This last comman is required
	}

	// Looping over key/value pairs using range
	daysinyear := 0
	for _, days := range monthnames {
		daysinyear += days
	}
	fmt.Println("Map:daysinyear", daysinyear, "Days in Feb", monthnames["Feb"])

	// Inserting/deleting values
	monthnames["January"] = 31
	delete(monthnames, "Jan")

	// Checking for existence
	value, ok := monthnames["Jan"]
	fmt.Printf("Map:Jan exists: value:%v ok:%v\n", value, ok)
	value, ok = monthnames["January"]
	fmt.Printf("Map:January exists: value:%v ok:%v\n", value, ok)

	// A map which returns functions
	incFunc := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		3: func() int { return 30 },
	}
	fmt.Println("Map:incFunc", incFunc[1]())
}

func main() {
	fmt.Println("Hello world", rand.Intn(100))

	variableDeclarations()
	functionsInGo()
	controlFlow()
	deferredFunctions()
	panicAndRecover()
	pointersInGo()
	structDataType()
	arrayDataType()
	moreOnSlices()
	mapDataType()
}
