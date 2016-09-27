FROM alpine:latest 

MAINTAINER sergeu novikau <box10@stim.by> 
#docker run -v "$PWD":/home/user  --rm enclist:latest
#sudo chmod -R 0777 /home/proect/enclist/
#RUN  mkdir /home/user 
RUN mkdir /home/user && addgroup user && adduser -D -G user user

ENV HOME /home/user

WORKDIR $HOME
USER user  

ADD encapp /home/user/

CMD "/home/user/encapp"


