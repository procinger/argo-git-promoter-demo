
# Manifest Hydration

To hydrate the manifests in this repository, run the following commands:

```shell

git clone git@github.com:procinger/argo-git-promoter-demo.git
# cd into the cloned directory
git checkout dab670760d1ec2734bafaf47a3b04137addfcfc6
kustomize build ./kustomize/overlays/test
```