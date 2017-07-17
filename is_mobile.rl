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
%% machine scanner;
%% write data;
  cs, p, pe,eof := 0, 0, len(data), len(data)
  %%{
    action windows {os = "windows"; mobile = false}
    action winPhone {os = "windows"}
    action mac {os = "mac"; mobile = false}
    action linux {os = "linux"; mobile = false;}
    action wii {os = "wii"}
    action playstation {os = "playstation"; mobile = false}
    action ipad {os = "ipad"; mobile = true; goto end;}
    action ipod {os = "ipod"; mobile = true; goto end;}
    action iphone {os = "iphone"; mobile = true; goto end;}
    action android {os = "android"; mobile = true; goto end;}
    action blackberry {os = "blackberry"; mobile = true; goto end;}
    action samsung {os = "samsung"; mobile = true; goto end;}
    main := any* graph* space* punct* ('windows nt'i @windows|'windows phone'i @winPhone|'macintosh'i @mac|'linux'i @linux|'wii'i @wii|'playstation'i @playstation|'ipad'i @ipad|'ipod'i @ipod|'iphone'i @iphone|'android'i @android|'blackberry'i @blackberry|'samsung'i @samsung)+ %{ return mobile, os };
    write init;
    write exec;
  }%%
  return false, ""
  end:
    return true, os
}
