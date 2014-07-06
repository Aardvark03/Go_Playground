Go_Playground
=============

Gogogogo! Let's go!


From the  Get started with Go tutorial
Workspace

A workspace in go has the following structure:
    workspace/
          bin # executable binaries
          pkg # compiled object files
          src # source code 

The bin and pkg folders are generated automatically by the go tool.

export GOPATH=$HOME/GoProjects/Go_Playground/workspace
set the path to the bin folder of your project to the path variable
export PATH=$PATH:$HOME/GoProjects/Go_Playground/workspace/bin

Choose a namespace 
It is possible to the github.com/username which seems to simplify things later on. An example given is the command go get.
The namespace is set by creating a folder within the src folder of your workspace.
src/github.com/Aardvark03/Go_Playground/

Now you can use the go install command with the full path of the new package to build and install the executable to your GOPATH - bin sub directory.
$ go install github.com/Aardvark03/Go_Playground/

Through this install procedure, it is possible to execute the binary by typing the path from the workspace to the bin directory to the console.
$  workspace/bin/Go_Playground // does execute the binary
