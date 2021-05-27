## Cisco IOS

Preare for NetCoCo establish to network device 


### Config Script

```
  # Create an account
  (config)# username <username> secret <secret>

  # Set enable secret
  (config)# enable secret <secret>

  # Generate crypto rsa algorithm
  (config)# crypto key generate rsa modulus 2048

  # Set SSH to version 2
  (config)# ip ssh version 2

  # OPEN SSH Service
  (config)# line vty 0 4
  (config-line)# login local
  (config-line)# transport input ssh

```