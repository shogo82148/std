// "go generate"をgolang.org/x/textで実行して生成されたコードです。編集しないでください。

package unicode

// Version はテーブルが派生されるUnicodeエディションです。
const Version = "15.0.0"

// CategoriesはUnicodeのカテゴリーテーブルの集合です。
var Categories = map[string]*RangeTable{
	"C":  C,
	"Cc": Cc,
	"Cf": Cf,
	"Cn": Cn,
	"Co": Co,
	"Cs": Cs,
	"L":  L,
	"LC": LC,
	"Ll": Ll,
	"Lm": Lm,
	"Lo": Lo,
	"Lt": Lt,
	"Lu": Lu,
	"M":  M,
	"Mc": Mc,
	"Me": Me,
	"Mn": Mn,
	"N":  N,
	"Nd": Nd,
	"Nl": Nl,
	"No": No,
	"P":  P,
	"Pc": Pc,
	"Pd": Pd,
	"Pe": Pe,
	"Pf": Pf,
	"Pi": Pi,
	"Po": Po,
	"Ps": Ps,
	"S":  S,
	"Sc": Sc,
	"Sk": Sk,
	"Sm": Sm,
	"So": So,
	"Z":  Z,
	"Zl": Zl,
	"Zp": Zp,
	"Zs": Zs,
}

// CategoryAliasesはカテゴリのエイリアスを標準的なカテゴリ名にマッピングします。
var CategoryAliases = map[string]string{
	"Cased_Letter":          "LC",
	"Close_Punctuation":     "Pe",
	"Combining_Mark":        "M",
	"Connector_Punctuation": "Pc",
	"Control":               "Cc",
	"Currency_Symbol":       "Sc",
	"Dash_Punctuation":      "Pd",
	"Decimal_Number":        "Nd",
	"Enclosing_Mark":        "Me",
	"Final_Punctuation":     "Pf",
	"Format":                "Cf",
	"Initial_Punctuation":   "Pi",
	"Letter":                "L",
	"Letter_Number":         "Nl",
	"Line_Separator":        "Zl",
	"Lowercase_Letter":      "Ll",
	"Mark":                  "M",
	"Math_Symbol":           "Sm",
	"Modifier_Letter":       "Lm",
	"Modifier_Symbol":       "Sk",
	"Nonspacing_Mark":       "Mn",
	"Number":                "N",
	"Open_Punctuation":      "Ps",
	"Other":                 "C",
	"Other_Letter":          "Lo",
	"Other_Number":          "No",
	"Other_Punctuation":     "Po",
	"Other_Symbol":          "So",
	"Paragraph_Separator":   "Zp",
	"Private_Use":           "Co",
	"Punctuation":           "P",
	"Separator":             "Z",
	"Space_Separator":       "Zs",
	"Spacing_Mark":          "Mc",
	"Surrogate":             "Cs",
	"Symbol":                "S",
	"Titlecase_Letter":      "Lt",
	"Unassigned":            "Cn",
	"Uppercase_Letter":      "Lu",
	"cntrl":                 "Cc",
	"digit":                 "Nd",
	"punct":                 "P",
}

// These variables have type *RangeTable.
var (
	Cc     = _Cc
	Cf     = _Cf
	Cn     = _Cn
	Co     = _Co
	Cs     = _Cs
	Digit  = _Nd
	Nd     = _Nd
	LC     = _LC
	Letter = _L
	L      = _L
	Lm     = _Lm
	Lo     = _Lo
	Lower  = _Ll
	Ll     = _Ll
	Mark   = _M
	M      = _M
	Mc     = _Mc
	Me     = _Me
	Mn     = _Mn
	Nl     = _Nl
	No     = _No
	Number = _N
	N      = _N
	Other  = _C
	C      = _C
	Pc     = _Pc
	Pd     = _Pd
	Pe     = _Pe
	Pf     = _Pf
	Pi     = _Pi
	Po     = _Po
	Ps     = _Ps
	Punct  = _P
	P      = _P
	Sc     = _Sc
	Sk     = _Sk
	Sm     = _Sm
	So     = _So
	Space  = _Z
	Z      = _Z
	Symbol = _S
	S      = _S
	Title  = _Lt
	Lt     = _Lt
	Upper  = _Lu
	Lu     = _Lu
	Zl     = _Zl
	Zp     = _Zp
	Zs     = _Zs
)

