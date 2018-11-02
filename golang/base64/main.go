package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	input := `aHR0cHM6Ly9tbWJpei5xcGljLmNuL21tYml6X2pwZy9vcTFQeW1SbDlEN1JkdUd0MHM1ZExTMUZRVkw1YXlNbkRqaWF0SWljdWljOURmaWNTZzZCRWphcXVtRW5UbE14cTV2ODZrUDdIZGljUVNtMFB5UXRZRGo1RWFRLzY0MD93eF9mbXQ9anBlZw==`
	uDec, _ := base64.URLEncoding.DecodeString(input)

	fmt.Println(string(uDec))
}
