FROM golang:1.6.2-onbuild
RUN echo "deb http://www.deb-multimedia.org jessie main non-free" >> /etc/apt/sources.list
RUN apt-get update
RUN apt-get -y --force-yes dist-upgrade
RUN apt-get install -y --force-yes ffmpeg