// ScriptsはUnicodeスクリプトテーブルの集合です。
var Scripts = map[string]*RangeTable{
	"Adlam":                  Adlam,
	"Ahom":                   Ahom,
	"Anatolian_Hieroglyphs":  Anatolian_Hieroglyphs,
	"Arabic":                 Arabic,
	"Armenian":               Armenian,
	"Avestan":                Avestan,
	"Balinese":               Balinese,
	"Bamum":                  Bamum,
	"Bassa_Vah":              Bassa_Vah,
	"Batak":                  Batak,
	"Bengali":                Bengali,
	"Bhaiksuki":              Bhaiksuki,
	"Bopomofo":               Bopomofo,
	"Brahmi":                 Brahmi,
	"Braille":                Braille,
	"Buginese":               Buginese,
	"Buhid":                  Buhid,
	"Canadian_Aboriginal":    Canadian_Aboriginal,
	"Carian":                 Carian,
	"Caucasian_Albanian":     Caucasian_Albanian,
	"Chakma":                 Chakma,
	"Cham":                   Cham,
	"Cherokee":               Cherokee,
	"Chorasmian":             Chorasmian,
	"Common":                 Common,
	"Coptic":                 Coptic,
	"Cuneiform":              Cuneiform,
	"Cypriot":                Cypriot,
	"Cypro_Minoan":           Cypro_Minoan,
	"Cyrillic":               Cyrillic,
	"Deseret":                Deseret,
	"Devanagari":             Devanagari,
	"Dives_Akuru":            Dives_Akuru,
	"Dogra":                  Dogra,
	"Duployan":               Duployan,
	"Egyptian_Hieroglyphs":   Egyptian_Hieroglyphs,
	"Elbasan":                Elbasan,
	"Elymaic":                Elymaic,
	"Ethiopic":               Ethiopic,
	"Georgian":               Georgian,
	"Glagolitic":             Glagolitic,
	"Gothic":                 Gothic,
	"Grantha":                Grantha,
	"Greek":                  Greek,
	"Gujarati":               Gujarati,
	"Gunjala_Gondi":          Gunjala_Gondi,
	"Gurmukhi":               Gurmukhi,
	"Han":                    Han,
	"Hangul":                 Hangul,
	"Hanifi_Rohingya":        Hanifi_Rohingya,
	"Hanunoo":                Hanunoo,
	"Hatran":                 Hatran,
	"Hebrew":                 Hebrew,
	"Hiragana":               Hiragana,
	"Imperial_Aramaic":       Imperial_Aramaic,
	"Inherited":              Inherited,
	"Inscriptional_Pahlavi":  Inscriptional_Pahlavi,
	"Inscriptional_Parthian": Inscriptional_Parthian,
	"Javanese":               Javanese,
	"Kaithi":                 Kaithi,
	"Kannada":                Kannada,
	"Katakana":               Katakana,
	"Kawi":                   Kawi,
	"Kayah_Li":               Kayah_Li,
	"Kharoshthi":             Kharoshthi,
	"Khitan_Small_Script":    Khitan_Small_Script,
	"Khmer":                  Khmer,
	"Khojki":                 Khojki,
	"Khudawadi":              Khudawadi,
	"Lao":                    Lao,
	"Latin":                  Latin,
	"Lepcha":                 Lepcha,
	"Limbu":                  Limbu,
	"Linear_A":               Linear_A,
	"Linear_B":               Linear_B,
	"Lisu":                   Lisu,
	"Lycian":                 Lycian,
	"Lydian":                 Lydian,
	"Mahajani":               Mahajani,
	"Makasar":                Makasar,
	"Malayalam":              Malayalam,
	"Mandaic":                Mandaic,
	"Manichaean":             Manichaean,
	"Marchen":                Marchen,
	"Masaram_Gondi":          Masaram_Gondi,
	"Medefaidrin":            Medefaidrin,
	"Meetei_Mayek":           Meetei_Mayek,
	"Mende_Kikakui":          Mende_Kikakui,
	"Meroitic_Cursive":       Meroitic_Cursive,
	"Meroitic_Hieroglyphs":   Meroitic_Hieroglyphs,
	"Miao":                   Miao,
	"Modi":                   Modi,
	"Mongolian":              Mongolian,
	"Mro":                    Mro,
	"Multani":                Multani,
	"Myanmar":                Myanmar,
	"Nabataean":              Nabataean,
	"Nag_Mundari":            Nag_Mundari,
	"Nandinagari":            Nandinagari,
	"New_Tai_Lue":            New_Tai_Lue,
	"Newa":                   Newa,
	"Nko":                    Nko,
	"Nushu":                  Nushu,
	"Nyiakeng_Puachue_Hmong": Nyiakeng_Puachue_Hmong,
	"Ogham":                  Ogham,
	"Ol_Chiki":               Ol_Chiki,
	"Old_Hungarian":          Old_Hungarian,
	"Old_Italic":             Old_Italic,
	"Old_North_Arabian":      Old_North_Arabian,
	"Old_Permic":             Old_Permic,
	"Old_Persian":            Old_Persian,
	"Old_Sogdian":            Old_Sogdian,
	"Old_South_Arabian":      Old_South_Arabian,
	"Old_Turkic":             Old_Turkic,
	"Old_Uyghur":             Old_Uyghur,
	"Oriya":                  Oriya,
	"Osage":                  Osage,
	"Osmanya":                Osmanya,
	"Pahawh_Hmong":           Pahawh_Hmong,
	"Palmyrene":              Palmyrene,
	"Pau_Cin_Hau":            Pau_Cin_Hau,
	"Phags_Pa":               Phags_Pa,
	"Phoenician":             Phoenician,
	"Psalter_Pahlavi":        Psalter_Pahlavi,
	"Rejang":                 Rejang,
	"Runic":                  Runic,
	"Samaritan":              Samaritan,
	"Saurashtra":             Saurashtra,
	"Sharada":                Sharada,
	"Shavian":                Shavian,
	"Siddham":                Siddham,
	"SignWriting":            SignWriting,
	"Sinhala":                Sinhala,
	"Sogdian":                Sogdian,
	"Sora_Sompeng":           Sora_Sompeng,
	"Soyombo":                Soyombo,
	"Sundanese":              Sundanese,
	"Syloti_Nagri":           Syloti_Nagri,
	"Syriac":                 Syriac,
	"Tagalog":                Tagalog,
	"Tagbanwa":               Tagbanwa,
	"Tai_Le":                 Tai_Le,
	"Tai_Tham":               Tai_Tham,
	"Tai_Viet":               Tai_Viet,
	"Takri":                  Takri,
	"Tamil":                  Tamil,
	"Tangsa":                 Tangsa,
	"Tangut":                 Tangut,
	"Telugu":                 Telugu,
	"Thaana":                 Thaana,
	"Thai":                   Thai,
	"Tibetan":                Tibetan,
	"Tifinagh":               Tifinagh,
	"Tirhuta":                Tirhuta,
	"Toto":                   Toto,
	"Ugaritic":               Ugaritic,
	"Vai":                    Vai,
	"Vithkuqi":               Vithkuqi,
	"Wancho":                 Wancho,
	"Warang_Citi":            Warang_Citi,
	"Yezidi":                 Yezidi,
	"Yi":                     Yi,
	"Zanabazar_Square":       Zanabazar_Square,
}

