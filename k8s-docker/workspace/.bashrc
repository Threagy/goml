export NVM_DIR="/usr/local/nvm"
[ -s "$NVM_DIR/nvm.sh" ] && . "$NVM_DIR/nvm.sh"  # This loads nvm

export GOPATH="/go"
export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin

export PS1='${debian_chroot:+($debian_chroot)}[\[\033[00;31m\]\u@\h\[\033[00m\]: \[\033[00;34m\]\W\[\033[00m\]]\$ '

alias ll='LC_COLLATE=C ls -alF --group-directories-first'