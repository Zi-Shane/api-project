package database

type Languages struct {
	Language_id int
	Language    string
}

type Country_languages struct {
	Country_id  int
	Language_id int
	Official    int
}

type Countries struct {
	Country_id    int
	Name          string
	Area          float32
	National_day  string
	Country_code2 string
	Country_code3 string
	Region_id     int
}

type CountryUesdLanguage struct {
	CountryName string
	Language    string
}
