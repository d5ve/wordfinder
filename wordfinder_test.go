package main

import (
	"reflect"
	"strings"
	"testing"
)

var dict string
var words []string

func TestLoadWords(t *testing.T) {
	words = LoadWords()
	if len(words) < 1000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
	if len(words) > 1000000 {
		t.Fatal("Less than 1000 words loaded from dictionary")
	}
	dict = strings.Join(words, " ")
}

func TestFindWords(t *testing.T) {

	wordsTests := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"z", []string{}},
		{"dgo", []string{"do", "dog", "go", "god", "o", "od", "og"}},
		{"eariotnslc", tencommonchars()},
	}

	for _, tt := range wordsTests {
		t.Run(tt.input, func(t *testing.T) {
			got := FindWords(tt.input, dict)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindWords('%s') got %s want %s", tt.input, got, tt.expected)
			}
			got = FindWords2(tt.input, dict)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindWords2('%s') got %s want %s", tt.input, got, tt.expected)
			}
			got = FindWords3(tt.input, words)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindWords3('%s') got %s want %s", tt.input, got, tt.expected)
			}
			got = FindWords4(tt.input, dict)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindWords4('%s') got %s want %s", tt.input, got, tt.expected)
			}
		})
	}

}

var got []string

func BenchmarkFindWords(b *testing.B) {
	words := LoadWords()
	dict := strings.Join(words, " ")
	for n := 0; n < b.N; n++ {
		got = FindWords("eariotnslc", dict)
	}
}
func BenchmarkFindWords2(b *testing.B) {
	words := LoadWords()
	dict := strings.Join(words, " ")
	for n := 0; n < b.N; n++ {
		got = FindWords2("eariotnslc", dict)
	}
}
func BenchmarkFindWords3(b *testing.B) {
	words := LoadWords()
	for n := 0; n < b.N; n++ {
		got = FindWords3("eariotnslc", words)
	}
}
func BenchmarkFindWords4(b *testing.B) {
	words := LoadWords()
	dict := strings.Join(words, " ")
	for n := 0; n < b.N; n++ {
		got = FindWords4("eariotnslc", dict)
	}
}

