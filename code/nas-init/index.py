# coding=utf-8
import os


def handler(event, context):
    if not os.path.exists("data/StableDiffusion"):
        os.system(
            "mkdir -p data/StableDiffusion")
        os.system(
            "wget https://huggingface.co/runwayml/stable-diffusion-v1-5/resolve/main/v1-5-pruned-emaonly.safetensors -O /data/StableDiffusion/v1-5-pruned-emaonly.safetensors")
    return "nas init"