// これらの変数は*RangeTable型です。
var (
	Adlam                  = _Adlam
	Ahom                   = _Ahom
	Anatolian_Hieroglyphs  = _Anatolian_Hieroglyphs
	Arabic                 = _Arabic
	Armenian               = _Armenian
	Avestan                = _Avestan
	Balinese               = _Balinese
	Bamum                  = _Bamum
	Bassa_Vah              = _Bassa_Vah
	Batak                  = _Batak
	Bengali                = _Bengali
	Bhaiksuki              = _Bhaiksuki
	Bopomofo               = _Bopomofo
	Brahmi                 = _Brahmi
	Braille                = _Braille
	Buginese               = _Buginese
	Buhid                  = _Buhid
	Canadian_Aboriginal    = _Canadian_Aboriginal
	Carian                 = _Carian
	Caucasian_Albanian     = _Caucasian_Albanian
	Chakma                 = _Chakma
	Cham                   = _Cham
	Cherokee               = _Cherokee
	Chorasmian             = _Chorasmian
	Common                 = _Common
	Coptic                 = _Coptic
	Cuneiform              = _Cuneiform
	Cypriot                = _Cypriot
	Cypro_Minoan           = _Cypro_Minoan
	Cyrillic               = _Cyrillic
	Deseret                = _Deseret
	Devanagari             = _Devanagari
	Dives_Akuru            = _Dives_Akuru
	Dogra                  = _Dogra
	Duployan               = _Duployan
	Egyptian_Hieroglyphs   = _Egyptian_Hieroglyphs
	Elbasan                = _Elbasan
	Elymaic                = _Elymaic
	Ethiopic               = _Ethiopic
	Georgian               = _Georgian
	Glagolitic             = _Glagolitic
	Gothic                 = _Gothic
	Grantha                = _Grantha
	Greek                  = _Greek
	Gujarati               = _Gujarati
	Gunjala_Gondi          = _Gunjala_Gondi
	Gurmukhi               = _Gurmukhi
	Han                    = _Han
	Hangul                 = _Hangul
	Hanifi_Rohingya        = _Hanifi_Rohingya
	Hanunoo                = _Hanunoo
	Hatran                 = _Hatran
	Hebrew                 = _Hebrew
	Hiragana               = _Hiragana
	Imperial_Aramaic       = _Imperial_Aramaic
	Inherited              = _Inherited
	Inscriptional_Pahlavi  = _Inscriptional_Pahlavi
	Inscriptional_Parthian = _Inscriptional_Parthian
	Javanese               = _Javanese
	Kaithi                 = _Kaithi
	Kannada                = _Kannada
	Katakana               = _Katakana
	Kawi                   = _Kawi
	Kayah_Li               = _Kayah_Li
	Kharoshthi             = _Kharoshthi
	Khitan_Small_Script    = _Khitan_Small_Script
	Khmer                  = _Khmer
	Khojki                 = _Khojki
	Khudawadi              = _Khudawadi
	Lao                    = _Lao
	Latin                  = _Latin
	Lepcha                 = _Lepcha
	Limbu                  = _Limbu
	Linear_A               = _Linear_A
	Linear_B               = _Linear_B
	Lisu                   = _Lisu
	Lycian                 = _Lycian
	Lydian                 = _Lydian
	Mahajani               = _Mahajani
	Makasar                = _Makasar
	Malayalam              = _Malayalam
	Mandaic                = _Mandaic
	Manichaean             = _Manichaean
	Marchen                = _Marchen
	Masaram_Gondi          = _Masaram_Gondi
	Medefaidrin            = _Medefaidrin
	Meetei_Mayek           = _Meetei_Mayek
	Mende_Kikakui          = _Mende_Kikakui
	Meroitic_Cursive       = _Meroitic_Cursive
	Meroitic_Hieroglyphs   = _Meroitic_Hieroglyphs
	Miao                   = _Miao
	Modi                   = _Modi
	Mongolian              = _Mongolian
	Mro                    = _Mro
	Multani                = _Multani
	Myanmar                = _Myanmar
	Nabataean              = _Nabataean
	Nag_Mundari            = _Nag_Mundari
	Nandinagari            = _Nandinagari
	New_Tai_Lue            = _New_Tai_Lue
	Newa                   = _Newa
	Nko                    = _Nko
	Nushu                  = _Nushu
	Nyiakeng_Puachue_Hmong = _Nyiakeng_Puachue_Hmong
	Ogham                  = _Ogham
	Ol_Chiki               = _Ol_Chiki
	Old_Hungarian          = _Old_Hungarian
	Old_Italic             = _Old_Italic
	Old_North_Arabian      = _Old_North_Arabian
	Old_Permic             = _Old_Permic
	Old_Persian            = _Old_Persian
	Old_Sogdian            = _Old_Sogdian
	Old_South_Arabian      = _Old_South_Arabian
	Old_Turkic             = _Old_Turkic
	Old_Uyghur             = _Old_Uyghur
	Oriya                  = _Oriya
	Osage                  = _Osage
	Osmanya                = _Osmanya
	Pahawh_Hmong           = _Pahawh_Hmong
	Palmyrene              = _Palmyrene
	Pau_Cin_Hau            = _Pau_Cin_Hau
	Phags_Pa               = _Phags_Pa
	Phoenician             = _Phoenician
	Psalter_Pahlavi        = _Psalter_Pahlavi
	Rejang                 = _Rejang
	Runic                  = _Runic
	Samaritan              = _Samaritan
	Saurashtra             = _Saurashtra
	Sharada                = _Sharada
	Shavian                = _Shavian
	Siddham                = _Siddham
	SignWriting            = _SignWriting
	Sinhala                = _Sinhala
	Sogdian                = _Sogdian
	Sora_Sompeng           = _Sora_Sompeng
	Soyombo                = _Soyombo
	Sundanese              = _Sundanese
	Syloti_Nagri           = _Syloti_Nagri
	Syriac                 = _Syriac
	Tagalog                = _Tagalog
	Tagbanwa               = _Tagbanwa
	Tai_Le                 = _Tai_Le
	Tai_Tham               = _Tai_Tham
	Tai_Viet               = _Tai_Viet
	Takri                  = _Takri
	Tamil                  = _Tamil
	Tangsa                 = _Tangsa
	Tangut                 = _Tangut
	Telugu                 = _Telugu
	Thaana                 = _Thaana
	Thai                   = _Thai
	Tibetan                = _Tibetan
	Tifinagh               = _Tifinagh
	Tirhuta                = _Tirhuta
	Toto                   = _Toto
	Ugaritic               = _Ugaritic
	Vai                    = _Vai
	Vithkuqi               = _Vithkuqi
	Wancho                 = _Wancho
	Warang_Citi            = _Warang_Citi
	Yezidi                 = _Yezidi
	Yi                     = _Yi
	Zanabazar_Square       = _Zanabazar_Square
)