// Results on my macos laptop for the ten most common letters, eariotnslc -
// giving 1631 results. Derived from the perl version.
func tencommonchars() []string {
	return []string{"a", "ace", "acer", "acerin", "acetin", "acetoin", "acetol", "acier", "acinose", "acis", "acle", "acne", "acoin", "acoine", "acone", "aconite", "acor", "acorn", "acre", "acrite", "acritol", "acrolein", "acron", "acrose", "act", "actin", "actine", "action", "actioner", "acton", "actor", "acts", "ae", "aeolic", "aeolis", "aeolist", "aeon", "aeonist", "aer", "aeric", "aero", "aes", "ai", "aiel", "ail", "aile", "aileron", "aint", "aion", "air", "aire", "airt", "aisle", "ait", "al", "alces", "alcine", "alco", "alcor", "ale", "alec", "alecost", "alectoris", "alectrion", "alen", "alert", "aletris", "alice", "alien", "alienor", "alin", "aline", "aliso", "alison", "alist", "alister", "alit", "alite", "aln", "alnico", "alnoite", "alo", "aloe", "aloetic", "aloin", "alois", "alone", "alose", "alsine", "also", "alt", "alter", "altern", "altin", "alto", "altrices", "altrose", "an", "ancestor", "ancile", "anerotic", "anes", "ani", "anice", "anil", "anile", "anis", "anise", "anisole", "anoetic", "anoil", "anole", "anoli", "anolis", "ansel", "anser", "ant", "ante", "antes", "anti", "antic", "anticor", "antler", "antoeci", "antre", "ao", "aorist", "aortic", "aotes", "ar", "arc", "arcite", "arcos", "arctos", "are", "arecolin", "areito", "arent", "ariel", "aries", "aril", "arion", "ariose", "ariot", "arise", "arisen", "arist", "aristol", "arite", "arles", "arline", "arn", "arne", "arni", "aro", "aroint", "arose", "arse", "arsenic", "arseno", "arsine", "arsino", "arsle", "arsoite", "arson", "arsonic", "arsonite", "art", "artel", "article", "artie", "as", "ascent", "asci", "asclent", "ascon", "ascot", "ase", "asiento", "asnort", "asor", "ast", "astelic", "aster", "asterin", "asterion", "astern", "astir", "astor", "at", "ate", "atelo", "aten", "ates", "ati", "atis", "atle", "atone", "atoner", "atonic", "atresic", "atroscine", "ca", "cain", "cairn", "cairo", "caite", "cal", "calite", "calor", "calorie", "caloris", "calorist", "calorite", "can", "cane", "canel", "canelo", "caner", "canis", "canistel", "canister", "canoe", "canoeist", "canso", "cant", "canter", "cantle", "canto", "cantor", "cantoris", "car", "care", "carest", "caret", "caries", "cariole", "carl", "carlet", "carlie", "carlin", "carline", "carlist", "carlo", "carlos", "carlot", "carls", "carneol", "carnose", "caro", "carol", "carole", "caroli", "carolin", "caroline", "carone", "carotin", "carse", "carsten", "cart", "carte", "cartel", "carton", "case", "casein", "casel", "caser", "casern", "casino", "caslon", "cast", "caste", "caster", "castle", "castor", "castorin", "cat", "cate", "cater", "cation", "catlin", "ce", "cearin", "ceil", "celation", "celia", "celosia", "celsia", "celsian", "celt", "celtis", "censor", "censorial", "cent", "cental", "centiar", "cento", "central", "ceorl", "ceral", "ceras", "cerasin", "ceration", "ceria", "cerin", "cerion", "cern", "cero", "cerotin", "certain", "certis", "certosina", "cest", "cestrian", "ceti", "cetin", "cetonia", "cine", "cinel", "cinter", "cion", "cise", "cist", "cista", "cistae", "cistern", "cisterna", "cisternal", "cit", "cite", "citer", "citole", "citral", "citrean", "citron", "claire", "clan", "clare", "claret", "clarin", "clarinet", "clarion", "clarionet", "clarist", "claro", "clart", "clat", "clean", "clear", "cleat", "client", "cline", "clint", "clio", "cliona", "clione", "clit", "clite", "clites", "cloister", "cloit", "clone", "close", "closen", "closer", "closet", "closter", "clot", "clote", "coal", "coaler", "coalite", "coan", "coarse", "coarsen", "coast", "coaster", "coat", "coater", "coati", "coatie", "coe", "coelar", "coelia", "coelian", "coelin", "coil", "coiler", "coin", "coiner", "cointer", "coir", "coistrel", "coital", "col", "cola", "colan", "colane", "colarin", "colate", "cole", "coli", "colias", "colin", "colinear", "colt", "colter", "con", "conal", "cone", "coner", "cones", "consertal", "conte", "conter", "contise", "contra", "contrail", "cor", "cora", "coraise", "coral", "coralist", "core", "corial", "corin", "corn", "cornea", "corneal", "cornel", "cornelia", "cornet", "corsaint", "corse", "corset", "corsie", "corsite", "corta", "cortes", "cortin", "cortina", "cos", "cosalite", "coseat", "coset", "cosine", "cost", "costa", "costal", "costar", "costean", "coster", "costrel", "cot", "cote", "cotesian", "cotise", "cotrine", "crain", "cran", "crane", "crants", "crate", "crea", "creant", "creat", "creation", "crena", "creolian", "creolin", "cresol", "cresolin", "crest", "creta", "cretan", "cretin", "cretion", "crile", "crin", "crinal", "crine", "crinet", "crinose", "cris", "crista", "cro", "croat", "crone", "cronet", "crosa", "crotal", "crotaline", "crotin", "ea", "ean", "ear", "earl", "earn", "east", "eat", "eats", "eciton", "eclair", "eclat", "ectal", "el", "elain", "elastic", "elastin", "elation", "elator", "eli", "elia", "elian", "elias", "elinor", "eliot", "elisor", "elon", "elric", "els", "elsa", "elsin", "elt", "en", "enact", "enactor", "enatic", "encist", "encoil", "eniac", "enlist", "enoil", "enol", "enolic", "enos", "enrol", "ens", "enstar", "entail", "ental", "entia", "entoil", "entosarc", "entrail", "entrails", "eoan", "eon", "eosin", "er", "era", "eral", "eranist", "erasion", "eria", "erian", "eric", "erica", "erical", "ernst", "eros", "erotic", "erotica", "erotical", "ers", "es", "esca", "escalin", "escolar", "escorial", "escort", "escrol", "estoc", "estrin", "estriol", "eta", "etalon", "etna", "i", "ian", "iao", "ice", "icon", "ie", "ila", "ileac", "ileon", "ilot", "in", "inca", "incase", "incast", "incest", "inclose", "increst", "inert", "inlet", "ino", "inro", "insea", "insect", "insecta", "insert", "inset", "insolate", "insole", "instar", "inter", "into", "io", "ion", "ione", "iota", "ira", "iran", "irascent", "irate", "ire", "irena", "iron", "irone", "is", "isle", "islet", "isleta", "islot", "iso", "isocrat", "isolate", "isotac", "israel", "ist", "istle", "it", "ita", "italon", "itea", "iten", "iter", "ito", "its", "la", "lac", "lace", "lacer", "lacet", "lacis", "lacto", "lactone", "lactose", "laet", "laeti", "laetic", "lai", "laic", "lain", "laine", "laiose", "lair", "lairstone", "lan", "lance", "lancer", "lances", "lancet", "lanciers", "lane", "lanose", "lant", "lao", "lar", "larcenist", "lari", "larin", "larine", "lars", "las", "laser", "lasi", "last", "laster", "lastre", "lat", "late", "laten", "later", "latices", "latin", "latiner", "lation", "latrine", "latris", "latro", "latron", "lea", "lean", "leant", "lear", "learn", "learnt", "least", "leat", "lection", "lector", "lei", "len", "lena", "lenca", "lenis", "leno", "lenora", "lens", "lent", "lentisc", "lentisco", "lento", "lentor", "leo", "leon", "leonis", "leonist", "leora", "ler", "lerot", "les", "lesion", "lest", "let", "leto", "li", "liar", "lias", "licensor", "licorn", "licorne", "lictor", "lie", "lien", "lienor", "lier", "lin", "lina", "line", "linea", "linear", "liner", "linet", "lino", "linos", "lint", "linter", "lion", "lionet", "lira", "lirate", "lire", "lis", "lisa", "lise", "list", "listen", "lister", "listera", "lit", "litas", "lite", "liter", "litra", "litsea", "lo", "loa", "loan", "loaner", "loca", "locarnist", "locarnite", "locate", "loci", "locrian", "locrine", "loin", "loir", "lois", "loiter", "lone", "lonicera", "lontar", "lora", "loran", "lorate", "lore", "loren", "lori", "loric", "lorica", "loricate", "lorien", "loris", "lorn", "lors", "lose", "loser", "lost", "lot", "lota", "lotase", "lote", "lotic", "lots", "na", "nace", "nacre", "nacrite", "nae", "nael", "nail", "nailer", "naio", "nair", "nais", "naos", "nar", "narcist", "narcose", "nares", "naric", "nasi", "nasrol", "nast", "nastic", "nat", "nate", "nates", "natr", "ne", "nea", "neal", "neat", "necator", "nectar", "nectria", "nei", "neil", "neist", "neo", "neri", "nerita", "neroic", "nesiot", "neslia", "nest", "nestor", "net", "neti", "ni", "nias", "nice", "nicol", "nicolas", "niels", "nil", "nile", "nilot", "nils", "niota", "nirles", "nit", "niter", "nito", "nitro", "no", "noa", "noel", "noetic", "noetics", "noil", "noiler", "noir", "noise", "nor", "nora", "norate", "noreast", "nori", "noria", "noric", "norie", "norite", "norse", "norsel", "nose", "noser", "nostic", "nostril", "not", "notal", "note", "noter", "notice", "noticer", "o", "oar", "oaric", "oast", "oat", "oaten", "ocean", "ocneria", "ocrea", "octan", "octane", "octans", "octine", "oe", "oecist", "oer", "oes", "oestrian", "oestrin", "oil", "oilcan", "oiler", "oint", "ole", "olea", "oleic", "olein", "olena", "olent", "on", "ona", "onca", "once", "oncia", "one", "oner", "onliest", "ons", "onset", "ontal", "ontaric", "or", "ora", "oracle", "oral", "oralist", "orant", "orate", "orc", "orca", "orcanet", "orcein", "orcin", "ore", "oreas", "orias", "oriel", "orient", "oriental", "orle", "orlean", "orleanist", "orleans", "orlet", "orna", "ornate", "ornis", "orsel", "ort", "ortalis", "os", "osc", "oscan", "oscar", "oscin", "oscine", "ose", "osela", "osier", "osteal", "ostein", "ostial", "ostic", "ostracine", "ostrea", "otarine", "otic", "otis", "ra", "race", "racist", "racon", "rail", "rain", "rais", "raise", "ran", "rance", "rancel", "rane", "rani", "ransel", "rant", "ras", "rase", "rasen", "rasion", "rastle", "rat", "rate", "ratel", "ratine", "ratio", "ration", "ratline", "re", "rea", "react", "reaction", "real", "realist", "reason", "recant", "recast", "recital", "recoal", "recoast", "recoat", "recoil", "recoin", "recon", "rect", "recta", "rectal", "recti", "rection", "recto", "rein", "reina", "reins", "reis", "reit", "rel", "relais", "relast", "relation", "reliant", "relic", "relict", "relist", "reloan", "relost", "relot", "renail", "renal", "rent", "rental", "reoil", "resail", "resalt", "rescan", "resin", "resina", "resinol", "reslot", "resoil", "rest", "restain", "restio", "ret", "retail", "retain", "retan", "retia", "retial", "retin", "retina", "retinal", "retinol", "ria", "rial", "riant", "ric", "rice", "rictal", "rie", "rile", "rine", "rinse", "rio", "riot", "rise", "risen", "rist", "rit", "rita", "rite", "ro", "roan", "roast", "roc", "rocta", "roe", "roi", "roil", "roist", "roit", "role", "ron", "roncet", "rone", "rosa", "rosal", "rosalie", "rosaline", "rose", "roseal", "rosel", "roset", "rosetan", "rosin", "rosinate", "rosine", "rostel", "rot", "rota", "rotal", "rotan", "rote", "rotse", "sa", "sac", "saco", "sacro", "sai", "saic", "sail", "sailer", "sailor", "sain", "saint", "sair", "saite", "sal", "sale", "salic", "salicorn", "salient", "saline", "salite", "salon", "salt", "salten", "salter", "saltern", "saltier", "saltine", "san", "sanct", "sane", "sanicle", "sant", "santir", "santo", "sao", "sar", "sarcine", "sarcle", "sarcoline", "sarcolite", "sare", "sari", "saron", "saronic", "sart", "sat", "sate", "satieno", "satin", "satine", "satire", "satron", "scale", "scaler", "scaloni", "scalt", "scan", "scant", "scantle", "scar", "scare", "scarlet", "scarn", "scart", "scat", "sceat", "scena", "scenario", "scent", "scian", "scient", "scintle", "scintler", "scion", "sciot", "sclate", "sclater", "scler", "sclera", "scleria", "sclerotia", "scolia", "scolite", "scone", "score", "scoria", "scoriae", "scorn", "scot", "scotale", "scote", "scoter", "scotia", "scrae", "scran", "scrat", "scrin", "scrine", "scrota", "scrotal", "se", "sea", "seal", "sean", "sear", "seat", "seatron", "sec", "secalin", "secant", "sect", "section", "sectional", "sector", "sectoral", "sectorial", "seit", "selictar", "selina", "selion", "selt", "sen", "senator", "senci", "senior", "senlac", "sent", "ser", "sera", "serai", "serail", "seral", "sercial", "seri", "serial", "serian", "seric", "serin", "serio", "seriola", "sero", "serolin", "seron", "serotina", "serotinal", "sert", "serta", "set", "seta", "setal", "seton", "si", "sia", "sial", "sic", "sice", "sicel", "sie", "siena", "sier", "sil", "silane", "sile", "silen", "silent", "silo", "silt", "sin", "sina", "sinae", "sinal", "since", "sine", "sinter", "sinto", "sintoc", "siol", "sion", "sir", "sire", "siren", "siroc", "sit", "sita", "sitao", "sitar", "site", "sla", "slae", "slain", "slainte", "slait", "slane", "slant", "slare", "slart", "slat", "slate", "slater", "slent", "slice", "slicer", "sline", "slirt", "slit", "slite", "sloan", "sloe", "slon", "slone", "slot", "slote", "snail", "snare", "snarl", "snirl", "snirt", "snirtle", "snite", "snore", "snort", "snortle", "snot", "so", "soar", "soc", "soce", "social", "societal", "socle", "soe", "soil", "sol", "sola", "solace", "solacer", "solan", "solar", "solate", "sole", "solea", "solen", "solent", "soler", "solera", "soli", "son", "sonar", "soneri", "sonic", "sonrai", "sora", "soral", "sore", "sori", "sorite", "sorn", "sort", "sortal", "sortie", "sot", "soter", "soterial", "sotie", "sotnia", "sri", "st", "stain", "stainer", "staio", "stair", "stale", "stan", "stance", "stane", "star", "stare", "starn", "starnel", "starnie", "steal", "stean", "stearic", "stearin", "stein", "stela", "stelai", "stelar", "sten", "stenar", "stencil", "steno", "stercolin", "steri", "steric", "sterin", "stern", "sterna", "sternal", "sterno", "stero", "sterol", "stile", "stine", "stion", "stir", "stoa", "stoic", "stoical", "stola", "stolae", "stole", "stolen", "stone", "stoner", "store", "stra", "strae", "strain", "stre", "stria", "striae", "strial", "striola", "striolae", "stroil", "strone", "ta", "tacso", "tae", "tael", "taen", "tai", "tail", "tailer", "tailor", "tain", "taino", "tairn", "taise", "tal", "talc", "talcer", "talcose", "tale", "taler", "tales", "tali", "talion", "talis", "talon", "talonic", "talose", "tan", "tancel", "tane", "tanier", "tano", "tanrec", "tao", "taos", "tar", "tare", "tari", "tarie", "tarin", "tarn", "taro", "taroc", "tars", "tarse", "tarsi", "tasco", "te", "tea", "teal", "tean", "tear", "tec", "teca", "tecali", "tecla", "teco", "tecon", "teian", "teil", "telar", "teli", "telic", "telson", "telsonic", "ten", "tenai", "tenio", "tenor", "tensor", "tera", "teras", "tercia", "tercio", "teri", "tern", "terna", "ternal", "tersion", "ti", "tiao", "tiar", "tic", "tical", "tice", "ticer", "tie", "tien", "tier", "til", "tile", "tiler", "tin", "tina", "tincal", "tine", "tinea", "tineal", "tino", "tinoceras", "tinosa", "tinsel", "tire", "tirl", "tirolean", "tisane", "tisar", "tlaco", "to", "toa", "tocsin", "toe", "toenail", "toi", "toil", "toiler", "toise", "tol", "tolan", "tolane", "tole", "ton", "tonal", "tone", "toner", "tonic", "tonsil", "tor", "tora", "toral", "toran", "torc", "torcel", "tore", "torenia", "toric", "torn", "tornal", "torse", "torsel", "torsile", "tra", "trace", "trail", "train", "tran", "trance", "treason", "trenail", "tri", "triace", "trial", "trias", "trica", "tricae", "trice", "tricosane", "triens", "trin", "trinal", "trine", "trinol", "trio", "triole", "triose", "troca", "troic", "tron", "trona", "tronc", "trone", "tsar", "tsia", "tsine", "tsoneca"}
}
