# Cisco SG300

Prepare SSH services for NetCoCo establish to network device

## Config Script

```text
  # Create an account
  (config)# username <username> privilege 15 password <secret>

  # Set enable secret
  (config)# enable password level 15 <secret>

  # Generate crypto rsa algorithm use 1024 - 2048 bit
  (config)# crypto key generate rsa

  # To use password authentication for incoming SSH sessions
  (config)# ip ssh password-auth
```

