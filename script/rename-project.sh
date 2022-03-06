#!/usr/bin/env bash

# Fail on any error
set -euo pipefail

OLD_PROJECT_NAME=$1
NEW_PROJECT_NAME=$2

# #代表删除从前往后最小匹配的内容
# %代表删除从后往前最小匹配的内容
# %/*代表取从头到最后一个slash之前的所有内容
# */代表去取从第一个slash之后的所有内容

function renameFilesAndDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*" -type f -path "*${OLD_PROJECT_NAME}*")
  do
    oldFile=${it}
    newFile=`echo ${it} | sed -e "s/${OLD_PROJECT_NAME}/${NEW_PROJECT_NAME}/gI"`

   newDir=${newFile%/*}
   mkdir -p ${newDir}; mv  ${oldFile} ${newFile}
  done
}

function rmFilesAndDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*" ! -path "*output*" -type d -path "*${OLD_PROJECT_NAME}*")
  do
    oldDir=${it}
    rm -rf ${oldDir}
  done

}

renameFilesAndDirectories 
rmFilesAndDirectories 

