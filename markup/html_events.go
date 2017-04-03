package g

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Window events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnAfterPrint(eval string) *ELEMENT { return ele.Attr("onafterprint", eval) }

func (ele *ELEMENT) OnBeforePrint(eval string) *ELEMENT { return ele.Attr("onbeforeprint", eval) }

func (ele *ELEMENT) OnBeforeUnload(eval string) *ELEMENT { return ele.Attr("onbeforeunload", eval) }

// Moved to html_sharedAttributes.go because of conflict
// func (ele *ELEMENT) OnError(eval string) *ELEMENT { return ele.Attr("onerror", eval) }

func (ele *ELEMENT) OnHashChange(eval string) *ELEMENT { return ele.Attr("onhashchange", eval) }

func (ele *ELEMENT) OnLoad(eval string) *ELEMENT { return ele.Attr("onload", eval) }

func (ele *ELEMENT) OnMessage(eval string) *ELEMENT { return ele.Attr("onmessage", eval) }

func (ele *ELEMENT) OnOffline(eval string) *ELEMENT { return ele.Attr("onoffline", eval) }

func (ele *ELEMENT) OnOnline(eval string) *ELEMENT { return ele.Attr("ononline", eval) }

func (ele *ELEMENT) OnPageHide(eval string) *ELEMENT { return ele.Attr("onpagehide", eval) }

func (ele *ELEMENT) OnPageShow(eval string) *ELEMENT { return ele.Attr("onpageshow", eval) }

func (ele *ELEMENT) OnPopState(eval string) *ELEMENT { return ele.Attr("onpopstate", eval) }

func (ele *ELEMENT) OnResize(eval string) *ELEMENT { return ele.Attr("onresize", eval) }

func (ele *ELEMENT) OnStorage(eval string) *ELEMENT { return ele.Attr("onstorage", eval) }

