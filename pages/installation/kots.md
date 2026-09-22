---
title: Existing Cluster (KOTS)
visible_when:
  entitlements:
    - isKotsInstallEnabled
---

# Existing Cluster (KOTS)

Install your application into an existing Kubernetes cluster using the
Replicated KOTS Admin Console. Select your network availability below. The
install steps update for online, proxy, or air gap.

## Requirements

- Kubernetes cluster with `kubectl` access
- If you install air gap, select **No outbound requests allowed (air gap)**
  under Configuration. Extra registry requirements for that path appear below.

<WhenNetwork mode="airgap">

- A compatible private image registry inside the air-gapped network. This
  portal does not provide the registry. One application uses one registry
  namespace; that namespace must exist before you push images. Amazon ECR
  does not use namespaces.
- Two accounts on **that registry** (not credentials from this portal or
  from your license):
  - **Read-write** for pushing Admin Console images. KOTS does not store or
    reuse these credentials.
  - **Read-only** for installing the Admin Console. KOTS persists these
    credentials as a pull secret in the cluster.

</WhenNetwork>

## Choose an installation

<PendingInstallSelector method="kots" />

<NewInstall method="kots" />

<InstanceName method="kots" />

## Configuration

Customize the options below. Online and proxy commands update from your
selections. Air gap uses the procedure on this page.

<NetworkAvailability installType="kots" />
<VersionSelector installType="kots" />

## Install

<InstallStep stepNumber={1} title="Download your license">

<LicenseDownload />

</InstallStep>

<WhenNetwork mode="online,proxy">

<Note>
The commands below are personalized to your selected installation. If you
switch installations or rename your instance, the commands will update
automatically.
</Note>

<KotsInstallAssets stepNumber={2} />

<InstallStep stepNumber={3} title="Workstation CLIs">

Step 2 already installs the KOTS CLI via `curl https://kots.io/install`. Download
a version-pinned KOTS CLI, Preflight CLI, or Support Bundle CLI here if you need
them. This list does not include the Admin Console bundle or the application
air gap bundle.

<KotsDownloadAssets cliOnly={true} />

</InstallStep>

</WhenNetwork>

<WhenNetwork mode="airgap">

<InstallStep stepNumber={2} title="Download the installation assets">

Select a version to download. Required: the KOTS CLI, the Admin Console bundle,
and the application air gap bundle. Optional: the Preflight and Support Bundle
CLIs.

<KotsDownloadAssets />

</InstallStep>

<InstallStep stepNumber={3} title="Install the KOTS CLI">

For air gap, use the KOTS CLI you downloaded from this page. Do not install it
with `curl https://kots.io/install` (that path is for online installs only).

The download is a `.tar.gz` (linux amd64 from this page). Extract it, rename
`kots` to `kubectl-kots`, and move it onto your PATH. The CLI version must match
the Admin Console bundle.

<CommandBlock>
tar xvf kots_linux_amd64.tar.gz
sudo mv kots /usr/local/bin/kubectl-kots
kubectl kots version
</CommandBlock>

</InstallStep>

<InstallStep stepNumber={4} title="Push Admin Console images">

Use the read-write account. KOTS does not store these credentials. Substitute
the placeholders before running. Use the Admin Console bundle file you
downloaded in step 2 (the filename is usually `kotsadm.tar.gz`, or
`kotsadm-nominio.tar.gz` for some teams).

<CommandBlock label="example">
kubectl kots admin-console push-images ./ADMIN_CONSOLE_BUNDLE.tar.gz REGISTRY_HOST \
  --registry-username RW_USERNAME \
  --registry-password RW_PASSWORD
</CommandBlock>

Substitute (from your registry, not from this page):

- `ADMIN_CONSOLE_BUNDLE.tar.gz` with the Admin Console bundle filename from
  your downloads.
- `REGISTRY_HOST` with the registry hostname, or host plus namespace. For
  example `private.registry.host` or `my-registry.example.com/my-namespace`.
- `RW_USERNAME` and `RW_PASSWORD` with a registry account that can push.

</InstallStep>

<InstallStep stepNumber={5} title="Install the Admin Console">

Use the same `REGISTRY_HOST` as the previous step. Use the **read-only**
account, not the read-write pair. KOTS stores the read-only credentials as an
imagePullSecret on Admin Console pods.

`--namespace` is the application slug, not Kubernetes `default`. Kubernetes
limits namespace names to 63 characters; if your app slug is longer, pick a
shorter namespace. When prompted, set the Admin Console password. That password
is not on this page.

If this customer is not on the `stable` channel, append `/` and the channel
slug to the app argument (for example `{{app.slug}}/beta`).

<CommandBlock label="example">
kubectl kots install {{app.slug}} \
  --kotsadm-registry REGISTRY_HOST \
  --registry-username RO_USERNAME \
  --registry-password RO_PASSWORD \
  --namespace {{app.slug}}
</CommandBlock>

Substitute (from your registry, not from this page):

- `REGISTRY_HOST` with the same host or host/namespace you used for
  `push-images`.
- `RO_USERNAME` and `RO_PASSWORD` with a registry account that can pull.
  Do not reuse the read-write pair here.

</InstallStep>

<InstallStep stepNumber={6} title="Upload the license and air gap bundle">

When install finishes it prints a port-forward to the Admin Console. Open
`http://localhost:8800` and log in with the password you set.

If you need to reopen the port-forward:

<CommandBlock>
kubectl kots admin-console -n {{app.slug}}
</CommandBlock>

In the Admin Console, upload your license, then upload the application
`.airgap` bundle. Complete configuration and preflight checks, then deploy.

</InstallStep>

</WhenNetwork>