// PropertiesはUnicodeプロパティテーブルの集合です。
var Properties = map[string]*RangeTable{
	"ASCII_Hex_Digit":                    ASCII_Hex_Digit,
	"Bidi_Control":                       Bidi_Control,
	"Dash":                               Dash,
	"Deprecated":                         Deprecated,
	"Diacritic":                          Diacritic,
	"Extender":                           Extender,
	"Hex_Digit":                          Hex_Digit,
	"Hyphen":                             Hyphen,
	"IDS_Binary_Operator":                IDS_Binary_Operator,
	"IDS_Trinary_Operator":               IDS_Trinary_Operator,
	"Ideographic":                        Ideographic,
	"Join_Control":                       Join_Control,
	"Logical_Order_Exception":            Logical_Order_Exception,
	"Noncharacter_Code_Point":            Noncharacter_Code_Point,
	"Other_Alphabetic":                   Other_Alphabetic,
	"Other_Default_Ignorable_Code_Point": Other_Default_Ignorable_Code_Point,
	"Other_Grapheme_Extend":              Other_Grapheme_Extend,
	"Other_ID_Continue":                  Other_ID_Continue,
	"Other_ID_Start":                     Other_ID_Start,
	"Other_Lowercase":                    Other_Lowercase,
	"Other_Math":                         Other_Math,
	"Other_Uppercase":                    Other_Uppercase,
	"Pattern_Syntax":                     Pattern_Syntax,
	"Pattern_White_Space":                Pattern_White_Space,
	"Prepended_Concatenation_Mark":       Prepended_Concatenation_Mark,
	"Quotation_Mark":                     Quotation_Mark,
	"Radical":                            Radical,
	"Regional_Indicator":                 Regional_Indicator,
	"Sentence_Terminal":                  Sentence_Terminal,
	"STerm":                              Sentence_Terminal,
	"Soft_Dotted":                        Soft_Dotted,
	"Terminal_Punctuation":               Terminal_Punctuation,
	"Unified_Ideograph":                  Unified_Ideograph,
	"Variation_Selector":                 Variation_Selector,
	"White_Space":                        White_Space,
}

