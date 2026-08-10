---
title: Linux (Embedded Cluster)
visible_when:
  entitlements:
    - isEmbeddedClusterDownloadEnabled
---

# Linux (Embedded Cluster)

Install your application on a Linux server using Embedded Cluster. Read the docs or select your deployment preferences.

## Requirements

Confirm the target host meets these prerequisites before installing:

- Ubuntu 20.04+ / RHEL 8+ / CentOS 8+
- 4 CPUs, 8GB RAM, 40GB disk minimum
- Root or sudo access

## Choose an installation

<PendingInstallSelector method="linux" />

<NewInstall method="linux" />

<InstanceName method="linux" />

## Configuration

Customize the options below. The install commands will update automatically based on your selections.

<NetworkAvailability installType="linux" />
<VersionSelector installType="linux" />

## Install

<Warning>
The install command must be run as root or with sudo. Running without elevated privileges will cause the installation to fail.
</Warning>

SSH into your target machine and run the following command.

<Note>
The commands below are personalized to your selected installation. If you switch installations or rename your instance, the commands will update automatically.
</Note>

<LinuxInstallAssets />

## Verify Installation

<Tip title="Admin Console Access">
The admin console URL is printed at the end of the installation. Bookmark it — you'll use it to manage updates, check status, and configure the application.
</Tip>

Once the installer completes, it will print the URL for the admin console. Open it in your browser to continue with application setup.

<CommandBlock>
# Check that all pods are running
kubectl get pods -A

# Access the admin console
echo "Admin Console: https://$(hostname):8800"
</CommandBlock>

## Post-Install

See the post-installation documentation for next steps including configuring TLS, setting up backups, and connecting to your identity provider.
