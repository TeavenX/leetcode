package main

import "bytes"

func main() {

}

const (
	READ_STATE          = "read"
	PRE_ANNO_STATE      = "pre_anno"
	BLOCK_ANNO_STATE    = "block_anno"
	BLOCK_PRE_END_STATE = "block_pre_end"
	LINE_ANNO_STATE     = "line_anno"
	LINE_PRE_END_STATE  = "line_pre_end_state"
)

type State interface {
	Enter(b rune)
	Transfer(b rune) string
}

type ReadState struct {
	b bytes.Buffer
}

func (s *ReadState) Enter(b rune) {
	s.b.WriteRune(b)
}

func (s *ReadState) Transfer(b rune) string {
	if b == '/' {
		return PRE_ANNO_STATE
	}
	return READ_STATE
}

type PreAnnoState struct{}

func (s *PreAnnoState) Enter(b rune) {}

func (s *PreAnnoState) Transfer(b rune) string {
	switch b {
	case '/':
		return LINE_ANNO_STATE
	case '*':
		return BLOCK_ANNO_STATE
	default:
		return READ_STATE
	}
}

type BlockAnnoState struct{}

func (s *BlockAnnoState) Enter(b rune) {}

func (s *BlockAnnoState) Transfer(b rune) string {
	if b == '*' {
		return BLOCK_PRE_END_STATE
	}
	return BLOCK_ANNO_STATE
}

type BlockPreEndState struct{}

func (s *BlockPreEndState) Enter(b rune) {}

func (s *BlockPreEndState) Transfer(b rune) string {
	if b == '/' {
		return READ_STATE
	}
	return BLOCK_ANNO_STATE
}

type LineAnnoState struct{}

func (s *LineAnnoState) Enter(b rune) {}

func (s *LineAnnoState) Transfer(b rune) string {
	if b == '/' {
		return LINE_PRE_END_STATE
	}
	return LINE_ANNO_STATE
}

type LinePreEndState struct{}

func (s *LinePreEndState) Enter(b rune) {}

func (s *LinePreEndState) Transfer(b rune) string {
	if b == 'n' {
		return READ_STATE
	}
	return LINE_ANNO_STATE
}

func removeComments(source []string) []string {
	ans := make([]string, 0)
	// todo
}

func removeComments(source []string) []string {
	ans := make([]string, 0)
	t := make([]byte, 0)
	inBlock := false
	for _, s := range source {
		n := len(s)
		for i := 0; i < n; i++ {
			if inBlock {
				if i+1 < n && s[i] == '*' && s[i+1] == '/' {
					inBlock = false
					i++
				}
			} else {
				if i+1 < n && s[i] == '/' && s[i+1] == '*' {
					inBlock = true
					i++
				} else if i+1 < n && s[i] == '/' && s[i+1] == '/' {
					break
				} else {
					t = append(t, s[i])
				}
			}
		}
		if !inBlock && len(t) > 0 {
			ans = append(ans, string(t))
			t = make([]byte, 0)
		}
	}
	return ans
}
