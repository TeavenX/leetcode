package main

func main() {

}

type AuthenticationManager struct {
	cache map[string]int
	ttl   int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		cache: make(map[string]int),
		ttl:   timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.cache[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if this.cache[tokenId] > currentTime {
		this.cache[tokenId] = currentTime + this.ttl
	} else {
		delete(this.cache, tokenId)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	ans := 0
	for _, v := range this.cache {
		if v > currentTime {
			ans++
		}
	}
	return ans
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
