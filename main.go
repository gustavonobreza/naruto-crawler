package main

import (
	"io/fs"
	"io/ioutil"
	"strings"
)

func main() {
	// characters := []string{"Minato_Namikaze", "Naruto_Uzumaki", "Nagato", "Sasuke_Uchiha", "Sakura_Haruno"}
	characters := []string{
		"Naruto Uzumaki",
		"Minato Namikaze",
		"Sasuke Uchiha",
		"Sakura Haruno",
		"Kakashi Hatake",
		"Hashirama Senju",
		"Tobirama Senju",
		"Madara Uchiha",
		"Maito Gai",
		"Hinata Hyuga",
		"Neji Hyuga",
		"Jiraiya",
		"Tsunade",
		"Orochimaru",
		"Nagato",
		"Sarada Uchiha",
		"Mitsuki",
		"Boruto Uzumaki",
		"Jigen",
		"Kaguya Ōtsutsuki",
		"Hagoromo Ōtsutsuki",
		"Hamura Ōtsutsuki",
		"Momoshiki Ōtsutsuki",
		"Itachi Uchiha",
		"Shikamaru Nara",
		"Ino Yamanaka",
		"Chōji Akimichi",
		"Rock Lee",
		"Kiba Inuzuka",
		"Asuma Sarutobi",
		"Shisui Uchiha",
		"Kurenai Yūhi",
		"Shino Aburame",
		"Obito Uchiha",
		"Deidara",
		"Sasori",
		"Gaara",
		"Sai",
		"Kushina Uzumaki",
		"Temari",
		"Konohamaru Sarutobi",
		"Kabuto Yakushi",
		"Tenten",
		"Killer B",
		"Hidan",
		"Danzō Shimura",
		"Chōji_Akimichi",
		"Yamato",
		"Karin",
		"Konan",
		"Hiruzen Sarutobi",
		"Anko Mitarashi",
		"Zabuza Momochi",
		"Kisame Hoshigaki",
		"Tayuya",
		"Iruka Umino",
		"Suigetsu Hōzuki",
		"Shizune",
		"Kakuzu",
		"Karui Akimichi",
		"Haku",
		"Jirōbō",
		"Kimimaro",
		"Kidōmaru",
		"Sakon",
		"Chōjūrō",
		"Ao",
		"Akatsuchi",
		"Baki",
		"Aoba Yamashiro",
		"Chōza Akimichi",
		"Dan_Katō",
		"Ebizō",
		"Genma_Shiranui",
		"Fugaku Uchiha",
		"Fū",
		"Fū_Yamanaka",
		"Hiashi_Hyūga",
		"Hizashi_Hyūga",
		"Hayate_Gekkō",
		"Ibiki Morino",
		"Jūgo",
		"Izumo Kamizuki",
		"Inoichi Yamanaka",
		"Kin_Tsuchi",
		"Kurotsuchi",
		"Kotetsu Hagane",
		"Dosu Kinuta",
		"Mikoto Uchiha",
		"Mei_Terumī",
		"Rin Nohara",
		"Sakumo Hatake",
		"Sora",
		"Ōnoki",
		"Mū",
		"Ishikawa",
		"Gamakichi",
		"Gamabunta"}

	for _, name := range characters {
		name = strings.ReplaceAll(name, " ", "_")
		urlBase := "https://naruto.fandom.com/wiki/"
		url := urlBase + name

		res, _ := Fetch(url)
		var str string

		for _, v := range res {
			str += v
		}
		ioutil.WriteFile(name+".txt", []byte(str), fs.ModeAppend)
	}

}
