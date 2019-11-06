#!/bin/bash 
set -e

template_name='go-webservice-template'
gopath_template_name='github.com/8bitdogs/'$template_name

new_name=$(basename $(pwd))
gopath=$(go env GOPATH)
p=$(pwd)

if [[ $p == *"$gopath"* ]]; then
    new_name=${p#"$gopath/src/"}
fi

echo $gopath_template_name
echo $new_name

echo 'applying new name='$new_name' in import path and another resources'

find * -type f \( ! -name '*.sh' \) -exec \
    sed -i "s#$gopath_template_name#$new_name#g" {} +

echo 'done. now you can remove this file'