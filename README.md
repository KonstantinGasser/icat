# A command line tool you probably don't need

## iCat
render images in you iTerm window or copy the base64 string of an image to a output file.
<br>(You could also print the base64 to the iTerm, in case you want to ??)<br>
inspired by Francesc Campoy and his video on io.Pipes.  

## Usage
- `$ icat path/to/image` -> display image in terminal window<br>
- `$ icat https://somepage.com/awesome_image` -> path can also be URL<br>
- !not implemented yet `$ icat username@server:22/home/username/img.png` -> display image from remote server (default port is 22)
- `$ icat -base64 path/to/image` -> write base64 of image to terminal (not quite useful I know)
- `$ icat -base64 -out path/to/output path/to/image` -> write base64 of image to output file

 
## Touch and feel
![](git_resources/icat-demo.gif)
