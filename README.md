# A command line tool you probably don't need

## iCat
have you ever been in the situation where you have an image on a server and don't really want to copy it to your machine just to see the content? Well this tool can help..it allows you to dispaly images in your terminal window (iTerm only sorry). No matter if the file is located on a remote server, your local file system or is accessible through an URL. (probs are you will not find yourself in this situation that often however if you do - well be sure there is a tool for everything)

## Usage
- `$ icat path/to/image` -> display image in terminal window<br>
- `$ icat https://somepage.com/awesome_image` -> path can also be URL<br>
- `$ icat sftp://username@server:1024/home/username/img.png` -> display image from remote server (default port is 22)
- `$ icat -base64 path/to/image` -> write base64 of image to terminal (not quite useful I know)
- `$ icat -base64 -out path/to/output path/to/image` -> write base64 of image to output file

<!-- 
## Touch and feel
![](git_resources/icat-demo.gif)
-->
