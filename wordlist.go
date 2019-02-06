package pgpwords

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

var wordsEven = [256]string{
	"aardvark", "absurd", "accrue", "acme", "adrift",
	"adult", "afflict", "ahead", "aimless", "Algol",
	"allow", "alone", "ammo", "ancient", "apple",
	"artist", "assume", "Athens", "atlas", "Aztec",
	"baboon", "backfield", "backward", "banjo", "beaming",
	"bedlamp", "beehive", "beeswax", "befriend", "Belfast",
	"berserk", "billiard", "bison", "blackjack", "blockade",
	"blowtorch", "bluebird", "bombast", "bookshelf", "brackish",
	"breadline", "breakup", "brickyard", "briefcase", "Burbank",
	"button", "buzzard", "cement", "chairlift", "chatter",
	"checkup", "chisel", "choking", "chopper", "Christmas",
	"clamshell", "classic", "classroom", "cleanup", "clockwork",
	"cobra", "commence", "concert", "cowbell", "crackdown",
	"cranky", "crowfoot", "crucial", "crumpled", "crusade",
	"cubic", "dashboard", "deadbolt", "deckhand", "dogsled",
	"dragnet", "drainage", "dreadful", "drifter", "dropper",
	"drumbeat", "drunken", "Dupont", "dwelling", "eating",
	"edict", "egghead", "eightball", "endorse", "endow",
	"enlist", "erase", "escape", "exceed", "eyeglass",
	"eyetooth", "facial", "fallout", "flagpole", "flatfoot",
	"flytrap", "fracture", "framework", "freedom", "frighten",
	"gazelle", "Geiger", "glitter", "glucose", "goggles",
	"goldfish", "gremlin", "guidance", "hamlet", "highchair",
	"hockey", "indoors", "indulge", "inverse", "involve",
	"island", "jawbone", "keyboard", "kickoff", "kiwi",
	"klaxon", "locale", "lockup", "merit", "minnow",
	"miser", "Mohawk", "mural", "music", "necklace",
	"Neptune", "newborn", "nightbird", "Oakland", "obtuse",
	"offload", "optic", "orca", "payday", "peachy",
	"pheasant", "physique", "playhouse", "Pluto", "preclude",
	"prefer", "preshrunk", "printer", "prowler", "pupil",
	"puppy", "python", "quadrant", "quiver", "quota",
	"ragtime", "ratchet", "rebirth", "reform", "regain",
	"reindeer", "rematch", "repay", "retouch", "revenge",
	"reward", "rhythm", "ribcage", "ringbolt", "robust",
	"rocker", "ruffled", "sailboat", "sawdust", "scallion",
	"scenic", "scorecard", "Scotland", "seabird", "select",
	"sentence", "shadow", "shamrock", "showgirl", "skullcap",
	"skydive", "slingshot", "slowdown", "snapline", "snapshot",
	"snowcap", "snowslide", "solo", "southward", "soybean",
	"spaniel", "spearhead", "spellbind", "spheroid", "spigot",
	"spindle", "spyglass", "stagehand", "stagnate", "stairway",
	"standard", "stapler", "steamship", "sterling", "stockman",
	"stopwatch", "stormy", "sugar", "surmount", "suspense",
	"sweatband", "swelter", "tactics", "talon", "tapeworm",
	"tempest", "tiger", "tissue", "tonic", "topmost",
	"tracker", "transit", "trauma", "treadmill", "Trojan",
	"trouble", "tumor", "tunnel", "tycoon", "uncut",
	"unearth", "unwind", "uproot", "upset", "upshot",
	"vapor", "village", "virus", "Vulcan", "waffle",
	"wallet", "watchword", "wayside", "willow", "woodlark",
	"Zulu"}

