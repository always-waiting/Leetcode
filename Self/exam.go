package Self

import (
	"fmt"
)

func GetSub(in []string) {
	num := len(in)
	for i := 0; i < num; i++ {
		for j := i + 1; j <= num; j++ {
			fmt.Println(in[i:j])
		}
	}
}

func Match(reg, ctx string) string {
	num := len(ctx)
	var idxCtx, idxReg, maxLen, matchIdx int
	var match bool
	for idxCtx < num {
		if ctx[idxCtx] == reg[idxReg] {
			if !match {
				match = true
				matchIdx = idxCtx
			}
			idxReg++
			idxCtx++
		} else {
			if match {
				if idxReg > maxLen {
					maxLen = idxReg
				}
				match = false
				idxCtx = matchIdx + idxReg
				idxReg = 0
			} else {
				idxCtx++
			}
		}
		fmt.Println(idxCtx)
	}
	return reg[0:maxLen]
}

func Slice(a []string) {
	a[0] = "a"
}

func Array(a [2]string) {
	a[0] = "a"
}
