# How to install


> First of all, check it completely according to requirement.

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

11. Run executable file.