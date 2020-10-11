# Created by kayxhding on 2020-10-11 21:04:00
#!/usr/bin/env bash

# Fail on any error
set -euo pipefail

echo " ==> Clean..."

OUT_PUT_PATH=output/bin
if [[ -d  ${OUT_PUT_PATH} ]]; then
  rm ${OUT_PUT_PATH}/sea >/dev/null 2>&1 || true
fi

echo "==> Clean finish"
