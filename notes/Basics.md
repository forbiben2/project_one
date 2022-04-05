Statements
    -Computer Program- sequence of instructions that tell a computer what to do
        -statement-type of instruction that cause a program to take an action usually ending in ;
    -In a high level language a single statement may compile many instructions
    -Types
        -Declaration
        -Jump
        -Expression
        -Compound
        -Selection
        -Iteration
        -Try Blocks
Functions
    -Function-A group of statements that Execute sequentially
    -the first function in a program is main()
Go Programming
-Strong statically typed language
-Key Features
    -Simplicity
    -Fast compile
    -Garbage collected language-an automatic memory management feature
    -Built-in concurrency
    -compiles to stand alone binaries-no reaching out to libraries, and version management at run time all dependencies are there
-Use golang.org and its play button to test things from other pages
Git
-VCS(Version control system)
    -most store info as a list of file based changes
-Git
    -more like snap shots of minature file systems
    -more compartmentalized reuse/copy unchanged items
    -works locally (offline)
    -git will know about all changes
        -makes it hard to loose data and good for experimenting
-Github
    -easy to contribute for open source
    -lots of documentation on git
    -share work
    -large storage area
    -tracks changes
    -integration options
    -great place to learn/share info
    -can work with files in browser
-Gitlab
    -some different tools
    -can host private servers
-Bitbucket
    -similar to github/gitlab
    -native support for mercurial repositories
        -mercurial repositories
            -dvcs-complete code base and full version history mirrored on all devices
            -tracked between computers
-Docker
    -DSL-Domain specific language
        -coupled software-as-a-service, and paltform-as-a-service
    -OS level virtualization-to develop and deliver software in containers
    -Containers
        -Builder-tool(s) to build a container
        -Engine-runs container
        -orchestration-tech used to manage multiple containers
        -quicker access
        -can have servers spawn and despawn quickly and automatically
        -strips down only using/showing neccessary processes
    -Docker Image
        -record of a container at a given poitn in time
        -cannot be changed
        -can be dupilicated, shared or delete
        -anatomy
            -base image-first layer built from scratch
            -parent image-can replace base image it is a reused image
            -layers-added to base image and has code to enable it to run in a container
            -container layer-docker image not only creates a container but a writable layer. this layer contains any changes.
            -docker manifest-contains a file to describe the image.
-Linux
    -spaces cause issues in terminal
    -Permissions
        -Ownership
            -user-owner/creator of a file
            -group-a set of people granted access to a file
                -two groups cannot own the same file (sub groups can work around this)
            -other-anyone else
        -Permissions
            -read-ability to read/view a file(r)
            -write-ability to alter a file(w)
            -execute-ability to run a file(x)
            -(-)- denotes no permission
            -terminal example- -rwxrwxrwx
                               -rwxr-x--x
        -Changing permissions
            -absolute
                0 ---
                1 --x
                2 -w-
                3 -wx
                4 r--
                5 r-x
                6 rw-
                7 rwx
            -symbolic
                +-add permission
                (-)- remove permisson
                =-set/override
                User types
                u-user
                g-group(/etc/group has a text file defining all groups on a system)
                o-other
                a-all
        -apt
            -advanced package tool
                -red hat-yum
                    yellow hat updater modified
            -use ex. sudo apt install [file or directory]
            
        
            

            
