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
%% machine scanner;
%% write data;
  cs, p, pe,eof := 0, 0, len(data), len(data)
  %%{
    action setEdge {browser = "edge"; version = extractVersion(fc, version, &lastversion);}
    action setFirefox {browser = "firefox"; version = extractVersion(fc, version, &lastversion);}
    action setMsie {browser = "explorer"; version = extractVersion(fc, version, &lastversion);}
    action setTrident {browser = "trident"; version = extractVersion(fc, version, &lastversion);}
    action setOpera {browser = "opera"; version = extractVersion(fc, version, &lastversion);}
    action setSafari {browser = "safari"; version = extractVersion(fc, version, &lastversion);}
    action setChrome {browser = "chrome"; version = extractChromeVersion(data, fpc);goto end_marker;}
    action setChromium {browser = "chromium";version = extractChromeVersion(data, fpc);goto end_marker;}
    action setUC {browser = "uc"; version = extractVersion(fc, version, &lastversion);}
    action setOmni {browser = "omniweb"; version = extractVersion(fc, version, &lastversion);}
    action setSeaMonkey {browser = "seamonkey"; version = extractVersion(fc, version, &lastversion);}
    action setPhantom {browser = "phantomjs"; goto end_marker;}
    action setFlock {browser = "flock"; version = extractVersion(fc, version, &lastversion);}
    action setEpiphany {browser = "epiphany"; version = extractVersion(fc, version, &lastversion);}

    safari = 'Safari/'i.([0-9 \.\-]+) @setSafari|('version/'i.([0-9 \.\-]+ ' Safari' @setSafari));
    firefox = 'firefox/'i.([0-9 \.\-]+ @setFirefox);
    ie = 'msie 'i.([0-9\.]+ [0-9] @setMsie)|'trident/'i.[0-9\.]+ @setTrident;
    chrome = 'chrome/'i.([0-9 \.\-]+ @setChrome)|('chromium/'i|'crios/'i)+.([0-9 \.\-]+ @setChromium);
    edge = 'Edge/'i.([0-9 \.\-]+ @setEdge);
    uc = 'UCBrowser/'i.[0-9 \.]+ @setUC;
    opera = 'version/'i.[0-9 \.\-]+ @setOpera|'OPR/'i.[0-9 \.\-]+ @setOpera;
    omni = 'omniweb/v'i.([0-9 \.\-]+ @setOmni);
    seamonkey = 'seamonkey/'i.([0-9 \.\-]+ @setSeaMonkey);
    phantom = 'phantomjs/'i.([0-9 \.\-]+ @setPhantom);
    flock = 'flock/'i.([0-9 \.\-]+ @setFlock);
    epiphany = 'epiphany/'i.([0-9 \.\-]+ @setEpiphany);

    main := any* graph* space* punct* (safari|firefox|ie|chrome|edge|uc|opera|omni|seamonkey|phantom|flock|epiphany)+ %{ goto end_marker };
    write init;
    write exec;
  }%%
  end_marker:
    if version == "" && lastversion != "" {
      version = lastversion
    }
    return browser, version
}