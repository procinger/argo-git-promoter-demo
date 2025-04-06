
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout d21e8792ab5a30901da731ceb71c1a9aa962f648
kustomize build ./kustomize/overlays/test
```