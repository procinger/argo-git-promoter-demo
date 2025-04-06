
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 3487e8926d4598ab1b9138d08fe9958dc46ff2a2
kustomize build ./kustomize/overlays/test
```