
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 2e72beb90d1f9a385a28468a263be286d36e8b57
kustomize build ./kustomize/overlays/dev
```