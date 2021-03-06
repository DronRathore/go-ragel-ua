
//line ./ua_version.rl:1
/*
  Package provides function to extract the version and browser name from user agent
*/
package ua

// Function takes userAgent string as input and returns browser name and version
func UaVersion(agent string) (string, string) {
  var bytes []byte = make([]byte, len(agent))
  copy(bytes[:], agent)
  return _uaVersion(bytes)
}

func extractVersion(fc byte, version string, lastversion *string) string {
  var str = string(fc)
  if str == " " {
    *lastversion = version
    return ""
  } else {
    return version + str
  }
}

func extractChromeVersion(data []byte, pos int) string {
  // extract chrome's version
  var version string
  for pos < len(data) && data[pos] != ' ' {
    version += string(data[pos])
    pos = pos + 1
  }
  return version
}

func _uaVersion (data []byte) (string, string) {
  var browser string
  var version string
  var lastversion string

//line ./ua_version.rl:38

//line src/ua/ua_version.go:43
var _scanner_actions []byte = []byte{
	0, 1, 0, 1, 1, 1, 2, 1, 3, 
	1, 4, 1, 5, 1, 6, 1, 7, 
	1, 8, 1, 9, 1, 10, 1, 11, 
	1, 12, 1, 13, 1, 14, 
}

var _scanner_key_offsets []int16 = []int16{
	0, 20, 44, 66, 88, 112, 134, 154, 
	174, 196, 218, 240, 262, 284, 306, 326, 
	348, 368, 394, 416, 438, 459, 484, 506, 
	528, 550, 570, 594, 614, 637, 660, 680, 
	706, 728, 748, 770, 794, 815, 840, 860, 
	884, 908, 928, 950, 972, 994, 1016, 1038, 
	1058, 1082, 1102, 1125, 1150, 1174, 1194, 1214, 
	1236, 1258, 1278, 1302, 1323, 1348, 1372, 1392, 
	1415, 1438, 1461, 1482, 1507, 1527, 1551, 1571, 
	1593, 1614, 1639, 1661, 1681, 1707, 1727, 1747, 
	1773, 1795, 1815, 1835, 1859, 1882, 1907, 1927, 
	1947, 1968, 1993, 2014, 2039, 2061, 2081, 2101, 
	2124, 2146, 2168, 2188, 2210, 2236, 2257, 2281, 
	2301, 2325, 2346, 2366, 2391, 2411, 2433, 2457, 
	2479, 2500, 2525, 2545, 2568, 2593, 2618, 2641, 
	2666, 2691, 2716, 2739, 2764, 2789, 2814, 2839, 
	2864, 2889, 2910, 2934, 2959, 2984, 
}

var _scanner_trans_keys []byte = []byte{
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 72, 
	77, 79, 80, 82, 83, 84, 85, 86, 
	99, 101, 102, 104, 109, 111, 112, 114, 
	115, 116, 117, 118, 67, 68, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	100, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 71, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	103, 109, 111, 112, 115, 116, 117, 118, 
	67, 69, 70, 73, 76, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 105, 
	108, 109, 111, 112, 115, 116, 117, 118, 
	67, 69, 70, 77, 79, 80, 82, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 114, 115, 116, 117, 118, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 67, 69, 
	70, 77, 78, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 110, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 73, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 105, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 87, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 119, 
	67, 69, 70, 72, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 104, 109, 
	111, 112, 115, 116, 117, 118, 65, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 97, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	78, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 110, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 65, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 97, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	65, 67, 69, 70, 73, 76, 77, 79, 
	80, 83, 84, 85, 86, 97, 99, 101, 
	102, 105, 108, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	82, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 114, 115, 116, 117, 118, 
	67, 69, 70, 73, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 105, 109, 
	111, 112, 115, 116, 117, 118, 47, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 32, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	45, 46, 48, 57, 67, 69, 70, 77, 
	79, 80, 82, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 114, 115, 116, 
	117, 118, 67, 69, 70, 73, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	105, 109, 111, 112, 115, 116, 117, 118, 
	67, 68, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 100, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 68, 69, 70, 77, 78, 
	79, 80, 83, 84, 85, 86, 99, 100, 
	101, 102, 109, 110, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 47, 67, 
	69, 70, 77, 79, 80, 82, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	114, 115, 116, 117, 118, 46, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 48, 57, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	66, 67, 69, 70, 72, 77, 79, 80, 
	82, 83, 84, 85, 86, 98, 99, 101, 
	102, 104, 109, 111, 112, 114, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	82, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 114, 115, 116, 117, 118, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 87, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 119, 67, 69, 70, 72, 77, 79, 
	80, 82, 83, 84, 85, 86, 99, 101, 
	102, 104, 109, 111, 112, 114, 115, 116, 
	117, 118, 47, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 32, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 45, 46, 48, 57, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 68, 69, 70, 
	77, 79, 80, 82, 83, 84, 85, 86, 
	99, 100, 101, 102, 109, 111, 112, 114, 
	115, 116, 117, 118, 67, 69, 70, 72, 
	73, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 104, 105, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	67, 69, 70, 72, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 104, 109, 
	111, 112, 115, 116, 117, 118, 65, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 97, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	78, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 110, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 89, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 121, 
	67, 69, 70, 77, 79, 80, 82, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 114, 115, 116, 117, 118, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 74, 77, 78, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 106, 109, 110, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 47, 65, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 97, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 32, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 45, 46, 48, 57, 65, 67, 
	68, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 97, 99, 100, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 67, 69, 
	70, 77, 78, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 110, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 75, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 107, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 67, 68, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 89, 99, 100, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 121, 47, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 32, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	45, 46, 48, 57, 65, 67, 69, 70, 
	73, 77, 79, 80, 83, 84, 85, 86, 
	97, 99, 101, 102, 105, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	32, 67, 68, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 100, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 46, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 48, 57, 46, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 48, 57, 47, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 32, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 45, 
	46, 48, 57, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 65, 
	67, 69, 70, 73, 77, 79, 80, 83, 
	84, 85, 86, 97, 99, 101, 102, 105, 
	109, 111, 112, 115, 116, 117, 118, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 67, 69, 70, 77, 78, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 110, 111, 112, 115, 116, 117, 
	118, 47, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 32, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 45, 46, 48, 57, 65, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 97, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 65, 67, 69, 70, 73, 76, 77, 
	79, 80, 83, 84, 85, 86, 97, 99, 
	101, 102, 105, 108, 109, 111, 112, 115, 
	116, 117, 118, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 67, 69, 70, 72, 75, 
	77, 79, 80, 82, 83, 84, 85, 86, 
	99, 101, 102, 104, 107, 109, 111, 112, 
	114, 115, 116, 117, 118, 67, 69, 70, 
	77, 79, 80, 82, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 114, 115, 
	116, 117, 118, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 67, 69, 70, 73, 77, 
	78, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 105, 109, 110, 111, 112, 115, 
	116, 117, 118, 47, 67, 68, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	100, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 32, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 45, 
	46, 48, 57, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 47, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	32, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 45, 46, 48, 
	57, 47, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 32, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 45, 46, 48, 57, 67, 
	69, 70, 73, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 105, 109, 111, 
	112, 115, 116, 117, 118, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 47, 65, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 97, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 67, 69, 70, 77, 
	79, 80, 82, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 114, 115, 116, 
	117, 118, 67, 69, 70, 73, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	105, 109, 111, 112, 115, 116, 117, 118, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 65, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 97, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 65, 67, 68, 69, 70, 77, 
	79, 80, 82, 83, 84, 85, 86, 97, 
	99, 100, 101, 102, 109, 111, 112, 114, 
	115, 116, 117, 118, 47, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 32, 46, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 48, 
	57, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 66, 67, 68, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 98, 99, 100, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 47, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 32, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 45, 46, 48, 57, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 67, 68, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 100, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 67, 69, 70, 73, 76, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	105, 108, 109, 111, 112, 115, 116, 117, 
	118, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 88, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 120, 47, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 32, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 45, 46, 48, 57, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 47, 67, 68, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 100, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	32, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 45, 46, 48, 
	57, 32, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 45, 46, 
	48, 57, 46, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 48, 
	57, 32, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 45, 46, 
	48, 57, 32, 67, 69, 70, 77, 79, 
	80, 83, 84, 85, 86, 99, 101, 102, 
	109, 111, 112, 115, 116, 117, 118, 45, 
	46, 48, 57, 32, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	45, 46, 48, 57, 46, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 48, 57, 32, 67, 69, 70, 77, 
	79, 80, 83, 84, 85, 86, 99, 101, 
	102, 109, 111, 112, 115, 116, 117, 118, 
	45, 46, 48, 57, 32, 67, 69, 70, 
	77, 79, 80, 83, 84, 85, 86, 99, 
	101, 102, 109, 111, 112, 115, 116, 117, 
	118, 45, 46, 48, 57, 32, 67, 69, 
	70, 77, 79, 80, 83, 84, 85, 86, 
	99, 101, 102, 109, 111, 112, 115, 116, 
	117, 118, 45, 46, 48, 57, 32, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 45, 46, 48, 57, 32, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 45, 46, 48, 57, 
	32, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 45, 46, 48, 
	57, 47, 67, 69, 70, 77, 79, 80, 
	83, 84, 85, 86, 99, 101, 102, 109, 
	111, 112, 115, 116, 117, 118, 32, 46, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 48, 57, 32, 67, 
	69, 70, 77, 79, 80, 83, 84, 85, 
	86, 99, 101, 102, 109, 111, 112, 115, 
	116, 117, 118, 45, 46, 48, 57, 32, 
	67, 69, 70, 77, 79, 80, 83, 84, 
	85, 86, 99, 101, 102, 109, 111, 112, 
	115, 116, 117, 118, 45, 46, 48, 57, 
	32, 67, 69, 70, 77, 79, 80, 83, 
	84, 85, 86, 99, 101, 102, 109, 111, 
	112, 115, 116, 117, 118, 45, 46, 48, 
	57, 
}

var _scanner_single_lengths []byte = []byte{
	20, 24, 22, 22, 24, 22, 20, 20, 
	22, 22, 22, 22, 22, 22, 20, 22, 
	20, 26, 22, 22, 21, 21, 22, 22, 
	22, 20, 24, 20, 23, 21, 20, 26, 
	22, 20, 22, 24, 21, 21, 20, 24, 
	24, 20, 22, 22, 22, 22, 22, 20, 
	24, 20, 23, 21, 24, 20, 20, 22, 
	22, 20, 24, 21, 21, 24, 20, 23, 
	21, 21, 21, 21, 20, 24, 20, 22, 
	21, 21, 22, 20, 26, 20, 20, 26, 
	22, 20, 20, 24, 23, 21, 20, 20, 
	21, 21, 21, 21, 22, 20, 20, 23, 
	22, 22, 20, 22, 26, 21, 22, 20, 
	24, 21, 20, 21, 20, 22, 24, 22, 
	21, 21, 20, 23, 21, 21, 21, 21, 
	21, 21, 21, 21, 21, 21, 21, 21, 
	21, 21, 22, 21, 21, 21, 
}

var _scanner_range_lengths []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 2, 0, 0, 
	0, 0, 0, 0, 0, 1, 0, 0, 
	0, 0, 0, 0, 0, 2, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 2, 0, 0, 0, 0, 
	0, 0, 0, 0, 2, 0, 0, 0, 
	1, 1, 0, 2, 0, 0, 0, 0, 
	0, 2, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 2, 0, 0, 
	0, 2, 0, 2, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 1, 0, 
	0, 0, 0, 2, 0, 0, 0, 0, 
	0, 2, 0, 0, 2, 2, 1, 2, 
	2, 2, 1, 2, 2, 2, 2, 2, 
	2, 0, 1, 2, 2, 2, 
}

