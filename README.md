## Ragel state machine based User Agent detection for golang
Faster than go regular expressions

```golang
package ua
```
The package provides the below API
```golang
import (
  userAgent "github.com/DronRathore/github-rage-ua/src/ua"
)
browser, version := userAgent.UaVersion(string)
var isBot bool = userAgent.IsBot(string)
isMobile, platform := userAgent.IsMobile(string)
```
__Note__: While compiling for production environment you can optimise the ragel output by running the build.sh with ```$GOENV=production``` and it will output the fully optimised ragel state machine.
```sh
# compile optimised ragel first
GOENV=production ./src/github.com/DronRathore/go-ragel-ua/build.sh
```