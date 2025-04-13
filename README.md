
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout 92786f31df3939a69a00f1dbaf5992ab55409d44
kustomize build ./kustomize/overlays/dev
```