# kubectl-ttl ![Build](https://github.com/TejaBeta/kubectl-ttl/workflows/Go/badge.svg) [![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](./LICENSE)

A kubectl plugin to add `time-to-live` option to Kubernetes resources.

## Story time (Why?)

The story behind the creation of this tool is quite straight forward, we just required a tool to somehow kill/clean resources after certain point of time. 

To facilitate the above mentioned process, `kubectl-ttl` came into existence just to help myself and others in the K8s community to use the tool and make life easier. 

## How?

Nothing complex, based on the command parameters the tool creates a **`job`** within the namespace with the appropriate role to delete the given resources at a particular time within the namespace.

## Usage

```

#Default ttl of 15mins for a particular resource
kubectl get deployments -o yaml | kubectl ttl

#Custom ttl for 10mins for a particular resource
kubectl get pods -o yaml | kubectl ttl -t 10

```

## Help menu

```

A tiny kubectl plugin to add time to live option
to k8s resources within a namespace.

Tool helps to create a job within the specified namespace
to kill/clean the resources after certain time.

Usage:
  kubectl-ttl [flags]

Flags:
  -h, --help        help for kubectl-ttl
  -t, --time uint   time in minutes to keep the resource alive (default 15)

```

## Similar Projects

- [kube-janitor](https://github.com/hjacobs/kube-janitor)

## How can I help?

All kind of contributions are always welcome 👏, if you find the tool useful star ⭐️ the project. Always happy to help with any issues or new feature requests. 

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.
