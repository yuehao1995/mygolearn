package main

import "fmt"

type ProviderType int

// The ProviderType of a member relative to the member API
const (
	FABRIC ProviderType = iota // MSP is of FABRIC type
	IDEMIX                     // MSP is of IDEMIX type
	OTHER                      // MSP is of OTHER TYPE

	// NOTE: as new types are added to this set,
	// the mspTypes array below must be extended
)
func main(){
	fmt.Println(FABRIC)
	fmt.Println(IDEMIX)
	fmt.Println(OTHER)
}
