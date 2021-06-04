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

NETCOCO_PORT="8080"

initial(){
   for f in templates/dist/js/app*.js
   do
   if [ -n "${COCO_HOSTNAME}" ]; then
      sed -i "s,localhost,${COCO_HOSTNAME},g" $f
   fi

   sed -i "s,8080,${NETCOCO_PORT},g" $f
   cat $f
   done

   if [ -n "${COCO_HOSTNAME}" ]; then
      sed -i "s,localhost,${COCO_HOSTNAME},g" templates/js/main.js 
   fi

   sed -i "s,8080,${NETCOCO_PORT},g" templates/js/main.js 
}


if [[ -f ./config.yaml ]]; then
   eval $(parse_yaml config.yaml "config_")
   NETCOCO_PORT=$config_netcoco_port
   initial
   ./netcoco-linux-amd64 -file=./config.yaml
else
   if [[ -n "$COCO_PORT" ]]; then
      NETCOCO_PORT=$COCO_PORT
   fi

   if [[ -z "$COCO_DB_HOSTNAME" ]]; then
      echo "Please enter COCO_DB_HOSTNAME environment "
      exit
   elif [[ -z "$COCO_DB_USERNAME" ]]; then
      echo "Please enter COCO_DB_USERNAME environment "
      exit
   elif [[ -z "$COCO_DB_PASSWORD" ]]; then
      echo "Please enter COCO_DB_PASSWORD environment "
      exit
   elif [[ -z "$COCO_DB_NAME" ]]; then
      echo "Please enter COCO_DB_NAME environment "
      exit
   elif [[ -z "$COCO_SEC_SALT" ]]; then
      echo "Please enter COCO_SEC_SALT environment "
      exit
   fi

   initial
   ./netcoco-linux-amd64 -e
fi

