
//line ./is_mobile.rl:1
/*
  To check whether a given user agent is a mobile agent and if than which one
*/
package ua

// Returns a bool flag whether the agent is mobile agent and also returns the platform name
func IsMobile(agent string) (bool, string) {
  var bytes []byte = make([]byte, len(agent))
  copy(bytes[:], agent)
  return _isMobile(bytes)
}
/*
  Ragel configuration to build an optimised regex
  using a Detriminstic Finite Automata machine
*/
func _isMobile(data []byte) (bool, string) {
  var os string
  var mobile bool = true

//line ./is_mobile.rl:20

//line src/ua/is_mobile.go:25
var _scanner_actions []byte = []byte{
	0, 2, 0, 12, 2, 1, 12, 2, 2, 
	12, 2, 3, 12, 2, 4, 12, 2, 
	5, 12, 2, 6, 12, 2, 7, 12, 
	2, 8, 12, 2, 9, 12, 2, 10, 
	12, 2, 11, 12, 
}

var _scanner_key_offsets []int16 = []int16{
	0, 16, 34, 50, 66, 82, 100, 116, 
	136, 152, 170, 188, 204, 220, 240, 258, 
	276, 294, 310, 328, 344, 362, 378, 396, 
	414, 432, 448, 466, 486, 506, 524, 542, 
	560, 578, 596, 614, 630, 646, 663, 681, 
	699, 717, 735, 753, 771, 787, 805, 821, 
	841, 857, 875, 893, 911, 927, 945, 963, 
	981, 997, 1017, 1035, 1051, 1069, 1087, 1105, 
	1123, 1139, 
}

var _scanner_trans_keys []byte = []byte{
	65, 66, 73, 76, 77, 80, 83, 87, 
	97, 98, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 78, 80, 83, 
	87, 97, 98, 105, 108, 109, 110, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 87, 97, 98, 105, 108, 109, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 87, 97, 98, 105, 108, 109, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 87, 97, 98, 105, 108, 109, 112, 
	115, 119, 65, 66, 73, 76, 77, 78, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	110, 112, 115, 119, 65, 66, 73, 76, 
	77, 80, 83, 87, 97, 98, 105, 108, 
	109, 112, 115, 119, 65, 66, 67, 73, 
	76, 77, 78, 80, 83, 87, 97, 98, 
	99, 105, 108, 109, 110, 112, 115, 119, 
	65, 66, 73, 76, 77, 80, 83, 87, 
	97, 98, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 78, 80, 83, 
	87, 97, 98, 105, 108, 109, 110, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 84, 87, 97, 98, 105, 108, 109, 
	112, 115, 116, 119, 65, 66, 73, 76, 
	77, 80, 83, 87, 97, 98, 105, 108, 
	109, 112, 115, 119, 65, 66, 73, 76, 
	77, 80, 83, 87, 97, 98, 105, 108, 
	109, 112, 115, 119, 65, 66, 73, 76, 
	77, 78, 80, 83, 87, 89, 97, 98, 
	105, 108, 109, 110, 112, 115, 119, 121, 
	65, 66, 68, 73, 76, 77, 80, 83, 
	87, 97, 98, 100, 105, 108, 109, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	82, 83, 87, 97, 98, 105, 108, 109, 
	112, 114, 115, 119, 65, 66, 73, 76, 
	77, 79, 80, 83, 87, 97, 98, 105, 
	108, 109, 111, 112, 115, 119, 65, 66, 
	73, 76, 77, 80, 83, 87, 97, 98, 
	105, 108, 109, 112, 115, 119, 65, 66, 
	68, 73, 76, 77, 80, 83, 87, 97, 
	98, 100, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 80, 83, 87, 
	97, 98, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 78, 80, 83, 
	87, 97, 98, 105, 108, 109, 110, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 87, 97, 98, 105, 108, 109, 112, 
	115, 119, 65, 66, 73, 76, 77, 80, 
	83, 85, 87, 97, 98, 105, 108, 109, 
	112, 115, 117, 119, 65, 66, 73, 76, 
	77, 78, 80, 83, 87, 97, 98, 105, 
	108, 109, 110, 112, 115, 119, 65, 66, 
	71, 73, 76, 77, 80, 83, 87, 97, 
	98, 103, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 80, 83, 87, 
	97, 98, 105, 108, 109, 112, 115, 119, 
	65, 66, 73, 76, 77, 78, 80, 83, 
	87, 97, 98, 105, 108, 109, 110, 112, 
	115, 119, 65, 66, 72, 73, 76, 77, 
	79, 80, 83, 87, 97, 98, 104, 105, 
	108, 109, 111, 112, 115, 119, 65, 66, 
	68, 73, 76, 77, 78, 80, 83, 87, 
	97, 98, 100, 105, 108, 109, 110, 112, 
	115, 119, 65, 66, 73, 76, 77, 79, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	111, 112, 115, 119, 65, 66, 73, 76, 
	77, 78, 80, 83, 87, 97, 98, 105, 
	108, 109, 110, 112, 115, 119, 65, 66, 
	69, 73, 76, 77, 80, 83, 87, 97, 
	98, 101, 105, 108, 109, 112, 115, 119, 
	65, 66, 68, 73, 76, 77, 80, 83, 
	87, 97, 98, 100, 105, 108, 109, 112, 
	115, 119, 65, 66, 68, 73, 76, 77, 
	80, 83, 87, 97, 98, 100, 105, 108, 
	109, 112, 115, 119, 65, 66, 73, 76, 
	77, 79, 80, 83, 87, 97, 98, 105, 
	108, 109, 111, 112, 115, 119, 65, 66, 
	73, 76, 77, 80, 83, 87, 97, 98, 
	105, 108, 109, 112, 115, 119, 65, 66, 
	73, 76, 77, 80, 83, 87, 97, 98, 
	105, 108, 109, 112, 115, 119, 32, 65, 
	66, 73, 76, 77, 80, 83, 87, 97, 
	98, 105, 108, 109, 112, 115, 119, 65, 
	66, 73, 76, 77, 78, 80, 83, 87, 
	97, 98, 105, 108, 109, 110, 112, 115, 
	119, 65, 66, 73, 76, 77, 80, 83, 
	84, 87, 97, 98, 105, 108, 109, 112, 
	115, 116, 119, 65, 66, 72, 73, 76, 
	77, 80, 83, 87, 97, 98, 104, 105, 
	108, 109, 112, 115, 119, 65, 66, 73, 
	76, 77, 79, 80, 83, 87, 97, 98, 
	105, 108, 109, 111, 112, 115, 119, 65, 
	66, 73, 76, 77, 78, 80, 83, 87, 
	97, 98, 105, 108, 109, 110, 112, 115, 
	119, 65, 66, 69, 73, 76, 77, 80, 
	83, 87, 97, 98, 101, 105, 108, 109, 
	112, 115, 119, 65, 66, 73, 76, 77, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	112, 115, 119, 65, 66, 73, 76, 77, 
	80, 83, 84, 87, 97, 98, 105, 108, 
	109, 112, 115, 116, 119, 65, 66, 73, 
	76, 77, 80, 83, 87, 97, 98, 105, 
	108, 109, 112, 115, 119, 65, 66, 73, 
	76, 77, 78, 80, 83, 84, 87, 97, 
	98, 105, 108, 109, 110, 112, 115, 116, 
	119, 65, 66, 73, 76, 77, 80, 83, 
	87, 97, 98, 105, 108, 109, 112, 115, 
	119, 65, 66, 73, 76, 77, 79, 80, 
	83, 87, 97, 98, 105, 108, 109, 111, 
	112, 115, 119, 65, 66, 73, 76, 77, 
	78, 80, 83, 87, 97, 98, 105, 108, 
	109, 110, 112, 115, 119, 65, 66, 73, 
	76, 77, 79, 80, 83, 87, 97, 98, 
	105, 108, 109, 111, 112, 115, 119, 65, 
	66, 73, 76, 77, 80, 83, 87, 97, 
	98, 105, 108, 109, 112, 115, 119, 65, 
	66, 72, 73, 76, 77, 80, 83, 87, 
	97, 98, 104, 105, 108, 109, 112, 115, 
	119, 65, 66, 73, 76, 77, 80, 83, 
	85, 87, 97, 98, 105, 108, 109, 112, 
	115, 117, 119, 65, 66, 73, 76, 77, 
	80, 83, 87, 88, 97, 98, 105, 108, 
	109, 112, 115, 119, 120, 65, 66, 73, 
	76, 77, 80, 83, 87, 97, 98, 105, 
	108, 109, 112, 115, 119, 65, 66, 67, 
	73, 76, 77, 78, 80, 83, 87, 97, 
	98, 99, 105, 108, 109, 110, 112, 115, 
	119, 65, 66, 73, 75, 76, 77, 80, 
	83, 87, 97, 98, 105, 107, 108, 109, 
	112, 115, 119, 65, 66, 73, 76, 77, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	112, 115, 119, 65, 66, 69, 73, 76, 
	77, 80, 83, 87, 97, 98, 101, 105, 
	108, 109, 112, 115, 119, 65, 66, 73, 
	76, 77, 80, 82, 83, 87, 97, 98, 
	105, 108, 109, 112, 114, 115, 119, 65, 
	66, 73, 76, 77, 80, 82, 83, 87, 
	97, 98, 105, 108, 109, 112, 114, 115, 
	119, 65, 66, 73, 76, 77, 80, 83, 
	87, 89, 97, 98, 105, 108, 109, 112, 
	115, 119, 121, 65, 66, 73, 76, 77, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	112, 115, 119, 65, 66, 73, 76, 77, 
	80, 83, 87, 97, 98, 105, 108, 109, 
	112, 115, 119, 
}

var _scanner_single_lengths []byte = []byte{
	16, 18, 16, 16, 16, 18, 16, 20, 
	16, 18, 18, 16, 16, 20, 18, 18, 
	18, 16, 18, 16, 18, 16, 18, 18, 
	18, 16, 18, 20, 20, 18, 18, 18, 
	18, 18, 18, 16, 16, 17, 18, 18, 
	18, 18, 18, 18, 16, 18, 16, 20, 
	16, 18, 18, 18, 16, 18, 18, 18, 
	16, 20, 18, 16, 18, 18, 18, 18, 
	16, 16, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 
}

var _scanner_index_offsets []int16 = []int16{
	0, 17, 36, 53, 70, 87, 106, 123, 
	144, 161, 180, 199, 216, 233, 254, 273, 
	292, 311, 328, 347, 364, 383, 400, 419, 
	438, 457, 474, 493, 514, 535, 554, 573, 
	592, 611, 630, 649, 666, 683, 701, 720, 
	739, 758, 777, 796, 815, 832, 851, 868, 
	889, 906, 925, 944, 963, 980, 999, 1018, 
	1037, 1054, 1075, 1094, 1111, 1130, 1149, 1168, 
	1187, 1204, 
}

var _scanner_indicies []byte = []byte{
	1, 2, 3, 4, 5, 6, 7, 8, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 9, 6, 
	7, 8, 1, 2, 3, 4, 5, 9, 
	6, 7, 8, 0, 1, 2, 3, 10, 
	5, 6, 7, 8, 1, 2, 3, 10, 
	5, 6, 7, 8, 0, 1, 2, 3, 
	4, 5, 11, 7, 8, 1, 2, 3, 
	4, 5, 11, 7, 8, 0, 1, 2, 
	12, 4, 5, 6, 7, 8, 1, 2, 
	12, 4, 5, 6, 7, 8, 0, 1, 
	2, 3, 4, 5, 13, 11, 7, 8, 
	1, 2, 3, 4, 5, 13, 11, 7, 
	8, 0, 14, 2, 3, 4, 5, 6, 
	7, 8, 14, 2, 3, 4, 5, 6, 
	7, 8, 0, 1, 2, 15, 3, 4, 
	5, 9, 6, 7, 8, 1, 2, 15, 
	3, 4, 5, 9, 6, 7, 8, 0, 
	1, 2, 16, 4, 5, 6, 7, 8, 
	1, 2, 16, 4, 5, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 17, 11, 
	7, 8, 1, 2, 3, 4, 5, 17, 
	11, 7, 8, 0, 1, 2, 3, 4, 
	5, 6, 7, 18, 8, 1, 2, 3, 
	4, 5, 6, 7, 18, 8, 0, 1, 
	2, 3, 19, 5, 6, 7, 8, 1, 
	2, 3, 19, 5, 6, 7, 8, 0, 
	20, 2, 12, 4, 5, 6, 7, 8, 
	20, 2, 12, 4, 5, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 9, 6, 
	7, 8, 21, 1, 2, 3, 4, 5, 
	9, 6, 7, 8, 21, 0, 1, 2, 
	22, 3, 4, 5, 6, 7, 8, 1, 
	2, 22, 3, 4, 5, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 6, 23, 
	7, 8, 1, 2, 3, 4, 5, 6, 
	23, 7, 8, 0, 1, 2, 3, 4, 
	5, 24, 6, 7, 8, 1, 2, 3, 
	4, 5, 24, 6, 7, 8, 0, 1, 
	2, 25, 4, 5, 6, 7, 8, 1, 
	2, 25, 4, 5, 6, 7, 8, 0, 
	1, 2, 26, 3, 4, 5, 11, 7, 
	8, 1, 2, 26, 3, 4, 5, 11, 
	7, 8, 0, 27, 2, 3, 4, 5, 
	6, 7, 8, 27, 2, 3, 4, 5, 
	6, 7, 8, 0, 1, 2, 3, 4, 
	28, 9, 6, 7, 8, 1, 2, 3, 
	4, 28, 9, 6, 7, 8, 0, 14, 
	2, 3, 4, 5, 6, 29, 8, 14, 
	2, 3, 4, 5, 6, 29, 8, 0, 
	27, 2, 3, 4, 5, 6, 7, 30, 
	8, 27, 2, 3, 4, 5, 6, 7, 
	30, 8, 0, 1, 2, 3, 4, 5, 
	31, 6, 7, 8, 1, 2, 3, 4, 
	5, 31, 6, 7, 8, 0, 1, 2, 
	32, 3, 4, 5, 6, 7, 8, 1, 
	2, 32, 3, 4, 5, 6, 7, 8, 
	0, 1, 2, 33, 4, 5, 6, 7, 
	8, 1, 2, 33, 4, 5, 6, 7, 
	8, 0, 1, 2, 34, 4, 5, 35, 
	11, 7, 8, 1, 2, 34, 4, 5, 
	35, 11, 7, 8, 0, 36, 2, 37, 
	3, 19, 5, 38, 6, 7, 8, 36, 
	2, 37, 3, 19, 5, 38, 6, 7, 
	8, 0, 1, 2, 39, 3, 4, 5, 
	9, 6, 7, 8, 1, 2, 39, 3, 
	4, 5, 9, 6, 7, 8, 0, 1, 
	2, 3, 4, 5, 40, 6, 7, 8, 
	1, 2, 3, 4, 5, 40, 6, 7, 
	8, 0, 1, 2, 3, 4, 5, 41, 
	6, 7, 8, 1, 2, 3, 4, 5, 
	41, 6, 7, 8, 0, 1, 2, 42, 
	3, 4, 5, 6, 7, 8, 1, 2, 
	42, 3, 4, 5, 6, 7, 8, 0, 
	1, 2, 43, 3, 4, 5, 6, 7, 
	8, 1, 2, 43, 3, 4, 5, 6, 
	7, 8, 0, 1, 2, 44, 3, 4, 
	5, 6, 7, 8, 1, 2, 44, 3, 
	4, 5, 6, 7, 8, 0, 1, 2, 
	3, 4, 5, 45, 6, 7, 8, 1, 
	2, 3, 4, 5, 45, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 6, 7, 
	46, 1, 2, 3, 4, 5, 6, 7, 
	46, 0, 1, 2, 33, 4, 5, 6, 
	47, 8, 1, 2, 33, 4, 5, 6, 
	47, 8, 0, 48, 27, 2, 3, 4, 
	5, 6, 7, 8, 27, 2, 3, 4, 
	5, 6, 7, 8, 0, 1, 2, 3, 
	4, 5, 49, 50, 7, 8, 1, 2, 
	3, 4, 5, 49, 50, 7, 8, 0, 
	1, 2, 3, 4, 5, 6, 7, 51, 
	8, 1, 2, 3, 4, 5, 6, 7, 
	51, 8, 0, 1, 2, 52, 3, 19, 
	5, 6, 7, 8, 1, 2, 52, 3, 
	19, 5, 6, 7, 8, 0, 1, 2, 
	3, 4, 5, 53, 6, 7, 8, 1, 
	2, 3, 4, 5, 53, 6, 7, 8, 
	0, 1, 2, 3, 4, 5, 54, 6, 
	7, 8, 1, 2, 3, 4, 5, 54, 
	6, 7, 8, 0, 1, 2, 55, 3, 
	4, 5, 6, 7, 8, 1, 2, 55, 
	3, 4, 5, 6, 7, 8, 0, 1, 
	2, 3, 4, 5, 6, 56, 8, 1, 
	2, 3, 4, 5, 6, 56, 8, 0, 
	27, 2, 3, 4, 5, 6, 7, 57, 
	8, 27, 2, 3, 4, 5, 6, 7, 
	57, 8, 0, 58, 2, 3, 4, 5, 
	6, 7, 8, 58, 2, 3, 4, 5, 
	6, 7, 8, 0, 1, 2, 3, 4, 
	5, 9, 6, 7, 59, 8, 1, 2, 
	3, 4, 5, 9, 6, 7, 59, 8, 
	0, 1, 2, 60, 4, 5, 6, 7, 
	8, 1, 2, 60, 4, 5, 6, 7, 
	8, 0, 1, 2, 3, 4, 5, 61, 
	11, 7, 8, 1, 2, 3, 4, 5, 
	61, 11, 7, 8, 0, 1, 2, 3, 
	4, 5, 62, 6, 7, 8, 1, 2, 
	3, 4, 5, 62, 6, 7, 8, 0, 
	1, 2, 3, 4, 5, 63, 6, 7, 
	8, 1, 2, 3, 4, 5, 63, 6, 
	7, 8, 0, 1, 2, 3, 4, 5, 
	6, 64, 8, 1, 2, 3, 4, 5, 
	6, 64, 8, 0, 27, 2, 65, 3, 
	4, 5, 6, 7, 8, 27, 2, 65, 
	3, 4, 5, 6, 7, 8, 0, 1, 
	2, 3, 4, 5, 6, 7, 66, 8, 
	1, 2, 3, 4, 5, 6, 7, 66, 
	8, 0, 1, 2, 3, 4, 5, 6, 
	7, 8, 67, 1, 2, 3, 4, 5, 
	6, 7, 8, 67, 0, 68, 2, 12, 
	4, 5, 6, 7, 8, 68, 2, 12, 
	4, 5, 6, 7, 8, 0, 1, 2, 
	69, 3, 4, 5, 9, 6, 7, 8, 
	1, 2, 69, 3, 4, 5, 9, 6, 
	7, 8, 0, 1, 2, 3, 70, 4, 
	5, 6, 7, 8, 1, 2, 3, 70, 
	4, 5, 6, 7, 8, 0, 1, 71, 
	3, 4, 5, 6, 7, 8, 1, 71, 
	3, 4, 5, 6, 7, 8, 0, 1, 
	2, 72, 3, 10, 5, 6, 7, 8, 
	1, 2, 72, 3, 10, 5, 6, 7, 
	8, 0, 1, 2, 3, 4, 5, 6, 
	73, 7, 8, 1, 2, 3, 4, 5, 
	6, 73, 7, 8, 0, 1, 2, 3, 
	4, 5, 6, 74, 7, 8, 1, 2, 
	3, 4, 5, 6, 74, 7, 8, 0, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	75, 1, 2, 3, 4, 5, 6, 7, 
	8, 75, 0, 1, 2, 3, 4, 5, 
	6, 7, 8, 1, 2, 3, 4, 5, 
	6, 7, 8, 0, 1, 2, 3, 4, 
	5, 11, 7, 8, 1, 2, 3, 4, 
	5, 11, 7, 8, 0, 
}

var _scanner_trans_targs []byte = []byte{
	0, 1, 2, 3, 4, 6, 11, 19, 
	25, 14, 56, 27, 5, 54, 7, 8, 
	9, 10, 51, 12, 13, 44, 15, 16, 
	17, 18, 64, 20, 21, 22, 23, 24, 
	64, 26, 65, 33, 28, 29, 32, 64, 
	30, 31, 64, 64, 34, 35, 36, 37, 
	38, 39, 40, 64, 41, 42, 43, 64, 
	45, 46, 47, 48, 49, 50, 64, 52, 
	53, 64, 55, 64, 57, 58, 59, 60, 
	61, 62, 63, 64, 
}

var _scanner_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 28, 0, 0, 0, 0, 0, 
	34, 0, 13, 0, 0, 0, 0, 19, 
	0, 0, 25, 22, 0, 0, 0, 0, 
	0, 0, 0, 1, 0, 0, 0, 4, 
	0, 0, 0, 0, 0, 0, 16, 0, 
	0, 7, 0, 10, 0, 0, 0, 0, 
	0, 0, 0, 31, 
}

