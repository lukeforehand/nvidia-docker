FROM cuda

RUN apt-get update && apt-get install -y --no-install-recommends \
        cuda-samples-$CUDA_PKG_VERSION && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/local/cuda/samples
RUN make -C 1_Utilities/deviceQuery
RUN make -C 1_Utilities/bandwidthTest
RUN make -C 0_Simple/matrixMulCUBLAS
RUN make -C 0_Simple/vectorAdd

