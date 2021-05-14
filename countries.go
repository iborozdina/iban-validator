package main

type countryFormat struct {
	Name       string
	BBANLength int
	BBANFormat string
}

var Countries = map[string]countryFormat{
	"AL": {Name: "Albania", BBANLength: 24, BBANFormat: `^\d{8}[A-Z0-9]{16}$`},         // AL 47 21211009 0000000235698741
	"AD": {Name: "Andorra", BBANLength: 20, BBANFormat: `^\d{8}[A-Z0-9]{12}$`},         // AD 12 0001 2030 200359100100
	"AT": {Name: "Austria", BBANLength: 16, BBANFormat: `^\d{16}$`},                    // AT 61 19043 00234573201
	"AZ": {Name: "Azerbaijan", BBANLength: 24, BBANFormat: `^[A-Z]{4}[A-Z0-9]{20}$`},   // AZ 21 NABZ 00000000137010001944
	"BH": {Name: "Bahrain", BBANLength: 18, BBANFormat: `^[A-Z]{4}[A-Z0-9]{14}$`},      // BH 67 BMAG 00001299123456
	"BY": {Name: "Belarus", BBANLength: 24, BBANFormat: `^[A-Z]{4}[A-Z0-9]{20}$`},      // BY 13 NBRB 3600900000002Z00AB00
	"BR": {Name: "Brazil", BBANLength: 25, BBANFormat: `^\d{23}[A-Z]{1}[A-Z0-9]{1}$`},  // BR 97 00360305 00001 0009795493 P 1
	"BG": {Name: "Bulgaria", BBANLength: 18, BBANFormat: `^[A-Z]{4}\d{6}[A-Z0-9]{8}$`}, // BG 80 BNBG 9661 10 20345678
	"CR": {Name: "Costa Rica", BBANLength: 18, BBANFormat: `^\d{18}$`},                 // CR 05 0 152 02001026284066
	"CY": {Name: "Cyprus", BBANLength: 24, BBANFormat: `^\d{8}[A-Z0-9]{16}$`},          // CY 17 002 00128 0000001200527600
	"DK": {Name: "Denmark", BBANLength: 14, BBANFormat: `^\d{14}$`},                    // DK 50 0040 0440116243
	"DO": {Name: "Dominican Republic", BBANLength: 24, BBANFormat: `^[A-Z]{4}\d{20}$`}, // DO 28 BAGR 00000001212453611324
	"EG": {Name: "Egypt", BBANLength: 25, BBANFormat: `^\d{25}$`},                      // EG 38 0019 0005 00000000263180002
	"EE": {Name: "Estonia", BBANLength: 16, BBANFormat: `^\d{16}$`},                    // EE 38 22 00 22102014568 5
	"FI": {Name: "Finland", BBANLength: 14, BBANFormat: `^\d{14}$`},                    // FI 21 123456 0000078 5
	"GE": {Name: "Georgia", BBANLength: 18, BBANFormat: `^[A-Z]{2}\d{16}$`},            // GE 29 NB 0000000101904917
	"DE": {Name: "Germany", BBANLength: 18, BBANFormat: `^\d{18}$`},                    // DE 89 37040044 0532013000
	"GR": {Name: "Greece", BBANLength: 23, BBANFormat: `^\d{7}[A-Z0-9]{16}$`},          // GR 16 011 0125 0000000012300695
	"GL": {Name: "Greenland", BBANLength: 14, BBANFormat: `^\d{14}$`},                  // GL 20 0040 0440116243
	"GT": {Name: "Guatemala", BBANLength: 24, BBANFormat: `^[A-Z]{4}[A-Z0-9]{20}$`},    // GT 82 TRAJ 01020000001210029690
	"HU": {Name: "Hungary", BBANLength: 24, BBANFormat: `^\d{24}$`},                    // HU 42 117 7301 6111110180000000 0
	"IS": {Name: "Iceland", BBANLength: 22, BBANFormat: `^\d{22}$`},                    // IS 14 0159 26 007654 5510730339
	"IE": {Name: "Ireland", BBANLength: 18, BBANFormat: `^[A-Z]{4}\d{14}$`},            // IE 29 AIBK 931152 12345678
	"IL": {Name: "Israel", BBANLength: 19, BBANFormat: `^\d{19}$`},                     // IL 62 010 800 0000099999999
	"IT": {Name: "Italy", BBANLength: 23, BBANFormat: `^[A-Z]{1}\d{10}[A-Z0-9]{12}$`},  // IT 60 X 05428 11101 000000123456
	"KZ": {Name: "Kazakhstan", BBANLength: 16, BBANFormat: `^\d{3}[A-Z0-9]{13}$`},      // KZ 86 125 KZT5004100100
	"LV": {Name: "Latvia", BBANLength: 17, BBANFormat: `^[A-Z]{4}\d{13}$`},             // LV 80 BANK 0000435195001
	"LI": {Name: "Liechtenstein", BBANLength: 17, BBANFormat: `^\d{5}[A-Z0-9]{12}$`},   // LI 21 08810 0002324013AA
	"LT": {Name: "Lithuania", BBANLength: 16, BBANFormat: `^\d{16}$`},                  // LT 12 10000 11101001000
	"LU": {Name: "Luxembourg", BBANLength: 16, BBANFormat: `^\d{3}[A-Z0-9]{13}$`},      // LU 28 001 9400644750000
	"MT": {Name: "Malta", BBANLength: 27, BBANFormat: `^[A-Z]{4}\d{5}[A-Z0-9]{18}$`},   // MT 84 MALT 01100 0012345MTLCAST001S
	"MD": {Name: "Moldova", BBANLength: 20, BBANFormat: `^[A-Z0-9]{20}$`},              // MD 24 AG 000225100013104168
	"MC": {Name: "Monaco", BBANLength: 23, BBANFormat: `^\d{10}[A-Z0-9]{11}\d{2}$`},    // MC 58 11222 00001 01234567890 30
	"ME": {Name: "Montenegro", BBANLength: 18, BBANFormat: `^\d{18}$`},                 // ME 25 505 0000123456789 51
	"NL": {Name: "Netherlands", BBANLength: 14, BBANFormat: `^[A-Z]{4}\d{10}$`},        // NL 91 ABNA 0417164300
	"NO": {Name: "Norway", BBANLength: 11, BBANFormat: `^\d{11}$`},                     // NO 93 8601 111794 7
	"PL": {Name: "Poland", BBANLength: 24, BBANFormat: `^\d{24}$`},                     // PL 61 109 0101 4 0000071219812874
	"PT": {Name: "Portugal", BBANLength: 21, BBANFormat: `^\d{21}$`},                   // PT 50 0001 0123 12345678901 92
	"QA": {Name: "Qatar", BBANLength: 25, BBANFormat: `^[A-Z]{4}[A-Z0-9]{21}$`},        // QA 58 DOHB 00001234567890ABCDEFG
	"RO": {Name: "Romania", BBANLength: 20, BBANFormat: `^[A-Z]{4}[A-Z0-9]{16}$`},      // RO 49 AAAA 1B31007593840000
	"SA": {Name: "Saudi Arabia", BBANLength: 20, BBANFormat: `^\d{2}[A-Z0-9]{18}$`},    // SA 03 80 000000608010167519
	"RS": {Name: "Serbia", BBANLength: 18, BBANFormat: `^\d{18}$`},                     // RS 35 260 0056010016113 79
	"SK": {Name: "Slovakia", BBANLength: 20, BBANFormat: `^\d{20}$`},                   // SK 31 1200 000019 8742637541
	"SI": {Name: "Slovenia", BBANLength: 15, BBANFormat: `^\d{15}$`},                   // SI 56 19 100 00001234 38
	"ES": {Name: "Spain", BBANLength: 20, BBANFormat: `^\d{20}$`},                      // ES 91 2100 0418 45 0200051332
	"SE": {Name: "Sweden", BBANLength: 20, BBANFormat: `^\d{20}$`},                     // SE 45 500 00000058398257466
	"CH": {Name: "Switzerland", BBANLength: 17, BBANFormat: `^\d{5}[A-Z0-9]{12}$`},     // CH 93 00762 011623852957
	"TR": {Name: "Turkey", BBANLength: 22, BBANFormat: `^\d{5}[A-Z0-9]{17}$`},          // TR 33 00061 0 0519786457841326
	"UA": {Name: "Ukraine", BBANLength: 25, BBANFormat: `^\d{6}[A-Z0-9]{19}$`},         // UA 21 399622 0000026007233566001
	"AE": {Name: "United Arab Emirates", BBANLength: 19, BBANFormat: `^\d{19}$`},       // AE 07 033 1234567890123456
	"GB": {Name: "United Kingdom", BBANLength: 18, BBANFormat: `^[A-Z]{4}\d{14}$`},     // GB 29 NWBK 601613 31926819
	"VA": {Name: "Vatican", BBANLength: 18, BBANFormat: `^\d{18}$`},                    // VA 59 001 123000012345678
}
