package main

var googleCountry = map[string]string{
	"com": "com",
	"ac":  "ac",
	"ad":  "ad",
	"ae":  "ae",
	"af":  "af",
	"ag":  "ag",
	"ai":  "ai",
	"al":  "al",
	"am":  "am",
	"ao":  "ao",
	"ar":  "ar",
	"as":  "as",
	"at":  "at",
	"au":  "au",
	"az":  "az",
	"ba":  "ba",
	"bd":  "bd",
	"be":  "be",
	"bf":  "bf",
	"bg":  "bg",
	"bh":  "bh",
	"bi":  "bi",
	"bj":  "bj",
	"bn":  "bn",
	"bo":  "bo",
	"br":  "br",
	"bs":  "bs",
	"bt":  "bt",
	"bw":  "bw",
	"by":  "by",
	"bz":  "bz",
	"ca":  "ca",
	"kh":  "kh",
	"cc":  "cc",
	"cd":  "cd",
	"cf":  "cf",
	"cat": "cat",
	"cg":  "cg",
	"ch":  "ch",
	"ci":  "ci",
	"ck":  "ck",
	"cl":  "cl",
	"cm":  "cm",
	"co":  "co",
	"cr":  "cr",
	"cu":  "cu",
	"cv":  "cv",
	"cy":  "cy",
	"cz":  "cz",
	"de":  "de",
	"dj":  "dj",
	"dk":  "dk",
	"dm":  "dm",
	"do":  "do",
	"dz":  "dz",
	"ec":  "ec",
	"ee":  "ee",
	"eg":  "eg",
	"es":  "es",
	"et":  "et",
	"fi":  "fi",
	"fj":  "fj",
	"fm":  "fm",
	"fr":  "fr",
	"ga":  "ga",
	"ge":  "ge",
	"gf":  "gf",
	"gg":  "gg",
	"gh":  "gh",
	"gi":  "gi",
	"gl":  "gl",
	"gm":  "gm",
	"gp":  "gp",
	"gr":  "gr",
	"gt":  "gt",
	"gy":  "gy",
	"hk":  "hk",
	"hn":  "hn",
	"hr":  "hr",
	"ht":  "ht",
	"hu":  "hu",
	"id":  "id",
	"iq":  "iq",
	"ie":  "ie",
	"il":  "il",
	"im":  "im",
	"in":  "in",
	"io":  "io",
	"is":  "is",
	"it":  "it",
	"je":  "je",
	"jm":  "jm",
	"jo":  "jo",
	"jp":  "jp",
	"ke":  "ke",
	"ki":  "ki",
	"kg":  "kg",
	"kr":  "kr",
	"kw":  "kw",
	"kz":  "kz",
	"la":  "la",
	"lb":  "lb",
	"lc":  "lc",
	"li":  "li",
	"lk":  "lk",
	"ls":  "ls",
	"lt":  "lt",
	"lu":  "lu",
	"lv":  "lv",
	"ly":  "ly",
	"ma":  "ma",
	"md":  "md",
	"me":  "me",
	"mg":  "mg",
	"mk":  "mk",
	"ml":  "ml",
	"mm":  "mm",
	"mn":  "mn",
	"ms":  "ms",
	"mt":  "mt",
	"mu":  "mu",
	"mv":  "mv",
	"mw":  "mw",
	"mx":  "mx",
	"my":  "my",
	"mz":  "mz",
	"na":  "na",
	"ne":  "ne",
	"nf":  "nf",
	"ng":  "ng",
	"ni":  "ni",
	"nl":  "nl",
	"no":  "no",
	"np":  "np",
	"nr":  "nr",
	"nu":  "nu",
	"nz":  "nz",
	"om":  "om",
	"pk":  "pk",
	"pa":  "pa",
	"pe":  "pe",
	"ph":  "ph",
	"pl":  "pl",
	"pg":  "pg",
	"pn":  "pn",
	"pr":  "pr",
	"ps":  "ps",
	"pt":  "pt",
	"py":  "py",
	"qa":  "qa",
	"ro":  "ro",
	"rs":  "rs",
	"ru":  "ru",
	"rw":  "rw",
	"sa":  "sa",
	"sb":  "sb",
	"sc":  "sc",
	"se":  "se",
	"sg":  "sg",
	"sh":  "sh",
	"si":  "si",
	"sk":  "sk",
	"sl":  "sl",
	"sn":  "sn",
	"sm":  "sm",
	"so":  "so",
	"st":  "st",
	"sr":  "sr",
	"sv":  "sv",
	"td":  "td",
	"tg":  "tg",
	"th":  "th",
	"tj":  "tj",
	"tk":  "tk",
	"tl":  "tl",
	"tm":  "tm",
	"to":  "to",
	"tn":  "tn",
	"tr":  "tr",
	"tt":  "tt",
	"tw":  "tw",
	"tz":  "tz",
	"ua":  "ua",
	"ug":  "ug",
	"uk":  "uk",
	"us":  "us",
	"uy":  "uy",
	"uz":  "uz",
	"vc":  "vc",
	"ve":  "ve",
	"vg":  "vg",
	"vi":  "vi",
	"vn":  "vn",
	"vu":  "vu",
	"ws":  "ws",
	"za":  "za",
	"zm":  "zm",
	"zw":  "zw",
}

