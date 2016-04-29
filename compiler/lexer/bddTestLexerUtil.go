package lexer

import (
//	"fmt"
	"strconv"
	"strings"
)

const (
	C_Feature = iota
	C_Scenario
	C_Assertion
)

var _emptystruct struct{} = struct{}{}

func (this BddTestLex) OnBddTest(item interface{}) {
//	fmt.Println("OnBddTest()")
	res := &BddTestParseRes{Lines:item}
//	switch t := item.(type) {
//		default:
//			fmt.Printf("OnBddTest() err, invalid type, %T", t)
//			return
//		case Assertion:
//			res.Type = C_Assertion
//	}
//	fmt.Println(this)
	
	if nil != this.OnBddTestParse {
		this.OnBddTestParse(res)
	}
	
	if nil != this.Itemchan {
		this.Itemchan <- res
	}
}

func (this *BddTestLex) peek() (bret byte) {
	bret = this.next()
	this.pos -= 1
	return
}

func (this *BddTestLex) next() (bret byte) {

	if this.pos < len(this.S) {
		bret = byte(this.S[this.pos])
		this.buf = append(this.buf, bret)
	} else {
		bret = 0
	}
	this.pos += 1
	return
}

func (this *BddTestLex) back() {
	if 0 < this.pos {
		this.pos -= 1
	}
	return
}

func (this BddTestLex) data() (bb []byte) {
	if this.pos < len(this.S) {
		bb = this.buf[:len(this.buf)-1]
	} else {
		bb = this.buf
	}
	return
}

func (this BddTestLex) Error(s string) {
	res := &BddTestParseRes{Error:s}
	if nil != this.OnBddTestParse {
		this.OnBddTestParse(res)
	}
}

func (this BddTestLex) getInt() (n int, err error) {
	s := string(this.data())
	if n, err = strconv.Atoi(strings.TrimSpace(s)); nil != err {
		this.Error(err.Error() + ";s=" + s)
	}
	return
}

func (this BddTestLex) getStr() (s string) {
	s = strings.TrimSpace(string(this.data()))
	return
}

func (this BddTestLex) getParam() (s string) {
	str := strings.TrimSpace(string(this.data()))
	s = str[1:len(str)-1]
	return
}

