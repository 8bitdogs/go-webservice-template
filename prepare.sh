#!/bin/bash 
set -e

template_name='go-webservice-template'
gopath_template_name='github.com/8bitdogs/'$template_name

new_name=$(basename $(pwd))

echo 'prepare README.md file'
echo '# '$new_name > README.md 
echo 'done'

gopath=$(go env GOPATH)
p=$(pwd)

if [[ $p == *"$gopath"* ]]; then
    new_name=${p#"$gopath/src/"}
fi

sedarg=(-i)
if [ "$(uname)" == "Darwin" ]; then
    sedarg=(-i '')
fi

echo 'applying new prefix '$new_name' in imports'

find * -type f \( ! -name '*.sh' \) -exec \
    sed "${sedarg[@]}" "s#$gopath_template_name#$new_name#g" {} +

echo 'done. now you can remove this file'
