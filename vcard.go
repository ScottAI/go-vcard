package go_vcard

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
A card is a set of properties,represent a contact
 */
type Card map[string][]*Property
/*
get first property depend on k
 */
func (c Card) Get(k string) *Property {
	props := c[k]
	if len(props) < 1{
		return nil
	}
	return props[0]
}
/*
Add property
 */
func (c Card) Add(key string,p *Property)  {
	c[key] = append(c[key],p)
}
/*
set a property
 */
func (c Card) Set(key string,p *Property)  {
	c[key] = []*Property{p}
}

/*
return the preferred property
 */
func (c Card) Pref(key string) *Property {
	props := c[key]
	if len(props) < 1{
		return nil
	}
	prop := props[0]
	max := 0
	for _,p := range props{
		n := 0
		if pref := p.GetFirstParamVal(ParamPref);pref != ""{
			n,_ = strconv.Atoi(pref)
		}else if p.IsHasType("pref"){
			//for Apple contact,add "pref" to TYPE param
			n = 1
		}

		if n > max{
			max = n
			prop = p
		}
	}
	return prop
}
/*
return first value of given property value
 */
func (c Card) Value(key string) [][]string {
	p := c.Get(key)
	if p == nil{
		return nil
	}
	return p.Value
}

func (c Card) AddValue(key string,vals [][]string)  {
	c.Add(key,&Property{Value:vals})
}

func (c Card) SetValue(key string,vals [][]string)  {
	c.Set(key,&Property{Value:vals})
}

func (c Card) PrefValue(key string) [][]string {
	p := c.Pref(key)
	if p == nil{
		return nil
	}
	return p.Value
}

func (c Card) Values(key string) [][][]string {
	props := c[key]
	if props == nil{
		return nil
	}

	vals := make([][][]string,len(props))
	for i,p := range props{
		vals[i] = p.Value
	}
	return vals
}

func (c Card) Kind() string {
	params := c.Value(PropKind)
	if params == nil{
		return ""
	}
	kind := strings.ToLower(params[0][0])
	if kind == ""{
		return KindIndividual
	}
	return kind
}

func (c Card) SetKind(kind string)  {
	c.SetValue(PropKind,[][]string{{kind}})
}

func (c Card) FormattedNames() []*Property {
	pnames := c[PropFN]
	if len(pnames) < 1{
		return []*Property{{Value:nil}}
	}
	return pnames
}

func (c Card) Names() []*Name {
	ps := c[PropN]
	if ps == nil{
		return nil
	}

	names := make([]*Name,len(ps))
	for i,p := range ps{
		names[i] = newName(p)
	}
	return names
}

func (c Card) Name() *Name {
	p := c.Pref(PropN)
	if p == nil{
		return nil
	}
	return newName(p)
}

func (c Card) AddName(n *Name)  {
	c.Add(PropN,n.property())
}

func (c Card) Gender() (gen string,identity string) {
	vals := c.Value(PropGender)
	sex := ""
	iden := ""
	if len(vals) == 2{
		sex = vals[0][0]
		iden = vals[1][0]
	}else if len(vals) == 1{
		sex = vals[0][0]
	}
	return sex,iden
}

func (c Card) SetGender(gen string,identity string)  {
	if gen == ""{
		return
	}
	if identity == ""{
		c.SetValue(PropGender,[][]string{{gen}})
	}else {
		c.SetValue(PropGender, [][]string{{gen},{identity}})
	}
}

func (c Card) Addresses() []*Address {
	adrs := c[PropAdr]
	if adrs == nil{
		return nil
	}

	addresses := make([]*Address,len(adrs))
	for i,adr := range adrs{
		addresses[i] = newAddress(adr)
	}
	return addresses
}

func (c Card) Address() *Address {
	adr := c.Pref(PropAdr)
	if adr == nil{
		return nil
	}
	return newAddress(adr)
}

func (c Card) AddAdress(adr *Address)  {
	c.Add(PropAdr,adr.property())
}

func (c Card) Categories() [][]string {
	return c.PrefValue(PropCategories)
}

func (c Card) SetCategories(categories [][]string)  {
	c.SetValue(PropCategories,categories)
}

func (c Card) Revision() (time.Time,error) {
	rev := c.Value(PropRev)
	if rev == nil{
		return time.Time{},nil
	}
	return time.Parse(timestampLayout,rev[0][0])
}

func (c Card) SetRevision(t time.Time)  {
	c.SetValue(PropRev,[][]string{{t.Format(timestampLayout)}})
}

func (c Card) ToV4() error {
	vers := c.Value(PropVersion)
	if len(vers) < 1{
		return fmt.Errorf("vcard:VCARD must have version infomation")
	}
	ver := vers[0][0]
	if strings.HasPrefix(ver,"4."){
		return nil
	}
	c.SetValue(PropVersion,[][]string{{"4.0"}})
	for k,props := range c{
		if strings.EqualFold(k,PropVersion){
			continue
		}
		for _,p := range props{
			if p.IsHasType("pref"){
				delete(p.Params,"pref")
				p.SetParam("pref","1")
			}
		}
	}
	return nil
}
