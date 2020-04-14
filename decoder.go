package go_vcard

import (
	"io"
	"text/scanner"
)

type Decoder struct {
	scan *scanner.Scanner
}

func NewDecoder(r io.Reader) *Decoder {
	var s scanner.Scanner
	s.Init(r)
	return &Decoder{&s}
}

func (dc *Decoder) Decode() (Card,error) {
	c := make(Card)
	prop := dc.ReadProp()
	for prop != nil{
		c[prop.Name] = append(c[prop.Name],prop)
		prop = dc.ReadProp()
	}
	return c,nil
}

func (dc *Decoder) ReadProp() *Property {
	if dc.scan.Peek() == scanner.EOF{
		return nil
	}
	group,name := dc.readGroupName()
	params := make(map[string][]string)
	if dc.scan.Peek() == ';'{
		params = dc.readParams()
	}
	dc.scan.Next()
	value := dc.readValues()
	return &Property{Group:group,Name:name,Params:params,Value:value}
}

func (dc *Decoder) readGroupName() (group,name string) {
	c := dc.scan.Peek()
	var buf []rune
	for c != scanner.EOF{
		if c == '.'{
			group = string(buf)
			buf = []rune{}
		}else if c == ';' || c == ':'{
			name = string(buf)
			return
		}else if c == '\n' || c == '\r'{
			//skip empty line in vcard
		}else{
			buf = append(buf,c)
		}
		dc.scan.Next()
		c = dc.scan.Peek()
	}
	return
}

func (dc *Decoder) readParams() (params map[string][]string) {
	lastChar := dc.scan.Peek()
	c := lastChar
	var buf []rune
	var name string
	var value string
	params = make(map[string][]string)
	var values []string
	for c != scanner.EOF{
		if c == ','{
			values = append(values,string(buf))
			buf = []rune{}
		} else if c == ';' || c == ':'{
			if name == ""{
				name = string(buf)
			} else {
				value = string(buf)
			}
			if name != ""{
				values = append(values,value)
				if _,ok := params[name];ok{
					params[name] = append(params[name],values...)
				} else {
					params[name] = values
				}
			}
			if c == ':'{
				return
			}
			buf = []rune{}
			values = []string{}
			name = ""
			value = ""
		} else if c == '='{
			name = string(buf)
			buf = []rune{}
		}else{
			buf = append(buf,c)
		}
		dc.scan.Next()
		c = dc.scan.Peek()
	}
	return 
}

func (dc *Decoder) readValues() (values [][]string) {
	lastChar := dc.scan.Next()
	c := lastChar
	var buf []rune
	escape := false
	var val []string
	for c != scanner.EOF{
		if c == '\n'{
			la := dc.scan.Peek()
			if la != 32 && la != 9{
				if len(buf) > 0{
					val = append(val,string(buf))
				}
				values = append(values,val)
				return
			}else{
				//unfold
				lastChar = la
				c = dc.scan.Next()
				for c == 32 || c == 9{
					c = dc.scan.Next()
				}
			}
		}
		if c == '\\'{
			escape = true
		}else if escape{
			if c == 'n' || c == 'N'{
				c = '\n'
			}
			buf = append(buf,c)
			escape = false
		}else if c == ','{
			if len(buf) > 0{
				val = append(val,string(buf))
				buf = []rune{}
			}
		}else if c == ';'{
			if len(buf) > 0{
				val = append(val,string(buf))
				buf = []rune{}
			}
			values = append(values,val)
			val = []string{}
		}else if c != '\n' && c != '\r'{
			buf = append(buf,c)
		}
		lastChar = c
		c = dc.scan.Next()
	}
	return
}
