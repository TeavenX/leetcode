package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	c := Constructor()
	//fmt.Println(c.encode("https://leetcode.csadom/probldsaems/design-tinwdsadSadsadsadsaDsaDsadsaDsaDsaDsadsawefdashbweryurl"))
	fmt.Println(c.hash("https://leetcode.csadom/probldsaems/design-tinwdsadSadsadsadsaDsaDsadsaDsaDsaDsadsawefdashbweryurl"))
	a := c.encode("https://leetcode.com/probldsaems/design-tinbweryurl")
	fmt.Println(c.decode(a))
}

type Codec struct {
	baseURL string
	cache   map[string]string
}

func Constructor() Codec {
	return Codec{
		baseURL: "http://tinyurl.com/",
		cache:   make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	key := this.hash(longUrl)
	this.cache[key] = longUrl
	return key
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	idx := strings.LastIndexByte(shortUrl, '/')
	//key, _ := strconv.Atoi(shortUrl[idx+1:])
	key := shortUrl[idx+1:]
	return this.cache[key]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func (this *Codec) hash(s string) string {
	urlsplit := strings.Split(s, "/")
	protocol := urlsplit[0]
	host := urlsplit[2]
	path := ""
	if len(urlsplit) > 3 {
		path = urlsplit[3]
	}
	result := ""
	base := protocol + host + path + host
	temp := 1
	for _, str := range base {
		n := int(str >> 1)
		if n > 0 {
			temp *= n
		}
	}
	temp %= 1e8
	for temp > 0 {
		a := temp % 10
		result += string(dict[(rand.Intn(len(dict))+a)%len(dict)])
		temp /= 10
	}
	return result
}
