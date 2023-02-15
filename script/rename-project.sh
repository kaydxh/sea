#!/usr/bin/env bash

# Fail on any error
set -euo pipefail
#set -o xtrace

PROJECT_ROOT_NAME=sea
OLD_PROJECT_NAME=$1
NEW_PROJECT_NAME=$2
OLD_PROJECT_DASH_NAME="${PROJECT_ROOT_NAME}-${OLD_PROJECT_NAME}"
NEW_PROJECT_DASH_NAME="${PROJECT_ROOT_NAME}-${NEW_PROJECT_NAME}"
OLD_PROJECT_JOINT_NAME="${PROJECT_ROOT_NAME}${OLD_PROJECT_NAME}"
NEW_PROJECT_JOINT_NAME="${PROJECT_ROOT_NAME}${NEW_PROJECT_NAME}"

OLD_GIT_REPOSITORY_NAME=$3 #github.com/kaydxh
NEW_GIT_REPOSITORY_NAME=$4 #git.code.oa.com/kaydxh

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

# sed 默认大小写敏感，/i参数不区分大小写
function renameFilesAndDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*" -type f -path "*${OLD_PROJECT_NAME}*")
  do
    renameProjectDir ${it} ${OLD_PROJECT_NAME} ${NEW_PROJECT_NAME}
    renameProjectDir ${it} ${OLD_PROJECT_DASH_NAME} ${NEW_PROJECT_DASH_NAME}
    renameProjectDir ${it} ${OLD_PROJECT_JOINT_NAME} ${NEW_PROJECT_JOINT_NAME}
  done

  rmDirectories 
}

# $1 file or dir name
# 老服务文件名/目录名 替换为 新服务文件名/目录名
function renameProjectDir {
    oldFile=$1
    #newFile=`echo ${oldFile} | sed -e "s/${OLD_PROJECT_NAME}/${NEW_PROJECT_NAME}/g"`
    newFile=`echo ${oldFile} | sed -e "s/$2/$3/g"`
    newDir=${newFile%/*}
    if [[ -f ${oldFile} ]]; then mkdir -p ${newDir}; mv -nv ${oldFile} ${newFile}; echo "${oldFile} ==> ${newFile}"; fi
}

function rmDirectories {
  for it in $(find . ! -path "*third_party*" ! -path "*node_modules*"  -type d -path "*${OLD_PROJECT_NAME}*")
  do
    oldDir=${it}
    rm -rf ${oldDir}
  done
}

function replaceContentOfFiles {
  UPPER_BEGIN_OLD_PROJECT_NAME=$(echo ${OLD_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${OLD_PROJECT_NAME:1}
  UPPER_BEGIN_NEW_PROJECT_NAME=$(echo ${NEW_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${NEW_PROJECT_NAME:1}
  #for it in $(grep -E -RIl --exclude-dir={*third_party*,*node_modules*,*output*} '${OLD_PROJECT_NAME}|${UPPER_BEGIN_OLD_PROJECT_NAME}' .)
  for it in $(grep -E -RIl --exclude-dir={*third_party*,*node_modules*,*output*} "${OLD_PROJECT_NAME}|${UPPER_BEGIN_OLD_PROJECT_NAME}" .)
  do
    echo "${it}"
    # skip soft link file
    if [[ -h ${it} ]];then
      continue
    fi

   sed -i "" "/Validate/!s/${OLD_PROJECT_NAME}/${NEW_PROJECT_NAME}/g" "${it}" 

   #首字母大写替换
   #将OLD_PROJECT_NAME变量值的首字母转化为大写，并保存在UPPER_BEGIN_OLD_PROJECT_NAME变量中
   #UPPER_BEGIN_OLD_PROJECT_NAME=$(echo ${OLD_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${OLD_PROJECT_NAME:1}
   #UPPER_BEGIN_NEW_PROJECT_NAME=$(echo ${NEW_PROJECT_NAME:0:1} | tr '[a-z]' '[A-Z]')${NEW_PROJECT_NAME:1}
   sed -i "" "s/${UPPER_BEGIN_OLD_PROJECT_NAME}/${UPPER_BEGIN_NEW_PROJECT_NAME}/g" "${it}"

   # replace git name
   sed -i "" "s/${UPPER_BEGIN_OLD_PROJECT_NAME} /${UPPER_BEGIN_NEW_PROJECT_NAME}/g" "${it}"

   # support by base 4.0
   #sed -i "" "s/${OLD_PROJECT_NAME^}/${NEW_PROJECT_NAME^}/g" "${it}"
  done
}

function replaceGitRespositoryNameOfFiles {
  for it in $(grep -RIl --exclude-dir={*third_party*,*node_modules*,*output*} --exclude=*rename-project.sh "${OLD_GIT_REPOSITORY_NAME}" .)
  do
    echo "${it}"
    # skip soft link file
    if [[ -h ${it} ]];then
      continue
    fi

    OLD_GIT_NAME=${OLD_GIT_REPOSITORY_NAME}/${OLD_PROJECT_NAME}
    NEW_GIT_NAME=${NEW_GIT_REPOSITORY_NAME}/${NEW_PROJECT_NAME}
    replaceString ${OLD_GIT_NAME} ${NEW_GIT_NAME}
    replaceString ${OLD_GIT_REPOSITORY_NAME} ${NEW_GIT_REPOSITORY_NAME}
  done
}


function replaceString {
    ESCAPED_OLD_GIT_NAME=$(printf '%s\n' "$1" | sed -e 's/[]\/$*.^[]/\\&/g')
    ESCAPED_NEW_GIT_NAME=$(printf '%s\n' "$2" | sed -e 's/[]\/$*.^[]/\\&/g')
    sed -i "" "s/${ESCAPED_OLD_GIT_NAME}/${ESCAPED_NEW_GIT_NAME}/g" "${it}"
   # support by base 4.0
   #sed -i "" "s/${OLD_PROJECT_NAME^}/${NEW_PROJECT_NAME^}/g" "${it}"
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
replaceGitRespositoryNameOfFiles 

