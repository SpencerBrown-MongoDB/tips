package main

import (
	"fmt"
	"regexp"
)

func main() {
	// rxStringDoesntDoPerl := `^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}(?<!-)([\w]{0,42}))$`
	rxString := `^([a-zA-Z0-9]([a-zA-Z0-9-]){0,21}([\w]{0,42}))$`
	rx := regexp.MustCompile(rxString)
	fmt.Println(rx.MatchString("my-cluster"))
	fmt.Println(rx.MatchString("my-my-clusterwithaveryvery-veryljkfdjkfdjkfdjkfdjkfdjkfdjkfdjfkdjfkdjfkdjfdkjfdkjfkdjfdkjfdkjfdkjkfdjkfdjkfd"))
	matches := rx.FindStringSubmatch("my-cluster-hello-there-x-x-x-x-x")
	fmt.Printf("%d matches:", len(matches))
	for _, match := range matches {
		fmt.Println(match)
	}
}
