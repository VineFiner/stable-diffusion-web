# from https://github.com/AbdBarho/stable-diffusion-webui-docker/blob/master/services/AUTOMATIC1111/Dockerfile
FROM registry.cn-shanghai.aliyuncs.com/vipas/stable-diffusion-web:v1

# model: https://huggingface.co/runwayml/stable-diffusion-v1-5/resolve/main/v1-5-pruned-emaonly.safetensors
# ADD data /data

# 下载模型
# RUN apt-get update && apt-get install -y wget && apt-get install -y unzip 
# RUN wget -O v1-5-pruned-emaonly.safetensors https://huggingface.co/runwayml/stable-diffusion-v1-5/resolve/main/v1-5-pruned-emaonly.safetensors
# RUN mv v1-5-pruned-emaonly.safetensors /stable-diffusion-webui/models/Stable-diffusion/v1-5-pruned-emaonly.safetensors

# 添加中文扩展
RUN wget -O main.zip https://github.com/dtlnor/stable-diffusion-webui-localization-zh_CN/archive/refs/heads/main.zip
RUN mv main.zip /stable-diffusion-webui/extensions/main.zip && cd /stable-diffusion-webui/extensions && unzip main.zip && rm -rf main.zip

# copy healthCheck flask
COPY healthCheck healthCheck
