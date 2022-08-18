## Aho-Corasick Automaton

Aho-Corasick Multi-pattern Matching Algorithm for Golang

Example:

``` go
package main

import (
	"fmt"
	ahocorasick "github.com/HansenH/Aho-Corasick"
)

func main() {
	dictionary := []string{"hello", "world", "世界", "google", "golang", "c++", "love", " ", "l", "!!"}
	s := "hello世界, hello google, hello world, I love golang!!!"
	ac := ahocorasick.NewACAutomaton(dictionary)
	fmt.Println(ac.FindAllIndex(s))
}
```

Output:
```
map[ :[12 18 26 32 39 41 46] !!:[53 54] golang:[47] google:[19] hello:[0 13 27] l:[42] love:[42] world:[33] 世界:[5]]
```

## LICENSE

MIT