var _scanner_index_offsets []int16 = []int16{
	0, 21, 46, 69, 92, 117, 140, 161, 
	182, 205, 228, 251, 274, 297, 320, 341, 
	364, 385, 412, 435, 458, 480, 504, 527, 
	550, 573, 594, 619, 640, 664, 687, 708, 
	735, 758, 779, 802, 827, 849, 873, 894, 
	919, 944, 965, 988, 1011, 1034, 1057, 1080, 
	1101, 1126, 1147, 1171, 1195, 1220, 1241, 1262, 
	1285, 1308, 1329, 1354, 1376, 1400, 1425, 1446, 
	1470, 1493, 1516, 1538, 1562, 1583, 1608, 1629, 
	1652, 1674, 1698, 1721, 1742, 1769, 1790, 1811, 
	1838, 1861, 1882, 1903, 1928, 1952, 1976, 1997, 
	2018, 2040, 2064, 2086, 2110, 2133, 2154, 2175, 
	2199, 2222, 2245, 2266, 2289, 2316, 2338, 2362, 
	2383, 2408, 2430, 2451, 2475, 2496, 2519, 2544, 
	2567, 2589, 2613, 2634, 2658, 2682, 2706, 2729, 
	2753, 2777, 2801, 2824, 2848, 2872, 2896, 2920, 
	2944, 2968, 2990, 3014, 3038, 3062, 
}

