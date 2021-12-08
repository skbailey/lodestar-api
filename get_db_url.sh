#!/bin/bash

ssm_db_path="/dev/lodestar/db/"
db_params=`aws ssm get-parameters-by-path --path ${ssm_db_path} | jq -c ' reduce .Parameters[] as $o ({}; .[$o.Name] = $o.Value)'`

host=$(jq -r '."/dev/lodestar/db/host"' <<< ${db_params})
port=$(jq -r '."/dev/lodestar/db/port"' <<< ${db_params})
name=$(jq -r '."/dev/lodestar/db/name"' <<< ${db_params})
password=$(jq -r '."/dev/lodestar/db/password"' <<< ${db_params})
username=$(jq -r '."/dev/lodestar/db/username"' <<< ${db_params})

echo "$host"
echo "$name"
echo "$port"
echo "$username"
echo "$password"
