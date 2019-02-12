# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH=/Users/wayneshen/.oh-my-zsh

# Set name of the theme to load. Optionally, if you set this to "random"
# it'll load a random theme each time that oh-my-zsh is loaded.
# See https://github.com/robbyrussell/oh-my-zsh/wiki/Themes
ZSH_THEME="robbyrussell"

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion. Case
# sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment the following line to disable bi-weekly auto-update checks.
# DISABLE_AUTO_UPDATE="true"

# Uncomment the following line to change how often to auto-update (in days).
# export UPDATE_ZSH_DAYS=13

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# The optional three formats: "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder
#export PATH="$HOME/.rbenv/bin:$PATH"

# Which plugins would you like to load? (plugins can be found in ~/.oh-my-zsh/plugins/*)
# Custom plugins may be added to ~/.oh-my-zsh/custom/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi
eval "$(rbenv init -)"

#mysql
export PATH="/usr/local/mysql/bin:$PATH"
export GOROOT="/usr/local/go"
export GOPATH="$HOME/go"
#source "/Users/wayneshen/.gvm/scripts/gvm"
#gvm use go1.9.2
export PATH="$GOPATH/bin:$PATH"
#export PATH="/usr/local/go/bin:$PATH"
export PATH="/usr/X11/bin:$PATH"
export PATH="$HOME/.rbenv/bin:$PATH"
# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# ssh
# export SSH_KEY_PATH="~/.ssh/rsa_id"
GO111MODULE=auto
alias aws:fe="ssh ubuntu@admin168.cow.bet" #production-operator-F2E
alias aws:runner="ssh ubuntu@gitlab-runner.cow.bet" #ci/cd
alias aws:w1="ssh ubuntu@w1.cow.bet" #staging-api
alias aws:w2="ssh ubuntu@w2.cow.bet" #production-api
alias aws:spider="ssh ubuntu@spider.cow.bet" #production-spider
alias aws:redis="ssh -f -N -L6979:afu.ze67a8.0001.apne1.cache.amazonaws.com:6379 ubuntu@w2.cow.bet"

alias zsh="source ~/.zshrc"
alias c:zsh="code ~/.zshrc"
alias gcpredis="gcloud compute ssh  network-forwarder -- -N -L 6386:10.0.0.11:6379"
alias gcpredis2="gcloud compute ssh  network-forwarder -- -N -L 6387:10.0.0.35:6379"
alias gcpredis3="gcloud compute ssh  network-forwarder -- -N -L 6388:10.0.0.43:6379"
alias gcpsql="gcloud compute ssh  network-forwarder -- -N -L 3307:10.63.16.3:3306"
alias gcpsql2="gcloud compute ssh  network-forwarder -- -N -L 3308:10.63.16.5:3306"
alias gcpsql3="gcloud compute ssh  network-forwarder -- -N -L 3309:10.63.16.7:3306"
alias gcplog="kubectl port-forward svc/kibana 5601:443 -n=logging"
alias nsqlook="nsqlookupd"
alias nsq="nsqd --lookupd-tcp-address=127.0.0.1:4160"
alias nsqui="nsqadmin --lookupd-http-address=127.0.0.1:4161"
alias libgomock="mockgen -source=./store/store.go -destination=./store/mock/store.go"

function kubeapply(){
  kubectl apply -f /Users/wayne/projects/doc-devops/kubernetes/fsbs/$1/libgo.yaml
}
function kubedel(){
    kubectl delete pods -n=$1 $2
}
function zsh(){
  source ~/.zshrc
}
function nsqlog(){
  nsq_to_file --topic=$1 --output-dir=/tmp --lookupd-http-address=127.0.0.1:4161
}
function testsvc(){
    go test -v gitlab.cow.bet/bkd_tool/libgo/services -run ^$1$
}
function testcover(){
    go test -cover $(go list ./... | grep -v /vendor/)
}
function testdata(){
    go test -v gitlab.cow.bet/bkd_tool/libgo/store/datastore -run ^$1$
}
function buildw(){
   cd /Users/wayneshen/go/src/github.com/WayneShenHH/toolsgo && go build -o ~/go/bin/waynego
}
function build(){
    cd /Users/wayneshen/go/src/gitlab.cow.bet/bkd_tool/libgo && go build -o ~/go/bin/libgo
}
function builds(){
  GO111MODULE=off swagger generate spec -o ~/go/bin/swagger.json
}
function sqldump(){
  mysqldump --user=libgo --host=127.0.0.1 --port=3307 --set-gtid-purged=OFF --password="z*GD0R^YrMpl^n$" sbodds $1 $2 $3 $4 > "dump$(date +"%Y%m%d").sql"
}
function sqlimport(){
    mysql --protocol=tcp --host=127.0.0.1 --user=root --port=3306 --password="123456" --default-character-set=utf8 --comments --database=sbodds  < $1
}
function golintck(){
  golint -set_exit_status=1 $(go list ./... | grep -v /vendor/)
}
function run(){
    /Users/wayneshen/go/bin/libgo $1
}
function server(){
    /Users/wayneshen/go/bin/libgo http:server
}
alias odds="redis-cli -p 6386 hget odds:$1 $2_$3"
alias odds2="redis-cli -p 6387 hget odds:$1 $2_$3"
alias odds3="redis-cli -p 6388 hget odds:$1 $2_$3"
function w(){
    /Users/wayneshen/go/bin/libgo worker:$1 $2
}
function tx(){
    /Users/wayneshen/go/bin/libgo tx:sync $1
}
function tool(){
    /Users/wayneshen/go/bin/waynego $1 $2
}
function boot(){
    code ~/Documents/Work.code-workspace
    c:waynego
    #c:libgo
    aws:redis
}

function openapp(){
    open /Applications/Notes.app
    open /Applications/Slack.app
    open /Applications/SourceTree.app
    open /Applications/rdm.app
    open /Applications/GitBook\ Editor.app/
    open /Applications/Sequel\ Pro.app/
    open /Applications/Skype.app
    #open /Users/wayneshen/Applications/Chrome Apps.localized/Default menkifleemblimdogmoihpfopnplikde.app/Contents/MacOS/app_mode_loader
}

function gitcp {
    cd /Users/wayneshen/go/src/github.com/WayneShenHH/toolsgo
    git add .
    git commit -m "$1"
    git push
    git status
    echo "https://github.com/WayneShenHH/toolsgo"
}