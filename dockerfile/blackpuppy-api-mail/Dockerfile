FROM ubuntu:14.04

EXPOSE 8001

# Add the server
RUN mkdir /blackpuppy-api-mail
ADD ./server /blackpuppy-api-mail/server
RUN chmod 755 /blackpuppy-api-mail/server

# Install supervisor and consul-template
RUN apt-get update
RUN echo y| apt-get install wget tar supervisor
RUN wget https://github.com/hashicorp/consul-template/releases/download/v0.10.0/consul-template_0.10.0_linux_386.tar.gz
RUN tar -xvzf consul-template_0.10.0_linux_386.tar.gz
RUN mv consul-template_0.10.0_linux_386/consul-template /bin/consul-template

# Configure consul-template supervisor and nginx
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf
ADD config.ctmpl /blackpuppy-api-mail/config.ctmpl
ADD config.json /blackpuppy-api-mail/config.json

CMD ["/usr/bin/supervisord", "-c","/etc/supervisor/conf.d/supervisord.conf"]
