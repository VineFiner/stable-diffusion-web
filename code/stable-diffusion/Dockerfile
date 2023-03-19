# from https://github.com/AbdBarho/stable-diffusion-webui-docker/blob/master/services/AUTOMATIC1111/Dockerfile
FROM registry.cn-shanghai.aliyuncs.com/vipas/stable-diffusion-web:v1

# model: https://huggingface.co/runwayml/stable-diffusion-v1-5/resolve/main/v1-5-pruned-emaonly.safetensors
# ADD data /data

# copy v1-5-pruned-emaonly.safetensors model
# COPY StableDiffusion/v1-5-pruned-emaonly.safetensors data/StableDiffusion/ 

# copy healthCheck flask
COPY healthCheck healthCheck
