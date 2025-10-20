#! /bin/bash

DIR="$(cd "$(dirname "$0")" && pwd)"

function increment_number() {
    local n=$1
    local m=$2
    local filename=$3
    # shellcheck disable=SC2155
    # shellcheck disable=SC2086
    local new_value=$(awk -v n=$n -v m=$m '{
        if (NR == n) {
            for (i=1; i<=NF; i++) {
                if (i == m) {
                    $i = $i + 1
                    print $i
                    exit
                }
            }
        }
    }' $filename)
    echo $new_value
    # shellcheck disable=SC2086
    awk "NR==${n}{\$${m}=${new_value}}1" $filename > ${DIR}/temp-built && mv ${DIR}/temp-built $filename
#    echo "$new_value"
}

case "$1" in
leo)
  verr=$(increment_number 1 1 ${DIR}/built)
  echo "$verr"
  ;;
leoctl)
  verr=$(increment_number 1 2 ${DIR}/built)
  echo "$verr"
  ;;
*)
  echo "path/version.sh leo/leoctl"
  ;;
esac