var language = map[string]string{
	"af":         "af",
	"ach":        "ach",
	"ak":         "ak",
	"am":         "am",
	"ar":         "ar",
	"az":         "az",
	"be":         "be",
	"bem":        "bem",
	"bg":         "bg",
	"bh":         "bh",
	"bn":         "bn",
	"br":         "br",
	"bs":         "bs",
	"ca":         "ca",
	"chr":        "chr",
	"ckb":        "ckb",
	"co":         "co",
	"crs":        "crs",
	"cs":         "cs",
	"cy":         "cy",
	"da":         "da",
	"de":         "de",
	"ee":         "ee",
	"el":         "el",
	"en":         "en",
	"eo":         "eo",
	"es":         "es",
	"es-419":     "es-419",
	"et":         "et",
	"eu":         "eu",
	"fa":         "fa",
	"fi":         "fi",
	"fo":         "fo",
	"fr":         "fr",
	"fy":         "fy",
	"ga":         "ga",
	"gaa":        "gaa",
	"gd":         "gd",
	"gl":         "gl",
	"gn":         "gn",
	"gu":         "gu",
	"ha":         "ha",
	"haw":        "haw",
	"hi":         "hi",
	"hr":         "hr",
	"ht":         "ht",
	"hu":         "hu",
	"hy":         "hy",
	"ia":         "ia",
	"id":         "id",
	"ig":         "ig",
	"is":         "is",
	"it":         "it",
	"iw":         "iw",
	"ja":         "ja",
	"jw":         "jw",
	"ka":         "ka",
	"kg":         "kg",
	"kk":         "kk",
	"km":         "km",
	"kn":         "kn",
	"ko":         "ko",
	"kri":        "kri",
	"ku":         "ku",
	"ky":         "ky",
	"la":         "la",
	"lg":         "lg",
	"ln":         "ln",
	"lo":         "lo",
	"loz":        "loz",
	"lt":         "lt",
	"lua":        "lua",
	"lv":         "lv",
	"mfe":        "mfe",
	"mg":         "mg",
	"mi":         "mi",
	"mk":         "mk",
	"ml":         "ml",
	"mn":         "mn",
	"mo":         "mo",
	"mr":         "mr",
	"ms":         "ms",
	"mt":         "mt",
	"ne":         "ne",
	"nl":         "nl",
	"nn":         "nn",
	"no":         "no",
	"nso":        "nso",
	"ny":         "ny",
	"nyn":        "nyn",
	"oc":         "oc",
	"om":         "om",
	"or":         "or",
	"pa":         "pa",
	"pcm":        "pcm",
	"ps":         "ps",
	"pt-BR":      "pt-BR",
	"pt-PT":      "pt-PT",
	"qu":         "qu",
	"rm":         "rm",
	"rn":         "rn",
	"ro":         "ro",
	"ru":         "ru",
	"rw":         "rw",
	"sd":         "sd",
	"sh":         "sh",
	"si":         "si",
	"sk":         "sk",
	"sl":         "sl",
	"sn":         "sn",
	"so":         "so",
	"sq":         "sq",
	"sr":         "sr",
	"sr-ME":      "sr-ME",
	"st":         "st",
	"su":         "su",
	"sv":         "sv",
	"sw":         "sw",
	"ta":         "ta",
	"te":         "te",
	"tg":         "tg",
	"th":         "th",
	"ti":         "ti",
	"tk":         "tk",
	"tl":         "tl",
	"tn":         "tn",
	"to":         "to",
	"tr":         "tr",
	"tt":         "tt",
	"tum":        "tum",
	"tw":         "tw",
	"ug":         "ug",
	"uk":         "uk",
	"ur":         "ur",
	"uz":         "uz",
	"vi":         "vi",
	"wo":         "wo",
	"xh":         "xh",
	"xx-bork":    "xx-bork",
	"xx-elmer":   "xx-elmer",
	"xx-hacker":  "xx-hacker",
	"xx-klingon": "xx-klingon",
	"xx-pirate":  "xx-pirate",
	"yi":         "yi",
	"yo":         "yo",
	"zh-CN":      "zh-CN",
	"zh-TW":      "zh-TW",
	"zu":         "zu",
}
