package main

type vector struct {
	start [2]int
	end   [2]int
	stepx int
	stepy int
	len   int
}

// var inputs = []vector{
// 	{
// 		start: [2]int{0, 9},
// 		end:   [2]int{5, 9},
// 	},
// 	{
// 		start: [2]int{8, 0},
// 		end:   [2]int{0, 8},
// 	},
// 	{
// 		start: [2]int{9, 4},
// 		end:   [2]int{3, 4},
// 	},
// 	{
// 		start: [2]int{2, 2},
// 		end:   [2]int{2, 1},
// 	},
// 	{
// 		start: [2]int{7, 0},
// 		end:   [2]int{7, 4},
// 	},
// 	{
// 		start: [2]int{6, 4},
// 		end:   [2]int{2, 0},
// 	},
// 	{
// 		start: [2]int{0, 9},
// 		end:   [2]int{2, 9},
// 	},
// 	{
// 		start: [2]int{3, 4},
// 		end:   [2]int{1, 4},
// 	},
// 	{
// 		start: [2]int{0, 0},
// 		end:   [2]int{8, 8},
// 	},
// 	{
// 		start: [2]int{5, 5},
// 		end:   [2]int{8, 2},
// 	},
// }

var inputs = []vector{
	{
		start: [2]int{194, 556},
		end:   [2]int{739, 556},
	},
	{
		start: [2]int{818, 920},
		end:   [2]int{818, 524},
	},
	{
		start: [2]int{340, 734},
		end:   [2]int{774, 300},
	},
	{
		start: [2]int{223, 511},
		end:   [2]int{146, 434},
	},
	{
		start: [2]int{841, 47},
		end:   [2]int{122, 766},
	},
	{
		start: [2]int{323, 858},
		end:   [2]int{859, 322},
	},
	{
		start: [2]int{277, 205},
		end:   [2]int{85, 205},
	},
	{
		start: [2]int{782, 901},
		end:   [2]int{782, 186},
	},
	{
		start: [2]int{969, 96},
		end:   [2]int{969, 648},
	},
	{
		start: [2]int{504, 971},
		end:   [2]int{989, 971},
	},
	{
		start: [2]int{926, 151},
		end:   [2]int{926, 480},
	},
	{
		start: [2]int{722, 895},
		end:   [2]int{722, 488},
	},
	{
		start: [2]int{15, 14},
		end:   [2]int{987, 986},
	},
	{
		start: [2]int{378, 486},
		end:   [2]int{267, 597},
	},
	{
		start: [2]int{732, 418},
		end:   [2]int{157, 418},
	},
	{
		start: [2]int{252, 515},
		end:   [2]int{257, 520},
	},
	{
		start: [2]int{61, 828},
		end:   [2]int{659, 230},
	},
	{
		start: [2]int{116, 652},
		end:   [2]int{893, 652},
	},
	{
		start: [2]int{827, 196},
		end:   [2]int{827, 564},
	},
	{
		start: [2]int{677, 515},
		end:   [2]int{677, 257},
	},
	{
		start: [2]int{380, 897},
		end:   [2]int{132, 897},
	},
	{
		start: [2]int{812, 959},
		end:   [2]int{812, 23},
	},
	{
		start: [2]int{989, 382},
		end:   [2]int{294, 382},
	},
	{
		start: [2]int{973, 89},
		end:   [2]int{81, 981},
	},
	{
		start: [2]int{292, 920},
		end:   [2]int{987, 225},
	},
	{
		start: [2]int{441, 394},
		end:   [2]int{441, 469},
	},
	{
		start: [2]int{502, 662},
		end:   [2]int{502, 213},
	},
	{
		start: [2]int{609, 570},
		end:   [2]int{609, 58},
	},
	{
		start: [2]int{559, 47},
		end:   [2]int{208, 47},
	},
	{
		start: [2]int{77, 192},
		end:   [2]int{277, 192},
	},
	{
		start: [2]int{229, 588},
		end:   [2]int{66, 588},
	},
	{
		start: [2]int{705, 363},
		end:   [2]int{705, 161},
	},
	{
		start: [2]int{944, 51},
		end:   [2]int{78, 917},
	},
	{
		start: [2]int{699, 889},
		end:   [2]int{699, 354},
	},
	{
		start: [2]int{90, 48},
		end:   [2]int{955, 913},
	},
	{
		start: [2]int{166, 491},
		end:   [2]int{24, 633},
	},
	{
		start: [2]int{154, 482},
		end:   [2]int{113, 441},
	},
	{
		start: [2]int{989, 989},
		end:   [2]int{10, 10},
	},
	{
		start: [2]int{421, 414},
		end:   [2]int{791, 44},
	},
	{
		start: [2]int{360, 272},
		end:   [2]int{966, 272},
	},
	{
		start: [2]int{264, 631},
		end:   [2]int{630, 631},
	},
	{
		start: [2]int{541, 50},
		end:   [2]int{541, 911},
	},
	{
		start: [2]int{17, 475},
		end:   [2]int{289, 203},
	},
	{
		start: [2]int{226, 772},
		end:   [2]int{697, 301},
	},
	{
		start: [2]int{163, 625},
		end:   [2]int{163, 513},
	},
	{
		start: [2]int{642, 971},
		end:   [2]int{642, 754},
	},
	{
		start: [2]int{975, 329},
		end:   [2]int{793, 329},
	},
	{
		start: [2]int{793, 878},
		end:   [2]int{793, 938},
	},
	{
		start: [2]int{10, 95},
		end:   [2]int{175, 95},
	},
	{
		start: [2]int{352, 903},
		end:   [2]int{352, 176},
	},
	{
		start: [2]int{914, 92},
		end:   [2]int{91, 915},
	},
	{
		start: [2]int{649, 768},
		end:   [2]int{649, 136},
	},
	{
		start: [2]int{347, 492},
		end:   [2]int{347, 977},
	},
	{
		start: [2]int{372, 839},
		end:   [2]int{372, 741},
	},
	{
		start: [2]int{587, 534},
		end:   [2]int{526, 534},
	},
	{
		start: [2]int{563, 936},
		end:   [2]int{102, 475},
	},
	{
		start: [2]int{126, 708},
		end:   [2]int{362, 708},
	},
	{
		start: [2]int{326, 869},
		end:   [2]int{326, 640},
	},
	{
		start: [2]int{358, 959},
		end:   [2]int{358, 408},
	},
	{
		start: [2]int{221, 99},
		end:   [2]int{221, 659},
	},
	{
		start: [2]int{572, 405},
		end:   [2]int{906, 71},
	},
	{
		start: [2]int{592, 664},
		end:   [2]int{687, 759},
	},
	{
		start: [2]int{618, 457},
		end:   [2]int{388, 687},
	},
	{
		start: [2]int{712, 850},
		end:   [2]int{245, 383},
	},
	{
		start: [2]int{981, 22},
		end:   [2]int{45, 958},
	},
	{
		start: [2]int{107, 884},
		end:   [2]int{340, 651},
	},
	{
		start: [2]int{17, 896},
		end:   [2]int{642, 896},
	},
	{
		start: [2]int{488, 135},
		end:   [2]int{851, 135},
	},
	{
		start: [2]int{54, 76},
		end:   [2]int{184, 76},
	},
	{
		start: [2]int{290, 596},
		end:   [2]int{290, 478},
	},
	{
		start: [2]int{468, 427},
		end:   [2]int{468, 316},
	},
	{
		start: [2]int{412, 434},
		end:   [2]int{412, 581},
	},
	{
		start: [2]int{899, 681},
		end:   [2]int{238, 20},
	},
	{
		start: [2]int{647, 231},
		end:   [2]int{542, 231},
	},
	{
		start: [2]int{54, 374},
		end:   [2]int{302, 622},
	},
	{
		start: [2]int{586, 555},
		end:   [2]int{13, 555},
	},
	{
		start: [2]int{875, 930},
		end:   [2]int{26, 81},
	},
	{
		start: [2]int{875, 115},
		end:   [2]int{127, 863},
	},
	{
		start: [2]int{863, 42},
		end:   [2]int{45, 860},
	},
	{
		start: [2]int{708, 862},
		end:   [2]int{100, 862},
	},
	{
		start: [2]int{190, 490},
		end:   [2]int{26, 654},
	},
	{
		start: [2]int{347, 944},
		end:   [2]int{711, 580},
	},
	{
		start: [2]int{259, 533},
		end:   [2]int{259, 516},
	},
	{
		start: [2]int{833, 790},
		end:   [2]int{891, 848},
	},
	{
		start: [2]int{556, 583},
		end:   [2]int{921, 948},
	},
	{
		start: [2]int{745, 929},
		end:   [2]int{745, 569},
	},
	{
		start: [2]int{111, 100},
		end:   [2]int{499, 100},
	},
	{
		start: [2]int{638, 903},
		end:   [2]int{525, 903},
	},
	{
		start: [2]int{726, 388},
		end:   [2]int{973, 388},
	},
	{
		start: [2]int{335, 504},
		end:   [2]int{638, 504},
	},
	{
		start: [2]int{628, 29},
		end:   [2]int{159, 29},
	},
	{
		start: [2]int{375, 406},
		end:   [2]int{200, 406},
	},
	{
		start: [2]int{12, 819},
		end:   [2]int{945, 819},
	},
	{
		start: [2]int{660, 330},
		end:   [2]int{318, 672},
	},
	{
		start: [2]int{436, 477},
		end:   [2]int{436, 988},
	},
	{
		start: [2]int{925, 41},
		end:   [2]int{464, 41},
	},
	{
		start: [2]int{868, 485},
		end:   [2]int{868, 109},
	},
	{
		start: [2]int{130, 859},
		end:   [2]int{979, 10},
	},
	{
		start: [2]int{895, 50},
		end:   [2]int{895, 568},
	},
	{
		start: [2]int{582, 943},
		end:   [2]int{582, 904},
	},
	{
		start: [2]int{589, 616},
		end:   [2]int{589, 590},
	},
	{
		start: [2]int{773, 359},
		end:   [2]int{441, 691},
	},
	{
		start: [2]int{396, 22},
		end:   [2]int{396, 730},
	},
	{
		start: [2]int{862, 947},
		end:   [2]int{30, 115},
	},
	{
		start: [2]int{573, 543},
		end:   [2]int{40, 10},
	},
	{
		start: [2]int{726, 743},
		end:   [2]int{726, 934},
	},
	{
		start: [2]int{360, 170},
		end:   [2]int{360, 187},
	},
	{
		start: [2]int{597, 287},
		end:   [2]int{982, 287},
	},
	{
		start: [2]int{537, 112},
		end:   [2]int{838, 112},
	},
	{
		start: [2]int{702, 683},
		end:   [2]int{151, 683},
	},
	{
		start: [2]int{770, 792},
		end:   [2]int{752, 792},
	},
	{
		start: [2]int{964, 60},
		end:   [2]int{896, 60},
	},
	{
		start: [2]int{53, 642},
		end:   [2]int{278, 642},
	},
	{
		start: [2]int{414, 871},
		end:   [2]int{798, 487},
	},
	{
		start: [2]int{96, 950},
		end:   [2]int{96, 983},
	},
	{
		start: [2]int{251, 65},
		end:   [2]int{289, 65},
	},
	{
		start: [2]int{797, 666},
		end:   [2]int{797, 200},
	},
	{
		start: [2]int{582, 157},
		end:   [2]int{582, 538},
	},
	{
		start: [2]int{107, 398},
		end:   [2]int{594, 885},
	},
	{
		start: [2]int{96, 66},
		end:   [2]int{806, 776},
	},
	{
		start: [2]int{124, 911},
		end:   [2]int{347, 911},
	},
	{
		start: [2]int{974, 538},
		end:   [2]int{974, 318},
	},
	{
		start: [2]int{45, 966},
		end:   [2]int{226, 785},
	},
	{
		start: [2]int{39, 960},
		end:   [2]int{827, 172},
	},
	{
		start: [2]int{163, 939},
		end:   [2]int{709, 939},
	},
	{
		start: [2]int{351, 540},
		end:   [2]int{351, 954},
	},
	{
		start: [2]int{656, 894},
		end:   [2]int{220, 458},
	},
	{
		start: [2]int{278, 314},
		end:   [2]int{278, 146},
	},
	{
		start: [2]int{637, 784},
		end:   [2]int{637, 283},
	},
	{
		start: [2]int{83, 690},
		end:   [2]int{899, 690},
	},
	{
		start: [2]int{16, 48},
		end:   [2]int{884, 916},
	},
	{
		start: [2]int{681, 865},
		end:   [2]int{310, 494},
	},
	{
		start: [2]int{333, 631},
		end:   [2]int{333, 832},
	},
	{
		start: [2]int{527, 652},
		end:   [2]int{527, 836},
	},
	{
		start: [2]int{352, 343},
		end:   [2]int{352, 623},
	},
	{
		start: [2]int{256, 316},
		end:   [2]int{479, 93},
	},
	{
		start: [2]int{450, 86},
		end:   [2]int{489, 86},
	},
	{
		start: [2]int{814, 834},
		end:   [2]int{814, 494},
	},
	{
		start: [2]int{406, 947},
		end:   [2]int{783, 947},
	},
	{
		start: [2]int{811, 643},
		end:   [2]int{318, 643},
	},
	{
		start: [2]int{240, 651},
		end:   [2]int{366, 651},
	},
	{
		start: [2]int{902, 618},
		end:   [2]int{303, 19},
	},
	{
		start: [2]int{843, 939},
		end:   [2]int{729, 939},
	},
	{
		start: [2]int{901, 149},
		end:   [2]int{131, 919},
	},
	{
		start: [2]int{365, 459},
		end:   [2]int{222, 459},
	},
	{
		start: [2]int{909, 218},
		end:   [2]int{426, 701},
	},
	{
		start: [2]int{746, 415},
		end:   [2]int{746, 199},
	},
	{
		start: [2]int{249, 327},
		end:   [2]int{807, 885},
	},
	{
		start: [2]int{760, 923},
		end:   [2]int{860, 923},
	},
	{
		start: [2]int{506, 259},
		end:   [2]int{506, 357},
	},
	{
		start: [2]int{933, 892},
		end:   [2]int{143, 892},
	},
	{
		start: [2]int{88, 589},
		end:   [2]int{88, 77},
	},
	{
		start: [2]int{597, 554},
		end:   [2]int{810, 554},
	},
	{
		start: [2]int{505, 574},
		end:   [2]int{505, 812},
	},
	{
		start: [2]int{211, 786},
		end:   [2]int{906, 91},
	},
	{
		start: [2]int{387, 238},
		end:   [2]int{480, 238},
	},
	{
		start: [2]int{394, 729},
		end:   [2]int{422, 757},
	},
	{
		start: [2]int{526, 436},
		end:   [2]int{526, 12},
	},
	{
		start: [2]int{660, 397},
		end:   [2]int{660, 290},
	},
	{
		start: [2]int{856, 469},
		end:   [2]int{176, 469},
	},
	{
		start: [2]int{653, 731},
		end:   [2]int{370, 731},
	},
	{
		start: [2]int{542, 241},
		end:   [2]int{542, 32},
	},
	{
		start: [2]int{471, 734},
		end:   [2]int{471, 384},
	},
	{
		start: [2]int{975, 468},
		end:   [2]int{783, 468},
	},
	{
		start: [2]int{25, 578},
		end:   [2]int{580, 578},
	},
	{
		start: [2]int{52, 632},
		end:   [2]int{551, 133},
	},
	{
		start: [2]int{150, 791},
		end:   [2]int{672, 791},
	},
	{
		start: [2]int{643, 348},
		end:   [2]int{643, 869},
	},
	{
		start: [2]int{893, 514},
		end:   [2]int{893, 422},
	},
	{
		start: [2]int{400, 463},
		end:   [2]int{335, 463},
	},
	{
		start: [2]int{564, 917},
		end:   [2]int{676, 917},
	},
	{
		start: [2]int{166, 433},
		end:   [2]int{166, 246},
	},
	{
		start: [2]int{798, 36},
		end:   [2]int{69, 765},
	},
	{
		start: [2]int{118, 977},
		end:   [2]int{882, 977},
	},
	{
		start: [2]int{718, 415},
		end:   [2]int{75, 415},
	},
	{
		start: [2]int{690, 807},
		end:   [2]int{690, 659},
	},
	{
		start: [2]int{163, 809},
		end:   [2]int{269, 809},
	},
	{
		start: [2]int{715, 238},
		end:   [2]int{715, 314},
	},
	{
		start: [2]int{970, 924},
		end:   [2]int{104, 58},
	},
	{
		start: [2]int{683, 762},
		end:   [2]int{683, 467},
	},
	{
		start: [2]int{554, 375},
		end:   [2]int{980, 801},
	},
	{
		start: [2]int{361, 130},
		end:   [2]int{361, 66},
	},
	{
		start: [2]int{879, 491},
		end:   [2]int{879, 843},
	},
	{
		start: [2]int{515, 700},
		end:   [2]int{515, 454},
	},
	{
		start: [2]int{465, 432},
		end:   [2]int{465, 444},
	},
	{
		start: [2]int{250, 239},
		end:   [2]int{216, 273},
	},
	{
		start: [2]int{894, 818},
		end:   [2]int{163, 818},
	},
	{
		start: [2]int{190, 790},
		end:   [2]int{190, 616},
	},
	{
		start: [2]int{384, 263},
		end:   [2]int{384, 84},
	},
	{
		start: [2]int{63, 875},
		end:   [2]int{851, 87},
	},
	{
		start: [2]int{154, 293},
		end:   [2]int{278, 417},
	},
	{
		start: [2]int{21, 592},
		end:   [2]int{883, 592},
	},
	{
		start: [2]int{372, 286},
		end:   [2]int{588, 70},
	},
	{
		start: [2]int{972, 447},
		end:   [2]int{972, 639},
	},
	{
		start: [2]int{838, 60},
		end:   [2]int{681, 60},
	},
	{
		start: [2]int{38, 366},
		end:   [2]int{38, 907},
	},
	{
		start: [2]int{746, 65},
		end:   [2]int{459, 65},
	},
	{
		start: [2]int{138, 640},
		end:   [2]int{66, 640},
	},
	{
		start: [2]int{536, 309},
		end:   [2]int{536, 109},
	},
	{
		start: [2]int{772, 634},
		end:   [2]int{772, 566},
	},
	{
		start: [2]int{43, 949},
		end:   [2]int{945, 47},
	},
	{
		start: [2]int{914, 85},
		end:   [2]int{395, 85},
	},
	{
		start: [2]int{25, 12},
		end:   [2]int{977, 964},
	},
	{
		start: [2]int{679, 455},
		end:   [2]int{679, 439},
	},
	{
		start: [2]int{420, 492},
		end:   [2]int{614, 492},
	},
	{
		start: [2]int{823, 658},
		end:   [2]int{823, 634},
	},
	{
		start: [2]int{45, 332},
		end:   [2]int{45, 943},
	},
	{
		start: [2]int{807, 344},
		end:   [2]int{807, 756},
	},
	{
		start: [2]int{634, 974},
		end:   [2]int{634, 892},
	},
	{
		start: [2]int{26, 26},
		end:   [2]int{988, 988},
	},
	{
		start: [2]int{628, 772},
		end:   [2]int{15, 772},
	},
	{
		start: [2]int{829, 614},
		end:   [2]int{550, 614},
	},
	{
		start: [2]int{513, 649},
		end:   [2]int{513, 369},
	},
	{
		start: [2]int{607, 923},
		end:   [2]int{607, 801},
	},
	{
		start: [2]int{809, 340},
		end:   [2]int{450, 699},
	},
	{
		start: [2]int{550, 193},
		end:   [2]int{666, 193},
	},
	{
		start: [2]int{175, 961},
		end:   [2]int{902, 234},
	},
	{
		start: [2]int{467, 146},
		end:   [2]int{500, 146},
	},
	{
		start: [2]int{543, 510},
		end:   [2]int{543, 626},
	},
	{
		start: [2]int{667, 52},
		end:   [2]int{667, 161},
	},
	{
		start: [2]int{635, 299},
		end:   [2]int{375, 299},
	},
	{
		start: [2]int{278, 807},
		end:   [2]int{904, 807},
	},
	{
		start: [2]int{269, 290},
		end:   [2]int{644, 290},
	},
	{
		start: [2]int{630, 268},
		end:   [2]int{630, 440},
	},
	{
		start: [2]int{241, 929},
		end:   [2]int{882, 288},
	},
	{
		start: [2]int{864, 907},
		end:   [2]int{360, 907},
	},
	{
		start: [2]int{455, 894},
		end:   [2]int{455, 265},
	},
	{
		start: [2]int{257, 43},
		end:   [2]int{257, 519},
	},
	{
		start: [2]int{414, 83},
		end:   [2]int{360, 83},
	},
	{
		start: [2]int{237, 64},
		end:   [2]int{237, 612},
	},
	{
		start: [2]int{260, 541},
		end:   [2]int{260, 927},
	},
	{
		start: [2]int{323, 909},
		end:   [2]int{323, 583},
	},
	{
		start: [2]int{929, 354},
		end:   [2]int{929, 695},
	},
	{
		start: [2]int{912, 914},
		end:   [2]int{40, 42},
	},
	{
		start: [2]int{579, 401},
		end:   [2]int{392, 401},
	},
	{
		start: [2]int{389, 222},
		end:   [2]int{895, 728},
	},
	{
		start: [2]int{831, 696},
		end:   [2]int{831, 707},
	},
	{
		start: [2]int{871, 304},
		end:   [2]int{212, 304},
	},
	{
		start: [2]int{207, 333},
		end:   [2]int{621, 333},
	},
	{
		start: [2]int{225, 897},
		end:   [2]int{355, 767},
	},
	{
		start: [2]int{883, 68},
		end:   [2]int{84, 867},
	},
	{
		start: [2]int{115, 397},
		end:   [2]int{115, 208},
	},
	{
		start: [2]int{889, 217},
		end:   [2]int{985, 217},
	},
	{
		start: [2]int{793, 402},
		end:   [2]int{250, 402},
	},
	{
		start: [2]int{555, 367},
		end:   [2]int{61, 861},
	},
	{
		start: [2]int{732, 954},
		end:   [2]int{466, 688},
	},
	{
		start: [2]int{39, 564},
		end:   [2]int{39, 481},
	},
	{
		start: [2]int{283, 816},
		end:   [2]int{346, 816},
	},
	{
		start: [2]int{383, 506},
		end:   [2]int{276, 506},
	},
	{
		start: [2]int{394, 661},
		end:   [2]int{394, 143},
	},
	{
		start: [2]int{988, 983},
		end:   [2]int{66, 61},
	},
	{
		start: [2]int{652, 638},
		end:   [2]int{652, 569},
	},
	{
		start: [2]int{185, 64},
		end:   [2]int{487, 64},
	},
	{
		start: [2]int{354, 935},
		end:   [2]int{251, 935},
	},
	{
		start: [2]int{201, 460},
		end:   [2]int{201, 552},
	},
	{
		start: [2]int{836, 285},
		end:   [2]int{836, 666},
	},
	{
		start: [2]int{878, 312},
		end:   [2]int{359, 831},
	},
	{
		start: [2]int{443, 684},
		end:   [2]int{887, 240},
	},
	{
		start: [2]int{221, 49},
		end:   [2]int{948, 776},
	},
	{
		start: [2]int{243, 959},
		end:   [2]int{22, 959},
	},
	{
		start: [2]int{573, 323},
		end:   [2]int{834, 323},
	},
	{
		start: [2]int{745, 734},
		end:   [2]int{456, 734},
	},
	{
		start: [2]int{594, 244},
		end:   [2]int{908, 244},
	},
	{
		start: [2]int{583, 360},
		end:   [2]int{578, 355},
	},
	{
		start: [2]int{288, 38},
		end:   [2]int{288, 364},
	},
	{
		start: [2]int{565, 339},
		end:   [2]int{251, 653},
	},
	{
		start: [2]int{215, 196},
		end:   [2]int{215, 476},
	},
	{
		start: [2]int{270, 705},
		end:   [2]int{586, 705},
	},
	{
		start: [2]int{749, 477},
		end:   [2]int{749, 658},
	},
	{
		start: [2]int{917, 838},
		end:   [2]int{511, 432},
	},
	{
		start: [2]int{935, 187},
		end:   [2]int{935, 381},
	},
	{
		start: [2]int{181, 190},
		end:   [2]int{323, 48},
	},
	{
		start: [2]int{399, 491},
		end:   [2]int{399, 779},
	},
	{
		start: [2]int{861, 798},
		end:   [2]int{91, 28},
	},
	{
		start: [2]int{160, 115},
		end:   [2]int{58, 115},
	},
	{
		start: [2]int{940, 68},
		end:   [2]int{940, 590},
	},
	{
		start: [2]int{806, 958},
		end:   [2]int{35, 187},
	},
	{
		start: [2]int{184, 538},
		end:   [2]int{438, 284},
	},
	{
		start: [2]int{283, 904},
		end:   [2]int{283, 114},
	},
	{
		start: [2]int{344, 935},
		end:   [2]int{222, 935},
	},
	{
		start: [2]int{435, 962},
		end:   [2]int{367, 962},
	},
	{
		start: [2]int{837, 768},
		end:   [2]int{837, 583},
	},
	{
		start: [2]int{100, 423},
		end:   [2]int{826, 423},
	},
	{
		start: [2]int{299, 172},
		end:   [2]int{465, 172},
	},
	{
		start: [2]int{130, 136},
		end:   [2]int{181, 187},
	},
	{
		start: [2]int{969, 759},
		end:   [2]int{55, 759},
	},
	{
		start: [2]int{936, 711},
		end:   [2]int{521, 711},
	},
	{
		start: [2]int{268, 619},
		end:   [2]int{349, 619},
	},
	{
		start: [2]int{946, 119},
		end:   [2]int{108, 957},
	},
	{
		start: [2]int{940, 25},
		end:   [2]int{10, 955},
	},
	{
		start: [2]int{867, 494},
		end:   [2]int{652, 279},
	},
	{
		start: [2]int{535, 202},
		end:   [2]int{321, 202},
	},
	{
		start: [2]int{876, 14},
		end:   [2]int{24, 866},
	},
	{
		start: [2]int{887, 208},
		end:   [2]int{887, 265},
	},
	{
		start: [2]int{129, 12},
		end:   [2]int{42, 12},
	},
	{
		start: [2]int{514, 800},
		end:   [2]int{940, 374},
	},
	{
		start: [2]int{722, 306},
		end:   [2]int{722, 418},
	},
	{
		start: [2]int{24, 928},
		end:   [2]int{935, 17},
	},
	{
		start: [2]int{798, 279},
		end:   [2]int{798, 293},
	},
	{
		start: [2]int{384, 701},
		end:   [2]int{193, 701},
	},
	{
		start: [2]int{100, 644},
		end:   [2]int{593, 644},
	},
	{
		start: [2]int{818, 48},
		end:   [2]int{216, 48},
	},
	{
		start: [2]int{51, 984},
		end:   [2]int{949, 86},
	},
	{
		start: [2]int{843, 494},
		end:   [2]int{843, 723},
	},
	{
		start: [2]int{809, 156},
		end:   [2]int{129, 836},
	},
	{
		start: [2]int{500, 38},
		end:   [2]int{656, 38},
	},
	{
		start: [2]int{311, 705},
		end:   [2]int{311, 101},
	},
	{
		start: [2]int{21, 850},
		end:   [2]int{21, 316},
	},
	{
		start: [2]int{530, 628},
		end:   [2]int{511, 628},
	},
	{
		start: [2]int{106, 366},
		end:   [2]int{415, 675},
	},
	{
		start: [2]int{542, 882},
		end:   [2]int{325, 665},
	},
	{
		start: [2]int{987, 937},
		end:   [2]int{987, 793},
	},
	{
		start: [2]int{926, 260},
		end:   [2]int{264, 922},
	},
	{
		start: [2]int{768, 149},
		end:   [2]int{914, 149},
	},
	{
		start: [2]int{548, 71},
		end:   [2]int{548, 812},
	},
	{
		start: [2]int{51, 946},
		end:   [2]int{812, 946},
	},
	{
		start: [2]int{430, 439},
		end:   [2]int{954, 963},
	},
	{
		start: [2]int{529, 301},
		end:   [2]int{133, 301},
	},
	{
		start: [2]int{282, 890},
		end:   [2]int{720, 890},
	},
	{
		start: [2]int{876, 231},
		end:   [2]int{336, 771},
	},
	{
		start: [2]int{489, 471},
		end:   [2]int{934, 471},
	},
	{
		start: [2]int{585, 174},
		end:   [2]int{100, 174},
	},
	{
		start: [2]int{284, 489},
		end:   [2]int{163, 489},
	},
	{
		start: [2]int{989, 983},
		end:   [2]int{33, 27},
	},
	{
		start: [2]int{31, 213},
		end:   [2]int{662, 213},
	},
	{
		start: [2]int{133, 832},
		end:   [2]int{559, 406},
	},
	{
		start: [2]int{730, 345},
		end:   [2]int{730, 194},
	},
	{
		start: [2]int{860, 288},
		end:   [2]int{736, 412},
	},
	{
		start: [2]int{110, 351},
		end:   [2]int{581, 351},
	},
	{
		start: [2]int{417, 151},
		end:   [2]int{77, 491},
	},
	{
		start: [2]int{674, 671},
		end:   [2]int{674, 711},
	},
	{
		start: [2]int{514, 867},
		end:   [2]int{514, 100},
	},
	{
		start: [2]int{885, 595},
		end:   [2]int{885, 680},
	},
	{
		start: [2]int{44, 31},
		end:   [2]int{928, 915},
	},
	{
		start: [2]int{969, 347},
		end:   [2]int{69, 347},
	},
	{
		start: [2]int{597, 227},
		end:   [2]int{357, 227},
	},
	{
		start: [2]int{347, 443},
		end:   [2]int{347, 216},
	},
	{
		start: [2]int{781, 736},
		end:   [2]int{781, 93},
	},
	{
		start: [2]int{968, 559},
		end:   [2]int{968, 81},
	},
	{
		start: [2]int{35, 93},
		end:   [2]int{232, 93},
	},
	{
		start: [2]int{273, 837},
		end:   [2]int{97, 837},
	},
	{
		start: [2]int{949, 833},
		end:   [2]int{748, 632},
	},
	{
		start: [2]int{712, 773},
		end:   [2]int{221, 773},
	},
	{
		start: [2]int{194, 884},
		end:   [2]int{978, 100},
	},
	{
		start: [2]int{217, 816},
		end:   [2]int{217, 861},
	},
	{
		start: [2]int{651, 122},
		end:   [2]int{71, 122},
	},
	{
		start: [2]int{166, 551},
		end:   [2]int{166, 892},
	},
	{
		start: [2]int{285, 193},
		end:   [2]int{883, 193},
	},
	{
		start: [2]int{858, 934},
		end:   [2]int{125, 201},
	},
	{
		start: [2]int{180, 190},
		end:   [2]int{577, 190},
	},
	{
		start: [2]int{491, 685},
		end:   [2]int{690, 486},
	},
	{
		start: [2]int{666, 598},
		end:   [2]int{337, 269},
	},
	{
		start: [2]int{455, 571},
		end:   [2]int{753, 571},
	},
	{
		start: [2]int{11, 769},
		end:   [2]int{11, 507},
	},
	{
		start: [2]int{391, 663},
		end:   [2]int{323, 595},
	},
	{
		start: [2]int{70, 740},
		end:   [2]int{70, 928},
	},
	{
		start: [2]int{205, 525},
		end:   [2]int{534, 854},
	},
	{
		start: [2]int{890, 851},
		end:   [2]int{151, 851},
	},
	{
		start: [2]int{382, 662},
		end:   [2]int{849, 195},
	},
	{
		start: [2]int{201, 870},
		end:   [2]int{201, 506},
	},
	{
		start: [2]int{549, 549},
		end:   [2]int{549, 528},
	},
	{
		start: [2]int{343, 172},
		end:   [2]int{601, 172},
	},
	{
		start: [2]int{22, 732},
		end:   [2]int{750, 732},
	},
	{
		start: [2]int{221, 689},
		end:   [2]int{881, 29},
	},
	{
		start: [2]int{628, 559},
		end:   [2]int{747, 559},
	},
	{
		start: [2]int{668, 879},
		end:   [2]int{437, 879},
	},
	{
		start: [2]int{712, 139},
		end:   [2]int{38, 139},
	},
	{
		start: [2]int{547, 322},
		end:   [2]int{905, 322},
	},
	{
		start: [2]int{872, 304},
		end:   [2]int{719, 304},
	},
	{
		start: [2]int{469, 604},
		end:   [2]int{389, 524},
	},
	{
		start: [2]int{256, 91},
		end:   [2]int{746, 91},
	},
	{
		start: [2]int{881, 548},
		end:   [2]int{641, 548},
	},
	{
		start: [2]int{683, 417},
		end:   [2]int{683, 800},
	},
	{
		start: [2]int{811, 917},
		end:   [2]int{646, 917},
	},
	{
		start: [2]int{578, 556},
		end:   [2]int{207, 185},
	},
	{
		start: [2]int{732, 343},
		end:   [2]int{260, 343},
	},
	{
		start: [2]int{86, 869},
		end:   [2]int{882, 73},
	},
	{
		start: [2]int{370, 587},
		end:   [2]int{765, 192},
	},
	{
		start: [2]int{649, 621},
		end:   [2]int{649, 165},
	},
	{
		start: [2]int{298, 339},
		end:   [2]int{298, 523},
	},
	{
		start: [2]int{131, 771},
		end:   [2]int{803, 99},
	},
	{
		start: [2]int{934, 791},
		end:   [2]int{934, 29},
	},
	{
		start: [2]int{782, 13},
		end:   [2]int{782, 741},
	},
	{
		start: [2]int{852, 808},
		end:   [2]int{852, 594},
	},
	{
		start: [2]int{390, 217},
		end:   [2]int{153, 217},
	},
	{
		start: [2]int{858, 980},
		end:   [2]int{94, 216},
	},
	{
		start: [2]int{832, 467},
		end:   [2]int{783, 418},
	},
	{
		start: [2]int{188, 49},
		end:   [2]int{981, 842},
	},
	{
		start: [2]int{438, 467},
		end:   [2]int{76, 829},
	},
	{
		start: [2]int{47, 911},
		end:   [2]int{164, 911},
	},
	{
		start: [2]int{670, 414},
		end:   [2]int{533, 414},
	},
	{
		start: [2]int{58, 61},
		end:   [2]int{740, 743},
	},
	{
		start: [2]int{264, 686},
		end:   [2]int{264, 799},
	},
	{
		start: [2]int{506, 300},
		end:   [2]int{64, 300},
	},
	{
		start: [2]int{509, 717},
		end:   [2]int{509, 952},
	},
	{
		start: [2]int{81, 819},
		end:   [2]int{81, 694},
	},
	{
		start: [2]int{512, 543},
		end:   [2]int{427, 543},
	},
	{
		start: [2]int{235, 78},
		end:   [2]int{788, 78},
	},
	{
		start: [2]int{952, 133},
		end:   [2]int{644, 133},
	},
	{
		start: [2]int{188, 302},
		end:   [2]int{695, 302},
	},
	{
		start: [2]int{272, 868},
		end:   [2]int{845, 295},
	},
	{
		start: [2]int{288, 413},
		end:   [2]int{704, 413},
	},
	{
		start: [2]int{774, 671},
		end:   [2]int{774, 24},
	},
	{
		start: [2]int{296, 932},
		end:   [2]int{296, 16},
	},
	{
		start: [2]int{99, 789},
		end:   [2]int{300, 789},
	},
	{
		start: [2]int{630, 560},
		end:   [2]int{630, 896},
	},
	{
		start: [2]int{328, 289},
		end:   [2]int{280, 289},
	},
	{
		start: [2]int{786, 772},
		end:   [2]int{294, 280},
	},
	{
		start: [2]int{437, 747},
		end:   [2]int{437, 110},
	},
	{
		start: [2]int{537, 709},
		end:   [2]int{42, 709},
	},
	{
		start: [2]int{655, 924},
		end:   [2]int{655, 117},
	},
	{
		start: [2]int{185, 65},
		end:   [2]int{963, 843},
	},
	{
		start: [2]int{70, 87},
		end:   [2]int{274, 87},
	},
	{
		start: [2]int{516, 727},
		end:   [2]int{183, 394},
	},
	{
		start: [2]int{322, 128},
		end:   [2]int{781, 587},
	},
	{
		start: [2]int{147, 278},
		end:   [2]int{482, 278},
	},
	{
		start: [2]int{188, 793},
		end:   [2]int{761, 793},
	},
	{
		start: [2]int{702, 441},
		end:   [2]int{702, 27},
	},
	{
		start: [2]int{686, 18},
		end:   [2]int{686, 275},
	},
	{
		start: [2]int{510, 254},
		end:   [2]int{510, 862},
	},
	{
		start: [2]int{666, 204},
		end:   [2]int{12, 204},
	},
	{
		start: [2]int{677, 63},
		end:   [2]int{677, 78},
	},
	{
		start: [2]int{868, 950},
		end:   [2]int{868, 110},
	},
	{
		start: [2]int{42, 845},
		end:   [2]int{739, 148},
	},
	{
		start: [2]int{343, 279},
		end:   [2]int{758, 279},
	},
	{
		start: [2]int{182, 792},
		end:   [2]int{727, 792},
	},
	{
		start: [2]int{346, 238},
		end:   [2]int{493, 238},
	},
	{
		start: [2]int{467, 493},
		end:   [2]int{467, 273},
	},
	{
		start: [2]int{823, 68},
		end:   [2]int{823, 886},
	},
	{
		start: [2]int{686, 302},
		end:   [2]int{39, 302},
	},
	{
		start: [2]int{984, 345},
		end:   [2]int{984, 936},
	},
	{
		start: [2]int{11, 480},
		end:   [2]int{11, 675},
	},
	{
		start: [2]int{989, 478},
		end:   [2]int{695, 772},
	},
	{
		start: [2]int{568, 235},
		end:   [2]int{535, 235},
	},
	{
		start: [2]int{203, 41},
		end:   [2]int{93, 41},
	},
	{
		start: [2]int{463, 569},
		end:   [2]int{304, 569},
	},
	{
		start: [2]int{909, 629},
		end:   [2]int{207, 629},
	},
	{
		start: [2]int{792, 678},
		end:   [2]int{792, 909},
	},
	{
		start: [2]int{486, 924},
		end:   [2]int{486, 948},
	},
	{
		start: [2]int{611, 79},
		end:   [2]int{611, 303},
	},
	{
		start: [2]int{762, 136},
		end:   [2]int{139, 759},
	},
	{
		start: [2]int{808, 872},
		end:   [2]int{726, 872},
	},
	{
		start: [2]int{22, 403},
		end:   [2]int{22, 401},
	},
	{
		start: [2]int{774, 134},
		end:   [2]int{369, 134},
	},
	{
		start: [2]int{131, 282},
		end:   [2]int{131, 849},
	},
	{
		start: [2]int{912, 245},
		end:   [2]int{912, 385},
	},
	{
		start: [2]int{338, 396},
		end:   [2]int{768, 396},
	},
	{
		start: [2]int{944, 978},
		end:   [2]int{20, 54},
	},
	{
		start: [2]int{623, 897},
		end:   [2]int{623, 10},
	},
	{
		start: [2]int{103, 402},
		end:   [2]int{207, 298},
	},
	{
		start: [2]int{39, 50},
		end:   [2]int{971, 50},
	},
	{
		start: [2]int{770, 423},
		end:   [2]int{882, 423},
	},
	{
		start: [2]int{195, 873},
		end:   [2]int{195, 40},
	},
	{
		start: [2]int{119, 659},
		end:   [2]int{119, 374},
	},
	{
		start: [2]int{678, 962},
		end:   [2]int{698, 962},
	},
	{
		start: [2]int{946, 64},
		end:   [2]int{946, 202},
	},
	{
		start: [2]int{790, 780},
		end:   [2]int{790, 66},
	},
	{
		start: [2]int{565, 21},
		end:   [2]int{614, 21},
	},
	{
		start: [2]int{617, 20},
		end:   [2]int{640, 20},
	},
	{
		start: [2]int{697, 773},
		end:   [2]int{697, 915},
	},
	{
		start: [2]int{467, 167},
		end:   [2]int{208, 167},
	},
	{
		start: [2]int{567, 713},
		end:   [2]int{567, 873},
	},
	{
		start: [2]int{120, 98},
		end:   [2]int{557, 98},
	},
	{
		start: [2]int{103, 395},
		end:   [2]int{103, 159},
	},
	{
		start: [2]int{148, 734},
		end:   [2]int{723, 159},
	},
	{
		start: [2]int{730, 949},
		end:   [2]int{730, 33},
	},
	{
		start: [2]int{322, 628},
		end:   [2]int{322, 272},
	},
	{
		start: [2]int{649, 57},
		end:   [2]int{44, 57},
	},
	{
		start: [2]int{261, 513},
		end:   [2]int{624, 513},
	},
	{
		start: [2]int{550, 414},
		end:   [2]int{738, 226},
	},
	{
		start: [2]int{774, 183},
		end:   [2]int{471, 486},
	},
	{
		start: [2]int{146, 659},
		end:   [2]int{146, 581},
	},
	{
		start: [2]int{599, 751},
		end:   [2]int{599, 320},
	},
	{
		start: [2]int{936, 225},
		end:   [2]int{226, 935},
	},
	{
		start: [2]int{378, 31},
		end:   [2]int{222, 187},
	},
	{
		start: [2]int{871, 691},
		end:   [2]int{502, 691},
	},
	{
		start: [2]int{195, 963},
		end:   [2]int{335, 963},
	},
	{
		start: [2]int{513, 465},
		end:   [2]int{382, 334},
	},
	{
		start: [2]int{620, 801},
		end:   [2]int{673, 801},
	},
	{
		start: [2]int{187, 428},
		end:   [2]int{318, 428},
	},
	{
		start: [2]int{572, 836},
		end:   [2]int{441, 836},
	},
	{
		start: [2]int{305, 398},
		end:   [2]int{305, 951},
	},
	{
		start: [2]int{978, 703},
		end:   [2]int{927, 703},
	},
	{
		start: [2]int{99, 219},
		end:   [2]int{846, 966},
	},
	{
		start: [2]int{952, 971},
		end:   [2]int{26, 45},
	},
	{
		start: [2]int{859, 775},
		end:   [2]int{859, 663},
	},
	{
		start: [2]int{144, 777},
		end:   [2]int{144, 390},
	},
	{
		start: [2]int{792, 859},
		end:   [2]int{441, 859},
	},
	{
		start: [2]int{513, 672},
		end:   [2]int{982, 203},
	},
	{
		start: [2]int{613, 342},
		end:   [2]int{671, 400},
	},
	{
		start: [2]int{802, 498},
		end:   [2]int{811, 498},
	},
	{
		start: [2]int{197, 240},
		end:   [2]int{197, 216},
	},
	{
		start: [2]int{45, 908},
		end:   [2]int{881, 72},
	},
	{
		start: [2]int{860, 573},
		end:   [2]int{12, 573},
	},
	{
		start: [2]int{817, 145},
		end:   [2]int{755, 83},
	},
	{
		start: [2]int{565, 562},
		end:   [2]int{660, 467},
	},
	{
		start: [2]int{918, 952},
		end:   [2]int{918, 111},
	},
	{
		start: [2]int{936, 174},
		end:   [2]int{936, 97},
	},
	{
		start: [2]int{630, 759},
		end:   [2]int{630, 89},
	},
	{
		start: [2]int{329, 762},
		end:   [2]int{608, 762},
	},
}
