#!/bin/bash


parse_yaml() {
   local prefix=$2
   local s='[[:space:]]*' w='[a-zA-Z0-9_]*' fs=$(echo @|tr @ '\034')
   sed -ne "s|^\($s\)\($w\)$s:$s\"\(.*\)\"$s\$|\1$fs\2$fs\3|p" \
        -e "s|^\($s\)\($w\)$s:$s\(.*\)$s\$|\1$fs\2$fs\3|p"  $1 |
   awk -F$fs '{
      indent = length($1)/2;
      vname[indent] = $2;
      for (i in vname) {if (i > indent) {delete vname[i]}}
      if (length($3) > 0) {
         vn=""; for (i=0; i<indent; i++) {vn=(vn)(vname[i])("_")}
         printf("%s%s%s=\"%s\"\n", "'$prefix'",vn, $2, $3);
      }
   }'
}

eval $(parse_yaml config.yaml "config_")

for f in templates/dist/js/app*.js
do
  if [ -n "${COCO_HOSTNAME}" ]; then
    sed -i "s,localhost,${COCO_HOSTNAME},g" $f
  fi

  sed -i "s,8080,${config_netcoco_port},g" $f
  cat $f
done

if [ -n "${COCO_HOSTNAME}" ]; then
    sed -i "s,localhost,${COCO_HOSTNAME},g" templates/js/main.js 
fi

sed -i "s,8080,${config_netcoco_port},g" templates/js/main.js 

./netcoco-linux-amd64 -file=./config.yaml

