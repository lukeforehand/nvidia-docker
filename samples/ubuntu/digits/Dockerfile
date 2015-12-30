FROM cuda

RUN apt-get update
RUN apt-get install -y wget

RUN CUDA_REPO_PKG=cuda-repo-ubuntu1404_7.5-18_amd64.deb && wget http://developer.download.nvidia.com/compute/cuda/repos/ubuntu1404/x86_64/$CUDA_REPO_PKG \
    && sudo dpkg -i $CUDA_REPO_PKG
RUN ML_REPO_PKG=nvidia-machine-learning-repo_4.0-2_amd64.deb && wget http://developer.download.nvidia.com/compute/machine-learning/repos/ubuntu1404/x86_64/$ML_REPO_PKG \
    && sudo dpkg -i $ML_REPO_PKG

RUN apt-get update && apt-get install -y digits

EXPOSE 34448

WORKDIR /usr/share/digits
CMD ["./digits-server"]

