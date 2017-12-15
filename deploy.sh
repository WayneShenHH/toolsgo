#!/bin/bash

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'
MACHINE=$1
ACTION=$2
CURRENT_COMMIT=`git rev-parse HEAD`

# User Setting
# ==========================
PROJECTPATH="$HOME/go/src/github.com/WayneShenHH/toolsgo"
REMOTE_BIN="/home/ubuntu/go/bin"
OS_SET="GOOS=linux GOARCH=amd64 GOARM=7"

# 第一組worker的機器及所有工作
if [ "$MACHINE" == "worker_1" ]; then
  REMOTE_SERVERS=("127.0.0.1")
  TASKS=("toolsgo")
  DAEMONS=("log" "tools")
fi
# ==========================

if [ "$1" == "" ] || [ "$2" == "" ] || [ "$1" == "help" ] || [ "$1" == "--help" ]; then
  echo ""
cat <<-EOF
Go Project Deployment Tool

Usage:

    ./deploy.sh [machine] [action]

Available machines:

    worker_1 - All workers
    worker_2 - Without variant_offers

Available actions:

    deploy - deploy all binary files to the target machine

Example:

    ./deploy.sh worker_1 deploy
    ./deploy.sh worker_2 deploy
    ./deploy.sh worker_1 deploy && ./deploy.sh worker_2 deploy

EOF
  exit 1
fi

echo 'Start Deployment'
echo "Compile Arguments: $OS_SET"
echo "Remote BIN Location: $REMOTE_BIN"
echo "Current Commit: $CURRENT_COMMIT"
echo ''

for TASK in "${TASKS[@]}"
do
  echo $TASK
  echo "[$TASK] Compiling"
  # if env $OS_SET go build -o "$PROJECTPATH/$TASK/$TASK" $PROJECTPATH/$TASK/*.go ; then
  printf "cd ${PROJECTPATH} && env ${OS_SET} go build\n"  
  if cd $PROJECTPATH && env $OS_SET go build ; then
    printf "[$TASK] ${GREEN}Compile Complete${NC}\n"
  else
    # printf "$PROJECTPATH/workers/$TASK/$TASK\n"
    printf "[$TASK] ${RED}Compile Failed${NC}\n"
    exit 1
  fi
  echo "[$TASK] swagger generating"
  if cd $PROJECTPATH && swagger generate  spec -o ./swagger.json ; then
    printf "[$TASK] ${GREEN}swagger generate complete${NC}\n"
  else
    printf "[$TASK] ${RED}swagger generate failed${NC}\n"
    exit 1
  fi
  for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
  do
    echo "[$TASK] Uploading to $REMOTE_SERVER"
    if scp $PROJECTPATH/$TASK ubuntu@$REMOTE_SERVER:$REMOTE_BIN/$TASK.tmp ; then
      printf "[$TASK] ${GREEN}Upload to $REMOTE_SERVER Complete${NC}\n"
    else
      printf "[$TASK] ${RED}Upload to $REMOTE_SERVER Failed${NC}\n"
      exit 1
    fi
  done

  for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
  do
    echo "[$TASK] Uploading swagger.json to $REMOTE_SERVER"
    if scp $PROJECTPATH/swagger.json ubuntu@$REMOTE_SERVER:$REMOTE_BIN/libgo-swagger.json ; then
      printf "[$TASK] ${GREEN}Upload to $REMOTE_SERVER Complete${NC}\n"
    else
      printf "[$TASK] ${RED}Upload to $REMOTE_SERVER Failed${NC}\n"
      exit 1
    fi
  done
done

echo 'Stopping All Services'
for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
do
  for DAEMON in "${DAEMONS[@]}"
  do
    ssh ubuntu@$REMOTE_SERVER "
      if [ -f /etc/systemd/system/multi-user.target.wants/$DAEMON.service ]; then
        sudo systemctl stop $DAEMON.service
      else
        printf \"[$DAEOMON] ${RED}Service does not exist. Make sure systemd is setup.${NC}\n\"
        echo '[$DAEOMON] Setup instruction: http://gitlab.cow.bet/bkd_tool/libgo'
      fi
    "
  done
done

echo 'Updating Tasks'
for TASK in "${TASKS[@]}"
do
  for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
  do
    ssh ubuntu@$REMOTE_SERVER "
      mv $REMOTE_BIN/$TASK.tmp $REMOTE_BIN/$TASK
    "
  done
done

echo 'Restarting All Services'
for DAEMON in "${DAEMONS[@]}"
do
  for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
  do
    ssh ubuntu@$REMOTE_SERVER "
      sudo systemctl start $DAEMON.service
    "
  done
done

echo "Updating Version Info"
touch $PROJECTPATH/libgo-version
echo '' > $PROJECTPATH/libgo-version
echo "commit: $CURRENT_COMMIT" >> $PROJECTPATH/libgo-version
echo "last updated: `date`" >> $PROJECTPATH/libgo-version

for REMOTE_SERVER in "${REMOTE_SERVERS[@]}"
do
  if scp $PROJECTPATH/libgo-version ubuntu@$REMOTE_SERVER:$REMOTE_BIN/libgo-version ; then
    printf "${GREEN}Version Info Updated${NC}\n"
  else
    printf "${RED}Version Info Update Failed${NC}\n"
    exit 1
  fi
done
rm $PROJECTPATH/libgo-version

printf "\n${GREEN}DONE!${NC}\n"
