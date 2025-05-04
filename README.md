
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 618ba74d37b6ccfffa169f23ddb8c1b8deae1c46
kustomize build ./kustomize/overlays/test
```