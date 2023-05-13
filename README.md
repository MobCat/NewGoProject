# NewGoProject
Make and setup a new golang project without and IDE on windows

# Why?
Well I have a really bad habit of not using IDEs. I just don't like the bloat. I make just about everything in Notepad++ or Sublime.<br>
So I needed a way to setup a new project without making all these files and running these commands by hand.<br>

# Setup
Download and "install" [UPX](https://upx.github.io/) for compiling and squishing exes<br>
(UPX is not needed and can be commented out, but I like it)<br>
Download the zip and extract it to where you have your projects stored.<br>
like Y:\projects\go\ <br>
Edit the `init.go` and `init.cmd` to your liking<br>
`init.go` Is your main.go template file. So any code or boiler plate stuff you add to every new project can be added here<br>
`init.cmd` is optional but, this is any post commands you may want to run<br>
The sample installs 2 libraries, but theses are just normal windows commands, you can place any windows commands you like in this file<br>
Like init the project to git, download stuff, do network things, anything you like.<br>
I said optional as there is a `Run init cmds? [Y/N]:` option, if you don't want to run extra cmds, just type `n`.. or anything that's not `y`<br>

# How to use
Simply just run the `newProject.exe` out of the root of your project folder.<br>
It will ask you for a project name (Names with spaces will just get cut off)<br>
Then ask if you want to run init cmds.<br>
Then it will make a new folder, copy the template files into new project folder.
`go mod init ProjectName` the project file<br>
Run the init cmds if you wanted to<br>
Then clean up after it's self.<br><br>

Now you can just edit your `main.go` however you normaly do, use `go run main.go` to test it,<br>
then you can use `build.bat` to compile your project into a compact exe.

# TODO:
Maybe make this more cmd friendly. like `newProject ProjectName`. If a command arg is found, don't ask the user for a project name, use the command arg.<br>
Give more options for setting up new projects with multiple templates the user can choose from.<br>
Pop your editor of choice for editing the `init.cmd` and `init.go` files.<br>
There's probably a way to make this platform agnostic but I only know how to windows, sorry.<br>
