# **go-jwt-cracker**

[![Golang](https://img.shields.io/static/v1?label=code&message=Golang&color=blue)](https://www.python.org/) ![License](https://img.shields.io/github/license/Evilran/go-jwt-cracker?style=plastic)

        _  _      _____      ____ ____  ____  ____ _  __ _____ ____
       / |/ \  /|/__ __\    /   _Y  __\/  _ \/   _Y |/ //  __//  __\
       | || |  ||  / \_____ |  / |  \/|| / \||  / |   / |  \  |  \/|
    /\_| || |/\||  | |\____\|  \_|    /| |-|||  \_|   \ |  /_ |    /
    \____/\_/  \|  \_/      \____|_/\_\\_/ \|\____|_|\_\\____\\_/\_\
    
                                                          -- Evi1ran

## Compile

```
$ go build -o bin/jwtcrack src/main/main.go
```

## Run

```
$ ./jwtcrack -t eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.cAOIAifu3fykvhkHpbuhbvtH807-Z2rI1FS3vX1XMjE -l 4
```

In the above example, the key is `Sn1f`. 

Usage
---

```
Usage: jwtcrack -t <token> [-a alphabet] [-l max_len]

Options:
  -a alphabet
    	set alphabet of secret (default "eariotnslcudpmhgbfywkvxzjqEARIOTNSLCUDPMHGBFYWKVXZJQ0123456789")
  -h	this help
  -l max length
    	set max length of secret (default 6)
  -t token
    	set token
```

## Important Things

No advantage at all, this program is not fast, and does not use multi-threading. But it can construct a dictionary through the alphabet itself (will not miss a possibility) for brute force.

## TODO

- Goroutine