package g



func (ele *ELEMENT) Background(s string) *ELEMENT {

  return ele.Style("background-image: url('"+s+"')")
}

func (ele *ELEMENT) BackgroundImage(url, size, pos string) *ELEMENT {

  return ele.Style("background-image: url('"+url+"');background-size:" + size + ";background-position:" + pos + ";background-repeat: no-repeat")
}

func (ele *ELEMENT) NgBackgroundImage(url, size, pos string) *ELEMENT {

	return ele.NgStyle("{'background-image': 'url('+" + url + "+')' }").Style("background-size:" + size + ";background-position:" + pos)
}