var _scanner_indicies []byte = []byte{
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	11, 4, 5, 6, 12, 7, 8, 9, 
	10, 1, 2, 3, 11, 4, 5, 6, 
	12, 7, 8, 9, 10, 0, 1, 13, 
	2, 3, 4, 5, 14, 7, 8, 9, 
	10, 1, 13, 2, 3, 4, 5, 14, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	15, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 15, 4, 5, 6, 7, 
	8, 9, 10, 0, 1, 2, 3, 16, 
	17, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 16, 17, 4, 5, 6, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	4, 5, 6, 18, 7, 8, 9, 10, 
	1, 2, 3, 4, 5, 6, 18, 7, 
	8, 9, 10, 0, 1, 2, 3, 4, 
	5, 6, 19, 8, 9, 10, 1, 2, 
	3, 4, 5, 6, 19, 8, 9, 10, 
	0, 1, 2, 3, 20, 5, 21, 7, 
	8, 9, 10, 1, 2, 3, 20, 5, 
	21, 7, 8, 9, 10, 0, 1, 2, 
	3, 4, 22, 5, 6, 19, 8, 9, 
	10, 1, 2, 3, 4, 22, 5, 6, 
	19, 8, 9, 10, 0, 1, 2, 3, 
	23, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 23, 4, 5, 6, 7, 
	8, 9, 10, 0, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 24, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 24, 0, 1, 2, 3, 25, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 25, 4, 5, 6, 7, 8, 9, 
	10, 0, 26, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 26, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	0, 1, 2, 3, 4, 27, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	27, 5, 6, 7, 8, 9, 10, 0, 
	1, 2, 3, 4, 5, 6, 7, 28, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	7, 28, 9, 10, 0, 29, 1, 30, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	29, 1, 30, 3, 4, 5, 6, 7, 
	8, 9, 10, 0, 1, 2, 31, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	31, 4, 5, 6, 7, 8, 9, 10, 
	0, 32, 1, 2, 3, 16, 17, 4, 
	5, 6, 7, 8, 9, 10, 32, 1, 
	2, 3, 16, 17, 4, 5, 6, 7, 
	8, 9, 10, 0, 1, 2, 3, 4, 
	5, 6, 33, 7, 8, 9, 10, 1, 
	2, 3, 4, 5, 6, 33, 7, 8, 
	9, 10, 0, 1, 2, 3, 34, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 34, 4, 5, 6, 7, 8, 9, 
	10, 0, 35, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 0, 
	36, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 36, 36, 0, 
	1, 2, 3, 4, 5, 6, 37, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 37, 7, 8, 9, 10, 0, 1, 
	2, 3, 38, 4, 5, 6, 7, 8, 
	9, 10, 1, 2, 3, 38, 4, 5, 
	6, 7, 8, 9, 10, 0, 1, 39, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 39, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 0, 1, 40, 3, 
	4, 5, 6, 7, 8, 9, 10, 1, 
	40, 3, 4, 5, 6, 7, 8, 9, 
	10, 0, 1, 13, 2, 3, 4, 41, 
	5, 14, 7, 8, 9, 10, 1, 13, 
	2, 3, 4, 41, 5, 14, 7, 8, 
	9, 10, 0, 1, 2, 3, 4, 5, 
	6, 7, 42, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 42, 9, 10, 0, 
	43, 1, 2, 3, 4, 5, 6, 37, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 37, 7, 8, 9, 10, 0, 
	44, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 44, 0, 45, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 45, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 0, 46, 1, 2, 3, 
	11, 4, 5, 6, 12, 7, 8, 9, 
	10, 46, 1, 2, 3, 11, 4, 5, 
	6, 12, 7, 8, 9, 10, 0, 1, 
	2, 3, 4, 5, 6, 47, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	47, 7, 8, 9, 10, 0, 1, 2, 
	3, 4, 48, 6, 7, 8, 9, 10, 
	1, 2, 3, 4, 48, 6, 7, 8, 
	9, 10, 0, 1, 2, 3, 20, 5, 
	21, 7, 8, 9, 10, 49, 1, 2, 
	3, 20, 5, 21, 7, 8, 9, 10, 
	49, 0, 1, 2, 3, 25, 4, 5, 
	6, 50, 7, 8, 9, 10, 1, 2, 
	3, 25, 4, 5, 6, 50, 7, 8, 
	9, 10, 0, 51, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	0, 52, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 52, 52, 
	0, 1, 53, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 53, 3, 4, 5, 
	6, 7, 8, 9, 10, 0, 1, 13, 
	2, 3, 4, 5, 14, 54, 7, 8, 
	9, 10, 1, 13, 2, 3, 4, 5, 
	14, 54, 7, 8, 9, 10, 0, 1, 
	2, 3, 25, 55, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 25, 55, 
	4, 5, 6, 7, 8, 9, 10, 0, 
	1, 2, 3, 4, 5, 56, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 56, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	57, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 57, 4, 5, 6, 7, 
	8, 9, 10, 0, 58, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 58, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 0, 1, 2, 3, 4, 59, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 4, 59, 5, 6, 7, 8, 9, 
	10, 0, 1, 2, 3, 4, 5, 6, 
	7, 28, 9, 10, 60, 1, 2, 3, 
	4, 5, 6, 7, 28, 9, 10, 60, 
	0, 1, 2, 3, 4, 61, 6, 37, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	61, 6, 37, 7, 8, 9, 10, 0, 
	1, 2, 3, 62, 5, 21, 7, 8, 
	9, 10, 1, 2, 3, 62, 5, 21, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	63, 4, 22, 5, 6, 19, 8, 9, 
	10, 1, 2, 3, 63, 4, 22, 5, 
	6, 19, 8, 9, 10, 0, 1, 2, 
	3, 4, 5, 6, 64, 8, 9, 10, 
	1, 2, 3, 4, 5, 6, 64, 8, 
	9, 10, 0, 65, 29, 1, 30, 3, 
	4, 5, 6, 7, 8, 9, 10, 29, 
	1, 30, 3, 4, 5, 6, 7, 8, 
	9, 10, 0, 66, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	66, 66, 0, 67, 1, 13, 2, 3, 
	4, 5, 14, 7, 8, 9, 10, 67, 
	1, 13, 2, 3, 4, 5, 14, 7, 
	8, 9, 10, 0, 1, 2, 3, 68, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 68, 5, 6, 7, 8, 9, 10, 
	0, 1, 2, 3, 4, 69, 6, 19, 
	8, 9, 10, 1, 2, 3, 4, 69, 
	6, 19, 8, 9, 10, 0, 1, 2, 
	3, 20, 70, 5, 21, 7, 8, 9, 
	10, 1, 2, 3, 20, 70, 5, 21, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	71, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 71, 4, 5, 6, 7, 
	8, 9, 10, 0, 1, 72, 3, 4, 
	5, 6, 7, 8, 9, 10, 1, 72, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	0, 1, 13, 2, 3, 4, 5, 14, 
	7, 8, 9, 10, 73, 1, 13, 2, 
	3, 4, 5, 14, 7, 8, 9, 10, 
	73, 0, 74, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 0, 
	75, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 75, 75, 0, 
	29, 1, 30, 3, 76, 4, 5, 6, 
	7, 8, 9, 10, 29, 1, 30, 3, 
	76, 4, 5, 6, 7, 8, 9, 10, 
	0, 1, 77, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 77, 3, 4, 5, 
	6, 7, 8, 9, 10, 0, 78, 1, 
	13, 2, 3, 4, 5, 14, 7, 8, 
	9, 10, 1, 13, 2, 3, 4, 5, 
	14, 7, 8, 9, 10, 0, 79, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 79, 0, 79, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 80, 0, 81, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 0, 82, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 82, 
	82, 0, 1, 2, 3, 4, 5, 6, 
	83, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 83, 8, 9, 10, 0, 29, 
	1, 30, 3, 84, 4, 5, 6, 7, 
	8, 9, 10, 29, 1, 30, 3, 84, 
	4, 5, 6, 7, 8, 9, 10, 0, 
	1, 2, 3, 4, 85, 6, 7, 8, 
	9, 10, 1, 2, 3, 4, 85, 6, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	20, 86, 5, 21, 7, 8, 9, 10, 
	1, 2, 3, 20, 86, 5, 21, 7, 
	8, 9, 10, 0, 87, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 0, 88, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 88, 
	88, 0, 29, 1, 30, 3, 4, 5, 
	6, 7, 8, 9, 10, 89, 1, 30, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	0, 1, 2, 31, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 90, 4, 5, 
	6, 7, 8, 9, 10, 0, 32, 1, 
	2, 3, 16, 17, 4, 5, 6, 7, 
	8, 9, 10, 91, 1, 2, 3, 16, 
	17, 4, 5, 6, 7, 8, 9, 10, 
	0, 1, 2, 3, 4, 92, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 92, 
	6, 7, 8, 9, 10, 0, 93, 2, 
	3, 20, 5, 21, 7, 8, 9, 10, 
	93, 2, 3, 20, 5, 21, 7, 8, 
	9, 10, 0, 1, 2, 3, 11, 94, 
	4, 5, 6, 12, 7, 8, 9, 10, 
	1, 2, 3, 11, 94, 4, 5, 6, 
	12, 7, 8, 9, 10, 0, 1, 2, 
	3, 4, 5, 6, 95, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 95, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	4, 96, 6, 7, 8, 9, 10, 1, 
	2, 3, 4, 96, 6, 7, 8, 9, 
	10, 0, 1, 2, 3, 97, 5, 21, 
	7, 8, 9, 10, 1, 2, 3, 97, 
	5, 21, 7, 8, 9, 10, 0, 1, 
	98, 3, 99, 4, 22, 5, 6, 19, 
	8, 9, 10, 1, 98, 3, 99, 4, 
	22, 5, 6, 19, 8, 9, 10, 0, 
	100, 1, 13, 2, 3, 4, 5, 14, 
	7, 8, 9, 10, 1, 13, 2, 3, 
	4, 5, 14, 7, 8, 9, 10, 0, 
	101, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 101, 101, 0, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	102, 10, 1, 2, 3, 4, 5, 6, 
	7, 8, 102, 10, 0, 45, 2, 3, 
	103, 5, 6, 7, 8, 9, 10, 45, 
	2, 3, 103, 5, 6, 7, 8, 9, 
	10, 0, 104, 1, 2, 3, 4, 5, 
	6, 19, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 19, 8, 9, 10, 0, 
	105, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 105, 105, 0, 
	106, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 0, 107, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 107, 107, 0, 1, 2, 
	3, 108, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 108, 4, 5, 6, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	4, 109, 6, 7, 8, 9, 10, 1, 
	2, 3, 4, 109, 6, 7, 8, 9, 
	10, 0, 1, 2, 3, 20, 5, 21, 
	110, 8, 9, 10, 1, 2, 3, 20, 
	5, 21, 110, 8, 9, 10, 0, 104, 
	29, 1, 30, 3, 4, 5, 6, 7, 
	8, 9, 10, 29, 1, 30, 3, 4, 
	5, 6, 7, 8, 9, 10, 0, 1, 
	2, 3, 4, 5, 6, 33, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	111, 7, 8, 9, 10, 0, 1, 2, 
	3, 34, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 112, 4, 5, 6, 
	7, 8, 9, 10, 0, 1, 2, 3, 
	4, 5, 6, 113, 8, 9, 10, 1, 
	2, 3, 4, 5, 6, 113, 8, 9, 
	10, 0, 29, 1, 114, 3, 4, 5, 
	6, 7, 8, 9, 10, 29, 1, 114, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	0, 67, 1, 13, 2, 3, 4, 5, 
	14, 115, 7, 8, 9, 10, 67, 1, 
	13, 2, 3, 4, 5, 14, 115, 7, 
	8, 9, 10, 0, 116, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 0, 117, 117, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	117, 0, 1, 118, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 118, 3, 4, 
	5, 6, 7, 8, 9, 10, 0, 119, 
	1, 13, 2, 3, 4, 5, 14, 7, 
	8, 9, 10, 119, 1, 13, 2, 3, 
	4, 5, 14, 7, 8, 9, 10, 0, 
	120, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 0, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 121, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 121, 0, 122, 1, 53, 3, 4, 
	5, 6, 7, 8, 9, 10, 1, 53, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	122, 122, 0, 1, 123, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 123, 3, 
	4, 5, 6, 7, 8, 9, 10, 0, 
	1, 13, 2, 124, 4, 5, 14, 7, 
	8, 9, 10, 1, 13, 2, 124, 4, 
	5, 14, 7, 8, 9, 10, 0, 1, 
	2, 3, 16, 17, 4, 125, 6, 7, 
	8, 9, 10, 1, 2, 3, 16, 17, 
	4, 125, 6, 7, 8, 9, 10, 0, 
	1, 2, 3, 20, 5, 21, 7, 8, 
	9, 10, 126, 1, 2, 3, 20, 5, 
	21, 7, 8, 9, 10, 126, 0, 127, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 0, 128, 1, 2, 
	3, 4, 5, 6, 7, 8, 9, 10, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 128, 128, 0, 1, 129, 3, 
	4, 5, 6, 7, 8, 9, 10, 1, 
	129, 3, 4, 5, 6, 7, 8, 9, 
	10, 0, 130, 1, 13, 2, 3, 4, 
	5, 14, 7, 8, 9, 10, 1, 13, 
	2, 3, 4, 5, 14, 7, 8, 9, 
	10, 0, 131, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 131, 
	131, 0, 36, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 36, 
	36, 0, 44, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 1, 2, 3, 
	4, 5, 6, 7, 8, 9, 10, 44, 
	0, 52, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 52, 52, 
	0, 66, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 66, 66, 
	0, 75, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 75, 75, 
	0, 79, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 1, 2, 3, 4, 
	5, 6, 7, 8, 9, 10, 80, 0, 
	82, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 82, 82, 0, 
	132, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 88, 88, 0, 
	132, 1, 2, 3, 4, 5, 6, 133, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 88, 88, 0, 
	101, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 101, 101, 0, 
	105, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 105, 105, 0, 
	107, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 107, 107, 0, 
	35, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 1, 2, 3, 4, 5, 
	6, 7, 8, 9, 10, 0, 117, 117, 
	1, 2, 3, 4, 5, 6, 7, 8, 
	9, 10, 1, 2, 3, 4, 5, 6, 
	7, 8, 9, 10, 117, 0, 122, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 122, 122, 0, 128, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 128, 128, 0, 131, 1, 
	2, 3, 4, 5, 6, 7, 8, 9, 
	10, 1, 2, 3, 4, 5, 6, 7, 
	8, 9, 10, 131, 131, 0, 
}

