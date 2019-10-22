/*
Name: main
Author: Evi1ran
Date Created: 2019/10/21
Description: Go-JWT-Cracker
*/

package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	h bool
	t string
	a string
	l int
)

type Jwt struct {
	Header           []byte
	Payload          []byte
	Signature        []byte
}

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&t, "t", "", "set `token`")
	flag.StringVar(&a, "a", "eariotnslcudpmhgbfywkvxzjqEARIOTNSLCUDPMHGBFYWKVXZJQ0123456789", "set `alphabet` of secret")
	flag.IntVar(&l, "l", 6, "set `max length` of secret")
	flag.Usage = usage
}

func computeHmac256(message string, secret string) []byte {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return h.Sum(nil)
}

/**
 * Check if the signature produced with "secret matches the original signature.
 * Return true if it matches, false otherwise
 */
func check(t Jwt, encrypt string, secret string) bool {
	return bytes.Equal(computeHmac256(encrypt, secret), t.Signature)
}

/**
 * Enumerate alphabet and generate a word list
 */
func enum(depth int, perm *[]string) {
	for i := 0; i <= len(a) - depth; i++ {
		result := make([]string, depth)
		result[0] = string(a[i])
		combine(result, i, 1, depth, perm)
	}
}

/**
 * Combine alphabet in word list
 */
func combine(result []string, rawIndex int, curIndex int, depth int, perm *[]string) {
	choice := len(a) - rawIndex + curIndex - depth
	for i := 0; i < choice; i++ {
		result[curIndex] = string(a[i+rawIndex+1])
		if curIndex + 1 == depth {
			num := strings.Join(result,"")
			*perm = append(*perm, num)
			continue
		} else {
			combine(result, rawIndex+i+1, curIndex+1, depth, perm)
		}
	}
}

/**
 * Create a permutation for signature
 */
func bruteSequential(a []byte, start int, buffer *[]string) {
	if start == len(a) {
		*buffer = append(*buffer, string(a))
	} else {
		for i := start; i < len(a); i++ {
			if i != start {
				tmp := a[start]
				a[start] = a[i]
				a[i] = tmp
			}
			bruteSequential(a, start + 1, buffer)
			if i != start {
				tmp := a[start]
				a[start] = a[i]
				a[i] = tmp
			}
		}
	}
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	if t == "" {
		_, _ = fmt.Fprintf(os.Stderr, `    _  _      _____      ____ ____  ____  ____ _  __ _____ ____ 
   / |/ \  /|/__ __\    /   _Y  __\/  _ \/   _Y |/ //  __//  __\
   | || |  ||  / \_____ |  / |  \/|| / \||  / |   / |  \  |  \/|
/\_| || |/\||  | |\____\|  \_|    /| |-|||  \_|   \ |  /_ |    /
\____/\_/  \|  \_/      \____|_/\_\\_/ \|\____|_|\_\\____\\_/\_\

                                                      -- Evi1ran
Usage: jwtcrack -t <token> [-a alphabet] [-l max_len]

`)
		return
	}

	if len(a) < l {
		fmt.Println("The length of alphabet must be bigger than max length!")
		return
	}

	// Split the JWT into header, payload and signature
	token := strings.Split(t, ".")
	tokenLen := len(token)
	if tokenLen != 3 {
		fmt.Println("It seems not like a token...")
		return
	}
	jwt := new(Jwt)
	jwt.Header, _ = base64.RawURLEncoding.DecodeString(token[0])
	jwt.Payload, _ = base64.RawURLEncoding.DecodeString(token[1])
	jwt.Signature, _ = base64.RawURLEncoding.DecodeString(token[2])
	// Recreate the part that is used to create the signature
	// Since it will always be the same
	encrypt := fmt.Sprintf("%s.%s", token[0], token[1])
	perm := make([]string, 0)
	buffer := make([]string, 0)
	for i := 1; i <= l; i++ {
		if i == 1 {
			for _, k := range a {
				perm = append(perm, string(k))
			}
		} else {
			enum(i, &perm)
		}
	}
	for _, v := range perm {
		bruteSequential([]byte(v), 0, &buffer)
	}

	for _, v := range buffer {
		if check(*jwt, encrypt, v) {
			fmt.Println("Secret is \"" + v + "\"")
			return
		}
	}
	fmt.Println("Not solution found :(")
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `go-jwt-cracker
Usage: jwtcrack -t <token> [-a alphabet] [-l max_len]

Options:
`)
	flag.PrintDefaults()
}

