package main

// Similar parenthesis based grouping is also allowed for "var" and "const" statements
import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// - Semicolons are optional at end of statements
// - Any identifier begining with a capital letter gets exported automatically
// - Sequence of evaluation: first all the imported packages are initialized, then
//   all variable definitions are evaluated and finally the "func init()" function
//   within the source file is evaluated. Each source file can have one (or more)
//   init() function.

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
	//
	// Other common slice types: []byte, []rune, []int

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

	str := "this is a test string"
	str = `this is a raw string without escape interpretation "\n"`
	// String to bytes
	byteStr := []byte(str)
	runeStr := []rune(str)
	fmt.Println("Variables:str", str)
	fmt.Println("Variables:byte_str", byteStr)
	fmt.Println("Variables:rune_str", runeStr)

	// Bytes to string after reversing the string
	for i, j := 0, len(byteStr)-1; i < j; i, j = i+1, j-1 {
		byteStr[i], byteStr[j] = byteStr[j], byteStr[i]
	}
	str = string(byteStr)
	fmt.Println("Variables:reversed", str)
	return
}

func functionsInGo() {
	// Functions cannot be nested but they can be assigned to variables.
	// Methods cannot be assigned to variables and so cannot be defined within a function.
	// Functions can be used before they are defined, within the same file.

	// Following are some cases in which functions can be nested and they can form
	// closures as well:
	// - A function is defined and assigned to a local variable within a parent function.
	// - The "go" keyword is used to execute a goroutine which is defined in-place within
	//   the parent function.
	// - The return statement is used to return a function to the caller, defined in place.

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
		// The blank identifier can be used to ignore a function's returned value.
		// It can also be used to import packages for side-effects only, i.e. for
		// invoking the package's init() function only, such as follows:
		//
		// import _ "net/http/pprof"
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
	// Case values can be function calls. And default can come first.
	switch os := runtime.GOOS; os {
	case "some other OS", "yes another OS": // Multiple cases can be separated by a comma
		fallthrough // Use fallthrough explicitly to fall through to next case
	default:
		fmt.Println("Ctrlflow:", os)
	case "darwin":
		fmt.Println("Ctrlflow:OS X")
	case "linux":
		fmt.Println("Ctrlflow:OS Linux")
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

	// Another example if switch similar to if
	shouldEscape := false
	ch := 'a'
	switch ch {
	case '?', ' ', '&', '=', '+', '%':
		shouldEscape = true
	}
	fmt.Println("Ctrlflow:shouldEscape", shouldEscape)

	// A Type Switch statement can be used to check the type
	var g interface{}
	g = "sample text"
	switch g.(type) {
	case int:
		fmt.Println("Ctrlflow:is int")
	case *int:
		fmt.Println("Ctrlflow:is *int")
	case string:
		fmt.Println("Ctrlflow:is string")
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

	fmt.Println("Pointers:", i, j, ptrToI, ptrToJ, *ptrToI, *ptrToJ)

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

	// Annotate struct with field names using %+v
	// Print a value in Go syntax using %#v
	// Print a quoted string using %q
	fmt.Printf("Struct:Vertex:v1 %+v %#v %q\n", v1, v1, "what")

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
	// Also note that "ok" can be defined multiple times in a multi-variable
	// assignment provided that at least one variable is newly defined and
	// that the type of the previous "ok" definition identical with the earlier
	// one. If this is done in an inner scope then the outer "ok" would be
	// shadowed by a newly created inner scoped "ok".
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

func makeAndNew() {
	// Remember that new(Type) returns a pointer to the zeroed value of Type.
	// And make(Type, len, cap) is used to allocate slices, maps and channels ONLY,
	// and returns an initialized (not zero'ed) Type itself instead of a pointer to it.

	var n *[]int = new([]int) // A nil slice is of not much use
	var m []int = make([]int, 10, 50)
	m[0] = 1
	m[1] = 2
	m[2] = 3

	*n = make([]int, 10) // We're making it unneccessarily complex. Idomatic: x := make([]int, 10)
	(*n)[0] = 10         // The brackets around (*n) are required.

	fmt.Println("Allocation:", n, m)
}

func constructorsInGo() {
	type StudentInfo struct {
		name string
		age  int
	}

	NewStudentInfo := func(name string, age int) *StudentInfo {
		if name == "" {
			return nil
		}
		// Below statement is equivalent to this:
		// var student *StudentInfo = new(StudentInfo)
		// student.name = name
		// student.age = age
		// return student
		return &StudentInfo{name: name, age: age}

		// Note in the above statement that it is valid in Go to return the address of
		// a local variable and this method of using composite literals is preferred
		// instead of using new().
	}

	var student *StudentInfo = NewStudentInfo("Fred", 10)
	fmt.Println("Constructors:", student)
}

func concurrencyAndChannels() {
	// A channel has a specific type or can be the empty interface for generic type.
	// Use make(chan int, 10) to create a channel which can buffer 10 items.
	// The default is 0, which means the sender blocks until the receiver receives.
	var numberChan chan int = make(chan int)

	waitAndPrint := func(str string, seconds int) {
		time.Sleep(time.Duration(seconds) * time.Second)
		fmt.Println("Concurrent:", str, "is ready", seconds)
		// Write to the channel
		numberChan <- len(str)
	}

	go waitAndPrint("Tea", 2)
	go waitAndPrint("Coffee", 1)

	fmt.Println("Concurrent: Waiting for tea/coffee")
	var bytesWritten int
	bytesWritten = <-numberChan
	bytesWritten += <-numberChan
	fmt.Println("Concurrent: Bytes sent", bytesWritten)
}

func moreOnChannels() {
	// - Channels have a specific type which can even be struct or the generic interface.
	// - When passing a channel to a function as a parameter, use these types:
	//   - func abc(<-chan int): For the read side of the channel.
	//   - func abc(chan<- int): For the write side of the channel.
	// - A reader can use select to monitor multiple channels.

	// Design patterns:
	//
	// - Infinite writer goroutine: A goroutine may write infinitely to a channel. The
	//   reader can read until it prefers and when the program terminates, it would
	//   cleanly exit despite the infinite goroutine.
	//
	// - Infinite writer goroutine with termination: Pass the goroutine another bool
	//   channel on which the caller will send a quit signal when he wants to terminate
	//   the infinite goroutine. The goroutine uses select to monitor the quit channel.
	//
	// - Finite writer goroutine: The goroutine writes data to the channel and then uses
	//   close(outchan) to close the channel. The reader can use either "range inchan" to
	//   read all values, or can use "value, ok := <-inchan" where "ok" would be false
	//   when the channel closes.
	//
	// - Async functions usually first create a channel which they return to the caller
	//   and then they run a goroutine which operates on the channel.
	//
	// - A channel can be used like a counting  semaphore due to its blocking nature:
	//   sem <- 1
	//   critical-section
	//   <- sem
	//
	// - Concurrency (structuring a program as independently executing components) is
	//   different from parallelism (executing calculations in parallel for efficiency
	//   on multiple CPUs). Go can handle both but is primarily a concurrent language.

	// Below is an example of Fibonacci using channels

	// dup3 duplicates the given channel into three separate channels.
	dup3 := func(in <-chan int) (<-chan int, <-chan int, <-chan int) {
		a, b, c := make(chan int, 2), make(chan int, 2), make(chan int, 2)
		// goroutine to duplicate data on incoming channel to the three new channels.
		go func() {
			for {
				x, ok := <-in
				if !ok {
					close(a)
					close(b)
					close(c)
					break
				}
				a <- x
				b <- x
				c <- x
			}
		}()
		return a, b, c
	}

	fib := func(count int) <-chan int {
		x := make(chan int, 2)
		a, b, out := dup3(x)
		go func() {
			x <- 0
			x <- 1
			<-b
			for ; count > 1; count = count - 1 {
				x <- <-a + <-b
			}
			close(x)
		}()
		return out
	}

	fmt.Print("Channels:")
	for num := range fib(7) {
		fmt.Print(num, " ")
	}
	fmt.Println()

	// Below is an example of a server/client model for handling concurrent requests.

	maxCPU := runtime.GOMAXPROCS(0)

	type Request struct {
		args        []int
		f           func([]int) int
		resultsChan chan int
	}

	requestHandler := func(inQueue <-chan *Request) {
		fmt.Println("***STARTED WORKER")
		for req := range inQueue {
			req.resultsChan <- req.f(req.args)
		}
		fmt.Println("***STOPPING WORKER")
	}

	serve := func(inQueue <-chan *Request, quit chan bool) {
		for i := 0; i < maxCPU; i++ {
			go requestHandler(inQueue)
		}
		fmt.Println("***WAITING QUIT")
		<-quit
	}

	sum := func(a []int) (x int) {
		for _, v := range a {
			x += v
		}
		return
	}

	queue := make(chan *Request, maxCPU)
	quit := make(chan bool)
	go serve(queue, quit)

	for i := 0; i < maxCPU; i++ {
		req := &Request{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, sum, make(chan int)}
		queue <- req
		fmt.Println("Channels:concurrent sum:", <-req.resultsChan)
	}

	quit <- true
	close(queue)
	time.Sleep(time.Duration(2) * time.Second)
}

func typeSwitchAndTypeAssertion() {
	// - Methods in Go can be defined for any custom type and not just structs.
	// - An interface is an abstract set of methods expected to be implemented by
	//   concrete types.

	// Using "type switch" to extract a generic value.
	type Stringer interface {
		String() string
	}
	var value interface{} = "hello"
	switch str := value.(type) {
	case string:
		fmt.Println("Types:string", str)
	case Stringer:
		fmt.Println("Types:stringer", str.String())
	}

	// Another possibility is to perform a direct conversion using a type assertion
	str, ok := value.(string)
	if ok {
		fmt.Println("Types:Conversion success", str)
	} else {
		fmt.Println("Types:Conversion failure")
	}
}

type Incrementer interface {
	Increment() int
}
type Decrementer interface {
	Decrement() int
}

// Directly embedding one interface into another causes all the child's methods to
// be inherited by the parent. See also struct-embedding below.
type IncrementerDecrementer interface {
	Incrementer
	Decrementer
}
type Counter int
type CompatibleCounter int

// Implement Incrementer and Decrementer interfaces for *Counter
func (ctr *Counter) Increment() int {
	*ctr++
	return int(*ctr)
}
func (ctr *Counter) Decrement() int {
	*ctr--
	return int(*ctr)
}

func methodsAndInterfaces() {
	// The Counter type implicitly satisfies the Incrementer interface.
	var ctrCounter Counter
	ctrCounter.Increment()

	// A *Counter type can be assigned directly to a variable of Incrementer type.
	// The compiler also verifies that Counter implements this interface.
	var ctrCounterPtr *Counter = new(Counter)
	var ctrIncrementer Incrementer = ctrCounterPtr
	ctrIncrementer.Increment()

	fmt.Println("Interfaces:ctrCounter", ctrCounter)
	fmt.Println("Interfaces:ctrCounterPtr", *ctrCounterPtr)
	fmt.Println("Interfaces:ctrIncrementer", ctrIncrementer)

	// Most interface conversions in Golang are explicit and are thus checked at
	// compile time. But sometimes the check is implicit. For example, the
	// encoding/json package takes a value which can optionally implement a
	// Marshaler interface. The package happily accepts values even if they don't
	// implement the interface. In such cases if the user intends to write a type
	// which explicitly implements an interface, then he can explicitly request
	// compiler type checking using the following construct involving the blank
	// identifier.
	//
	// var _ json.Marshaler = (*RawMessage)(nil)
	//
	// This is basically converting a nil pointer to a RawMessage type pointer and
	// assigning it to a Marshaler interface variable, thus explicitly invoking
	// type checking.

	// Methods are functions associated with specific types or interfaces.
	// An interface is the defintion of a set of methods. An interface only defines
	// methods and does not define data/variables.

	// When defining methods of a type, the type is usually a pointer if it is a
	// concrete type such as an int, string, struct, etc. But if it is a type which
	// implicitly holds a reference to another object and which causes a pass by
	// reference behavior instead of a pass by value, then the type is used instead
	// of a pointer, for example, types such as slices, maps and channels are used
	// directly when specified as a receiver (the type associated with a method).

	// Methods can only be defined for locally defined types and not for built-in
	// types.

	// A struct-embedding example
	func() {
		type Job struct {
			Command string
			// When embedding one struct into another then if the name is skipped then
			// the parent struct automatically inherits all methods of the child
			// struct. This has an implicit field name of "Logger" which must be
			// initialized with a valid value when a variable of this type is created.
			// This way no intermeidate forwarding methods need to be defined.
			*log.Logger
		}

		newJob := func(command string, logger *log.Logger) *Job {
			return &Job{command, logger}
		}
		job := newJob("dir", log.New(os.Stdout, "Interfaces:Job:", log.Ldate))
		// The Println method is inherited from log.Logger.
		job.Println("Job created")
	}()
}

func communicationInGo() {
	cat := func(filename string) int {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		buf := make([]byte, 256)
		count := 0
		for {
			n, err := f.Read(buf)
			if n != 0 {
				//os.Stdout.Write(buf[:n])
				fmt.Print("Comm:cat:", string(buf[:n]))
				count += n
			}
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		return count
	}

	catBuffered := func(filename string) int {
		f, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		rd := bufio.NewReader(f)
		wr := bufio.NewWriter(os.Stdout)
		defer wr.Flush()

		count := 0
		for {
			line, err := rd.ReadString('\n')
			if line != "" {
				fmt.Print("Comm:catbuf:", line)
				count++
			}
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		return count
	}

	if len(os.Args) == 2 {
		cat(os.Args[1])
	} else {
		cat("/etc/hosts")
	}
	catBuffered("/etc/resolv.conf")

	runCommand := func(cmdname string, cmdargs ...string) {
		cmd := exec.Command(cmdname, cmdargs...)
		out, err := cmd.Output()
		if out != nil && len(out) != 0 {
			os.Stdout.Write(out)
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	runCommand("ls", "-l")

	// Networking can be used by importing the "net" package and using net.Dial:
	// conn, err := Dial("tcp", "192.0.32.10:80")

	httpGet := func(url string) (string, error) {
		r, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		data, err := ioutil.ReadAll(r.Body)
		dataStr := string(data)
		r.Body.Close()
		if err == nil {
			fmt.Println(dataStr)
		}
		return dataStr, err
	}

	httpGet("http://www.google.com/robots.txt")
}

func main() {
	fmt.Println("Hello world", rand.Intn(100))

	// Getting help:
	// $ go doc package-name/package-name symbol-name
	// $ go doc package-name.symbol-name
	//
	// List all go packages:
	// $ go list ...
	//
	// Source of built-in packages can be found here: /usr/share/go/src/

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
	makeAndNew()
	constructorsInGo()
	concurrencyAndChannels()
	moreOnChannels()
	typeSwitchAndTypeAssertion()
	methodsAndInterfaces()
	communicationInGo()
}
