# Installation

> First of all, check it completely according to requirement.

1. Clone NetCoCo-UI.

   ```console
    # git clone https://github.com/mrzack99s/netcoco-ui
   ```

2. Install dependency.
   - Yarn \(Recommend\)

      ```console
      # yarn install
      ```

   - NPM

      ```console
      # npm install
      ```

3. Change base URL of API at file **.env.production**.

   ```text
    VUE_APP_VERSION=<NetCoCo-version>
    VUE_APP_ROOT_API=http://<Domain or Hostname>:8080
   ```

4. Build NetCoCo-UI with production.
   - Yarn \(Recommend\)

      ```console
      # yarn build --mode production
      ```

   - NPM

      ```console
      # npm build --mode production
      ```

5. Make **NetCoCo** directory and create sub-directory
   - Linux
   ```
      |usr
         |share
            |netcoco
               |templates
                  |dist
                  | .......
               |config.yaml
               |netcoco-linux-amd64
   ```
   > Make directory with <b>sudo mkdir /usr/share/netcoco</b>

   - Windows
   ```
      |NetCoCo
         |templates
            |dist
            | .......
         |config.yaml
         |netcoco-windows.exe
   ```
6. Download templates and executable file from [Github releases](https://github.com/mrzack99s/netcoco/releases) 
7. Untar or unzip templates
8. Copy **/dist** of NetCoCo-UI to **/templates** of NetCoCo.
9. Create **config.yaml** in NetCoCo directory with executable file. (In linux: /usr/share/netcoco)

    ```yaml
     netcoco:
         port: <port-number> #(default: 8080)
         db:
             sql:
             hostname: <db-hostname>:<db-port>
             username: <db-username>
             password: <db-password>
             db_name: <db-name>
         security:
             salt: <salt-in-hex-32digit>
    ```
10. Edit **main.js** in NetCoCo directory. (In linux: /usr/share/netcoco/templates/js)

    ```text
      const BASE_API_URL = 'http://<Hostname or Domain name>:<port>/v1.0-alpha/secure/api';
    ```

11. On Linux, must install the netcoco service
      - Download service file
      ```console
         # sudo curl -L https://raw.githubusercontent.com/mrzack99s/netcoco/main/netcoco.service -o /etc/systemd/system/netcoco.service
      ```
      - Enable and start service
      ```console
         # sudo systemctl enable --now netcoco
      ```
      - Enable sysctl ping
      ```console
        # sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
      ```
12. On Windows, run executable file.