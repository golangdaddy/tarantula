package g

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  HTML Audio / Video Methods - http://www.w3schools.com/tags/ref_av_dom.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) AddTextTrack(eval string) *ELEMENT { return ele.Attr("addTextTrack()", eval) }

func (ele *ELEMENT) CanPlayType(eval string) *ELEMENT { return ele.Attr("canPlayType()", eval) }

func (ele *ELEMENT) Load(eval string) *ELEMENT { return ele.Attr("load()", eval) }


// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  HTML Audio / Video Propertios - http://www.w3schools.com/tags/ref_av_dom.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) AudioTracks(eval string) *ELEMENT { return ele.Attr("audioTracks", eval) }

func (ele *ELEMENT) Autoplay(eval string) *ELEMENT { return ele.Attr("autoplay", eval) }

func (ele *ELEMENT) Buffered(eval string) *ELEMENT { return ele.Attr("buffered", eval) }

func (ele *ELEMENT) Controller(eval string) *ELEMENT { return ele.Attr("controller", eval) }

func (ele *ELEMENT) Controls(eval string) *ELEMENT { return ele.Attr("controls", eval) }

func (ele *ELEMENT) CrossOrigin(eval string) *ELEMENT { return ele.Attr("crossOrigin", eval) }

func (ele *ELEMENT) CurrentSrc(eval string) *ELEMENT { return ele.Attr("currentSrc", eval) }

func (ele *ELEMENT) CurrentTime(eval string) *ELEMENT { return ele.Attr("currentTime", eval) }

func (ele *ELEMENT) DefaultMuted(eval string) *ELEMENT { return ele.Attr("defaultMuted", eval) }

func (ele *ELEMENT) DefaultPlaybackRate(eval string) *ELEMENT { return ele.Attr("defaultPlaybackRate", eval) }

func (ele *ELEMENT) Duration(eval string) *ELEMENT { return ele.Attr("duration", eval) }

func (ele *ELEMENT) Loop(eval string) *ELEMENT { return ele.Attr("loop", eval) }

func (ele *ELEMENT) MediaGroup(eval string) *ELEMENT { return ele.Attr("mediaGroup", eval) }

func (ele *ELEMENT) Muted(eval string) *ELEMENT { return ele.Attr("muted", eval) }

func (ele *ELEMENT) NetworkState(eval string) *ELEMENT { return ele.Attr("networkState", eval) }

func (ele *ELEMENT) Paused(eval string) *ELEMENT { return ele.Attr("paused", eval) }

func (ele *ELEMENT) PlaybackRate(eval string) *ELEMENT { return ele.Attr("playbackRate", eval) }

func (ele *ELEMENT) Played(eval string) *ELEMENT { return ele.Attr("played", eval) }

func (ele *ELEMENT) Preload(eval string) *ELEMENT { return ele.Attr("preload", eval) }

func (ele *ELEMENT) ReadyState(eval string) *ELEMENT { return ele.Attr("readyState", eval) }

func (ele *ELEMENT) Seekable(eval string) *ELEMENT { return ele.Attr("seekable", eval) }

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) Src(eval string) *ELEMENT { return ele.Attr("src", eval) }

func (ele *ELEMENT) StartDate(eval string) *ELEMENT { return ele.Attr("startDate", eval) }

func (ele *ELEMENT) TextTracks(eval string) *ELEMENT { return ele.Attr("textTracks", eval) }

func (ele *ELEMENT) VideoTracks(eval string) *ELEMENT { return ele.Attr("videoTracks", eval) }

func (ele *ELEMENT) Volume(eval string) *ELEMENT { return ele.Attr("volume", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  HTML Audio / Video Events - http://www.w3schools.com/tags/ref_av_dom.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Abort(eval string) *ELEMENT { return ele.Attr("abort", eval) }

func (ele *ELEMENT) Canplay(eval string) *ELEMENT { return ele.Attr("canplay", eval) }

func (ele *ELEMENT) Canplaythrough(eval string) *ELEMENT { return ele.Attr("canplaythrough", eval) }

func (ele *ELEMENT) Durationchange(eval string) *ELEMENT { return ele.Attr("durationchange", eval) }

func (ele *ELEMENT) Emptied(eval string) *ELEMENT { return ele.Attr("emptied", eval) }

func (ele *ELEMENT) Loadeddata(eval string) *ELEMENT { return ele.Attr("loadeddata", eval) }

func (ele *ELEMENT) Loadedmetadata(eval string) *ELEMENT { return ele.Attr("loadedmetadata", eval) }

func (ele *ELEMENT) Loadstart(eval string) *ELEMENT { return ele.Attr("loadstart", eval) }

func (ele *ELEMENT) Playing(eval string) *ELEMENT { return ele.Attr("playing", eval) }

func (ele *ELEMENT) Progress(eval string) *ELEMENT { return ele.Attr("progress", eval) }

func (ele *ELEMENT) Ratechange(eval string) *ELEMENT { return ele.Attr("ratechange", eval) }

func (ele *ELEMENT) Seeked(eval string) *ELEMENT { return ele.Attr("seeked", eval) }

func (ele *ELEMENT) Stalled(eval string) *ELEMENT { return ele.Attr("stalled", eval) }

func (ele *ELEMENT) Suspend(eval string) *ELEMENT { return ele.Attr("suspend", eval) }

func (ele *ELEMENT) Timeupdate(eval string) *ELEMENT { return ele.Attr("timeupdate", eval) }

func (ele *ELEMENT) Volumechange(eval string) *ELEMENT { return ele.Attr("volumechange", eval) }

func (ele *ELEMENT) Waiting(eval string) *ELEMENT { return ele.Attr("waiting", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
//  HTML Audio / Video Shared Attributes - http://www.w3schools.com/tags/ref_av_dom.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) Ended(eval string) *ELEMENT { return ele.Attr("ended", eval) }

func (ele *ELEMENT) Error(eval string) *ELEMENT { return ele.Attr("error", eval) }

func (ele *ELEMENT) Pause(eval string) *ELEMENT { return ele.Attr("pause", eval) }

func (ele *ELEMENT) Play(eval string) *ELEMENT { return ele.Attr("play", eval) }

func (ele *ELEMENT) Seeking(eval string) *ELEMENT { return ele.Attr("seeking", eval) }
