package go_vcard

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
func (c Card) Set()  {
	
}

