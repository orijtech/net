// generated by go run gen.go; DO NOT EDIT

package publicsuffix

const version = "subset of publicsuffix.org's effective_tld_names.dat, hg revision 05b11a8d1ace (2012-11-09)"

const (
	nodesBitsChildren   = 9
	nodesBitsNodeType   = 2
	nodesBitsTextOffset = 15
	nodesBitsTextLength = 6

	childrenBitsWildcard = 1
	childrenBitsHi       = 14
	childrenBitsLo       = 14
)

const (
	nodeTypeNormal     = 0
	nodeTypeException  = 1
	nodeTypeParentOnly = 2
)

// numTLD is the number of top level domains.
const numTLD = 8

// Text is the combined text of all labels.
const text = "clubafukuchiyamashinacionakagyorgamecongresodelalengua3govgvin-a" +
	"ddretinagaokakyotambainelip6irisakyotanabeducityjetjoyoyamazakit" +
	"ajpblogspotkizuridebizwkumiyamakyotangobiernoelectronicomilkyoto" +
	"minamiyamashiromiyazurnantanational-library-scotlandmukobenlschi" +
	"gashiyamaizurujitawarapromocionetseikameokamodxn--czrw28british-" +
	"libraryawatarparliamentwazukayabe164xn--p1aidvxn--uc0atvxn--zf0a" +
	"o64a"

