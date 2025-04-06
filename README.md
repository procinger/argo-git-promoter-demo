
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 4b68b97780c4a828e8bd3fb67dc59884f4f726b9
kustomize build ./kustomize/overlays/dev
```