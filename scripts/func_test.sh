#!/bin/bash

FOLDERS=($(ls -d events/*/))

PRE_NAME=Smaily

test_lambda() {
  for folder in "${FOLDERS[@]}"; do
    (
      folder_name=$(basename "${folder}")
      aws lambda invoke --function-name ${PRE_NAME}-${folder_name} --payload file://events/${folder_name}/request/input.json --cli-binary-format raw-in-base64-out ./events/${folder_name}/response/output.json
      echo -e "\n"
      cat ./events/${folder_name}/response/output.json
      echo -e "\n"
    )
  done
}

test_lambda