## WebShot 

WebShot is an easy-to-use CLI for taking screenshots of webpages.
It is based on [Rod](https://github.com/go-rod/rod).

## Installation

`go install github.com/gusanmaz/webshot/cmd`

## Usage

######  Taking full webpage screenshot as an image

`webshot  -url https://www.atlasobscura.com -output atlas.png`

###### Taking full webpage screenshot with a specific width

`webshot  -url https://www.atlasobscura.com -width 800 -output atlas2.png`

###### Taking full webpage screenshot as a PDF file

`webshot  -url https://www.atlasobscura.com -type pdf -output atlas.pdf`

Unfortunately for now fo most webpages PDF rendering produces mediocre results

###### Taking full webpage screenshot as an HTML file

`webshot  -url https://www.atlasobscura.com -type html -output atlas.html`

For now this command only outputs HTML content of the given URL. Other resources referenced in the URL such as images are not captured.

###### Taking screenshot of an infinite scroll webpage

For infinite scroll webpages you need to specify screenshot height and use infinite flag with true value.

` webshot -url https://www.flickr.com/search/?text=antalya -height 10000 -type image -output antalya.png -infinite true
`
You may also modify default values of stepheight and steptime flags. WebShot scrolls infinite page downwards stepheight pixels every steptime milliseconds.

`webshot  -steptime 300 -stepheight 100 -url https://www.flickr.com/search/?text=winter -height 50000 -type image -o winter.png`

## Author

Güvenç Usanmaz

## License 

MIT License