var _scanner_trans_targs []byte = []byte{
	0, 1, 2, 4, 6, 7, 11, 15, 
	22, 30, 38, 80, 92, 3, 40, 114, 
	5, 77, 108, 61, 8, 35, 9, 10, 
	103, 12, 13, 14, 46, 16, 52, 17, 
	18, 19, 20, 21, 117, 23, 24, 25, 
	26, 27, 28, 29, 118, 31, 32, 33, 
	34, 98, 36, 37, 119, 39, 68, 41, 
	42, 43, 44, 45, 66, 47, 48, 49, 
	50, 51, 120, 53, 54, 55, 56, 57, 
	58, 59, 60, 121, 62, 63, 64, 65, 
	122, 67, 123, 69, 70, 71, 72, 73, 
	124, 75, 76, 96, 78, 79, 90, 81, 
	82, 83, 84, 86, 85, 126, 87, 88, 
	89, 127, 91, 128, 93, 94, 95, 97, 
	129, 99, 100, 101, 102, 130, 104, 105, 
	106, 107, 131, 109, 110, 111, 112, 113, 
	132, 115, 116, 133, 125, 74, 
}

var _scanner_trans_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 11, 0, 0, 0, 
	0, 0, 0, 0, 7, 0, 0, 0, 
	0, 0, 0, 0, 9, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 23, 0, 0, 0, 0, 0, 
	0, 0, 0, 21, 0, 0, 0, 0, 
	5, 0, 27, 0, 0, 0, 0, 0, 
	9, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 13, 0, 0, 
	0, 15, 0, 25, 0, 0, 0, 0, 
	11, 0, 0, 0, 0, 17, 0, 0, 
	0, 0, 19, 0, 0, 0, 0, 0, 
	3, 0, 0, 1, 9, 0, 
}

