# A command line tool you probably don't need

## iCat
render images in you iTerm window or copy the base64 string of an image to a output file.
<br>(You could also print the base64 to the iTerm, in case you want to ??)<br>
inspired by Francesc Campoy and his video on io.Pipes.  

## Usage
`$ icat view path/to/image` -> display image in terminal window<br>
`$ icat view https://somepage.com/awesome_image` -> path can also be URL<br>
`$ icat view path/to/image --base64` -> display base64 text file as rendered image in terminal window<br>
`$ icat base64 path/to/image` -> write base64 of image to terminal (not quite useful I know)<br>
`$ icat base64 path/to/image --out path/to/output` -> write base64 of image to output file
`$ icat view username@server:/home/username/img.png --remote` -> display image from remote server (default port is 22)
`$ icat view username@server:/home/username/img.png --remote -p 31415` -> remote call with custom port

## Touch and feel
![](git_resources/icat-demo.gif)
