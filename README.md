
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout cefaeb4373dbf6756f6f46a754648eed9194312e
kustomize build ./kustomize/overlays/dev
```