var _scanner_eof_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 29, 29, 29, 
	29, 29, 29, 29, 29, 29, 29, 29, 
	29, 29, 29, 29, 29, 29, 
}

const scanner_start int = 0
const scanner_first_final int = 117
const scanner_error int = -1

const scanner_en_main int = 0


//line ./ua_version.rl:39
  cs, p, pe,eof := 0, 0, len(data), len(data)
  
//line src/ua/ua_version.go:970
	{
	cs = scanner_start
	}

//line src/ua/ua_version.go:975
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
//line ./ua_version.rl:41
browser = "edge"; version = extractVersion(data[p], version, &lastversion);
		case 1:
//line ./ua_version.rl:42
browser = "firefox"; version = extractVersion(data[p], version, &lastversion);
		case 2:
//line ./ua_version.rl:43
browser = "explorer"; version = extractVersion(data[p], version, &lastversion);
		case 3:
//line ./ua_version.rl:44
browser = "trident"; version = extractVersion(data[p], version, &lastversion);
		case 4:
//line ./ua_version.rl:45
browser = "opera"; version = extractVersion(data[p], version, &lastversion);
		case 5:
//line ./ua_version.rl:46
browser = "safari"; version = extractVersion(data[p], version, &lastversion);
		case 6:
//line ./ua_version.rl:47
browser = "chrome"; version = extractChromeVersion(data, p);goto end_marker;
		case 7:
