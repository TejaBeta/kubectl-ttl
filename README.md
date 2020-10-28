# kubectl-ttl

**WIP**

A kubectl plugin to add `time-to-live` option to Kubernetes resources to run experiments

## Story time (Why?)

The story behind the creation of this tool is quite straight forward, we just required a tool to somehow kill/clean resources after certain point of time. For some reasons original kubernetes doesn't support `time-to-live` option to **most of the resources**(there is an option to add ttl to jobs). 

To facilitate the above mentioned process, `kubectl-ttl` came into existence just to help myself and others in the K8s community to use the tool and make life easier. 

## How?

Nothing complex, based on the command parameters the tool creates a **`job`** within the namespace with the appropriate role to delete the given resources within a namespace.

## How can I help?

All kind of contributions are always welcome 👏, if you find the tool useful star ⭐️ the project. Always happy to help with any issues or new feature requests. 

## License
This project is distributed under the [Apache License, Version 2.0](http://www.apache.org/licenses/LICENSE-2.0), see [LICENSE](./LICENSE) for more information.