package go_vcard

import "strings"

type Property struct {
	Group string
	Name string
	Params map[string][]string
	Value [][]string //first seperate by ';',secondly seperate by ','
}

func (p *Property) GetValueTextList() []string {
	var tl []string
	for _,sts := range p.Value{
		for _,st := range sts{
			tl = append(tl,st)
		}
	}
	return tl
}
/*
get value first value,if not return ""
 */
func (p *Property) GetValueFirstText() string {
	if len(p.Value)>0 && len(p.Value[0])>0{
		return p.Value[0][0]
	}
	return ""
}

/*
get params first,if not return ""
 */
func (p *Property) GetFirstParamVal(param string) string {
	if len(p.Params[param])>0{
		return p.Params[param][0]
	}
	return ""
}

/*
Add a param,the key is existed
 */
func (p *Property) AddParam(key,val string)  {
	p.Params[key] = append(p.Params[key],val)
}

/*
set a param,new a param
 */
func (p *Property) SetParam(key,val string)  {
	p.Params[key] = []string{val}
}

/*
get list of param types
 */
func (p *Property) GetParamTypeList() []string {
	types := p.Params[ParamType]
	list := make([]string,len(types))
	for i,t := range types{
		list[i] = strings.ToLower(t)
	}
	return list
}

/*
judge has a type?
 */
func (p *Property) IsHasType(t string) bool {
	for _,tt := range p.Params[ParamType]{
		if strings.EqualFold(t,tt){
			return true
		}
	}
	return false
}

type Name struct {
	*Property

	FamilyName string
	GivenName	string
	AdditionalName string
	HonorificPrefix string
	HonorificSuffix string
}

func newName(p *Property) *Name {
	vals := p.Value[0]
	return &Name{
		Property:        p,
		FamilyName:      vals[0],
		GivenName:       vals[1],
		AdditionalName:  vals[2],
		HonorificPrefix: vals[3],
		HonorificSuffix: vals[4],
	}
}

func (n *Name) property() *Property {
	if n.Property == nil{
		n.Property = new(Property)
		n.Property.Value = [][]string{{
			n.FamilyName,
			n.GivenName,
			n.AdditionalName,
			n.HonorificPrefix,
			n.HonorificSuffix,
		}}
	}else{
		n.Property.Value = append(n.Property.Value,[]string{
			n.FamilyName,
			n.GivenName,
			n.AdditionalName,
			n.HonorificPrefix,
			n.HonorificSuffix,
		})
	}

	return n.Property
}

type Address struct {
	*Property

	PostOfficeBox string
	ExtendedAddress string//e.g apartment or suite no
	StreetAddress string
	Locality	string
	Region string
	PostalCode string
	Country string
}

func newAddress(p *Property) *Address {
	vals := p.Value[0]
	return &Address{
		Property:        p,
		PostOfficeBox:   vals[0],
		ExtendedAddress: vals[1],
		StreetAddress:   vals[2],
		Locality:        vals[3],
		Region:          vals[4],
		PostalCode:      vals[5],
		Country:         vals[6],
	}
}

func (a *Address) property() *Property {
	if a.Property == nil {
		a.Property = new(Property)
	}
	a.Property.Value = append(a.Property.Value,[]string{
		a.PostOfficeBox,
		a.ExtendedAddress,
		a.StreetAddress,
		a.Locality,
		a.Region,
		a.PostalCode,
		a.Country,
	})
	return  a.Property
}