var wordsOdd = [256]string{
	"adroitness", "adviser", "aftermath", "aggregate", "alkali",
	"almighty", "amulet", "amusement", "antenna", "applicant",
	"Apollo", "armistice", "article", "asteroid", "Atlantic",
	"atmosphere", "autopsy", "Babylon", "backwater", "barbecue",
	"belowground", "bifocals", "bodyguard", "bookseller",
	"borderline", "bottomless", "Bradbury", "bravado",
	"Brazilian", "breakaway", "Burlington", "businessman",
	"butterfat", "Camelot", "candidate", "cannonball",
	"Capricorn", "caravan", "caretaker", "celebrate",
	"cellulose", "certify", "chambermaid", "Cherokee",
	"Chicago", "clergyman", "coherence", "combustion",
	"commando", "company", "component", "concurrent",
	"confidence", "conformist", "congregate", "consensus",
	"consulting", "corporate", "corrosion", "councilman",
	"crossover", "crucifix", "cumbersome", "customer", "Dakota",
	"decadence", "December", "decimal", "designing", "detector",
	"detergent", "determine", "dictator", "dinosaur",
	"direction", "disable", "disbelief", "disruptive",
	"distortion", "document", "embezzle", "enchanting",
	"enrollment", "enterprise", "equation", "equipment",
	"escapade", "Eskimo", "everyday", "examine", "existence",
	"exodus", "fascinate", "filament", "finicky", "forever",
	"fortitude", "frequency", "gadgetry", "Galveston",
	"getaway", "glossary", "gossamer", "graduate", "gravity",
	"guitarist", "hamburger", "Hamilton", "handiwork",
	"hazardous", "headwaters", "hemisphere", "hesitate",
	"hideaway", "holiness", "hurricane", "hydraulic",
	"impartial", "impetus", "inception", "indigo", "inertia",
	"infancy", "inferno", "informant", "insincere", "insurgent",
	"integrate", "intention", "inventive", "Istanbul",
	"Jamaica", "Jupiter", "leprosy", "letterhead", "liberty",
	"maritime", "matchmaker", "maverick", "Medusa", "megaton",
	"microscope", "microwave", "midsummer", "millionaire",
	"miracle", "misnomer", "molasses", "molecule", "Montana",
	"monument", "mosquito", "narrative", "nebula", "newsletter",
	"Norwegian", "October", "Ohio", "onlooker", "opulent",
	"Orlando", "outfielder", "Pacific", "pandemic", "Pandora",
	"paperweight", "paragon", "paragraph", "paramount",
	"passenger", "pedigree", "Pegasus", "penetrate",
	"perceptive", "performance", "pharmacy", "phonetic",
	"photograph", "pioneer", "pocketful", "politeness",
	"positive", "potato", "processor", "provincial",
	"proximate", "puberty", "publisher", "pyramid", "quantity",
	"racketeer", "rebellion", "recipe", "recover", "repellent",
	"replica", "reproduce", "resistor", "responsive",
	"retraction", "retrieval", "retrospect", "revenue",
	"revival", "revolver", "sandalwood", "sardonic", "Saturday",
	"savagery", "scavenger", "sensation", "sociable",
	"souvenir", "specialist", "speculate", "stethoscope",
	"stupendous", "supportive", "surrender", "suspicious",
	"sympathy", "tambourine", "telephone", "therapist",
	"tobacco", "tolerance", "tomorrow", "torpedo", "tradition",
	"travesty", "trombonist", "truncated", "typewriter",
	"ultimate", "undaunted", "underfoot", "unicorn", "unify",
	"universe", "unravel", "upcoming", "vacancy", "vagabond",
	"vertigo", "Virginia", "visitor", "vocalist", "voyager",
	"warranty", "Waterloo", "whimsical", "Wichita",
	"Wilmington", "Wyoming", "yesteryear", "Yucatan"}

// File is a structure conttaining information on the file that is to be processed
type File struct {
	Filename        string
	Hashmethod      string
	Binchecksumdata []byte
}

// Binchecksum will create a binary checksum of a file
func (pgpfn *File) Binchecksum() []byte {
	if len(pgpfn.Binchecksumdata) > 0 {
		return pgpfn.Binchecksumdata
	}

	h := getHashObject(pgpfn.Hashmethod)
	f, err := os.Open(pgpfn.Filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to open file: %s", pgpfn.Filename))
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		panic(fmt.Sprintf("Unable to read file: %s", pgpfn.Filename))
	}

	pgpfn.Binchecksumdata = h.Sum(nil)
	return pgpfn.Binchecksumdata
}

// Hexchecksum will create a hexidecemal output of a files checksum
func (pgpfn *File) Hexchecksum() string {
	chk := pgpfn.Binchecksum()
	return hex.EncodeToString(chk)
}

// getHashObject will in time select and return the correct hash object
// For now i only require sha256
func getHashObject(hashtype string) hash.Hash {
	return sha256.New()
}

// WordList will iterate over the hex string and select the correct words
func (pgpfn File) WordList() []string {
	binchecksum := pgpfn.Binchecksum()
	var returndata []string

	for i := 0; i < len(binchecksum); i++ {
		if isOdd(i) {
			// fmt.Printf("%s ", wordsOdd[int(binchecksum[i])])
			returndata = append(returndata, wordsOdd[int(binchecksum[i])])

		} else {
			// fmt.Printf("%s ", wordsEven[int(binchecksum[i])])
			returndata = append(returndata, wordsEven[int(binchecksum[i])])

		}
	}
	return returndata
}

func isOdd(i int) bool {
	if i == 0 {
		return false
	}

	if i%2 == 0 {
		return false
	}

	return true
}
