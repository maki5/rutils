package strings

type inflection struct {
	rule        string
	replacement string
}

type inflections struct{}

func (i *inflections) plural() []inflection {
	return []inflection{
		inflection{rule: `$`, replacement: "s"},
		inflection{rule: `(?i)s$`, replacement: "s"},
		inflection{rule: `(?i)^(ax|test)is$`, replacement: "\\1es"},
		inflection{rule: `(?i)(octop|vir)us$`, replacement: "\\1i"},
		inflection{rule: `(?i)(octop|vir)i$`, replacement: "\\1i"},
		inflection{rule: `(?i)(alias|status)$`, replacement: "\\1es"},
		inflection{rule: `(?i)(bu)s$`, replacement: "\\1ses"},
		inflection{rule: `(?i)(buffal|tomat)o$`, replacement: "\\1oes"},
		inflection{rule: `(?i)([ti])um$`, replacement: "\\1a"},
		inflection{rule: `(?i)([ti])a$`, replacement: "\\1a"},
		inflection{rule: `(?i)sis$`, replacement: "ses"},
		inflection{rule: `(?i)(?:([^f])fe|([lr])f)$`, replacement: "\\1\\2ves"},
		inflection{rule: `(?i)(hive)$`, replacement: "\\1s"},
		inflection{rule: `(?i)([^aeiouy]|qu)y$`, replacement: "\\1ies"},
		inflection{rule: `(?i)(x|ch|ss|sh)$`, replacement: "\\1es"},
		inflection{rule: `(?i)(matr|vert|ind)(?:ix|ex)$`, replacement: "\\1ices"},
		inflection{rule: `(?i)^(m|l)ouse$`, replacement: "\\1ice"},
		inflection{rule: `(?i)^(m|l)ice$`, replacement: "\\1ice"},
		inflection{rule: `(?i)^(ox)$`, replacement: "\\1en"},
		inflection{rule: `(?i)^(oxen)$`, replacement: "\\1"},
		inflection{rule: `(?i)(quiz)$`, replacement: "\\1zes"},
	}
}

func (i *inflections) singular() []inflection {
	return []inflection{}
}

func (i *inflections) irregular() []inflection {
	return []inflection{}
}

func (i *inflections) uncountable() []string {
	return []string{"equipment", "information", "rice", "money", "species", "series", "fish", "sheep", "jeans", "police"}
}

// inflect.singular(/s$/i, "")
// inflect.singular(/(ss)$/i, '\1')
// inflect.singular(/(n)ews$/i, '\1ews')
// inflect.singular(/([ti])a$/i, '\1um')
// inflect.singular(/((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$/i, '\1sis')
// inflect.singular(/(^analy)(sis|ses)$/i, '\1sis')
// inflect.singular(/([^f])ves$/i, '\1fe')
// inflect.singular(/(hive)s$/i, '\1')
// inflect.singular(/(tive)s$/i, '\1')
// inflect.singular(/([lr])ves$/i, '\1f')
// inflect.singular(/([^aeiouy]|qu)ies$/i, '\1y')
// inflect.singular(/(s)eries$/i, '\1eries')
// inflect.singular(/(m)ovies$/i, '\1ovie')
// inflect.singular(/(x|ch|ss|sh)es$/i, '\1')
// inflect.singular(/^(m|l)ice$/i, '\1ouse')
// inflect.singular(/(bus)(es)?$/i, '\1')
// inflect.singular(/(o)es$/i, '\1')
// inflect.singular(/(shoe)s$/i, '\1')
// inflect.singular(/(cris|test)(is|es)$/i, '\1is')
// inflect.singular(/^(a)x[ie]s$/i, '\1xis')
// inflect.singular(/(octop|vir)(us|i)$/i, '\1us')
// inflect.singular(/(alias|status)(es)?$/i, '\1')
// inflect.singular(/^(ox)en/i, '\1')
// inflect.singular(/(vert|ind)ices$/i, '\1ex')
// inflect.singular(/(matr)ices$/i, '\1ix')
// inflect.singular(/(quiz)zes$/i, '\1')
// inflect.singular(/(database)s$/i, '\1')

// inflect.irregular("person", "people")
// inflect.irregular("man", "men")
// inflect.irregular("child", "children")
// inflect.irregular("sex", "sexes")
// inflect.irregular("move", "moves")
// inflect.irregular("zombie", "zombies")

// inflect.uncountable(%w(equipment information rice money species series fish sheep jeans police))
