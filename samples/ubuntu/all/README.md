# docker notes

## must be at least kernel v3.13
```
uname -a
```

## install docker apt repo key
```
sudo apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys 58118E89F3A912897C070ADBF76221572C52609D
```

## add docker apt repo
```
sudo vi /etc/apt/sources.list.d/docker.list
```

## update apt
```
sudo apt-get update
```

## install docker
```
sudo apt-get install docker-engine
sudo service docker start
```

## add docker group to user and logout/login
```
sudo usermod -aG docker <user>
exit
```

## verify installation with hello-world
```
sudo docker run hello-world
```

## clone nvidia-docker
```
git clone https://github.com/NVIDIA/nvidia-docker
cd nvidia-docker
```

## build cuda runtime image
```
docker build -t cuda:7.5-runtime ubuntu/cuda/7.5/runtime
docker build -t cuda:7.5-devel ubuntu/cuda/7.5/devel
docker tag cuda:7.5-devel cuda
```

## build cuDNN runtime image
```
docker build -t cuda:7.0-runtime ubuntu/cuda/7.0/runtime
docker build -t cuda:7.0-devel ubuntu/cuda/7.0/devel
docker build -t cuda:7.0-cudnn2-devel ubuntu/cuda/7.0/devel/cudnn2
docker tag cuda:7.0-cudnn2-devel cuda:cudnn-devel
```

## run device query example
* comma-delimit the GPUS i.e. 0,1
```
docker build -t device_query samples/ubuntu/deviceQuery
GPU=0 ./nvidia-docker run device_query
```

## what is ./nvidia-docker ?
* it is a wrapper for docker-cli that configures the nvidia devices and library volumes for the container
* i.e.
```
docker run \
--device=/dev/nvidiactl \
--device=/dev/nvidia-uvm \
--device=/dev/nvidia0 \
-v /usr/lib/x86_64-linux-gnu/libcuda.so.1:/usr/lib/x86_64-linux-gnu/libcuda.so.1 \
...
-it cuda_samples
```

# custom docker image example

## copy a Dockerfile
```
mkdir samples/ubuntu/all && cp samples/ubuntu/deviceQuery/Dockerfile samples/ubuntu/all
```

## custom modify a Dockerfile
```
vi samples/ubuntu/all/Dockerfile

# modify
WORKDIR /usr/local/cuda/samples
RUN make -C 1_Utilities/deviceQuery
RUN make -C 1_Utilities/bandwidthTest
RUN make -C 0_Simple/matrixMulCUBLAS
RUN make -C 0_Simple/vectorAdd
```
* reference https://github.com/lukeforehand/nvidia-docker/blob/master/samples/ubuntu/all/Dockerfile

## build custom image and run
```
docker build -t cuda_samples samples/ubuntu/all
GPU=0 ./nvidia-docker run -it cuda_samples
```

## try these sample programs
```
nvidia-smi -q
1_Utilities/deviceQuery/deviceQuery
1_Utilities/bandwidthTest/bandwidthTest --mode=shmoo
0_Simple/matrixMulCUBLAS/matrixMulCUBLAS -sizemult=10
0_Simple/vectorAdd/vectorAdd
```

## other resources
* [this gist](https://gist.github.com/lukeforehand/2f041e6badc78fba8c83)
* [docker installation](https://docs.docker.com/v1.5/installation/ubuntulinux/)
* [nvidia-docker](https://github.com/NVIDIA/nvidia-docker)
* [my fork of nvidia-docker](https://github.com/lukeforehand/nvidia-docker)

## what to do next
* create a container registry to version and share containers with your team
* learn how environment variables can be used to customize a container for specific hardware

