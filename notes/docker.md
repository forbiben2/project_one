-Containers
    Parent/Base Image
        -first layer of a docker file is a parent/base iamge, and all others build upon it.
        -many PARENT images are availabe
        -Containers can share base images/parent images
    -BASE images are an empty layer to modify for more advanced users
    -A container image becomes a container at run time
    -deployable on any OS nearly seamlessly(software system agnostic)
    -specialized programs
    -containers run one specific task but can be networked
    -do not require hypervisor
    -run native on all machines
    -can share files efficiently
    -File>Image>Container
        -files instruct on how its built
        -image names it and shows location of files
        -container becomes one at run time
-Docker Commands?
    -FROM-specify a parent image
    -WORKDIR-to set working directory
    -RUN-to install any application required for your container
    -COPY- copies files or directories from a specified location
    -ADD-same as copy but can handle remote URLS/compressed files
    -ENTRYPOINT-starting point;unspecified is /bin/sh -c
    -CMD-command the  container executes
    -EXPOSE-defines port to access container
    -LABEL-add meta data to an image
Container vs virtual machine
    Container
    -absttraction is at app layer
    -package code/dependenicies together
    -multiple on a single machine
    -share OS kernal with many containers
    -can handle more apps while needing fewer VMs
    -files are 10s of MBs
    Virtual Machine
    -abstrcation hardware level
    -multiple on one machine with hypervisor?
    -each contains full OS,apps,and libraries
    -10s of GBs each