// nodes is the list of nodes. Each node is represented as a uint32, which
// encodes the node's children (as an index into the children array), wildcard
// bit, node type and text.
//
// In the //-comment after each node's data, the nodes indexes of the children
// are formatted as (n0x1234-n0x1256), with * denoting the wildcard bit. The
// nodeType is printed as + for normal, ! for exception, and o for parent-only
// nodes that have children but don't match a domain label in their own right.
//
// The layout within the uint32, from MSB to LSB, is:
//	[ 0 bits] unused
//	[ 9 bits] children index
//	[ 2 bits] nodeType
//	[15 bits] text index
//	[ 6 bits] text length
var nodes = [...]uint32{
	0x01001242, // n0x0000 c0x0002 (n0x0008-n0x000e)  + ao
	0x01c03a02, // n0x0001 c0x0003 (n0x000e-n0x0018)* o ar
	0x02c052c4, // n0x0002 c0x0005 (n0x0019-n0x001f)  o arpa
	0x03002042, // n0x0003 c0x0006 (n0x001f-n0x0021)  + jp
	0x04805582, // n0x0004 c0x0009 (n0x0041-n0x004f)  + tw
	0x05400182, // n0x0005 c0x000a (n0x004f-n0x005a)* o uk
	0x00005908, // n0x0006 c0x0000 (---------------)  + xn--p1ai
	0x00c02542, // n0x0007 c0x0001 (---------------)* o zw
	0x00000902, // n0x0008 c0x0000 (---------------)  + co
	0x00001a42, // n0x0009 c0x0000 (---------------)  + ed
	0x00000e82, // n0x000a c0x0000 (---------------)  + gv
	0x00001b42, // n0x000b c0x0000 (---------------)  + it
	0x00002142, // n0x000c c0x0000 (---------------)  + og
	0x00002082, // n0x000d c0x0000 (---------------)  + pb
	0x02402d83, // n0x000e c0x0004 (n0x0018-n0x0019)  o com
	0x00200913, // n0x000f c0x0000 (---------------)  ! congresodelalengua3
	0x00201a44, // n0x0010 c0x0000 (---------------)  ! educ
	0x00202953, // n0x0011 c0x0000 (---------------)  ! gobiernoelectronico
	0x00200885, // n0x0012 c0x0000 (---------------)  ! mecon
	0x002004c6, // n0x0013 c0x0000 (---------------)  ! nacion
	0x00202d03, // n0x0014 c0x0000 (---------------)  ! nic
	0x00204589, // n0x0015 c0x0000 (---------------)  ! promocion
	0x00201086, // n0x0016 c0x0000 (---------------)  ! retina
	0x00200083, // n0x0017 c0x0000 (---------------)  ! uba
	0x000020c8, // n0x0018 c0x0000 (---------------)  + blogspot
	0x00005804, // n0x0019 c0x0000 (---------------)  + e164
	0x00000f07, // n0x001a c0x0000 (---------------)  + in-addr
	0x00001643, // n0x001b c0x0000 (---------------)  + ip6
	0x00001704, // n0x001c c0x0000 (---------------)  + iris
	0x00002383, // n0x001d c0x0000 (---------------)  + uri
	0x00003503, // n0x001e c0x0000 (---------------)  + urn
	0x03c03d84, // n0x001f c0x0007 (n0x0021-n0x0022)* o kobe
	0x04002ec5, // n0x0020 c0x0008 (n0x0022-n0x0041)  + kyoto
	0x00201b04, // n0x0021 c0x0000 (---------------)  ! city
	0x00005705, // n0x0022 c0x0000 (---------------)  + ayabe
	0x0000014b, // n0x0023 c0x0000 (---------------)  + fukuchiyama
	0x00003f8b, // n0x0024 c0x0000 (---------------)  + higashiyama
	0x00002403, // n0x0025 c0x0000 (---------------)  + ide
	0x00001543, // n0x0026 c0x0000 (---------------)  + ine
	0x00001cc4, // n0x0027 c0x0000 (---------------)  + joyo
	0x00004907, // n0x0028 c0x0000 (---------------)  + kameoka
	0x00004a44, // n0x0029 c0x0000 (---------------)  + kamo
	0x00001f44, // n0x002a c0x0000 (---------------)  + kita
	0x000022c4, // n0x002b c0x0000 (---------------)  + kizu
	0x000025c8, // n0x002c c0x0000 (---------------)  + kumiyama
	0x00001348, // n0x002d c0x0000 (---------------)  + kyotamba
	0x00001849, // n0x002e c0x0000 (---------------)  + kyotanabe
	0x000027c8, // n0x002f c0x0000 (---------------)  + kyotango
	0x000041c7, // n0x0030 c0x0000 (---------------)  + maizuru
	0x00003006, // n0x0031 c0x0000 (---------------)  + minami
	0x0000300f, // n0x0032 c0x0000 (---------------)  + minamiyamashiro
	0x000033c6, // n0x0033 c0x0000 (---------------)  + miyazu
	0x00003d04, // n0x0034 c0x0000 (---------------)  + muko
	0x0000118a, // n0x0035 c0x0000 (---------------)  + nagaokakyo
	0x00000607, // n0x0036 c0x0000 (---------------)  + nakagyo
	0x00003586, // n0x0037 c0x0000 (---------------)  + nantan
	0x00001d89, // n0x0038 c0x0000 (---------------)  + oyamazaki
	0x000017c5, // n0x0039 c0x0000 (---------------)  + sakyo
	0x00004845, // n0x003a c0x0000 (---------------)  + seika
	0x00001906, // n0x003b c0x0000 (---------------)  + tanabe
	0x00004343, // n0x003c c0x0000 (---------------)  + uji
	0x00004349, // n0x003d c0x0000 (---------------)  + ujitawara
	0x000055c6, // n0x003e c0x0000 (---------------)  + wazuka
	0x00000309, // n0x003f c0x0000 (---------------)  + yamashina
	0x00005186, // n0x0040 c0x0000 (---------------)  + yawata
	0x000020c8, // n0x0041 c0x0000 (---------------)  + blogspot
	0x00000004, // n0x0042 c0x0000 (---------------)  + club
	0x00002d83, // n0x0043 c0x0000 (---------------)  + com
	0x00002484, // n0x0044 c0x0000 (---------------)  + ebiz
	0x00001a43, // n0x0045 c0x0000 (---------------)  + edu
	0x00000804, // n0x0046 c0x0000 (---------------)  + game
	0x00000dc3, // n0x0047 c0x0000 (---------------)  + gov
	0x00005ac3, // n0x0048 c0x0000 (---------------)  + idv
	0x00002e03, // n0x0049 c0x0000 (---------------)  + mil
	0x00004783, // n0x004a c0x0000 (---------------)  + net
	0x00000783, // n0x004b c0x0000 (---------------)  + org
	0x00004b8b, // n0x004c c0x0000 (---------------)  + xn--czrw28b
	0x00005b8a, // n0x004d c0x0000 (---------------)  + xn--uc0atv
	0x00005e0c, // n0x004e c0x0000 (---------------)  + xn--zf0ao64a
	0x002020c2, // n0x004f c0x0000 (---------------)  ! bl
	0x00204e0f, // n0x0050 c0x0000 (---------------)  ! british-library
	0x05c00902, // n0x0051 c0x000b (n0x005a-n0x005b)  o co
	0x00201c03, // n0x0052 c0x0000 (---------------)  ! jet
	0x00204ac3, // n0x0053 c0x0000 (---------------)  ! mod
	0x002036d9, // n0x0054 c0x0000 (---------------)  ! national-library-scotland
	0x00201583, // n0x0055 c0x0000 (---------------)  ! nel
	0x00202d03, // n0x0056 c0x0000 (---------------)  ! nic
	0x00203e83, // n0x0057 c0x0000 (---------------)  ! nls
	0x0020534a, // n0x0058 c0x0000 (---------------)  ! parliament
	0x00c03f03, // n0x0059 c0x0001 (---------------)* o sch
	0x000020c8, // n0x005a c0x0000 (---------------)  + blogspot
}

// children is the list of nodes' children, and the wildcard bit. If a node
// has no children then their children index will be 0 or 1, depending on the
// wildcard bit.
//
// The layout within the uint32, from MSB to LSB, is:
//	[ 3 bits] unused
//	[ 1 bits] wildcard bit
//	[14 bits] high nodes index (exclusive) of children
//	[14 bits] low nodes index (inclusive) of children
var children = [...]uint32{
	0x00000000, // c0x0000 (---------------)
	0x10000000, // c0x0001 (---------------)*
	0x00038008, // c0x0002 (n0x0008-n0x000e)
	0x1006000e, // c0x0003 (n0x000e-n0x0018)*
	0x00064018, // c0x0004 (n0x0018-n0x0019)
	0x0007c019, // c0x0005 (n0x0019-n0x001f)
	0x0008401f, // c0x0006 (n0x001f-n0x0021)
	0x10088021, // c0x0007 (n0x0021-n0x0022)*
	0x00104022, // c0x0008 (n0x0022-n0x0041)
	0x0013c041, // c0x0009 (n0x0041-n0x004f)
	0x1016804f, // c0x000a (n0x004f-n0x005a)*
	0x0016c05a, // c0x000b (n0x005a-n0x005b)
}