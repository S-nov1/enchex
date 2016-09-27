# 1. Step
Build source : go build strcontent.go or in docker 
docker run --rm -v "$PWD":/usr/src/encapp -w /usr/src/encapp golang:alpine go build -v
# 2. Step
Build application in docker container: docker build -t enclist .
# 3. Step
Paste source "*.txt" files in current dir.
# 4. Step
Run application: docker run -v "$PWD":/home/user  --rm enclist:latest
# 5. Step
Read "*.hxf" files in current dir.
