package ucl

import (
	"testing"
)

func TestParseJSON(t *testing.T) {
	good := []string{
		`{}`,
		`[]`,
		`{"1":"2"}`,
		`{"a":123}`,
		`{"a":1,"b":"321"}`,
		`{"a":123.45}`,
		`{"a":123.45e67}`,
		`{"a":-123}`,
		`{"a":-123.45}`,
		`{"a":-123.45e-67}`,
		`{"a":["b"]}`,
		`{"a":{"b":{"c":123}}}`,
		`[[[123,"1 ", "1", 123 , {}]]]`,
		`[
  {
    "_id": "55465dc9b9c7cb2b64e51a78",
    "index": 0,
    "guid": "27d221d9-b3b1-4ca4-98db-13be69735527",
    "isActive": false,
    "balance": "$3,923.49",
    "picture": "http://placehold.it/32x32",
    "age": 21,
    "eyeColor": "green",
    "name": "Tammi Byers",
    "gender": "female",
    "company": "ISOLOGIX",
    "email": "tammibyers@isologix.com",
    "phone": "+1 (857) 429-2118",
    "address": "248 Garfield Place, Allentown, Iowa, 9158",
    "about": "Et veniam labore voluptate non quis excepteur consequat cupidatat adipisicing deserunt minim. Occaecat tempor ut enim consequat veniam mollit nulla anim non Lorem labore mollit nulla nulla. Cillum id labore sint Lorem cupidatat ex incididunt nostrud anim cupidatat irure ex ut. Eu dolore sit ipsum ut aliquip. Irure occaecat pariatur deserunt excepteur sit incididunt. Cillum deserunt laborum sint officia sunt sit do duis cillum eu officia aliquip.\r\n",
    "registered": "2014-12-20T17:00:54 -00:00",
    "latitude": 18.416923,
    "longitude": -157.590654,
    "tags": [
      "voluptate",
      "adipisicing",
      "voluptate",
      "labore",
      "velit",
      "amet",
      "deserunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Chandra Holland"
      },
      {
        "id": 1,
        "name": "Lopez Dixon"
      },
      {
        "id": 2,
        "name": "Hunter Blair"
      }
    ],
    "greeting": "Hello, Tammi Byers! You have 5 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "55465dc9249878df7ff44b64",
    "index": 1,
    "guid": "b5fb118f-5028-44d3-bab8-734c97c76690",
    "isActive": true,
    "balance": "$3,319.96",
    "picture": "http://placehold.it/32x32",
    "age": 40,
    "eyeColor": "brown",
    "name": "Anderson Carrillo",
    "gender": "male",
    "company": "BITREX",
    "email": "andersoncarrillo@bitrex.com",
    "phone": "+1 (946) 494-3836",
    "address": "191 Joval Court, Glenville, Kansas, 8973",
    "about": "Pariatur sunt voluptate adipisicing id. Minim magna eiusmod incididunt excepteur laboris enim cupidatat exercitation cupidatat labore nostrud aute culpa. Exercitation officia culpa qui deserunt ex cupidatat id cillum ut nulla adipisicing. Consequat aliqua dolor duis cillum laboris cupidatat deserunt mollit irure incididunt amet est adipisicing. Velit duis culpa tempor commodo proident occaecat aliqua ullamco et. Reprehenderit cupidatat id mollit ad adipisicing velit excepteur cillum aliquip reprehenderit nisi.\r\n",
    "registered": "2014-07-03T12:21:53 -01:00",
    "latitude": 13.523948,
    "longitude": -152.716486,
    "tags": [
      "dolore",
      "officia",
      "ad",
      "pariatur",
      "aliquip",
      "adipisicing",
      "non"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Lakisha Parker"
      },
      {
        "id": 1,
        "name": "Kathryn Hogan"
      },
      {
        "id": 2,
        "name": "Jannie Hewitt"
      }
    ],
    "greeting": "Hello, Anderson Carrillo! You have 9 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "55465dc92664f0bb4494d40f",
    "index": 2,
    "guid": "2fbaf240-30c1-4572-a8dd-be2af9bc4f09",
    "isActive": true,
    "balance": "$3,902.37",
    "picture": "http://placehold.it/32x32",
    "age": 20,
    "eyeColor": "blue",
    "name": "Finley Holt",
    "gender": "male",
    "company": "BIOSPAN",
    "email": "finleyholt@biospan.com",
    "phone": "+1 (849) 528-2802",
    "address": "823 Dobbin Street, Tuskahoma, Connecticut, 7581",
    "about": "Culpa amet fugiat esse eu tempor laboris excepteur cillum veniam elit duis. Veniam commodo id sit ea. Veniam consequat et occaecat ipsum aute aliqua excepteur officia sit eu velit nisi exercitation. Commodo quis adipisicing in id nostrud culpa exercitation voluptate aliqua. Culpa labore do reprehenderit ex nisi pariatur est exercitation commodo laborum laboris magna.\r\n",
    "registered": "2014-01-18T06:05:06 -00:00",
    "latitude": 69.974056,
    "longitude": 122.034273,
    "tags": [
      "eu",
      "aliquip",
      "Lorem",
      "ullamco",
      "qui",
      "eiusmod",
      "incididunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Bette Martinez"
      },
      {
        "id": 1,
        "name": "Russo Padilla"
      },
      {
        "id": 2,
        "name": "Lowery Gay"
      }
    ],
    "greeting": "Hello, Finley Holt! You have 4 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "55465dc9c1d4c25e43af2b9d",
    "index": 3,
    "guid": "3252086a-16fc-4c2a-8142-dbf79c0450a9",
    "isActive": true,
    "balance": "$2,637.43",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "green",
    "name": "Watkins Langley",
    "gender": "male",
    "company": "COMTRACT",
    "email": "watkinslangley@comtract.com",
    "phone": "+1 (851) 588-3283",
    "address": "992 Vanderbilt Street, Hebron, Minnesota, 5100",
    "about": "Aute exercitation laborum esse dolor reprehenderit dolor do mollit. Duis dolore excepteur in ea. Proident et voluptate magna et nisi nulla eu anim nulla laborum duis. Nostrud ea tempor sunt aliquip voluptate ea eu consectetur laborum et esse.\r\n",
    "registered": "2015-04-22T00:20:55 -01:00",
    "latitude": -63.728092,
    "longitude": -53.169103,
    "tags": [
      "nulla",
      "anim",
      "consectetur",
      "laborum",
      "ipsum",
      "velit",
      "culpa"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Madelyn Sampson"
      },
      {
        "id": 1,
        "name": "Tabitha Gonzalez"
      },
      {
        "id": 2,
        "name": "Avis Graves"
      }
    ],
    "greeting": "Hello, Watkins Langley! You have 3 unread messages.",
    "favoriteFruit": "apple"
  },
  {
    "_id": "55465dc9bf9eadc76bdbccad",
    "index": 4,
    "guid": "1628e647-69f2-4566-adb4-e12b50f3adcc",
    "isActive": true,
    "balance": "$3,185.68",
    "picture": "http://placehold.it/32x32",
    "age": 25,
    "eyeColor": "blue",
    "name": "Judith Cardenas",
    "gender": "female",
    "company": "COMTRAK",
    "email": "judithcardenas@comtrak.com",
    "phone": "+1 (966) 421-3187",
    "address": "226 Moore Place, Lewis, Nebraska, 6307",
    "about": "Laborum laborum deserunt est deserunt proident proident labore dolore. Fugiat eu incididunt ea elit. Eu ad exercitation Lorem esse eiusmod et officia ipsum ullamco mollit ea. Qui qui commodo aute veniam occaecat consequat. Nulla voluptate duis consectetur et labore voluptate adipisicing voluptate exercitation tempor pariatur officia consectetur commodo. Officia occaecat quis dolore nisi elit aute pariatur do dolore qui. Amet reprehenderit nulla id est consectetur qui do adipisicing anim fugiat.\r\n",
    "registered": "2014-07-08T14:05:42 -01:00",
    "latitude": 2.945309,
    "longitude": -43.139882,
    "tags": [
      "id",
      "sunt",
      "officia",
      "minim",
      "ea",
      "laboris",
      "sunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Hendricks Wheeler"
      },
      {
        "id": 1,
        "name": "Freda Reed"
      },
      {
        "id": 2,
        "name": "Margie Guy"
      }
    ],
    "greeting": "Hello, Judith Cardenas! You have 5 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "55465dc9da6138b6ed2cbce1",
    "index": 5,
    "guid": "9aaf8c4b-5a9e-4006-a1cc-dfa44c688594",
    "isActive": true,
    "balance": "$2,851.69",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "brown",
    "name": "Gomez Mejia",
    "gender": "male",
    "company": "EMTRAK",
    "email": "gomezmejia@emtrak.com",
    "phone": "+1 (815) 598-3881",
    "address": "940 John Street, Rodanthe, Georgia, 3888",
    "about": "Occaecat excepteur consequat proident pariatur duis fugiat ut. Ad labore officia adipisicing deserunt laborum magna irure id est qui proident. Non adipisicing nulla est ex irure laborum qui exercitation officia. Ipsum laborum id duis aliqua laborum minim in consequat qui nostrud aliqua.\r\n",
    "registered": "2015-01-21T19:45:12 -00:00",
    "latitude": -66.368742,
    "longitude": -175.852809,
    "tags": [
      "commodo",
      "occaecat",
      "dolor",
      "eu",
      "nisi",
      "dolor",
      "dolor"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Monroe Alexander"
      },
      {
        "id": 1,
        "name": "Theresa Farley"
      },
      {
        "id": 2,
        "name": "Paul Armstrong"
      }
    ],
    "greeting": "Hello, Gomez Mejia! You have 6 unread messages.",
    "favoriteFruit": "banana"
  }
]`,
	}
	for _, s := range good {
		if err := parse_json([]rune(s)); err != nil {
			t.Errorf("Failed to parse '%s': %s", s, len(s), err)
		}
	}
	bad := []string{
		`""`,
		`["""]`,
		`{}{}`,
		`{{}}`,
		`{[]:{}}`,
		`{{{{{`,
		`}{`,
		`][`,
		`"a"a"`,
	}
	for _, s := range bad {
		err := parse_json([]rune(s))
		if err == nil {
			t.Errorf("Parse succeeded on invalid JSON document '%s'", s)
		} else {
			t.Logf("%s: %s", s, err)
		}
	}
}
