# Manifest installation

### Requirements
*   Kubectl version >= 1.18.x
---
### How to install

1.  Download **manifest.yaml** file
    *  In linux
        ```console
        # curl -LO https://raw.githubusercontent.com/mrzack99s/netcoco/main/manifest.yaml
        ``` 
2.  Change the value in **manifest.yaml**
3.  Deploy NetCoCo with **manifest.yaml**
    ```console
    # kubectl apply -f manifest.yaml
    ```
---
### How to uninstall
*   Undeploy NetCoCo with **manifest.yaml**
    ```console
    # kubectl delete -f manifest.yaml
    ```