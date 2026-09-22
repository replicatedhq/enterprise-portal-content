---
id: 2026-09-21-kots-online-cli-downloads
title: Workstation CLIs on the default KOTS online page
published_at: 2026-09-21T00:00:00Z
impact: recommended
summary: The default Existing Cluster (KOTS) online and proxy path includes a CLI-only download list (KOTS, Preflight, Support Bundle) after the generated install commands. Air gap downloads stay in the air gap branch. Do not unwrap the air gap tag. Adopt this page only after the portal build that supports cliOnly (vandoor #10549).
affects:
  - installation
  - kots
---

The default page mounts `<KotsDownloadAssets cliOnly={true} />` on the online/proxy branch. Copy leads with Preflight and Support Bundle CLIs (those are not curl-dependent). Customers who cannot run `curl https://kots.io/install` download the KOTS CLI here, extract it, put `kubectl-kots` on PATH, then run the `kubectl kots install` line from step 2. The generated curl command remains step 2 for environments that can use it. This list does not include the Admin Console bundle or the application air gap bundle.

The air gap branch is unchanged: full `<KotsDownloadAssets />` inside `<WhenNetwork mode="airgap">`. Do not move that tag out of the wrap. `cliOnly` keys off the mount, so a dual-entitled customer who selected **online** does not get air-gap installer files in that flow.

Do not adopt this page until the portal lists `cliOnly` on `KotsDownloadAssets` (vandoor #10549). If you ship the page first, an older portal may strip the prop and show Admin Console and air-gap bundles on the online branch for customers whose license also has air gap.
