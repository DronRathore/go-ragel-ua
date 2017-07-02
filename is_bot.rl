/*
  To checks whether the user agent is bot or not
*/
package ua

// Return a boolean flag
func IsBot(agent string) bool {
  var bytes []byte = make([]byte, len(agent))
  copy(bytes[:], agent)
  return _isBot(bytes)
}
/*
  Ragel configuration to build an optimised regex
  using a Detriminstic Finite Automata machine
*/
func _isBot(data []byte) bool {

%% machine scanner;
%% write data;
  cs, p, pe := 0, 0, len(data)
  %%{
    main := any* graph* space* punct* ('\\+https:\\/\\/developers.google.com\\/\\+\\/web\\/snippet\\/'|'adsbot'|'baiduspider'|'bingbot'|'cloudflare'|'crawler'|'embedly'|'exabot'|'facebookexternalhit'|'google'|'googlebot'|'gurujibot'|'heritrix'|'linkedinbot'|'msnbot'|'pingdom'|'rtlnieuws'|'slackbot'|'slurp'|'spbot'|'telegrambot'|'test certificate'|'testing'|'tiabot'|'tumblr '|'twitterbot'|'yandexbot'|'apex'|'applebot'|'duckduckbot'|'facebot'|'flipboard'|'gsa-crawler'|'ia_archiver'|'pinterest'|'skypeuripreview'|'odklbot'|'archive.org_bot'|'ltx71'|'guzzlehttp'|'vkshare'|'discordbot'|'whatsapp')+ @{ return true };
    write init;
    write exec;
  }%%
  return false
}
