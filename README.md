# NetCoCo \| Network Automation Application

NetCoCo is a network automation application for manage network device and config network device without CLI Configuration. NetCoCo can generate network audits with your freeform topology through the use of a Topology Builder.

## Feature

* Topology Builder
* Network Monitor
* L2 Switch configure via GUI
  * Vlan Setting
  * Interface setting

## Support Device

* Cisco IOS
* Cisco SG300
* Cisco SG350

## Requirement

* NetCoCo-UI  [go to repository](https://github.com/mrzack99s/netcoco-ui)

## Prepare System

* Install NodeJs 14.x for NetCoCo-UI

## How to install

1. Clone NetCoCo-UI.

   ```text
    git clone https://github.com/mrzack99s/netcoco-ui
   ```

2. Install dependency.
3. Yarn \(Recommend\)

   ```text
   yarn install
   ```

4. NPM

   ```text
   npm install
   ```

5. Change base URL of API at file **.env.production**.

   ```text
    VUE_APP_ROOT_API=http://<Domain or Hostname>:8080
   ```

6. Build NetCoCo-UI with production.
7. Yarn \(Recommend\)

   ```text
   yarn build --mode production
   ```

8. NPM

   ```text
   npm build --mode production
   ```

9. Copy **/dist** of NetCoCo-UI to **/templates** of NetCoCo.
10. Create **config.yaml** in NetCoCo. with

    ```text
     netcoco:
         db:
             sql:
             hostname: <db-hostname>:<db-port>
             username: <db-username>
             password: <db-password>
             db_name: <db-name>
         security:
             salt: <salt-in-hex-32digit>
    ```

## License

Copyright \(c\) 2020 - Chatdanai Phakaket

Licensed under the Apache License, Version 2.0 \(the "License"\); you may not use this file except in compliance with the License. You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

