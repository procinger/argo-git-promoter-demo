
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout e200da7c1362ebbcde8f02c55cda000abd4c8d2b
kustomize build ./kustomize/overlays/dev
```