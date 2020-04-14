package go_vcard

//https://tools.ietf.org/html/rfc6350#section-5
//Property parameters
const(
	ParamLanguage = "LANGUAGE"
	ParamValue = "VALUE"
	ParamPref = "PREF"
	ParamAltid = "ALTID"
	ParamPid = "PID"
	ParamType = "TYPE"
	ParamMediatype = "MEDIATYPE"
	ParamCalscale = "CALSCALE"
	ParamSortAs = "SORT-AS"
	ParamGEO = "GEO"
	ParamTZ = "TZ"
)

const timestampLayout = "20060102T150405Z"

//https://tools.ietf.org/html/rfc6350#section-6
//vcard properties

const(
	//standard vcard properties
	/*
	Value type:  text
	Cardinality:  1
	 */
	PropBegin = "BEGIN"

	/*
	Value type:  text
	Cardinality:  1
	 */
	PropEnd = "END"

	/*
	Purpose:  To identify the source of directory information contained
	      in the content type.
	Value type:  uri
	Cardinality:  *
	ABNF:
	     SOURCE-param = "VALUE=uri" / pid-param / pref-param / altid-param
	                  / mediatype-param / any-param
	     SOURCE-value = URI
	   Examples:

	     SOURCE:ldap://ldap.example.com/cn=Babs%20Jensen,%20o=Babsco,%20c=US

	     SOURCE:http://directory.example.com/addressbooks/jdoe/
	      Jean%20Dupont.vcf
	 */
	PropSource = "SOURCE"

	/*
	Purpose:  To specify the kind of object the vCard represents.
	Value type:  A single text value.
	Cardinality:  *1
	the value may be on of the : individual group  org  location
	ABNF:
	     KIND-param = "VALUE=text" / any-param
	     KIND-value = "individual" / "group" / "org" / "location"
	                / iana-token / x-name
	   Example:
	      This represents someone named Jane Doe working in the marketing
	      department of the North American division of ABC Inc.
	         BEGIN:VCARD
	         VERSION:4.0
	         KIND:individual
	         FN:Jane Doe
	         ORG:ABC\, Inc.;North American Division;Marketing
	         END:VCARD
	   This represents the department itself, commonly known as ABC
	   Marketing.
	         BEGIN:VCARD
	         VERSION:4.0
	         KIND:org
	         FN:ABC Marketing
	         ORG:ABC\, Inc.;North American Division;Marketing
	         END:VCARD
	 */
	PropKind = "KIND"

	/*
	Purpose:  To include extended XML-encoded vCard data in a plain
	      vCard.
	Value type:  A single text value.
	Cardinality:  *

	ABNF:
	     XML-param = "VALUE=text" / altid-param
	     XML-value = text
	 */
	PropXML = "XML"


	//Identification Properties
	/*
	These types are used to capture information associated with the
	   identification and naming of the entity associated with the vCard.
	 */

	/*
	Purpose:  To specify the formatted text corresponding to the name of
	      the object the vCard represents.
	   Value type:  A single text value.
	   Cardinality:  1
	   ABNF:
	     FN-param = "VALUE=text" / type-param / language-param / altid-param
	              / pid-param / pref-param / any-param
	     FN-value = text
	   Example:
	         FN:Mr. John Q. Public\, Esq.
	 */
	PropFN = "FN"

	/*
	Purpose:  To specify the components of the name of the object the
	      vCard represents.
	Value type:  A single structured text value.  Each component can have
	      multiple values.
	Cardinality:  *1
	ABNF:
	     N-param = "VALUE=text" / sort-as-param / language-param
	             / altid-param / any-param
	     N-value = list-component 4(";" list-component)
	Examples:
			N:Public;John;Quinlan;Mr.;Esq.
			N:Stevenson;John;Philip,Paul;Dr.;Jr.,M.D.,A.C.P.
	 */
	PropN = "N"

	/*
	Purpose:  To specify the text corresponding to the nickname of the
	      object the vCard represents.
	Value type:  One or more text values separated by a COMMA character
	      (U+002C).
	Cardinality:  *
	ABNF:
	     NICKNAME-param = "VALUE=text" / type-param / language-param
	                    / altid-param / pid-param / pref-param / any-param
	     NICKNAME-value = text-list
	Examples:
	             NICKNAME:Robbie
	             NICKNAME:Jim,Jimmie
	             NICKNAME;TYPE=work:Boss
	*/
	PropNickName = "NICKNAME"

	/*
	Purpose:  To specify an image or photograph information that
	      annotates some aspect of the object the vCard represents.
	   Value type:  A single URI.
	   Cardinality:  *
	ABNF:
	     PHOTO-param = "VALUE=uri" / altid-param / type-param
	                 / mediatype-param / pref-param / pid-param / any-param
	     PHOTO-value = URI
	Examples:
	       PHOTO:http://www.example.com/pub/photos/jqpublic.gif
	       PHOTO:data:image/jpeg;base64,MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhv
	        AQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bm
	        ljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0
	        <...remainder of base64-encoded data...>
	 */
	PropPhoto = "PHOTO"

	/*
	Purpose:  To specify the birth date of the object the vCard
	      represents.
	Value type:  The default is a single date-and-or-time value.  It can
	      also be reset to a single text value.
	Cardinality:  *1
	ABNF:
	     BDAY-param = BDAY-param-date / BDAY-param-text
	     BDAY-value = date-and-or-time / text
	       ; Value and parameter MUST match.

	     BDAY-param-date = "VALUE=date-and-or-time"
	     BDAY-param-text = "VALUE=text" / language-param

	     BDAY-param =/ altid-param / calscale-param / any-param
	       ; calscale-param can only be present when BDAY-value is
	       ; date-and-or-time and actually contains a date or date-time.
	Examples:
	             BDAY:19960415
	             BDAY:--0415
	             BDAY;19531015T231000Z
	             BDAY;VALUE=text:circa 1800
	 */
	PropBday = "BDAY"

	/*
	Purpose:  The date of marriage, or equivalent, of the object the
	      vCard represents.
	   Value type:  The default is a single date-and-or-time value.  It can
	      also be reset to a single text value.
	   Cardinality:  *1
	ABNF:
	     ANNIVERSARY-param = "VALUE=" ("date-and-or-time" / "text")
	     ANNIVERSARY-value = date-and-or-time / text
	       ; Value and parameter MUST match.
	     ANNIVERSARY-param =/ altid-param / calscale-param / any-param
	       ; calscale-param can only be present when ANNIVERSARY-value is
	       ; date-and-or-time and actually contains a date or date-time.
	Examples:
		ANNIVERSARY:19960415
	 */
	PropAnniversary = "ANNIVERSARY"

	/*
	   Purpose:  To specify the components of the sex and gender identity of
	      the object the vCard represents.
	   Value type:  A single structured value with two components.  Each
	      component has a single text value.
	   Cardinality:  *1
	   ABNF:
	                   GENDER-param = "VALUE=text" / any-param
	                   GENDER-value = sex [";" text]
	                   sex = "" / "M" / "F" / "O" / "N" / "U"
	   Examples:
	     GENDER:M
	     GENDER:F
	     GENDER:M;Fellow
	     GENDER:F;grrrl
	     GENDER:O;intersex
	     GENDER:;it's complicated
	 */
	PropGender = "GENDER"


	//Delivery Addressing Properties
	/*
	These types are concerned with information related to the delivery
	   addressing or label for the vCard object.
	 */

	/*
	   Purpose:  To specify the components of the delivery address for the
	      vCard object.
	   Value type:  A single structured text value, separated by the
	      SEMICOLON character (U+003B).
	   Cardinality:  *
	ABNF:
	     label-param = "LABEL=" param-value
	     ADR-param = "VALUE=text" / label-param / language-param
	               / geo-parameter / tz-parameter / altid-param / pid-param
	               / pref-param / type-param / any-param
	     ADR-value = ADR-component-pobox ";" ADR-component-ext ";"
	                 ADR-component-street ";" ADR-component-locality ";"
	                 ADR-component-region ";" ADR-component-code ";"
	                 ADR-component-country
	     ADR-component-pobox    = list-component
	     ADR-component-ext      = list-component
	     ADR-component-street   = list-component
	     ADR-component-locality = list-component
	     ADR-component-region   = list-component
	     ADR-component-code     = list-component
	     ADR-component-country  = list-component

	Example: In this example, the post office box and the extended
	   address are absent.
	     ADR;GEO="geo:12.3457,78.910";LABEL="Mr. John Q. Public, Esq.\n
	      Mail Drop: TNE QB\n123 Main Street\nAny Town, CA  91921-1234\n
	      U.S.A.":;;123 Main Street;Any Town;CA;91921-1234;U.S.A.
	 */
	PropAdr = "ADR"

	//Communications Properties
	/*
	These properties describe information about how to communicate with
	   the object the vCard represents.
	 */

	/*
		Purpose:  To specify the telephone number for telephony communication
	      with the object the vCard represents.
	   Value type:  By default, it is a single free-form text value (for
	      backward compatibility with vCard 3), but it SHOULD be reset to a
	      URI value.  It is expected that the URI scheme will be "tel", as
	      specified in [RFC3966], but other schemes MAY be used.
	   Cardinality:  *
		ABNF:

	     TEL-param = TEL-text-param / TEL-uri-param
	     TEL-value = TEL-text-value / TEL-uri-value
	       ; Value and parameter MUST match.

	     TEL-text-param = "VALUE=text"
	     TEL-text-value = text

	     TEL-uri-param = "VALUE=uri" / mediatype-param
	     TEL-uri-value = URI

	     TEL-param =/ type-param / pid-param / pref-param / altid-param
	                / any-param

	     type-param-tel = "text" / "voice" / "fax" / "cell" / "video"
	                    / "pager" / "textphone" / iana-token / x-name
	       ; type-param-tel MUST NOT be used with a property other than TEL.
	   Example:

	     TEL;VALUE=uri;PREF=1;TYPE="voice,home":tel:+1-555-555-5555;ext=5555
	     TEL;VALUE=uri;TYPE=home:tel:+33-01-23-45-67
	 */
	PropTel = "TEL"

	/*
	   Purpose:  To specify the electronic mail address for communication
	      with the object the vCard represents.
	   Value type:  A single text value.
	   Cardinality:  *
	   ABNF:
	     EMAIL-param = "VALUE=text" / pid-param / pref-param / type-param
	                 / altid-param / any-param
	     EMAIL-value = text
	   Example:
	           EMAIL;TYPE=work:jqpublic@xyz.example.com
	           EMAIL;PREF=1:jane_doe@example.com
	 */
	PropEmail = "EMAIL"

	/*
	   Purpose:  To specify the URI for instant messaging and presence
	      protocol communications with the object the vCard represents.
	   Value type:  A single URI.
	   Cardinality:  *
	   ABNF:
	     IMPP-param = "VALUE=uri" / pid-param / pref-param / type-param
	                / mediatype-param / altid-param / any-param
	     IMPP-value = URI
	   Example:
	       IMPP;PREF=1:xmpp:alice@example.com
	 */
	PropImpp = "IMPP"

	/*
	   Purpose:  To specify the language(s) that may be used for contacting
	      the entity associated with the vCard.
	   Value type:  A single language-tag value.
	   Cardinality:  *
	   ABNF:
	     LANG-param = "VALUE=language-tag" / pid-param / pref-param
	                / altid-param / type-param / any-param
	     LANG-value = Language-Tag
	   Example:
	       LANG;TYPE=work;PREF=1:en
	       LANG;TYPE=work;PREF=2:fr
	       LANG;TYPE=home:fr
	 */
	PropLang = "LANG"

	//Geographical Properties
	/*
	These properties are concerned with information associated with
	   geographical positions or regions associated with the object the
	   vCard represents.
	*/

	/*
		Purpose:  To specify information related to the time zone of the
	      object the vCard represents.
	   Value type:  The default is a single text value.  It can also be
	      reset to a single URI or utc-offset value.
	   Cardinality:  *
	   ABNF:
	     TZ-param = "VALUE=" ("text" / "uri" / "utc-offset")
	     TZ-value = text / URI / utc-offset
	       ; Value and parameter MUST match.
	     TZ-param =/ altid-param / pid-param / pref-param / type-param
	               / mediatype-param / any-param
	   Examples:
	     TZ:Raleigh/North America
	     TZ;VALUE=utc-offset:-0500
	       ; Note: utc-offset format is NOT RECOMMENDED.
	 */
	PropTZ = "TZ"

	/*
		Purpose:  To specify information related to the global positioning of
	      the object the vCard represents.
	   Value type:  A single URI.
	   Cardinality:  *
	   ABNF:
	     GEO-param = "VALUE=uri" / pid-param / pref-param / type-param
	               / mediatype-param / altid-param / any-param
	     GEO-value = URI
	   Example:
	           GEO:geo:37.386013,-122.082932
	 */
	PropGEO = "GEO"


	//Organizational Properties
	/*
	These properties are concerned with information associated with
	   characteristics of the organization or organizational units of the
	   object that the vCard represents.
	 */

	/*
	   Purpose:  To specify the position or job of the object the vCard
	      represents.
	   Value type:  A single text value.
	   Cardinality:  *
	   ABNF:
	     TITLE-param = "VALUE=text" / language-param / pid-param
	                 / pref-param / altid-param / type-param / any-param
	     TITLE-value = text
	   Example:
	           TITLE:Research Scientist
	 */
	PropTiTle = "TITLE"

	/*
	   Purpose:  To specify the function or part played in a particular
	      situation by the object the vCard represents.
	   Value type:  A single text value.
	   Cardinality:  *
	   ABNF:
	     ROLE-param = "VALUE=text" / language-param / pid-param / pref-param
	                / type-param / altid-param / any-param
	     ROLE-value = text
	   Example:
	           ROLE:Project Leader
	*/
	PropRole = "ROLE"

	/*
	   Purpose:  To specify a graphic image of a logo associated with the
	      object the vCard represents.
	   Value type:  A single URI.
	   Cardinality:  *
	   ABNF:
	     LOGO-param = "VALUE=uri" / language-param / pid-param / pref-param
	                / type-param / mediatype-param / altid-param / any-param
	     LOGO-value = URI
	   Examples:
	     LOGO:http://www.example.com/pub/logos/abccorp.jpg

	     LOGO:data:image/jpeg;base64,MIICajCCAdOgAwIBAgICBEUwDQYJKoZIhvc
	      AQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bm
	      ljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0
	      <...the remainder of base64-encoded data...>
	 */
	PropLogo = "LOGO"

	/*
	   Purpose:  To specify the organizational name and units associated
	      with the vCard.
	   Value type:  A single structured text value consisting of components
	      separated by the SEMICOLON character (U+003B).
	   Cardinality:  *
	   ABNF:
	     ORG-param = "VALUE=text" / sort-as-param / language-param
	               / pid-param / pref-param / altid-param / type-param
	               / any-param
	     ORG-value = component *(";" component)
	   Example: A property value consisting of an organizational name,
	   organizational unit #1 name, and organizational unit #2 name.

	           ORG:ABC\, Inc.;North American Division;Marketing
	*/
	PropOrg = "ORG"

	/*
	   Purpose:  To include a member in the group this vCard represents.
	   Value type:  A single URI.  It MAY refer to something other than a
	      vCard object.  For example, an email distribution list could
	      employ the "mailto" URI scheme [RFC6068] for efficiency.
	   Cardinality:  *
	   ABNF:
	     MEMBER-param = "VALUE=uri" / pid-param / pref-param / altid-param
	                  / mediatype-param / any-param
	     MEMBER-value = URI
	   Examples:

	     BEGIN:VCARD
	     VERSION:4.0
	     KIND:group
	     FN:The Doe family
	     MEMBER:urn:uuid:03a0e51f-d1aa-4385-8a53-e29025acd8af
	     MEMBER:urn:uuid:b8767877-b4a1-4c70-9acc-505d3819e519
	     END:VCARD
	     BEGIN:VCARD
	     VERSION:4.0
	     FN:John Doe
	     UID:urn:uuid:03a0e51f-d1aa-4385-8a53-e29025acd8af
	     END:VCARD
	     BEGIN:VCARD
	     VERSION:4.0
	     FN:Jane Doe
	     UID:urn:uuid:b8767877-b4a1-4c70-9acc-505d3819e519
	     END:VCARD

	     BEGIN:VCARD
	     VERSION:4.0
	     KIND:group
	     FN:Funky distribution list
	     MEMBER:mailto:subscriber1@example.com
	     MEMBER:xmpp:subscriber2@example.com
	     MEMBER:sip:subscriber3@example.com
	     MEMBER:tel:+1-418-555-5555
	     END:VCARD
	 */
	PropMember = "MEMBER"

	/*
	   Purpose:  To specify a relationship between another entity and the
	      entity represented by this vCard.
	   Value type:  A single URI.  It can also be reset to a single text
	      value.  The text value can be used to specify textual information.
	   Cardinality:  *
	   ABNF:
	     RELATED-param = RELATED-param-uri / RELATED-param-text
	     RELATED-value = URI / text
	       ; Parameter and value MUST match.

	     RELATED-param-uri = "VALUE=uri" / mediatype-param
	     RELATED-param-text = "VALUE=text" / language-param

	     RELATED-param =/ pid-param / pref-param / altid-param / type-param
	                    / any-param

	     type-param-related = related-type-value *("," related-type-value)
	       ; type-param-related MUST NOT be used with a property other than
	       ; RELATED.
	     related-type-value = "contact" / "acquaintance" / "friend" / "met"
	                        / "co-worker" / "colleague" / "co-resident"
	                        / "neighbor" / "child" / "parent"
	                        / "sibling" / "spouse" / "kin" / "muse"
	                        / "crush" / "date" / "sweetheart" / "me"
	                        / "agent" / "emergency"
	   Examples:
	   RELATED;TYPE=friend:urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6
	   RELATED;TYPE=contact:http://example.com/directory/jdoe.vcf
	   RELATED;TYPE=co-worker;VALUE=text:Please contact my assistant Jane
	    Doe for any inquiries.
	 */
	PropRelated = "RELATED"

	//Explanatory Properties
	/*
	   These properties are concerned with additional explanations, such as
	   that related to informational notes or revisions specific to the
	   vCard.
	 */

	/*
	   Purpose:  To specify application category information about the
	      vCard, also known as "tags".
	   Value type:  One or more text values separated by a COMMA character
	      (U+002C).
	   Cardinality:  *
	   ABNF:

	     CATEGORIES-param = "VALUE=text" / pid-param / pref-param
	                      / type-param / altid-param / any-param
	     CATEGORIES-value = text-list

	   Example:
	           CATEGORIES:TRAVEL AGENT
	           CATEGORIES:INTERNET,IETF,INDUSTRY,INFORMATION TECHNOLOGY
	 */
	PropCategories = "CATEGORIES"

	/*
	   Purpose:  To specify supplemental information or a comment that is
	      associated with the vCard.
	   Value type:  A single text value.
	   Cardinality:  *
	   ABNF:
	     NOTE-param = "VALUE=text" / language-param / pid-param / pref-param
	                / type-param / altid-param / any-param
	     NOTE-value = text
	   Example:
	           NOTE:This fax number is operational 0800 to 1715
	             EST\, Mon-Fri.
	 */
	PropNote = "NOTE"

	/*
	   Purpose:  To specify the identifier for the product that created the
	      vCard object.
	   Type value:  A single text value.
	   Cardinality:  *1
	   ABNF:
	     PRODID-param = "VALUE=text" / any-param
	     PRODID-value = text
	   Example:

	           PRODID:-//ONLINE DIRECTORY//NONSGML Version 1//EN
	 */
	PropProdid = "PRODID"

	/*
	   Purpose:  To specify revision information about the current vCard.
	   Value type:  A single timestamp value.
	   Cardinality:  *1
	   ABNF:
	     REV-param = "VALUE=timestamp" / any-param
	     REV-value = timestamp
	   Example:
	           REV:19951031T222710Z
	 */
	PropRev = "REV"

	/*
	   Purpose:  To specify a digital sound content information that
	      annotates some aspect of the vCard.  This property is often used
	      to specify the proper pronunciation of the name property value of
	      the vCard.
	   Value type:  A single URI.
	   Cardinality:  *
	   ABNF:
	     SOUND-param = "VALUE=uri" / language-param / pid-param / pref-param
	                 / type-param / mediatype-param / altid-param
	                 / any-param
	     SOUND-value = URI
	   Example:
	     SOUND:CID:JOHNQPUBLIC.part8.19960229T080000.xyzMail@example.com

	     SOUND:data:audio/basic;base64,MIICajCCAdOgAwIBAgICBEUwDQYJKoZIh
	      AQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05ldHNjYXBlIENvbW11bm
	      ljYXRpb25zIENvcnBvcmF0aW9uMRwwGgYDVQQLExNJbmZvcm1hdGlvbiBTeXN0
	      <...the remainder of base64-encoded data...>
	 */
	PropSound = "SOUND"

	/*
	   Purpose:  To specify a value that represents a globally unique
	      identifier corresponding to the entity associated with the vCard.
	   Value type:  A single URI value.  It MAY also be reset to free-form
	      text.
	   Cardinality:  *1
	   ABNF:
	     UID-param = UID-uri-param / UID-text-param
	     UID-value = UID-uri-value / UID-text-value
	       ; Value and parameter MUST match.

	     UID-uri-param = "VALUE=uri"
	     UID-uri-value = URI

	     UID-text-param = "VALUE=text"
	     UID-text-value = text

	     UID-param =/ any-param

	   Example:
	           UID:urn:uuid:f81d4fae-7dec-11d0-a765-00a0c91e6bf6
	 */
	PropUid = "UID"

	/*
	   Purpose:  To give a global meaning to a local PID source identifier.
	   Value type:  A semicolon-separated pair of values.  The first field
	      is a small integer corresponding to the second field of a PID
	      parameter instance.  The second field is a URI.  The "uuid" URN
	      namespace defined in [RFC4122] is particularly well suited to this
	      task, but other URI schemes MAY be used.
	   Cardinality:  *
	   ABNF:
	     CLIENTPIDMAP-param = any-param
	     CLIENTPIDMAP-value = 1*DIGIT ";" URI
	   Example:
	     TEL;PID=3.1,4.2;VALUE=uri:tel:+1-555-555-5555
	     EMAIL;PID=4.1,5.2:jdoe@example.com
	     CLIENTPIDMAP:1;urn:uuid:3df403f4-5924-4bb7-b077-3c711d9eb34b
	     CLIENTPIDMAP:2;urn:uuid:d89c9c7a-2e1b-4832-82de-7e992d95faa5
	 */
	PropClientPidmap = "CLIENTPIDMAP"

	/*
	   Purpose:  To specify a uniform resource locator associated with the
	      object to which the vCard refers.  Examples for individuals
	      include personal web sites, blogs, and social networking site
	      identifiers.
	   Cardinality:  *
	   Value type:  A single uri value.
	   ABNF:
	     URL-param = "VALUE=uri" / pid-param / pref-param / type-param
	               / mediatype-param / altid-param / any-param
	     URL-value = URI
	   Example:
	           URL:http://example.org/restaurant.french/~chezchic.html
	 */
	PropUrl = "URL"

	/*
	   Purpose:  To specify the version of the vCard specification used to
	      format this vCard.
	   Value type:  A single text value.
	   Cardinality:  1
	   ABNF:
	     VERSION-param = "VALUE=text" / any-param
	     VERSION-value = "4.0"
	   Example:
	           VERSION:4.0
	 */
	PropVersion = "VERSION"

	//Security Properties
	/*
	   These properties are concerned with the security of communication
	   pathways or access to the vCard.
	 */

	/*
	   Purpose:  To specify a public key or authentication certificate
	      associated with the object that the vCard represents.
	   Value type:  A single URI.  It can also be reset to a text value.
	   Cardinality:  *
	   ABNF:
	     KEY-param = KEY-uri-param / KEY-text-param
	     KEY-value = KEY-uri-value / KEY-text-value
	       ; Value and parameter MUST match.

	     KEY-uri-param = "VALUE=uri" / mediatype-param
	     KEY-uri-value = URI

	     KEY-text-param = "VALUE=text"
	     KEY-text-value = text

	     KEY-param =/ altid-param / pid-param / pref-param / type-param
	                / any-param
	   Examples:
	     KEY:http://www.example.com/keys/jdoe.cer

	     KEY;MEDIATYPE=application/pgp-keys:ftp://example.com/keys/jdoe

	     KEY:data:application/pgp-keys;base64,MIICajCCAdOgAwIBAgICBE
	      UwDQYJKoZIhvcNAQEEBQAwdzELMAkGA1UEBhMCVVMxLDAqBgNVBAoTI05l
	      <... remainder of base64-encoded data ...>
	*/
	PropKey = "KEY"

	//Calendar Properties
	/*
	These properties are further specified in [RFC2739].
	 */

	/*
	   Purpose:  To specify the URI for the busy time associated with the
	      object that the vCard represents.
	   Value type:  A single URI value.
	   Cardinality:  *
	   ABNF:
	     FBURL-param = "VALUE=uri" / pid-param / pref-param / type-param
	                 / mediatype-param / altid-param / any-param
	     FBURL-value = URI
	   Examples:
	     FBURL;PREF=1:http://www.example.com/busy/janedoe
	     FBURL;MEDIATYPE=text/calendar:ftp://example.com/busy/project-a.ifb
	 */
	PropFBurl = "FBURL"

	/*
	   Purpose:  To specify the calendar user address [RFC5545] to which a
	      scheduling request [RFC5546] should be sent for the object
	      represented by the vCard.
	   Value type:  A single URI value.
	   Cardinality:  *
	   ABNF:
	     CALADRURI-param = "VALUE=uri" / pid-param / pref-param / type-param
	                     / mediatype-param / altid-param / any-param
	     CALADRURI-value = URI
	   Example:
	     CALADRURI;PREF=1:mailto:janedoe@example.com
	     CALADRURI:http://example.com/calendar/jdoe
	 */
	PropCaladruri = "CALADRURI"

	/*
	   Purpose:  To specify the URI for a calendar associated with the
	      object represented by the vCard.
	   Value type:  A single URI value.
	   Cardinality:  *
	   ABNF:
	     CALURI-param = "VALUE=uri" / pid-param / pref-param / type-param
	                  / mediatype-param / altid-param / any-param
	     CALURI-value = URI
	   Examples:
	     CALURI;PREF=1:http://cal.example.com/calA
	     CALURI;MEDIATYPE=text/calendar:ftp://ftp.example.com/calA.ics
	 */
	PropCalUri = "CALURI"

)

const(
	KindIndividual = "individual"
	KindGroup = "group"
	KindOrg = "org"
	KindLocation = "location"
)

const (
	SexUnspecified  = ""
	SexFemale       = "F"
	SexMale         = "M"
	SexOther        = "O"
	SexNone         = "N"
	SexUnknown      = "U"
)




