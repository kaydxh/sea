#!/usr/bin/env bash

# Fail on any error
set -euo pipefail

OLD_PROJECT_NAME=$1
NEW_PROJECT_NAME=$2

function usage() {
    echo >&2 "Usage: $0 OLD_PROJECT_NAME NEW_PROJECT_NAME"
}


function checkParams {
  if [[ -z "${OLD_PROJECT_NAME}" ]]  || [[ -z "${NEW_PROJECT_NAME}" ]]; then
    usage
    exit 1
  fi

}


# #代表删除从前往后最小匹配的内容
# %代表删除从后往前最小匹配的内容
# %/*代表取从头到最后一个slash之前的所有内容
# */代表去取从第一个slash之后的所有内容

function renameFilesAndDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*" -type f -path "*${OLD_PROJECT_NAME}*")
  do
    oldFile=${it}
    newFile=`echo ${it} | sed -e "s/${OLD_PROJECT_NAME}/${NEW_PROJECT_NAME}/g"`

   newDir=${newFile%/*}
   mkdir -p ${newDir}; mv  ${oldFile} ${newFile}
  done

  rmDirectories 
}

function rmDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*"  -type d -path "*${OLD_PROJECT_NAME}*")
  do
    oldDir=${it}
    rm -rf ${oldDir}
  done
}

function replaceContentOfFiles {
  for it in $(grep -RIl --exclude-dir={*third_party*,*node_modules*,*output*} ${OLD_PROJECT_NAME} .)
  do
    echo "${it}"
    # skip soft link file
    if [[ -h ${it} ]];then
      continue
    fi

   sed -i "" "s/${OLD_PROJECT_NAME}/${NEW_PROJECT_NAME}/g" "${it}"

   UPPER_BEGIN_OLD_PROJECT_NAME=$(echo ${OLD_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${OLD_PROJECT_NAME:1}
   UPPER_BEGIN_NEW_PROJECT_NAME=$(echo ${NEW_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${NEW_PROJECT_NAME:1}
   sed -i "" "s/${UPPER_BEGIN_OLD_PROJECT_NAME}/${UPPER_BEGIN_NEW_PROJECT_NAME}/g" "${it}"

   # support by base 4.0
   # sed -i "" "s/${OLD_PROJECT_NAME^}/${NEW_PROJECT_NAME^}/g" "${it}"
  done
}

function replaceProjectRootName {
  if [[ -d ../${OLD_PROJECT_NAME} ]]; then
    mv ../${OLD_PROJECT_NAME} ../${NEW_PROJECT_NAME}
  fi
}

checkParams 
renameFilesAndDirectories 
replaceContentOfFiles 
replaceProjectRootName 

