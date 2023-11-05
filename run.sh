#!/bin/sh
if [ $# -gt 0 ]; then
  _govs_prompt $@
  return
fi
_govs_prompt
selected_goroot=$(cat ~/.govs/goroot)
path=$(cat ~/.govs/path)
if [ -z $selected_goroot ]; then
  return
fi
if [ -z $path ]; then
  return
fi

export GOROOT=$selected_goroot
export PATH=$path
go version
