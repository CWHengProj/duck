#Installation Script
## TODO: input params for platform to support mac and windows

##  downloads the latest release
curl -Lo duck https://github.com/CWHengProj/duck/releases/latest/download/duck-linux-amd64

## Identify current filePath
FILEPATH=$(pwd)

## create alias and append to bashrc
echo "alias duck=\"${FILEPATH}/duck\"" >> ~/.bashrc
source ~/.bashrc;

## invoke duck
echo "giving duck execute permissions";
chmod +x duck;
echo "running duck.. use "duck" next time to invoke duck (restart terminal if it does not call duck properly)";
${FILEPATH}/duck < /dev/tty;
