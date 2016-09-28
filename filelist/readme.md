# 1. Step
Build container from source: bash build.sh
# 2. Step
Paste source "*.txt" files in current dir.
# 3. Step
Run application: docker run -v "$PWD":/home/user  --rm enclist:latest
# 4. Step
Read "*.hxf" files in current dir.