const scanner_start int = 0
const scanner_first_final int = 64
const scanner_error int = -1

const scanner_en_main int = 0


//line ./is_mobile.rl:21
  cs, p, pe := 0, 0, len(data)
  
//line src/ua/is_mobile.go:422
	{
	cs = scanner_start
	}

//line src/ua/is_mobile.go:427
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
_resume:
	_keys = int(_scanner_key_offsets[cs])
	_trans = int(_scanner_index_offsets[cs])

	_klen = int(_scanner_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _scanner_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_scanner_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _scanner_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _scanner_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_scanner_indicies[_trans])
	cs = int(_scanner_trans_targs[_trans])

	if _scanner_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_scanner_trans_actions[_trans])
	_nacts = uint(_scanner_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _scanner_actions[_acts-1] {
		case 0:
//line ./is_mobile.rl:23
os = "windows"; mobile = false
		case 1:
//line ./is_mobile.rl:24
os = "windows"
		case 2:
//line ./is_mobile.rl:25
os = "mac"; mobile = false
		case 3:
//line ./is_mobile.rl:26
os = "linux"; mobile = false
		case 4:
//line ./is_mobile.rl:27
os = "wii"
		case 5:
//line ./is_mobile.rl:28
os = "playstation"; mobile = false
		case 6:
//line ./is_mobile.rl:29
os = "ipad"
		case 7:
//line ./is_mobile.rl:30
os = "ipod"
		case 8:
//line ./is_mobile.rl:31
os = "iphone"
		case 9:
//line ./is_mobile.rl:32
os = "android"
		case 10:
//line ./is_mobile.rl:33
os = "blackberry"
		case 11:
//line ./is_mobile.rl:34
os = "samsung"
		case 12:
//line ./is_mobile.rl:35
 return mobile, os 
//line src/ua/is_mobile.go:542
		}
	}

_again:
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	}

//line ./is_mobile.rl:38

  return false, ""
}