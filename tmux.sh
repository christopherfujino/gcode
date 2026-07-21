#!/usr/bin/env bash

set -euo pipefail

SESSION='__gcode_go__'
BUILD_CMD='go run . ./heart_0.15mm_PLA_MK3S_2h19m.gcode'

cd "$(dirname $(realpath ${BASH_SOURCE[0]} ) )"

tmux new-session -s "$SESSION" -d
tmux send-keys -t "$SESSION" 'nvim .' Enter
# specific to Chris' tmux.conf, on split the focus jumps to new pane
tmux split-window -h
tmux send-keys -t "$SESSION" "ipc-runner-server -- '${BUILD_CMD}'" Enter

# move left
tmux select-pane -L

exec tmux attach-session -t "$SESSION"
