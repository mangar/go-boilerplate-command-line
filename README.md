# go-boilerplate-command-line


This is a basic project for go command line program.
Features:
- configuration at home directory
- help
- install and uninstall 



## Configuration

Most of the configuraitons are in ```base```directory.

Some of them are listed here and some others you may have to hack by yourself.




### Debug mode

Export an environment variable: ```export FlagDebug=true``` and then run the program. 

The ```log.LogDebug``` uses this flag to print the log in command line.



### Logs

There are different logs in base project with friendly notation. Open the file ```log.go```and check the different methods.

The log.LogDebug need to have an extra parameter. This parameter can be fixed: ```config.FlagDebug```` 



- - - 



## Command line parameters


### Install / Uninstall

The installation process is: 1) create the project directory (~/.go-boilerplate); 2) temp dir and 3) basic config file.

```go run go-boilerplate.go install```

```go run go-boilerplate.go uninstall```



### Help

Prints all commands available.

```go run go-boilerplate.go help```


### Version

Prints the build date and time.

```go run go-boilerplate.go version```



### Config


Prints the config file

```go run go-boilerplate.go config```



- - - 





## Extra checks

Look for ```@BOILERPLATE``` in all code and check what you have to change in order configure parts of the software.


## Making

    make
    make install


## Extra packages - for development

    go get gopkg.in/kyokomi/emoji.v1


