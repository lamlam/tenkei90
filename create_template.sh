#!/bin/zsh

QNUM=$1
if [ -z $QNUM ];then
  echo "Usage: ./create_template.sh {NUMBER_OF_QUESTION}"
  exit 1
fi

if [ -d $QNUM ];then
  echo "$QNUM is already exist"
  exit 1
fi

echo "create $QNUM directory"
mkdir $QNUM

cd $QNUM
cp ../XXX/XXX.go $QNUM.go
cp ../XXX/XXX_test.go ${QNUM}_test.go
cp ../XXX/sample_input.txt sample_input.txt
cp ../XXX/sample_result.txt sample_result.txt

go mod init github.com/lamlam/goatcoder/tenkei90/$QNUM