func (ele *ELEMENT) OnUnload(eval string) *ELEMENT { return ele.Attr("onunload", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Form events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnBlur(eval string) *ELEMENT { return ele.Attr("onblur", eval) }

func (ele *ELEMENT) OnChange(eval string) *ELEMENT { return ele.Attr("onchange", eval) }

func (ele *ELEMENT) OnContextMenu(eval string) *ELEMENT { return ele.Attr("oncontextmenu", eval) }

func (ele *ELEMENT) OnFocus(eval string) *ELEMENT { return ele.Attr("onfocus", eval) }

func (ele *ELEMENT) OnInput(eval string) *ELEMENT { return ele.Attr("oninput", eval) }

func (ele *ELEMENT) OnInvalid(eval string) *ELEMENT { return ele.Attr("oninvalid", eval) }

func (ele *ELEMENT) OnReset(eval string) *ELEMENT { return ele.Attr("onreset", eval) }

func (ele *ELEMENT) OnSearch(eval string) *ELEMENT { return ele.Attr("onsearch", eval) }

func (ele *ELEMENT) OnSelect(eval string) *ELEMENT { return ele.Attr("onselect", eval) }

func (ele *ELEMENT) OnSubmit(eval string) *ELEMENT { return ele.Attr("onsubmit", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Keyboard events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnKeyDown(eval string) *ELEMENT { return ele.Attr("onkeydown", eval) }

func (ele *ELEMENT) OnKeyPress(eval string) *ELEMENT { return ele.Attr("onkeypress", eval) }

func (ele *ELEMENT) OnKeyUp(eval string) *ELEMENT { return ele.Attr("onkeyup", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Mouse events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnClick(eval string) *ELEMENT { return ele.Attr("onclick", eval) }

func (ele *ELEMENT) OnDblClick(eval string) *ELEMENT { return ele.Attr("ondblclick", eval) }

func (ele *ELEMENT) OnDrag(eval string) *ELEMENT { return ele.Attr("ondrag", eval) }

func (ele *ELEMENT) OnDrageNd(eval string) *ELEMENT { return ele.Attr("ondragend", eval) }

func (ele *ELEMENT) OnDragEnter(eval string) *ELEMENT { return ele.Attr("ondragenter", eval) }

func (ele *ELEMENT) OnDragLeave(eval string) *ELEMENT { return ele.Attr("ondragleave", eval) }

func (ele *ELEMENT) OnDragOver(eval string) *ELEMENT { return ele.Attr("ondragover", eval) }

func (ele *ELEMENT) OnDragStart(eval string) *ELEMENT { return ele.Attr("ondragstart", eval) }

func (ele *ELEMENT) OnDrop(eval string) *ELEMENT { return ele.Attr("ondrop", eval) }

func (ele *ELEMENT) OnMouseDown(eval string) *ELEMENT { return ele.Attr("onmousedown", eval) }

func (ele *ELEMENT) OnMouseMove(eval string) *ELEMENT { return ele.Attr("onmousemove", eval) }

func (ele *ELEMENT) OnMouseOut(eval string) *ELEMENT { return ele.Attr("onmouseout", eval) }

func (ele *ELEMENT) OnMouseOver(eval string) *ELEMENT { return ele.Attr("onmouseover", eval) }

func (ele *ELEMENT) OnMouseUp(eval string) *ELEMENT { return ele.Attr("onmouseup", eval) }

func (ele *ELEMENT) OnMouseWheel(eval string) *ELEMENT { return ele.Attr("onmousewheel", eval) }

func (ele *ELEMENT) OnScroll(eval string) *ELEMENT { return ele.Attr("onscroll", eval) }

func (ele *ELEMENT) OnWheel(eval string) *ELEMENT { return ele.Attr("onwheel", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Clipboard events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnCopy(eval string) *ELEMENT { return ele.Attr("oncopy", eval) }

func (ele *ELEMENT) OnCut(eval string) *ELEMENT { return ele.Attr("oncut", eval) }

func (ele *ELEMENT) OnPaste(eval string) *ELEMENT { return ele.Attr("onpaste", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Media events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnAbort(eval string) *ELEMENT { return ele.Attr("onabort", eval) }

func (ele *ELEMENT) OnCanPlay(eval string) *ELEMENT { return ele.Attr("oncanplay", eval) }

func (ele *ELEMENT) OnCanPlayThrough(eval string) *ELEMENT { return ele.Attr("oncanplaythrough", eval) }

func (ele *ELEMENT) OnCueChange(eval string) *ELEMENT { return ele.Attr("oncuechange", eval) }

func (ele *ELEMENT) OnDurationChange(eval string) *ELEMENT { return ele.Attr("ondurationchange", eval) }

func (ele *ELEMENT) OnEmptied(eval string) *ELEMENT { return ele.Attr("onemptied", eval) }

func (ele *ELEMENT) OnEnded(eval string) *ELEMENT { return ele.Attr("onended", eval) }

func (ele *ELEMENT) OnLoadedData(eval string) *ELEMENT { return ele.Attr("onloadeddata", eval) }

func (ele *ELEMENT) OnLoadedMetadata(eval string) *ELEMENT { return ele.Attr("onloadedmetadata", eval) }

func (ele *ELEMENT) OnLoadStart(eval string) *ELEMENT { return ele.Attr("onloadstart", eval) }

func (ele *ELEMENT) OnPause(eval string) *ELEMENT { return ele.Attr("onpause", eval) }

func (ele *ELEMENT) OnPlay(eval string) *ELEMENT { return ele.Attr("onplay", eval) }

func (ele *ELEMENT) OnPlaying(eval string) *ELEMENT { return ele.Attr("onplaying", eval) }

func (ele *ELEMENT) OnProgress(eval string) *ELEMENT { return ele.Attr("onprogress", eval) }

func (ele *ELEMENT) OnRateChange(eval string) *ELEMENT { return ele.Attr("onratechange", eval) }

func (ele *ELEMENT) OnSeeked(eval string) *ELEMENT { return ele.Attr("onseeked", eval) }

func (ele *ELEMENT) OnSeeking(eval string) *ELEMENT { return ele.Attr("onseeking", eval) }

func (ele *ELEMENT) OnStalled(eval string) *ELEMENT { return ele.Attr("onstalled", eval) }

func (ele *ELEMENT) OnSuspend(eval string) *ELEMENT { return ele.Attr("onsuspend", eval) }

func (ele *ELEMENT) OnTimeUpdate(eval string) *ELEMENT { return ele.Attr("ontimeupdate", eval) }

func (ele *ELEMENT) OnVolumeChange(eval string) *ELEMENT { return ele.Attr("onvolumechange", eval) }

func (ele *ELEMENT) OnWaiting(eval string) *ELEMENT { return ele.Attr("onwaiting", eval) }

// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---
// Media events - http://www.w3schools.com/tags/ref_eventattributes.asp
// --- --- --- --- --- --- --- --- --- --- --- --- --- --- --- ---

func (ele *ELEMENT) OnShow(eval string) *ELEMENT { return ele.Attr("onshow", eval) }

func (ele *ELEMENT) OnToggle(eval string) *ELEMENT { return ele.Attr("ontoggle", eval) }
