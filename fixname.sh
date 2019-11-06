#!/bin/bash 
set -e

tpname='webservice-template'
fname=$(basename $(pwd))

echo 'applying new name='$fname' in import path and another resources'

find * -type f \( ! -name '*.sh' \) -exec \
    sed -i "s/$tpname/$fname/g" {} +

echo 'done. now you can remove this file'