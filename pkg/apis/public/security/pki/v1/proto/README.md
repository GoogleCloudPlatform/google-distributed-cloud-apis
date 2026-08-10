# PKI v1 Protobuf API

This directory contains the protocol buffer definitions for the Certificate Authority Service (CAS) API.

## API Design Guidelines

Please refer to the [Google API Improvement Proposals (AIP)](https://google.aip.dev/1) for standard protobuf API design patterns, naming conventions, and style guidelines.

## Code Generation

To generate the Go API client and server code from the `.proto` definitions, run the following Bazel command:

```bash
bazel run //pkg/apis/public/security/pki/v1/proto:pki_go_proto_gen_srcs
```
