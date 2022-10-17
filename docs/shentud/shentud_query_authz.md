## shentud query authz

Querying commands for the authz module

```
shentud query authz [flags]
```

### Options

```
  -h, --help   help for authz
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID
      --home string         directory for config and data (default "/home/ylee/.shentud")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

### SEE ALSO

* [shentud query](shentud_query.md)	 - Querying subcommands
* [shentud query authz grantee-grants](shentud_query_authz_grantee-grants.md)	 - query authorization grants granted to a grantee
* [shentud query authz granter-grants](shentud_query_authz_granter-grants.md)	 - query authorization grants granted by granter
* [shentud query authz grants](shentud_query_authz_grants.md)	 - query grants for a granter-grantee pair and optionally a msg-type-url

###### Auto generated by spf13/cobra on 18-Aug-2022