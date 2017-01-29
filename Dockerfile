FROM centos:7
LABEL maintainer "oreuta@gmail.com"
EXPOSE 9999 
RUN mkdir /root/ui
COPY ui /root/ui
COPY euler10 /root/
CMD /root/euler10 -dir='/root/ui' -port=9999 