//line ./ua_version.rl:48
browser = "chromium";version = extractChromeVersion(data, p);goto end_marker;
		case 8:
//line ./ua_version.rl:49
browser = "uc"; version = extractVersion(data[p], version, &lastversion);
		case 9:
//line ./ua_version.rl:50
browser = "omniweb"; version = extractVersion(data[p], version, &lastversion);
		case 10:
//line ./ua_version.rl:51
browser = "seamonkey"; version = extractVersion(data[p], version, &lastversion);
		case 11:
//line ./ua_version.rl:52
browser = "phantomjs"; goto end_marker;
		case 12:
//line ./ua_version.rl:53
browser = "flock"; version = extractVersion(data[p], version, &lastversion);
		case 13:
//line ./ua_version.rl:54
browser = "epiphany"; version = extractVersion(data[p], version, &lastversion);
//line src/ua/ua_version.go:1093
		}
	}

_again:
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		__acts := _scanner_eof_actions[cs]
		__nacts := uint(_scanner_actions[__acts]); __acts++
		for ; __nacts > 0; __nacts-- {
			__acts++
			switch _scanner_actions[__acts-1] {
			case 14:
//line ./ua_version.rl:69
 goto end_marker 
//line src/ua/ua_version.go:1112
			}
		}
	}

	}

//line ./ua_version.rl:72

  end_marker:
    if version == "" && lastversion != "" {
      version = lastversion
    }
    return browser, version
}