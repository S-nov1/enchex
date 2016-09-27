FROM alpine:latest 

MAINTAINER sergeu novikau <box10@stim.by> 

#RUN  mkdir /home/user 
RUN mkdir /home/user && addgroup user && adduser -D -G user user

ENV HOME /home/user

WORKDIR $HOME
USER user  

ADD myapp /home/user/ 

CMD ["/home/user/myapp","/home/user/engtext.txt","/home/user/out_engtext.txt"]

