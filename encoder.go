package go_vcard

import (
	"errors"
	"io"
	"sort"
	"strings"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{writer:w}
}
func (ec *Encoder) Encode(c Card) error {
	begin := "BEGIN:VCARD\r\n"
	if _,err := io.WriteString(ec.writer,begin);err != nil{
		return err
	}
	version := c.Get(PropVersion)
	if version == nil{
		return errors.New("VCARD: Version property missing")
	}
	ec.WriteProperty(version)
	var keys []string
	for k := range c{
		keys = append(keys,k)
	}
	sort.Strings(keys)
	for _,k := range keys{
		props := c[k]
		if strings.EqualFold(k,PropVersion){
			continue
		}
		for _,p := range props{
			ec.WriteProperty(p)
		}
	}
	end := "END:VCARD\r\n"
	io.WriteString(ec.writer,end)
	return nil
}
func (ec *Encoder) WriteProperty(prop *Property)  {
	if prop.Group != ""{
		io.WriteString(ec.writer,prop.Group)
		io.WriteString(ec.writer,".")
	}
	io.WriteString(ec.writer,prop.Name)
	if prop.Params != nil{
		for key,vals := range prop.Params{
			io.WriteString(ec.writer,";")
			io.WriteString(ec.writer,key)
			if len(vals) > 0{
				io.WriteString(ec.writer,"=")
				for vi := 0;vi < len(vals);vi++{
					io.WriteString(ec.writer,vals[vi])
					if vi+1 < len(vals){
						io.WriteString(ec.writer,",")
					}
				}
			}
		}
	}
	io.WriteString(ec.writer,":")
	for si:=0;si < len(prop.Value);si++{
		for vi := 0; vi < len(prop.Value[si]);vi++{

		}
	}
}

func (ec *Encoder) WriteValue(val string)  {
	i := 0
	for _,c := range val{
		if i == 76{
			io.WriteString(ec.writer,"\n  ")
			i = 0
		}
		var e string
		switch c {
		case '\r':
			e = `\r`
		case '\n':
			e = `\n`
		case ';':
			e = `\;`
		case ':':
			e = `\:`
		case ',':
			e = `\,`
		default:
			e = string(c)
			
		}
		io.WriteString(ec.writer,e)
		i++
	}
}
