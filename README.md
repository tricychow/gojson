# gojson
gojson包提供json数据打平成key-value结构，和key-value结构组合成json结构。

Js2kv

```
package main

import (
	"fmt"
	"github.com/tricychow/gojson"
)
const s = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
func main(){
	x, _ := gojson.Js2kv(s, "_")
}
// result:{"age":47,"name_first":"Janet","name_last":"Prichard"}
```

Kv2js

```
package main

import (
	"fmt"
	"github.com/tricychow/gojson"
)
const s = `{"age":47,"name_first":"Janet","name_last":"Prichard"}`
func main(){
	x, _ := gojson.Kv2js(s, "_")
	fmt.Println(x)
}
// result:{"age":47,"name":{"first":"Janet","last":"Prichard"}}
```

