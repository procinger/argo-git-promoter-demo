
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 7aa2b0741151fe685d87d69b7ed2ad6c08044be6
kustomize build ./kustomize/overlays/dev
```