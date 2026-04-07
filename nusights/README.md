nusights-api module provides a REST client for Insights endpoint.

This module is blind copy of `mesosphere/nkp-pulse` project.
Link : https://github.com/mesosphere/nkp-pulse/tree/53696f514f6cd3a288749047f25ae6bae26d0fe5/generated/nusights-api
The reason for duplicating the code is that difference of github organisation
between the NDK/CSI and NKP-Pulse project. Ideally such a common client code
should stay in some central repository & for time being it's being parked at
`prism go client`.


To generate the client code from swagger file, invoke:
`swagger generate client --spec=swagger.json --target .`



---

Folder layout

```
nusights/
├── client/                          # package client
│   ├── nutanix_insights_r_e_s_t_a_p_is_client.go   # top-level API client
│   ├── nutanix_insights_client_with_retry.go         # retry + factory helpers
│   └── config/                      # package config (sub-API)
│       ├── config_client.go                         # config API client
│       ├── put_config_data_parameters.go
│       └── put_config_data_responses.go
├── models/                          # shared DTOs
│   ├── entity.go
│   ├── entities_collection.go
│   └── ...
├── swagger.json
├── go.mod
└── README.md
```

---