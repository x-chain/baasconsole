

grep image: ibm_cloud_deploy.yaml | awk -F"image:" '{print $2}' | xargs -I@ docker pull @