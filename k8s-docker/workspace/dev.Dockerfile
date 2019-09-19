############################
# dev image
############################
FROM goml/base-go-proto:v0.1 as dev

RUN apt-get update
RUN apt-get install -y openssh-server

RUN mkdir /var/run/sshd

RUN echo 'root:root' |chpasswd

RUN sed -ri 's/^#?PermitRootLogin\s+.*/PermitRootLogin yes/' /etc/ssh/sshd_config
RUN sed -ri 's/UsePAM yes/#UsePAM yes/g' /etc/ssh/sshd_config

RUN mkdir /root/.ssh

RUN apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

EXPOSE 22

COPY ./workspace/.bashrc /root/.bashrc
COPY ./workspace/run.sh /tmp

ENTRYPOINT ["/tmp/run.sh"]
