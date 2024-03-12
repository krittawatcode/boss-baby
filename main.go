package main

import "fmt"

func main() {
	fmt.Println(bossBaby("SRSSRRR"))  // Outputs: Good boy
	fmt.Println(bossBaby("RSSRR"))    // Outputs: Bad boy
	fmt.Println(bossBaby("SSSRRRRS")) // Outputs: Bad boy
	fmt.Println(bossBaby("SSRR"))     // Outputs: Good boy
	fmt.Println(bossBaby("SRRSSR"))   // Outputs: Bad boy
}