// これらの変数は*RangeTable型を持っています。
var (
	ASCII_Hex_Digit                    = _ASCII_Hex_Digit
	Bidi_Control                       = _Bidi_Control
	Dash                               = _Dash
	Deprecated                         = _Deprecated
	Diacritic                          = _Diacritic
	Extender                           = _Extender
	Hex_Digit                          = _Hex_Digit
	Hyphen                             = _Hyphen
	IDS_Binary_Operator                = _IDS_Binary_Operator
	IDS_Trinary_Operator               = _IDS_Trinary_Operator
	Ideographic                        = _Ideographic
	Join_Control                       = _Join_Control
	Logical_Order_Exception            = _Logical_Order_Exception
	Noncharacter_Code_Point            = _Noncharacter_Code_Point
	Other_Alphabetic                   = _Other_Alphabetic
	Other_Default_Ignorable_Code_Point = _Other_Default_Ignorable_Code_Point
	Other_Grapheme_Extend              = _Other_Grapheme_Extend
	Other_ID_Continue                  = _Other_ID_Continue
	Other_ID_Start                     = _Other_ID_Start
	Other_Lowercase                    = _Other_Lowercase
	Other_Math                         = _Other_Math
	Other_Uppercase                    = _Other_Uppercase
	Pattern_Syntax                     = _Pattern_Syntax
	Pattern_White_Space                = _Pattern_White_Space
	Prepended_Concatenation_Mark       = _Prepended_Concatenation_Mark
	Quotation_Mark                     = _Quotation_Mark
	Radical                            = _Radical
	Regional_Indicator                 = _Regional_Indicator
	STerm                              = _Sentence_Terminal
	Sentence_Terminal                  = _Sentence_Terminal
	Soft_Dotted                        = _Soft_Dotted
	Terminal_Punctuation               = _Terminal_Punctuation
	Unified_Ideograph                  = _Unified_Ideograph
	Variation_Selector                 = _Variation_Selector
	White_Space                        = _White_Space
)

// CaseRangesは、非自己マッピングを持つすべての文字の大文字小文字変換に関するテーブルです。
var CaseRanges = _CaseRanges

// FoldCategoryはカテゴリ名を、カテゴリ内のコードポイントと単純な大文字小文字変換で等価なカテゴリ外のコードポイントのテーブルにマッピングします。
// カテゴリ名のエントリが存在しない場合、そのようなポイントは存在しません。
var FoldCategory = map[string]*RangeTable{
	"L":  foldL,
	"Ll": foldLl,
	"Lt": foldLt,
	"Lu": foldLu,
	"M":  foldM,
	"Mn": foldMn,
}

// FoldScriptはスクリプト名をスクリプト内のコードポイントに対して単純なケースフォールディングで等価なスクリプト外のコードポイントのテーブルにマッピングします。
// スクリプト名のエントリが存在しない場合、そのようなポイントは存在しません。
var FoldScript = map[string]*RangeTable{
	"Common":    foldCommon,
	"Greek":     foldGreek,
	"Inherited": foldInherited,
}
