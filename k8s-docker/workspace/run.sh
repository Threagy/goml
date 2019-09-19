#!/bin/bash

#########################################
# Handle SIG TERM
WORKER_PID=''

handle_sig_term(){
    echo "[Shell] SIGTERM received"
    kill -TERM $WORKER_PID
    wait $WORKER_PID
}

trap 'handle_sig_term' TERM

#########################################
# Custom Logic

# ssh config
if [ ! -d "~/.ssh" ]; then
  mkdir ~/.ssh > /dev/null
fi

cp -r /root/.temp/.ssh /root
chown -R root:root /root/.ssh

cp -r /root/.temp/.gitconfig /root/.gitconfig
chown root:root /root/.gitconfig

export TERM=xterm
exec /usr/sbin/sshd -D & WORKER_PID=$!
wait $WORKER_PID